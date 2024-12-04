package database

import (
	"context"
)

type DatabaseConnection interface {
	QueryWithContext(ctx context.Context, stmt string, args ...any) ([]Row, error)
	ExecWithContext(ctx context.Context, stmt string, args ...any) error
}

type Row struct {
	Columns map[string]interface{}
}
