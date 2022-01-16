package postgres

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	envDatabaseURI = "BNPL_DB_URL"
)

// PostgresDB stores user and merchant entities.
type PostgresDB struct {
	conn *sql.DB
}

// newPostgresDB creates new PostgresDB object.
func newPostgresDB() (*PostgresDB, error) {
	var uri, ok = os.LookupEnv(envDatabaseURI)
	if !ok {
		return nil, fmt.Errorf("%s is not set in env", envDatabaseURI)
	}

	var db, err = sql.Open("postgres", uri)
	if err != nil {
		return nil, fmt.Errorf("failed to open postgres connection: %v", err)
	}

	return &PostgresDB{
		conn: db,
	}, nil
}
