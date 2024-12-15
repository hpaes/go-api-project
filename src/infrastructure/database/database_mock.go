package database

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockDatabaseAdapter struct {
	mock.Mock
}

func NewMockDatabaseConnection() *MockDatabaseAdapter {
	return &MockDatabaseAdapter{}
}

func (m *MockDatabaseAdapter) ExecWithContext(ctx context.Context, query string, args ...interface{}) error {
	arguments := m.Called(ctx, query, args)
	return arguments.Error(0)
}

func (m *MockDatabaseAdapter) QueryWithContext(ctx context.Context, query string, args ...interface{}) (Rows, error) {
	arguments := m.Called(ctx, query, args)
	rows := arguments.Get(0).(Rows)
	return rows, arguments.Error(1)
}

type MockRows struct {
	mock.Mock
}

func (m *MockRows) Columns() ([]string, error) {
	args := m.Called()
	return args.Get(0).([]string), args.Error(1)
}

func (m *MockRows) Next() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MockRows) Scan(dest ...interface{}) error {
	args := m.Called(dest...)
	return args.Error(0)
}

func (m *MockRows) Err() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockRows) Close() error {
	args := m.Called()
	return args.Error(0)
}
