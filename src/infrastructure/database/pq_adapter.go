// Package database provides a PostgreSQL adapter for database operations.
package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type PqAdapter struct {
	connection *sql.DB
}

func NewPqAdapter() (*PqAdapter, error) {
	connStr := os.Getenv("GO_DATABASE_URL")
	if connStr == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable is not set")
	}
	pqConn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = pqConn.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return &PqAdapter{
		connection: pqConn,
	}, nil
}

func (p *PqAdapter) QueryWithContext(ctx context.Context, stmt string, args ...any) (*sql.Rows, error) {
	preparedStmt, err := p.connection.PrepareContext(ctx, stmt)
	if err != nil {
		return nil, err
	}
	rows, err := preparedStmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (p *PqAdapter) ExecWithContext(ctx context.Context, stmt string, args ...any) error {
	preparedStmt, err := p.connection.PrepareContext(ctx, stmt)
	if err != nil {
		return err
	}
	defer preparedStmt.Close()

	_, err = preparedStmt.ExecContext(ctx, args...)
	if err != nil {
		return err
	}
	return nil
}
