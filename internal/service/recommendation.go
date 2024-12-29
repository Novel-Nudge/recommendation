package service

import (
	"context"
	"fmt"
	"recommendation/internal/config"
	"recommendation/internal/domain"
	"recommendation/internal/repository/postgres"
	"recommendation/pkg/recommendation"
	"recommendation/pkg/vector"
)

type RecommendationService struct {
	// Repository implementations
	ur postgres.UserRepository
	br postgres.BookRepository
	ar postgres.ActionsRepository

	config config.RecommendationConfig
}

func NewRecommendationService(
	ur postgres.UserRepository,
	br postgres.BookRepository,
	ar postgres.ActionsRepository,
	config config.RecommendationConfig,
) *RecommendationService {
	return &RecommendationService{
		ur:     ur,
		br:     br,
		ar:     ar,
		config: config,
	}
}

// GetRecommendations Generates n recommendations for a given user id
func (s *RecommendationService) GetRecommendations(ctx context.Context, userID string, limit int) ([]domain.FinalRecommendation, error) {

	// Grab the user vec for collab and content filtering
	userVec, err := s.ur.FetchUserVec(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve user vector %s", err)
	}

	// Search for similar book embeddings using the user vector and map to pkg model
	contBooks, err := s.br.FindNearestBooks(ctx, userVec, limit)
	if err != nil {
		return nil, fmt.Errorf("Failed similarity search for books %s", err)
	}

	// Find the nearest n users using similarity search
	nearUsers, err := s.ur.FindNearestUsers(ctx, userVec, userID, limit)
	if err != nil {
		return nil, fmt.Errorf("Failed finding close users %s", err)
	}

	// Lookup the 10 nearest users, select actions that are likes and map to pkg model
	nearUserActions, err := s.ar.FetchBatchUserActions(ctx, nearUsers.UserIDs, domain.ActionLikeDB, 10)
	if err != nil {
		return nil, fmt.Errorf("Failed retrieving multi user actions %s", err)
	}

	// Map the content and collab to the recommendation pkg
	pkgContent := mapContToPkg(contBooks)
	pkgCollab := mapCollabToPkg(nearUserActions, nearUsers.Mapping)

	// Combine the collab/content filter recommendations and map to return model
	config := recommendation.ScoringConfig{
		ContentWeight:   s.config.ContentWeight,
		CollabWeight:    s.config.CollabWeight,
		DistanceDecay:   s.config.DistanceDecay,
		MinContentScore: s.config.MinContentScore,
		MinCollabScore:  s.config.MinCollabScore,
	}
	recommender := recommendation.NewRecommender(&config)
	recommendedBooks := recommender.CombineRecommendations(pkgContent, pkgCollab)
	if err != nil {
		return nil, fmt.Errorf("Failed to combine recommendations %s", err)
	}

	// Map to the domain model and return
	finalRecommendations := mapRecommendationsToDomain(recommendedBooks)
	return finalRecommendations, nil
}

func (s *RecommendationService) UpdateUserFeats(ctx context.Context, userID string) error {
	// Retrieve the initial user features and map to domain model
	features, err := s.ur.FetchUserFeats(ctx, userID)
	if err != nil {
		return fmt.Errorf("Failed to retrieve user features %s", err)
	}

	// Retrieve actions and map to domain model
	userActions, err := s.ar.FetchUserActions(ctx, userID, s.config.UpdateLength)
	if err != nil {
		return fmt.Errorf("Failed retrieving user actions %s", err)
	}

	// Create a new instance of the feature updater
	config := recommendation.UpdateConfig{
		ContentLR:       s.config.ContentLR,
		CollabLR:        s.config.CollabLR,
		EmbeddingLR:     s.config.EmbeddingLR,
		TimeDecay:       s.config.TimeDecay,
		ConfidenceDecay: s.config.ConfidenceDecay,
		MinConfidence:   s.config.MinConfidence,
		ContentWeight:   s.config.ContentWeight,
		CollabWeight:    s.config.ContentWeight,
		L2Reg:           s.config.L2Reg,
	}
	featureUpdater := recommendation.NewFeatureUpdater(&config)

	// Map models to pkg and update the features
	pkgFeatures := toPkgFeatures(features)
	pkgActions := toPkgUserActions(userActions)
	err = featureUpdater.UpdateFeatures(&pkgFeatures, pkgActions)
	if err != nil {
		return fmt.Errorf("Failed to calculate new features %s", err)
	}

	// Map pkg back to domain and update features
	updatedFeatures := toDomainFeatures(pkgFeatures)
	err = s.ur.UpdateUserFeats(ctx, userID, updatedFeatures)
	if err != nil {
		return fmt.Errorf("Failed to update the user features %s", err)
	}

	// Create the vector and update the value in the db using the new updated features
	vecPkgFeatures := toVecPkgFeatures(updatedFeatures)
	v := vector.NewVector(s.config.VectorSize)
	v.FromFeatures(vecPkgFeatures)

	err = s.ur.UpdateVec(ctx, userID, v.Get())
	if err != nil {
		return fmt.Errorf("Failed to update vector %s", err)
	}

	return nil
}
