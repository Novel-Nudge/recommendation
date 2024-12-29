package postgres

import (
	"context"
	"database/sql"
	"recommendation/internal/domain"
)

type BookRepository struct {
	db *sql.DB
}

func (r *BookRepository) FindNearestBooks(ctx context.Context, searchVec []float64, n int) ([]domain.NearestBook, error) {
	return nil, nil
}
