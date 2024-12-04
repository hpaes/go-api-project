package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidAccount(t *testing.T) {
	account, err := NewAccount("", "John Doe", "123.456.789-09", "johnDoe@gmail.com", "ABC-1B34", true, false)
	assert.NoError(t, err)
	assert.NotNil(t, account)

	assert.Equal(t, "John Doe", account.Name.Value)
	assert.Equal(t, "123.456.789-09", account.Cpf.Value)
	assert.Equal(t, "johnDoe@gmail.com", account.Email.Value)
	assert.Equal(t, "ABC-1B34", account.CarPlate.Value)
	assert.True(t, account.IsPassenger)
	assert.False(t, account.IsDriver)
}

func TestRestoreAccount(t *testing.T) {
	account, err := NewAccount("", "John Doe", "123.456.789-09", "johnDoe@gmail.com", "ABC-1B34", true, false)
	assert.NoError(t, err)
	assert.NotNil(t, account)

	restoredAccount, err := NewAccount(account.AccountId, account.Name.Value, account.Cpf.Value, account.Email.Value, account.CarPlate.Value, account.IsPassenger, account.IsDriver)
	assert.NoError(t, err)
	assert.NotNil(t, restoredAccount)
}
func TestInvalidName(t *testing.T) {
	acc, err := NewAccount("", "John Doe1", "123.456.789-09", "johnDoe1@gmail.com", "ABC-1B34", true, false)
	assert.Errorf(t, err, "Invalid name")
	assert.Nil(t, acc)
}

func TestInvalidCpf(t *testing.T) {
	acc, err := NewAccount("", "John Doe", "123.456.789-0", "johnDoe@gmail.com", "ABC-1B34", true, false)
	assert.Errorf(t, err, "Invalid cpf")
	assert.Nil(t, acc)
}

func TestInvalidEmails(t *testing.T) {
	tests := []struct {
		email string
	}{
		{"johnDoe@gmail"},
		{"johnDoe"},
		{"johnDoe@"},
	}

	for _, tt := range tests {
		acc, err := NewAccount("", "John Doe", "123.456.789-09", tt.email, "ABC-1234", true, false)
		assert.Errorf(t, err, "Invalid email")
		assert.Nil(t, acc)
	}
}

func TestInvalidCarPlates(t *testing.T) {
	tests := []struct {
		carPlate string
	}{
		{"ABC-123"},
		{"ABC-12345"},
		{"ABC-1234"},
		{"ABC1234"},
	}

	for _, tt := range tests {
		acc, err := NewAccount("", "John Doe", "123.456.789-09", "email@email.com", tt.carPlate, true, false)
		assert.Errorf(t, err, "Invalid car plate")
		assert.Nil(t, acc)
	}
}

func TestInvalidNameRestoreAccount(t *testing.T) {
	acc, err := NewAccount("123", "John Doe1", "123.456.789-09", "email@email.com", "ABC-1D45", true, false)
	assert.Errorf(t, err, "Invalid name")
	assert.Nil(t, acc)
}

func TestInvalidCpfRestoreAccount(t *testing.T) {
	acc, err := NewAccount("123", "John Doe", "123.456.789-31", "email@email.com", "ABC-1D45", true, false)
	assert.Errorf(t, err, "Invalid cpf")
	assert.Nil(t, acc)
}

func TestInvalidEmailsRestoreAccount(t *testing.T) {
	tests := []struct {
		email string
	}{
		{"johnDoe@gmail"},
		{"johnDoe"},
		{"johnDoe@"},
	}

	for _, tt := range tests {
		acc, err := NewAccount("123", "John Doe", "123.456.789-09", tt.email, "ABC-1234", true, false)
		assert.Errorf(t, err, "Invalid email")
		assert.Nil(t, acc)
	}
}

func TestInvalidCarPlatesRestoreAccount(t *testing.T) {
	tests := []struct {
		carPlate string
	}{
		{"ABC-123"},
		{"ABC-12345"},
		{"ABC-1234"},
		{"ABC1234"},
	}

	for _, tt := range tests {
		acc, err := NewAccount("123", "John Doe", "123.456.789-09", "email@email.com", tt.carPlate, true, false)
		assert.Errorf(t, err, "Invalid car plate")
		assert.Nil(t, acc)
	}
}
