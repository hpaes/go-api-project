package usecase

import (
	"context"
	"testing"

	"github.com/hpaes/go-api-project/src/infrastructure/database"
	"github.com/hpaes/go-api-project/src/infrastructure/logger"
	"github.com/hpaes/go-api-project/src/infrastructure/repository"
	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T) (*SignupUseCase, *GetAccountUseCase, func()) {
	// setup db connection
	dbConnection, err := database.NewPqAdapter()
	if err != nil {
		panic(err)
	}
	// create acc repository
	accountRepository := repository.NewAccountRepository(dbConnection)

	// create logger
	logHandler := &logger.ConsoleLogger{}

	// create signup use case
	signupUseCase := NewSignupUseCase(accountRepository, logHandler)

	// create get account use case
	getAccountUseCase := NewAccountUseCase(accountRepository, logHandler)

	cleanup := func() {
		err := dbConnection.ExecWithContext(context.Background(), "DELETE FROM brq_golang.account")
		assert.NoError(t, err)
	}
	return signupUseCase, getAccountUseCase, cleanup
}

func TestSignupUseCase_Execute_ValidAccount(t *testing.T) {

	signupUseCase, getAccountUseCase, cleanup := setup(t)
	defer cleanup()

	ctx := context.Background()

	inputSignup := SignupInput{
		Name:        "John Doe",
		Email:       "JohnDoe@email.com",
		Cpf:         "123.456.789-09",
		CarPlate:    "ABC-1B34",
		IsPassenger: true,
		IsDriver:    false,
	}

	outputSignup, err := signupUseCase.Execute(ctx, inputSignup)
	assert.NoError(t, err)
	assert.NotNil(t, outputSignup)
	outputGetAccount, err := getAccountUseCase.Execute(ctx, outputSignup.AccountId)
	assert.NoError(t, err)
	assert.NotNil(t, outputGetAccount)

	assert.Equal(t, outputSignup.AccountId, outputGetAccount.AccountId)
	assert.Equal(t, inputSignup.Name, outputGetAccount.Name)
	assert.Equal(t, inputSignup.Email, outputGetAccount.Email)
}

func TestSignupUseCase_Execute_WhenAccountAlreadyExistis(t *testing.T) {

	signupUseCase, _, cleanup := setup(t)
	defer cleanup()

	ctx := context.Background()

	inputSignup := SignupInput{
		Name:        "John Doe",
		Email:       "JohnDoe@email.com",
		Cpf:         "123.456.789-09",
		CarPlate:    "ABC-1B34",
		IsPassenger: true,
		IsDriver:    false,
	}

	outputSignup, err := signupUseCase.Execute(ctx, inputSignup)
	assert.NoError(t, err)
	assert.NotNil(t, outputSignup)

	_, err = signupUseCase.Execute(ctx, inputSignup)
	assert.EqualError(t, err, "account already exists")
}
