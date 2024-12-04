package database

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockDatabaseConnection struct {
	mock.Mock
}

func NewMockDatabaseConnection() *MockDatabaseConnection {
	return &MockDatabaseConnection{}
}

func (m *MockDatabaseConnection) ExecWithContext(ctx context.Context, query string, args ...interface{}) error {
	arguments := m.Called(ctx, query, args)
	return arguments.Error(0)
}

func (m *MockDatabaseConnection) QueryWithContext(ctx context.Context, query string, args ...interface{}) ([]Row, error) {
	arguments := m.Called(ctx, query, args)
	rows := arguments.Get(0).(func() []Row)()
	return rows, arguments.Error(1)
}
