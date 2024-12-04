package repository

import (
	"context"
	"testing"

	"github.com/hpaes/go-api-project/src/core/domain"
	"github.com/hpaes/go-api-project/src/core/domain/valueObjects"
	"github.com/hpaes/go-api-project/src/infrastructure/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSaveAccountRepository(t *testing.T) {
	connection := database.NewMockDatabaseConnection()
	repo := NewAccountRepository(connection)

	account := &domain.Account{
		AccountId: "1",
		Name: valueObjects.Name{
			Value: "John Doe",
		},
		Email: valueObjects.Email{
			Value: "johnDoe@email.com",
		},
		Cpf: valueObjects.Cpf{
			Value: "123.456.789-09",
		},
		CarPlate: valueObjects.CarPlate{
			Value: "ABC-1B34",
		},
		IsPassenger: true,
		IsDriver:    false,
	}

	ctx := context.TODO()
	connection.On("ExecWithContext", ctx, mock.Anything, mock.Anything).Return(nil)

	err := repo.Save(ctx, account)

	assert.NoError(t, err)
	connection.AssertExpectations(t)
}

func TestAccountRepository_GetById(t *testing.T) {
	connection := database.NewMockDatabaseConnection()
	repo := NewAccountRepository(connection)

	expectedAccount := &domain.Account{
		AccountId: "123",
		Name: valueObjects.Name{
			Value: "John Doe",
		},
		Email: valueObjects.Email{
			Value: "johnDoe@email.com",
		},
		Cpf: valueObjects.Cpf{
			Value: "123.456.789-09",
		},
		CarPlate: valueObjects.CarPlate{
			Value: "ABC-1B34",
		},
		IsPassenger: true,
		IsDriver:    false,
	}

	rows := []database.Row{
		{
			Columns: map[string]interface{}{
				"account_id":   "123",
				"name":         "John Doe",
				"cpf":          "123.456.789-09",
				"email":        "johnDoe@email.com",
				"car_plate":    "ABC-1B34",
				"is_passenger": true,
				"is_driver":    false,
			},
		},
	}

	connection.On("QueryWithContext", mock.Anything, mock.Anything, mock.Anything).Return(func() []database.Row { return rows }, nil)

	ctx := context.TODO()
	account, err := repo.GetById(ctx, "123")

	assert.NoError(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, expectedAccount.AccountId, account.AccountId)
	assert.Equal(t, expectedAccount.Name, account.Name)
	assert.Equal(t, expectedAccount.Email, account.Email)
	assert.Equal(t, expectedAccount.Cpf, account.Cpf)
	assert.Equal(t, expectedAccount.CarPlate, account.CarPlate)
	assert.Equal(t, expectedAccount.IsPassenger, account.IsPassenger)
	assert.Equal(t, expectedAccount.IsDriver, account.IsDriver)
	connection.AssertExpectations(t)
}

func TestAccountRepository_GetByEmail(t *testing.T) {
	connection := database.NewMockDatabaseConnection()
	repo := NewAccountRepository(connection)

	expectedAccount := &domain.Account{
		AccountId: "123",
		Name: valueObjects.Name{
			Value: "John Doe",
		},
		Email: valueObjects.Email{
			Value: "johnDoe@email.com",
		},
		Cpf: valueObjects.Cpf{
			Value: "123.456.789-09",
		},
		CarPlate: valueObjects.CarPlate{
			Value: "ABC-1B34",
		},
		IsPassenger: true,
		IsDriver:    false,
	}

	rows := []database.Row{
		{
			Columns: map[string]interface{}{
				"account_id":   "123",
				"name":         "John Doe",
				"cpf":          "123.456.789-09",
				"email":        "johnDoe@email.com",
				"car_plate":    "ABC-1B34",
				"is_passenger": true,
				"is_driver":    false,
			},
		},
	}

	connection.On("QueryWithContext", mock.Anything, mock.Anything, mock.Anything).Return(func() []database.Row { return rows }, nil)

	ctx := context.TODO()
	account, err := repo.GetByEmail(ctx, "johnDoe@email.com")

	assert.NoError(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, expectedAccount.AccountId, account.AccountId)
	assert.Equal(t, expectedAccount.Name, account.Name)
	assert.Equal(t, expectedAccount.Email, account.Email)
	assert.Equal(t, expectedAccount.Cpf, account.Cpf)
	assert.Equal(t, expectedAccount.CarPlate, account.CarPlate)
	assert.Equal(t, expectedAccount.IsPassenger, account.IsPassenger)
	assert.Equal(t, expectedAccount.IsDriver, account.IsDriver)
	connection.AssertExpectations(t)
}
