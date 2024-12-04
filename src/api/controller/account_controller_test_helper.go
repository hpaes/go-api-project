package controller

import (
	"context"

	"github.com/hpaes/go-api-project/src/core/application/usecase"
	"github.com/stretchr/testify/mock"
)

func NewMockLogger() *mockLogger {
	return &mockLogger{}
}

type mockLogger struct {
	mock.Mock
}

func (ml *mockLogger) LogInformation(format string, args ...interface{}) {
	ml.Called(format, args)
}

func (ml *mockLogger) LogError(format string, args ...interface{}) {
	ml.Called(format, args)
}

type mockSignUpUseCase struct {
	mock.Mock
}

func NewMockSignup() *mockSignUpUseCase {
	return &mockSignUpUseCase{}
}

type mockGetAccountUseCase struct {
	mock.Mock
}

func NewMockGetAccount() *mockGetAccountUseCase {
	return &mockGetAccountUseCase{}
}

func (sm *mockSignUpUseCase) Execute(ctx context.Context, input usecase.SignupInput) (*usecase.SignupOutput, error) {
	args := sm.Called(ctx, input)
	return args.Get(0).(*usecase.SignupOutput), args.Error(1)
}

func (gm *mockGetAccountUseCase) Execute(ctx context.Context, accountId string) (*usecase.GetAccountOutput, error) {
	args := gm.Called(ctx, accountId)
	return args.Get(0).(*usecase.GetAccountOutput), args.Error(1)
}
