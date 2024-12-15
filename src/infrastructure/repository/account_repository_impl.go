package repository

import (
	"context"

	"github.com/hpaes/go-api-project/src/core/domain"
	"github.com/hpaes/go-api-project/src/infrastructure/database"
)

type (
	accountRepository struct {
		connection database.DatabaseAdapter
	}
	AccountRepository interface {
		GetById(ctx context.Context, id string) (*domain.Account, error)
		GetByEmail(ctx context.Context, email string) (*domain.Account, error)
		Save(ctx context.Context, account *domain.Account) error
	}
)

func NewAccountRepository(connection database.DatabaseAdapter) AccountRepository {
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
	defer rows.Close()

	if !rows.Next() {
		return &domain.Account{}, nil
	}

	var accountId, name, cpf, emailValue, carPlate string
	var isPassenger, isDriver bool
	if err := rows.Scan(&accountId, &name, &cpf, &emailValue, &carPlate, &isPassenger, &isDriver); err != nil {
		return &domain.Account{}, err
	}

	acc, err := domain.NewAccount(accountId, name, cpf, emailValue, carPlate, isPassenger, isDriver)
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
	defer rows.Close()

	if !rows.Next() {
		return &domain.Account{}, nil
	}

	var accountId, name, cpf, emailValue, carPlate string
	var isPassenger, isDriver bool
	if err := rows.Scan(&accountId, &name, &cpf, &emailValue, &carPlate, &isPassenger, &isDriver); err != nil {
		return &domain.Account{}, err
	}

	acc, err := domain.NewAccount(accountId, name, cpf, emailValue, carPlate, isPassenger, isDriver)
	if err != nil {
		return &domain.Account{}, err
	}

	return acc, nil
}
