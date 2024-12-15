package database

import (
	"context"
)

type DatabaseAdapter interface {
	QueryWithContext(ctx context.Context, stmt string, args ...any) (Rows, error)
	ExecWithContext(ctx context.Context, stmt string, args ...any) error
}
