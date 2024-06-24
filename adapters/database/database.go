package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/georgysavva/scany/v2/sqlscan"
)

type Database[T any] struct {
	pool *sql.DB
}

func New[T any](pool *sql.DB) *Database[T] {
	return &Database[T]{
		pool: pool,
	}
}

func (db *Database[T]) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.pool.Exec(query, args...)
}

func (db *Database[T]) QueryAll(query string, args ...interface{}) ([]*T, error) {
	output := new([]*T)
	err := sqlscan.Select(context.Background(), db.pool, output, query, args...)
	if errors.Is(err, sql.ErrNoRows) {
		err = nil
	}

	return *output, err
}

func (db *Database[T]) QueryOne(query string, args ...interface{}) (*T, error) {
	output := new(T)
	err := sqlscan.Get(context.Background(), db.pool, output, query, args...)
	if errors.Is(err, sql.ErrNoRows) {
		err = nil
	}

	return output, err
}

func (db *Database[T]) QueryCount(query string, args ...interface{}) (int, error) {
	row := db.pool.QueryRow(query, args...)

	var count int
	err := row.Scan(&count)
	if errors.Is(err, sql.ErrNoRows) {
		err = nil
	}

	return count, err
}
