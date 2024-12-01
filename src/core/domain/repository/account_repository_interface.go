package domain_repository

import (
	"context"

	"github.com/hpaes/go-api-project/src/core/domain"
)

type AccountRepository interface {
	GetById(ctx context.Context, id string) (*domain.Account, error)
	GetByEmail(ctx context.Context, email string) (*domain.Account, error)
	Save(ctx context.Context, account *domain.Account) error
}
