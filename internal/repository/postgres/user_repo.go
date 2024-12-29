package postgres

import (
	"context"
	"database/sql"
	"recommendation/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

func (r *UserRepository) FindNearestUsers(ctx context.Context, searchVec []float64, userID string, n int) (domain.NearestUsers, error) {
	return domain.NearestUsers{}, nil
}

func (r *UserRepository) FetchUserFeats(ctx context.Context, userID string) (domain.Features, error) {
	return domain.Features{}, nil
}

func (r *UserRepository) FetchUserVec(ctx context.Context, userID string) ([]float64, error) {
	return nil, nil
}

func (r *UserRepository) UpdateUserFeats(ctx context.Context, userID string, newFeats domain.Features) error {
	return nil
}

func (r *UserRepository) UpdateVec(ctx context.Context, userID string, embedding []float64) error {
	return nil
}
