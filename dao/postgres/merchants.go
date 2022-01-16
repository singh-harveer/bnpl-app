package postgres

import (
	"bnpl/service"
	"context"
	"fmt"
	"time"
)

const (
	merchantTable = "merchants"
)

var (
	_ service.MerchantManager = (*PostgresDB)(nil)
)

func (db *PostgresDB) AddMerchant(ctx context.Context, merchant *service.Merchant) error {
	var query = `INSERT INTO ` + merchantTable + `(
		name,
		email,
		discount,
		total_payment,
		created_at,
		updated_at
	)
	VALUES($1, $2, $3, $4, $5, $6)
	RETURNING id`

	var now = time.Now().UTC()

	var result, err = db.conn.ExecContext(ctx, query, merchant.Name,
		merchant.Email,
		merchant.Discount,
		merchant.TotalPayment,
		now, now)
	if err != nil {
		return fmt.Errorf("failed to add merchant: %w", err)
	}

	if row, _ := result.RowsAffected(); row > 0 {
		var lastInsertedID int64
		lastInsertedID, err = result.LastInsertId()
		if err != nil {
			return fmt.Errorf("failed to get last inserted id: %v", err)
		}

		merchant.ID = service.ID(lastInsertedID)
	}

	return nil
}

func (db *PostgresDB) DeleteMerchant(ctx context.Context, merchantID service.ID) error {
	var query = `DELETE FROM ` + merchantTable + ` WHERE id=$1`

	var _, err = db.conn.ExecContext(ctx, query, merchantID)
	if err != nil {
		return fmt.Errorf("failed to delete merchant: %w\n", err)
	}

	return nil
}

func (db *PostgresDB) GetAllMerchants(ctx context.Context) ([]service.Merchant, error) {
	var query = `SELECT id, name, email,
	discount, total_payment,
	created_at, updated_at 
	FROM ` + merchantTable

	var rows, err = db.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieves merchants: %w\n", err)
	}
	defer rows.Close()

	var result []service.Merchant
	for rows.Next() {
		var merchant service.Merchant

		rows.Scan(&merchant.ID, &merchant.Name, &merchant.Email,
			&merchant.Discount, &merchant.TotalPayment,
			&merchant.CreatedAt, &merchant.UpdatedAt)
		result = append(result, merchant)
	}

	return result, nil
}

func (db *PostgresDB) GetMerchantByID(ctx context.Context, id service.ID) (service.Merchant, error) {
	var rows, err = db.conn.QueryContext(ctx, `
	SELECT id, name, email,
	discount, total_payment,
	created_at, updated_at
	FROM `+merchantTable+` 
	WHERE id=$1
	`, id)
	if err != nil {
		return service.Merchant{}, fmt.Errorf("failed to retreive merchant: %w\n", err)
	}
	defer rows.Close()

	var merchant service.Merchant
	if rows.Next() {
		rows.Scan(&merchant.ID, &merchant.Name, &merchant.Email,
			&merchant.Discount, &merchant.TotalPayment,
			&merchant.CreatedAt, &merchant.UpdatedAt)
	}

	return merchant, nil
}

func NewLocalDBMerchantManager() (service.MerchantManager, error) {
	return newPostgresDB()
}
