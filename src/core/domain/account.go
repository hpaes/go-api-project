package domain

import (
	"github.com/google/uuid"
	"github.com/hpaes/go-api-project/src/core/domain/valueObjects"
)

type Account struct {
	AccountId   string                `json:"accountId"`
	Name        valueObjects.Name     `json:"name"`
	Cpf         valueObjects.Cpf      `json:"cpf"`
	Email       valueObjects.Email    `json:"email"`
	CarPlate    valueObjects.CarPlate `json:"carPlate"`
	IsPassenger bool                  `json:"isPassenger"`
	IsDriver    bool                  `json:"isDriver"`
}

func newAccount(accountId string, name valueObjects.Name, cpf valueObjects.Cpf, email valueObjects.Email, carPlate valueObjects.CarPlate, isPassenger bool, isDriver bool) *Account {
	account := &Account{
		AccountId:   accountId,
		Name:        name,
		Cpf:         cpf,
		Email:       email,
		CarPlate:    carPlate,
		IsPassenger: isPassenger,
		IsDriver:    isDriver,
	}

	return account
}

func CreateAccount(name string, cpf string, email string, carPlate string, isPassenger bool, isDriver bool) (*Account, error) {
	accountId := uuid.New().String()

	nameObj, err := valueObjects.NewName(name)
	if err != nil {
		return nil, err
	}

	cpfObj, err := valueObjects.NewCpf(cpf)
	if err != nil {
		return nil, err
	}

	emailObj, err := valueObjects.NewEmail(email)
	if err != nil {
		return nil, err
	}

	carPlateObj, err := valueObjects.NewCarPlate(carPlate)
	if err != nil {
		return nil, err
	}

	return newAccount(
		accountId,
		*nameObj,
		*cpfObj,
		*emailObj,
		*carPlateObj,
		isPassenger,
		isDriver,
	), nil
}

func RestoreAccount(accountId string, name string, cpf string, email string, carPlate string, isPassenger bool, isDriver bool) (*Account, error) {
	nameObj, err := valueObjects.NewName(name)
	if err != nil {
		return nil, err
	}

	cpfObj, err := valueObjects.NewCpf(cpf)
	if err != nil {
		return nil, err
	}

	emailObj, err := valueObjects.NewEmail(email)
	if err != nil {
		return nil, err
	}

	carPlateObj, err := valueObjects.NewCarPlate(carPlate)
	if err != nil {
		return nil, err
	}

	return newAccount(
		accountId,
		*nameObj,
		*cpfObj,
		*emailObj,
		*carPlateObj,
		isPassenger,
		isDriver,
	), nil
}
