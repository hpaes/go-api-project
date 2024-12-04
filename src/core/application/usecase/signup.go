package usecase

import (
	"context"

	"github.com/hpaes/go-api-project/src/core/domain"
	"github.com/hpaes/go-api-project/src/infrastructure/logger"
	"github.com/hpaes/go-api-project/src/infrastructure/repository"
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

type (
	SignUp interface {
		Execute(ctx context.Context, input SignupInput) (*SignupOutput, error)
	}

	signUp struct {
		accountRepository repository.AccountRepository
		logger            logger.LogHandler
	}
)

func NewSignupUseCase(accountRepository repository.AccountRepository, logger logger.LogHandler) SignUp {
	return &signUp{
		accountRepository: accountRepository,
		logger:            logger,
	}
}

func (su *signUp) Execute(ctx context.Context, input SignupInput) (*SignupOutput, error) {
	su.logger.LogInformation("Executing SignupUseCase for email: %s", input.Email)

	err := su.validateInput(input)
	if err != nil {
		return nil, err
	}

	existingAccount, err := su.accountRepository.GetByEmail(ctx, input.Email)
	if err != nil {
		su.logger.LogError("Error getting account by email: %v", err)
		return nil, errors.Wrap(err, "error getting account by email")
	}
	if existingAccount != nil && existingAccount.AccountId != "" {
		su.logger.LogError("Account already exists for email: %s", input.Email)
		return nil, errors.New("account already exists")
	}

	account, err := domain.NewAccount("", input.Name, input.Cpf, input.Email, input.CarPlate, input.IsPassenger, input.IsDriver)
	if err != nil {
		su.logger.LogError("Error creating account: %v", err)
		return nil, errors.Wrap(err, "error creating account")
	}

	err = su.accountRepository.Save(ctx, account)
	if err != nil {
		su.logger.LogError("Error saving account: %v", err)
		return nil, errors.Wrap(err, "error saving account")
	}

	su.logger.LogInformation("Account created successfully: %s", account.AccountId)
	return &SignupOutput{AccountId: account.AccountId}, nil
}

func (su *signUp) validateInput(input SignupInput) error {

	switch {
	case input.Email == "":
		su.logger.LogError("Invalid input: email is required")
		return errors.New("email is required")
	case input.Name == "":
		su.logger.LogError("Invalid input: name is required")
		return errors.New("name is required")

	case input.Cpf == "":
		su.logger.LogError("Invalid input: cpf is required")
		return errors.New("cpf is required")
	default:
		return nil
	}
}
