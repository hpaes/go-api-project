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

func NewAccount(accountId string, name string, cpf string, email string, carPlate string, isPassenger bool, isDriver bool) (*Account, error) {

	if accountId == "" {
		accountId = uuid.New().String()
	}

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

	account := &Account{
		AccountId:   accountId,
		Name:        *nameObj,
		Cpf:         *cpfObj,
		Email:       *emailObj,
		CarPlate:    *carPlateObj,
		IsPassenger: isPassenger,
		IsDriver:    isDriver,
	}

	return account, nil
}
