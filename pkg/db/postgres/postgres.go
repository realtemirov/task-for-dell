package postgres

import (
	"fmt"
	"time"

	_ "github.com/jackc/pgx/stdlib" // pgx driver
	"github.com/jmoiron/sqlx"
	"github.com/realtemirov/task-for-dell/config"
)

const (
	maxOpenConnections    = 60
	connectionMaxLifetime = 120
	maxIdleConnections    = 30
	connectionMaxIdleTime = 60
)

// NewPostgresDB returns a new postgres DB instance.
func NewPostgresDB(cfg *config.Config) (*sqlx.DB, error) {

	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Postgres.Username,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.DBName,
		cfg.Postgres.SSLMode,
	)

	db, err := sqlx.Connect(cfg.Postgres.PgDriver, connectionString)
	if err != nil {
		fmt.Println("error connecting to db", connectionString)
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConnections)
	db.SetConnMaxLifetime(connectionMaxLifetime * time.Second)
	db.SetMaxIdleConns(maxIdleConnections)
	db.SetConnMaxIdleTime(connectionMaxIdleTime * time.Second)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
