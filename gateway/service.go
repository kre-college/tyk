package gateway

import (
	"context"
	"github.com/TykTechnologies/tyk/user"
)

type Service struct {
	db *DB
}

func NewService(db *DB) *Service {
	return &Service{
		db: db,
	}
}

func (r *Service) CreatePolicy(ctx context.Context, policy *user.Policy) error {
	err := r.db.CreatePolicy(ctx, policy)
	if err != nil {
		return err
	}
	return nil
}
