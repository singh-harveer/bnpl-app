package postgres

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // to import postgres driver
)

const (
	driverName = "postgres"
)

type postgresDB struct {
	conn *sql.DB
}

func newPostgresDB(ctx context.Context, uri string) (*postgresDB, error) {
	var connection, err = sql.Open(driverName, uri)
	if err != nil {
		return nil, fmt.Errorf("failed to open database :%v", err)
	}

	return &postgresDB{
		conn: connection,
	}, nil
}
