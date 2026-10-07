package db

import (
	"context"
	"database/sql"
	"embed"
	"path-api/internal/utils"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func Connect(ctx context.Context) (*Queries, *pgxpool.Pool, *sql.DB, error) {
	pool, err := pgxpool.New(ctx, utils.GetDatabaseURL())
	if err != nil {
		return nil, nil, nil, err
	}

	dbConn := stdlib.OpenDBFromPool(pool)

	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		return nil, nil, nil, err
	}
	if err := goose.Up(dbConn, "migrations"); err != nil {
		return nil, nil, nil, err
	}

	queries := New(pool)

	return queries, pool, dbConn, nil
}
