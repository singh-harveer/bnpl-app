package postgres

import (
	"bnpl/service"
	"context"
	"fmt"
	"log"
)

var (
	_ service.Reporter = (*postgresDB)(nil)
)

// TotalDue retrieves all due amount for all users.
func (db *postgresDB) TotalDue(ctx context.Context) ([]service.User, error) {
	var query = `SELECT name, due_amount
			FROM ` + userTable + `
			WHERE due_amount > 0
			`

	var rows, err = db.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve users due amount details: %w", err)
	}

	var result []service.User
	for rows.Next() {
		log.Println("inside Next")
		var user service.User
		var err = rows.Scan(&user.Name, &user.DueAmount)
		if err != nil {
			return nil, fmt.Errorf("failed to Unmarshal user due amount details: %w", err)
		}
		result = append(result, user)
	}

	return result, nil
}

// AllUserAtCreditLimit retrieves all due amount for all users.
func (db *postgresDB) AllUserAtCreditLimit(ctx context.Context) ([]service.User, error) {
	var query = `SELECT name, due_amount
			FROM ` + userTable + `
			WHERE credit_limit = 0`

	var rows, err = db.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve users due amount details: %w", err)
	}

	var result []service.User
	for rows.Next() {
		var user service.User
		var err = rows.Scan(&user.Name, &user.DueAmount)
		if err != nil {
			return nil, fmt.Errorf("failed to Unmarshal user due amount details: %w", err)
		}
		result = append(result, user)
	}

	return result, nil
}

// NewReportManager creates NewReportManager's postgres implementation object.
func NewReportManager(ctx context.Context, uri string) (service.Reporter, error) {
	return newPostgresDB(ctx, uri)
}
