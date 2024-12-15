package usecase

import (
	"context"
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/hpaes/go-api-project/src/core/domain"
	"github.com/hpaes/go-api-project/src/core/domain/valueObjects"
	"github.com/hpaes/go-api-project/src/infrastructure/database"
	"github.com/hpaes/go-api-project/src/infrastructure/logger"
	"github.com/hpaes/go-api-project/src/infrastructure/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupGetAccount(t *testing.T) (*database.PqAdapter, GetAccount, func()) {
	// setup db connection
	dbConnection, err := database.NewPqAdapter()
	if err != nil {
		log.Fatal(err)
	}
	// create acc repository
	accountRepository := repository.NewAccountRepository(dbConnection)

	// create logger
	logger := logger.NewConsoleLogger()
	// create get account use case
	getAccountUseCase := NewAccountUseCase(accountRepository, logger)

	cleanup := func() {
		err := dbConnection.ExecWithContext(context.Background(), "DELETE FROM brq_golang.account")
		assert.NoError(t, err)
	}
	return dbConnection, getAccountUseCase, cleanup
}
func TestGetAccount(t *testing.T) {
	t.SkipNow()
	dbConnection, getAccountUseCase, cleanup := setupGetAccount(t)
	defer cleanup()

	ctx := context.Background()

	acc, err := domain.NewAccount("", "John Doe", "123.456.789-09", "johnDoe@emai.com", "ABC-1B34", true, false)
	assert.NoError(t, err)
	assert.NotNil(t, acc)

	err = dbConnection.ExecWithContext(ctx, "INSERT INTO brq_golang.account (account_id, name, cpf, email, car_plate, is_passenger, is_driver) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		acc.AccountId, acc.Name.Value, acc.Cpf.Value, acc.Email.Value, acc.CarPlate.Value, acc.IsPassenger, acc.IsDriver)
	assert.NoError(t, err)

	output, err := getAccountUseCase.Execute(ctx, acc.AccountId)
	assert.NoError(t, err)

	assert.Equal(t, acc.AccountId, output.AccountId)
	assert.Equal(t, acc.Name.Value, output.Name)
	assert.Equal(t, acc.Cpf.Value, output.Cpf)
	assert.Equal(t, acc.Email.Value, output.Email)
	assert.Equal(t, acc.CarPlate.Value, output.CarPlate)
	assert.Equal(t, acc.IsPassenger, output.IsPassenger)
	assert.Equal(t, acc.IsDriver, output.IsDriver)
}

func TestGetAccountNotFound(t *testing.T) {
	t.SkipNow()
	_, getAccountUseCase, cleanup := setupGetAccount(t)
	defer cleanup()

	ctx := context.Background()

	output, err := getAccountUseCase.Execute(ctx, uuid.New().String())
	assert.NoError(t, err)
	assert.Nil(t, output)
}

func TestGetAccountUsecaseWithMock(t *testing.T) {
	r := repository.NewMockAccountRepository()
	logger := logger.NewConsoleLogger()
	uc := NewAccountUseCase(r, logger)

	expectedOuput := &domain.Account{
		AccountId:   "123",
		Name:        valueObjects.Name{Value: "John Doe"},
		Cpf:         valueObjects.Cpf{Value: "123.456.789-09"},
		Email:       valueObjects.Email{Value: "johnDoe@email.com"},
		CarPlate:    valueObjects.CarPlate{Value: "ABC-1B34"},
		IsPassenger: true,
		IsDriver:    false,
	}

	r.On("GetById", mock.Anything, mock.Anything).Return(expectedOuput, nil)
	ctx := context.TODO()
	output, err := uc.Execute(ctx, "123")
	assert.NoError(t, err)
	assert.NotNil(t, output)
}

func TestGetAccountUsecaseWithMockWhenAccountNotFound(t *testing.T) {
	r := repository.NewMockAccountRepository()
	logger := logger.NewConsoleLogger()
	uc := NewAccountUseCase(r, logger)

	r.On("GetById", mock.Anything, mock.Anything).Return(&domain.Account{}, nil)
	ctx := context.TODO()
	output, err := uc.Execute(ctx, "123")
	assert.NoError(t, err)
	assert.Nil(t, output)
}
