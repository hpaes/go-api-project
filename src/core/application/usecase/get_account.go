package usecase

import (
	"context"

	"github.com/hpaes/go-api-project/src/infrastructure/logger"
	"github.com/hpaes/go-api-project/src/infrastructure/repository"
	"github.com/pkg/errors"
)

type (
	GetAccount interface {
		Execute(ctx context.Context, accountId string) (*GetAccountOutput, error)
	}
	getAccount struct {
		accountRepository repository.AccountRepository
		logger            logger.LogHandler
	}
)

func NewAccountUseCase(accountRepository repository.AccountRepository, logger logger.LogHandler) GetAccount {
	return &getAccount{
		accountRepository: accountRepository,
		logger:            logger,
	}
}

func (ga *getAccount) Execute(ctx context.Context, accountId string) (*GetAccountOutput, error) {
	ga.logger.LogInformation("Executing GetAccountUseCase for accountId: %s", accountId)
	account, err := ga.accountRepository.GetById(ctx, accountId)
	if err != nil {
		ga.logger.LogError("Error getting account: %v", err)
		return nil, errors.Wrap(err, "error getting account by id")
	}
	if account.AccountId == "" {
		ga.logger.LogInformation("Account not found")
		return nil, nil
	}
	ga.logger.LogInformation("Successfully retrieved account: %s", accountId)

	return &GetAccountOutput{
		AccountId:   account.AccountId,
		Name:        account.Name.Value,
		Cpf:         account.Cpf.Value,
		Email:       account.Email.Value,
		CarPlate:    account.CarPlate.Value,
		IsPassenger: account.IsPassenger,
		IsDriver:    account.IsDriver,
	}, nil

}

type GetAccountOutput struct {
	AccountId   string `json:"accountId"`
	Name        string `json:"name"`
	Cpf         string `json:"cpf"`
	Email       string `json:"email"`
	CarPlate    string `json:"carPlate"`
	IsPassenger bool   `json:"isPassenger"`
	IsDriver    bool   `json:"isDriver"`
}
