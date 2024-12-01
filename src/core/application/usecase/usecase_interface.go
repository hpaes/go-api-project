package usecase

import (
	"context"
)

type GetAccount interface {
	Execute(ctx context.Context, accountId string) (*GetAccountOutput, error)
}

type SignUp interface {
	Execute(ctx context.Context, input SignupInput) (*SignupOutput, error)
}
