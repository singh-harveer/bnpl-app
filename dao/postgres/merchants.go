package postgres

import (
	"bnpl/service"
	"context"
	"fmt"
)

const (
	merchantTable = "merchants"
)

var (
	_ service.MerchantManager = (*postgresDB)(nil)
)

// Add create new merchants.
func (db *postgresDB) AddMerchant(ctx context.Context, merchant *service.Merchant) error {
	var query = `INSERT INTO ` + merchantTable + `(
		name, email,
		discount,
		total_payment,
		created_at, updated_at
		) VALUES($1, $2, $3, $4, $5, $6)`

	var _, err = db.conn.ExecContext(ctx, query,
		merchant.Name, merchant.Email,
		merchant.Discount, merchant.TotalPayment,
		merchant.CreatedAt, merchant.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to insert merchant details: %w", err)
	}

	return nil
}

// DeleteMerchantByName delete merchant by name.
func (db *postgresDB) DeleteMerchantByName(ctx context.Context, name string) error {
	// TODO- implementations
	return nil
}

// GetAllMerchants retrieves all merchants.
func (db *postgresDB) GetAllMerchants(ctx context.Context) ([]service.Merchant, error) {
	// TODO- implementations
	return nil, nil
}

// GetMerchantByName retrieves merchant by namee.
func (db *postgresDB) GetMerchantByName(ctx context.Context, name string) (service.Merchant, error) {
	var query = `SELECT id, name, email,
			discount, total_payment,
			created_at, updated_at
			FROM ` + merchantTable + `
			WHERE name= $1`

	var rows, err = db.conn.QueryContext(ctx, query,
		name,
	)
	if err != nil {
		return service.Merchant{}, fmt.Errorf("failed to retrieve merchant: %w", err)
	}

	var merchant service.Merchant
	if rows.Next() {
		var err = rows.Scan(&merchant.ID, &merchant.Name,
			&merchant.Email, &merchant.Discount,
			&merchant.TotalPayment, &merchant.CreatedAt,
			&merchant.UpdatedAt)
		if err != nil {
			return service.Merchant{}, fmt.Errorf("failed to Unmarshal merchant details: %w", err)
		}
	}

	return merchant, nil
}

// Discount retrieves marchant's discount by name.
func (db *postgresDB) Discount(ctx context.Context, merchantName string) (service.Merchant, error) {
	var query = `SELECT name,
			discount
			FROM ` + merchantTable + `
			WHERE name= $1
			`

	var rows, err = db.conn.QueryContext(ctx, query, merchantName)
	if err != nil {
		return service.Merchant{}, fmt.Errorf("failed to retrieve merchant discount: %w", err)
	}

	var merchant service.Merchant
	if rows.Next() {
		var err = rows.Scan(&merchant.Name, &merchant.Discount)
		if err != nil {
			return service.Merchant{}, fmt.Errorf("failed to Unmarshal merchant discount details: %w", err)
		}
	}

	return merchant, nil
}

// NewMerchantManager creates MerchantManager's postgres implementation object.
func NewMerchantManager(ctx context.Context, uri string) (service.MerchantManager, error) {
	return newPostgresDB(ctx, uri)
}
