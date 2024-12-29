package postgres

import (
	"context"
	"database/sql"
	"recommendation/internal/domain"
)

type ActionsRepository struct {
	db *sql.DB
}

func (r *ActionsRepository) FetchBatchUserActions(ctx context.Context, userID []string, actionType domain.ActionType, n int) ([]domain.UserAction, error) {
	return nil, nil
}

func (r *ActionsRepository) FetchUserActions(ctx context.Context, userID string, n int) ([]domain.UserAction, error) {
	return nil, nil
}
