package storage

import (
	"context"
	"database/sql"
)

type QueryDatabase interface {
	QueryRow(query string, args ...interface{}) *sql.Row
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

type TransacationDatabase interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

type DisposableDatabase interface {
	Close() error
}

type Database interface {
	QueryDatabase
	TransacationDatabase
	DisposableDatabase
}
