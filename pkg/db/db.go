package db

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"io.github.virsavik/alchemist/pkg/app"
)

// Connect connects to the database
func Connect(cfg app.AppConfig) (*sql.DB, error) {
	sqlDB, err := sql.Open("postgres", cfg.PGURL)
	if err != nil {
		return nil, err
	}

	// sqlDB.SetMaxOpenConns(cfg.PgMaxOpenConns)
	// sqlDB.SetMaxIdleConns(cfg.PgMaxIdleConns)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Ping the db to check connection with context timeout
	if err = sqlDB.PingContext(ctx); err != nil {
		return nil, err
	}

	return sqlDB, err
}
