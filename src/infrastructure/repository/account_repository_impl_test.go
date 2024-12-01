package repository

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/hpaes/go-api-project/src/core/domain"
	"github.com/hpaes/go-api-project/src/infrastructure/database"
	"github.com/stretchr/testify/assert"
)

func randomEmail() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("user%d@example.com", r.Int())
}

func setup(t *testing.T) (*database.PqAdapter, func()) {
	databaseConnection, err := database.NewPqAdapter()
	assert.NoError(t, err)

	cleanup := func() {
		err := databaseConnection.ExecWithContext(context.Background(), "DELETE FROM brq_golang.account")
		assert.NoError(t, err)
	}

	return databaseConnection, cleanup
}

func TestSaveAccountRepository(t *testing.T) {
	databaseConnection, cleanup := setup(t)
	defer cleanup()

	accountRepo := NewAccountRepository(databaseConnection)
	email := randomEmail()
	acc, err := domain.CreateAccount("John Doe", "123.456.789-09", email, "ABC-1B34", true, false)
	assert.NoError(t, err)

	ctx := context.Background()
	t.Run("Save", func(t *testing.T) {
		err := accountRepo.Save(ctx, acc)
		assert.NoError(t, err)
	})
}

func TestGetByEmailAccountRepository(t *testing.T) {
	databaseConnection, cleanup := setup(t)
	defer cleanup()

	accountRepo := NewAccountRepository(databaseConnection)
	email := randomEmail()
	acc, err := domain.CreateAccount("John Doe", "123.456.789-09", email, "ABC-1B34", true, false)
	assert.NoError(t, err)
	err = accountRepo.Save(context.TODO(), acc)
	assert.NoError(t, err)

	retrievedAcc, err := accountRepo.GetByEmail(context.TODO(), acc.Email.Value)
	assert.NoError(t, err)
	assert.NotNil(t, retrievedAcc)
	assert.Equal(t, acc, retrievedAcc)
}

func TestGetByIdAccountRepository(t *testing.T) {
	databaseConnection, cleanup := setup(t)
	defer cleanup()

	accountRepo := NewAccountRepository(databaseConnection)
	email := randomEmail()
	acc, err := domain.CreateAccount("John Doe", "123.456.789-09", email, "ABC-1B34", true, false)
	assert.NoError(t, err)
	err = accountRepo.Save(context.TODO(), acc)
	assert.NoError(t, err)

	retrievedAcc, err := accountRepo.GetById(context.TODO(), acc.AccountId)
	assert.NoError(t, err)
	assert.NotNil(t, retrievedAcc)
	assert.Equal(t, acc, retrievedAcc)
}
