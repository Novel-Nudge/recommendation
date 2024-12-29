package service

import (
	"recommendation/internal/domain"
	"recommendation/pkg/recommendation"
	"recommendation/pkg/vector"
)

func mapContToPkg(contBooks []domain.NearestBook) []recommendation.ContentMatch {
	contentMatches := make([]recommendation.ContentMatch, len(contBooks))

	for i, book := range contBooks {
		contentMatches[i] = recommendation.ContentMatch{
			BookID:    book.BookID,
			Distance:  book.Distance,
			Embedding: book.Embedding,
		}
	}
	return contentMatches
}

func mapCollabToPkg(collabBooks []domain.UserAction, userMapping map[string]float64) []recommendation.CollabMatch {
	collabMatches := make([]recommendation.CollabMatch, len(collabBooks))

	for i, book := range collabBooks {
		collabMatches[i] = recommendation.CollabMatch{
			BookID:        book.BookID,
			UserDistance:  userMapping[book.UserID],
			UserBookScore: book.FinalScore,
		}
	}
	return collabMatches
}

func mapRecommendationsToDomain(recommendedBooks []recommendation.FinalRecommendation) []domain.FinalRecommendation {
	// Finally map the pkg model to the domain model
	finalRecommendations := make([]domain.FinalRecommendation, len(recommendedBooks))
	for i, book := range recommendedBooks {
		finalRecommendations[i] = domain.FinalRecommendation{
			BookID:      book.BookID,
			FinalScore:  book.FinalScore,
			ContentRank: book.ContentRank,
			CollabRank:  book.CollabRank,
		}
	}

	return finalRecommendations
}

// Batch conversion helpers
func toPkgUserActions(actions []domain.UserAction) []recommendation.UserAction {
	pkgActions := make([]recommendation.UserAction, len(actions))
	for i, a := range actions {
		pkgActions[i] = recommendation.UserAction{
			BookID:            a.BookID,
			FinalScore:        a.FinalScore,
			ContentRank:       (a.ContentRank),
			CollabRank:        (a.CollabRank),
			Action:            toPkgActionType(a.Action),
			OverlappingTopics: a.OverlappingTopics,
			BookVector:        a.BookVector,
		}
	}
	return pkgActions
}

func toPkgActionType(action domain.ActionType) int {
	switch action {
	case domain.ActionLikeDB:
		return 1
	case domain.ActionReadDB:
		return 0
	case domain.ActionDislikeDB:
		return -1
	default:
		return 0
	}
}

// Map domain Features to pkg Features
func toVecPkgFeatures(f domain.Features) vector.Features {
	return vector.Features{

		// Amount of times we've updated the user features
		Updates: f.Updates,

		// Core traits
		Complexity: f.Complexity,
		Emotion:    f.Emotion,
		Pace:       f.Pace,
		// Genre weights
		LiteraryFiction:   f.LiteraryFiction,
		Mystery:           f.Mystery,
		Romance:           f.Romance,
		ScienceFiction:    f.ScienceFiction,
		Fantasy:           f.Fantasy,
		Thriller:          f.Thriller,
		HistoricalFiction: f.HistoricalFiction,
		Contemporary:      f.Contemporary,
		// Mood weights
		Uplifting:   f.Uplifting,
		Dark:        f.Dark,
		Mysterious:  f.Mysterious,
		Romantic:    f.Romantic,
		Humorous:    f.Humorous,
		Serious:     f.Serious,
		Melancholic: f.Melancholic,
		Inspiring:   f.Inspiring,
		// Additional characteristics
		ReadingLevel:    f.ReadingLevel,
		VocabularyLevel: f.VocabularyLevel,
		Experimental:    f.Experimental,
	}
}

// Map domain Features to pkg Features
func toPkgFeatures(f domain.Features) recommendation.Features {
	return recommendation.Features{

		// Amount of times we've updated the user features
		Updates: f.Updates,

		// Core traits
		Complexity: f.Complexity,
		Emotion:    f.Emotion,
		Pace:       f.Pace,
		// Genre weights
		LiteraryFiction:   f.LiteraryFiction,
		Mystery:           f.Mystery,
		Romance:           f.Romance,
		ScienceFiction:    f.ScienceFiction,
		Fantasy:           f.Fantasy,
		Thriller:          f.Thriller,
		HistoricalFiction: f.HistoricalFiction,
		Contemporary:      f.Contemporary,
		// Mood weights
		Uplifting:   f.Uplifting,
		Dark:        f.Dark,
		Mysterious:  f.Mysterious,
		Romantic:    f.Romantic,
		Humorous:    f.Humorous,
		Serious:     f.Serious,
		Melancholic: f.Melancholic,
		Inspiring:   f.Inspiring,
		// Additional characteristics
		ReadingLevel:    f.ReadingLevel,
		VocabularyLevel: f.VocabularyLevel,
		Experimental:    f.Experimental,
	}
}

// Map pkg Features to domain Features
func toDomainFeatures(f recommendation.Features) domain.Features {
	return domain.Features{
		// Amount of times we've updated the user features
		Updates: f.Updates,

		// Core traits
		Complexity: f.Complexity,
		Emotion:    f.Emotion,
		Pace:       f.Pace,
		// Genre weights
		LiteraryFiction:   f.LiteraryFiction,
		Mystery:           f.Mystery,
		Romance:           f.Romance,
		ScienceFiction:    f.ScienceFiction,
		Fantasy:           f.Fantasy,
		Thriller:          f.Thriller,
		HistoricalFiction: f.HistoricalFiction,
		Contemporary:      f.Contemporary,
		// Mood weights
		Uplifting:   f.Uplifting,
		Dark:        f.Dark,
		Mysterious:  f.Mysterious,
		Romantic:    f.Romantic,
		Humorous:    f.Humorous,
		Serious:     f.Serious,
		Melancholic: f.Melancholic,
		Inspiring:   f.Inspiring,
		// Additional characteristics
		ReadingLevel:    f.ReadingLevel,
		VocabularyLevel: f.VocabularyLevel,
		Experimental:    f.Experimental,
	}
}
