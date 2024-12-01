package repository

import (
	"context"

	"github.com/hpaes/go-api-project/src/core/domain"
	"github.com/hpaes/go-api-project/src/infrastructure/database"
)

type AccountRepository struct {
	connection database.DatabaseConnection
}

func NewAccountRepository(connection database.DatabaseConnection) *AccountRepository {
	return &AccountRepository{
		connection: connection,
	}
}

func (ar *AccountRepository) Save(ctx context.Context, account *domain.Account) error {
	stmt := `INSERT INTO brq_golang.account (account_id, name, cpf, email, car_plate, is_passenger, is_driver) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	err := ar.connection.ExecWithContext(ctx, stmt, account.AccountId, account.Name.Value, account.Cpf.Value, account.Email.Value, account.CarPlate.Value, account.IsPassenger, account.IsDriver) //Exec(stmt, account.AccountId, account.Name.Value, account.Cpf.Value, account.Email.Value, account.CarPlate.Value, account.IsPassenger, account.IsDriver)
	if err != nil {
		return err
	}
	return nil
}

func (ar *AccountRepository) GetById(ctx context.Context, id string) (*domain.Account, error) {
	stmt := `SELECT account_id, name, cpf, email, car_plate, is_passenger, is_driver FROM brq_golang.account WHERE account_id = $1`
	rows, err := ar.connection.QueryWithContext(ctx, stmt, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	var account domain.Account
	err = rows.Scan(&account.AccountId, &account.Name.Value, &account.Cpf.Value, &account.Email.Value, &account.CarPlate.Value, &account.IsPassenger, &account.IsDriver)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (ar *AccountRepository) GetByEmail(ctx context.Context, email string) (*domain.Account, error) {
	stmt := `SELECT account_id, name, cpf, email, car_plate, is_passenger, is_driver FROM brq_golang.account WHERE email = $1`
	rows, err := ar.connection.QueryWithContext(ctx, stmt, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, nil
	}

	var account domain.Account
	err = rows.Scan(&account.AccountId, &account.Name.Value, &account.Cpf.Value, &account.Email.Value, &account.CarPlate.Value, &account.IsPassenger, &account.IsDriver)
	if err != nil {
		return nil, err
	}

	acc, err := domain.RestoreAccount(account.AccountId, account.Name.Value, account.Cpf.Value, account.Email.Value, account.CarPlate.Value, account.IsPassenger, account.IsDriver)
	if err != nil {
		return nil, err
	}
	return acc, nil
}
