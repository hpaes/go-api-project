package usecase

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/hpaes/go-api-project/src/core/domain"
	"github.com/hpaes/go-api-project/src/infrastructure/database"
	"github.com/hpaes/go-api-project/src/infrastructure/logger"
	"github.com/hpaes/go-api-project/src/infrastructure/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setup(t *testing.T) (SignUp, GetAccount, func()) {
	// setup db connection
	dbConnection, err := database.NewPqAdapter()
	if err != nil {
		panic(err)
	}
	// create acc repository
	accountRepository := repository.NewAccountRepository(dbConnection)

	// create logger
	logHandler := logger.NewConsoleLogger()

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
	t.SkipNow()
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

var input = SignupInput{
	Name:        "John Doe",
	Email:       "johnDoe@email.com",
	Cpf:         "123.456.789-09",
	CarPlate:    "ABC-1B34",
	IsPassenger: true,
	IsDriver:    false,
}

func TestSignupUseCase_Execute_WhenAccountAlreadyExistis(t *testing.T) {
	t.SkipNow()
	signupUseCase, _, cleanup := setup(t)
	defer cleanup()

	ctx := context.Background()

	outputSignup, err := signupUseCase.Execute(ctx, input)
	assert.NoError(t, err)
	assert.NotNil(t, outputSignup)

	_, err = signupUseCase.Execute(ctx, input)
	assert.EqualError(t, err, "account already exists")
}

func TestSignupUsecaseWithMock(t *testing.T) {
	r := repository.NewMockAccountRepository()
	logHandler := logger.NewConsoleLogger()
	uc := NewSignupUseCase(r, logHandler)

	r.On("Save", mock.Anything, mock.Anything).Return(nil)
	r.On("GetByEmail", mock.Anything, mock.Anything).Return(&domain.Account{}, nil)

	ctx := context.TODO()
	output, err := uc.Execute(ctx, input)

	assert.NotNil(t, output)
	assert.NoError(t, err)
}

func TestSignupUsecaseWithMockWhenAccountAlreadyExists(t *testing.T) {
	r := repository.NewMockAccountRepository()
	logHandler := logger.NewConsoleLogger()
	uc := NewSignupUseCase(r, logHandler)

	errorMsg := "account already exists"
	r.On("GetByEmail", mock.Anything, mock.Anything).Return(&domain.Account{AccountId: uuid.NewString()}, nil)
	r.On("Save", mock.Anything, mock.Anything).Return(fmt.Errorf(errorMsg))

	ctx := context.TODO()

	output, err := uc.Execute(ctx, input)
	assert.EqualError(t, err, errorMsg)
	assert.Nil(t, output)
}

func TestSignupUsecaseWithMockWithInvalidPameters(t *testing.T) {

	tt := []struct {
		name  string
		input SignupInput
		err   error
	}{
		{
			name: "empty name",
			input: SignupInput{
				Name:        "",
				Email:       "johnDoe@email.com",
				Cpf:         "123.456.789-09",
				CarPlate:    "ABC-1B34",
				IsPassenger: true,
				IsDriver:    false,
			},
			err: fmt.Errorf("name is required"),
		},
		{
			name: "empty email",
			input: SignupInput{
				Name:        "John Doe",
				Email:       "",
				Cpf:         "123.456.789-09",
				CarPlate:    "ABC-1B34",
				IsPassenger: true,
				IsDriver:    false,
			},
			err: fmt.Errorf("email is required"),
		},
		{
			name: "empty cpf",
			input: SignupInput{
				Name:        "John Doe",
				Email:       "johnDoe@email.com",
				Cpf:         "",
				CarPlate:    "ABC-1B34",
				IsPassenger: true,
				IsDriver:    false,
			},
			err: fmt.Errorf("cpf is required"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			r := repository.NewMockAccountRepository()
			logHandler := logger.NewConsoleLogger()
			uc := NewSignupUseCase(r, logHandler)
			output, err := uc.Execute(context.TODO(), tc.input)
			assert.Nil(t, output)
			assert.EqualError(t, err, tc.err.Error())
		})
	}
}
