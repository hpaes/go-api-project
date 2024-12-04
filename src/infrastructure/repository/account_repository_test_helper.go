package repository

import (
	"context"

	"github.com/hpaes/go-api-project/src/core/domain"
	"github.com/stretchr/testify/mock"
)

type accountRepositoryMock struct {
	mock.Mock
}

func NewMockAccountRepository() *accountRepositoryMock {
	return &accountRepositoryMock{}
}

func (ar *accountRepositoryMock) Save(ctx context.Context, account *domain.Account) error {
	args := ar.Called(ctx, account)
	return args.Error(0)
}

func (ar *accountRepositoryMock) GetById(ctx context.Context, id string) (*domain.Account, error) {
	args := ar.Called(ctx, id)
	return args.Get(0).(*domain.Account), args.Error(1)
}

func (ar *accountRepositoryMock) GetByEmail(ctx context.Context, email string) (*domain.Account, error) {
	args := ar.Called(ctx, email)
	return args.Get(0).(*domain.Account), args.Error(1)
}
