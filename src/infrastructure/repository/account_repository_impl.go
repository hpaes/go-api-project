package repository

import (
	"context"

	"github.com/hpaes/go-api-project/src/core/domain"
	"github.com/hpaes/go-api-project/src/core/domain/valueObjects"
	"github.com/hpaes/go-api-project/src/infrastructure/database"
)

type (
	accountRepository struct {
		connection database.DatabaseConnection
	}
	AccountRepository interface {
		GetById(ctx context.Context, id string) (*domain.Account, error)
		GetByEmail(ctx context.Context, email string) (*domain.Account, error)
		Save(ctx context.Context, account *domain.Account) error
	}
)

func NewAccountRepository(connection database.DatabaseConnection) AccountRepository {
	return &accountRepository{
		connection: connection,
	}
}

func (ar *accountRepository) Save(ctx context.Context, account *domain.Account) error {
	stmt := `INSERT INTO brq_golang.account (account_id, name, cpf, email, car_plate, is_passenger, is_driver) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	err := ar.connection.ExecWithContext(ctx, stmt, account.AccountId, account.Name.Value, account.Cpf.Value, account.Email.Value, account.CarPlate.Value, account.IsPassenger, account.IsDriver)
	if err != nil {
		return err
	}
	return nil
}

func (ar *accountRepository) GetById(ctx context.Context, id string) (*domain.Account, error) {
	stmt := `SELECT account_id, name, cpf, email, car_plate, is_passenger, is_driver FROM brq_golang.account WHERE account_id = $1`
	rows, err := ar.connection.QueryWithContext(ctx, stmt, id)
	if err != nil {
		return &domain.Account{}, err
	}
	if rows == nil {
		return &domain.Account{}, nil
	}

	row := rows[0]

	account := &domain.Account{
		AccountId:   string(row.Columns["account_id"].([]uint8)),
		Name:        valueObjects.Name{Value: row.Columns["name"].(string)},
		Cpf:         valueObjects.Cpf{Value: row.Columns["cpf"].(string)},
		Email:       valueObjects.Email{Value: row.Columns["email"].(string)},
		CarPlate:    valueObjects.CarPlate{Value: row.Columns["car_plate"].(string)},
		IsPassenger: row.Columns["is_passenger"].(bool),
		IsDriver:    row.Columns["is_driver"].(bool),
	}

	acc, err := domain.NewAccount(account.AccountId, account.Name.Value, account.Cpf.Value, account.Email.Value, account.CarPlate.Value, account.IsPassenger, account.IsDriver)
	if err != nil {
		return &domain.Account{}, err
	}

	return acc, nil
}

func (ar *accountRepository) GetByEmail(ctx context.Context, email string) (*domain.Account, error) {
	stmt := `SELECT account_id, name, cpf, email, car_plate, is_passenger, is_driver FROM brq_golang.account WHERE email = $1`
	rows, err := ar.connection.QueryWithContext(ctx, stmt, email)
	if err != nil {
		return &domain.Account{}, err
	}
	if rows == nil {
		return &domain.Account{}, nil
	}

	row := rows[0]

	account := &domain.Account{
		AccountId:   string(row.Columns["account_id"].([]uint8)),
		Name:        valueObjects.Name{Value: row.Columns["name"].(string)},
		Cpf:         valueObjects.Cpf{Value: row.Columns["cpf"].(string)},
		Email:       valueObjects.Email{Value: row.Columns["email"].(string)},
		CarPlate:    valueObjects.CarPlate{Value: row.Columns["car_plate"].(string)},
		IsPassenger: row.Columns["is_passenger"].(bool),
		IsDriver:    row.Columns["is_driver"].(bool),
	}

	acc, err := domain.NewAccount(account.AccountId, account.Name.Value, account.Cpf.Value, account.Email.Value, account.CarPlate.Value, account.IsPassenger, account.IsDriver)
	if err != nil {
		return &domain.Account{}, err
	}

	return acc, nil
}
