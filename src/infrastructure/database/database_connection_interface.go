package database

import (
	"context"
	"database/sql"
)

type DatabaseConnection interface {
	QueryWithContext(ctx context.Context, stmt string, args ...any) (*sql.Rows, error)
	ExecWithContext(ctx context.Context, stmt string, args ...any) error
}
