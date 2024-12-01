package usecase

import (
	"context"

	"github.com/hpaes/go-api-project/src/core/domain"
	"github.com/hpaes/go-api-project/src/core/domain/logger"
	domain_repository "github.com/hpaes/go-api-project/src/core/domain/repository"
	"github.com/pkg/errors"
)

type SignupInput struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Cpf         string `json:"cpf"`
	CarPlate    string `json:"carPlate"`
	IsPassenger bool   `json:"isPassenger"`
	IsDriver    bool   `json:"isDriver"`
}

type SignupOutput struct {
	AccountId string `json:"accountId"`
}

type SignupUseCase struct {
	accountRepository domain_repository.AccountRepository
	logger            logger.LogHandler
}

func NewSignupUseCase(accountRepository domain_repository.AccountRepository, logger logger.LogHandler) *SignupUseCase {
	return &SignupUseCase{
		accountRepository: accountRepository,
		logger:            logger,
	}
}

func (su *SignupUseCase) Execute(ctx context.Context, input SignupInput) (*SignupOutput, error) {
	su.logger.LogInformation("Executing SignupUseCase for email: %s", input.Email)
	existingAccount, err := su.accountRepository.GetByEmail(ctx, input.Email)
	if err != nil {
		su.logger.LogError("Error getting account: %v", err)
		return nil, errors.Wrap(err, "error getting account by email")
	}
	if existingAccount != nil {
		su.logger.LogError("Account already exists")
		return nil, errors.New("account already exists")
	}

	account, err := domain.CreateAccount(input.Name, input.Cpf, input.Email, input.CarPlate, input.IsPassenger, input.IsDriver)
	if err != nil {
		su.logger.LogError("Error creating account: %v", err)
		return nil, errors.Wrap(err, "error creating account")
	}

	su.logger.LogInformation("Saving account: %s", account.AccountId)
	su.accountRepository.Save(ctx, account)
	return &SignupOutput{AccountId: account.AccountId}, nil
}
