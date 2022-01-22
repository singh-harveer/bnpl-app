package postgres

import (
	"bnpl/service"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

const (
	transactionTable = "transactions"
)

var (
	_ service.TransactionManager = (*postgresDB)(nil)
)

// Create creates new transaction.
func (db *postgresDB) Create(ctx context.Context, txn service.Transaction) error {
	var tx, err = db.conn.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Defer commit/rollback of transaction. If commit is true, the
	// transaction will be committed, otherwise it will be rolled back.
	var commit bool
	defer func() {
		if commit {
			var tErr = tx.Commit()
			if tErr != nil {
				err = fmt.Errorf("failed to commit transaction: %w", err)
			} else {
				tx.Rollback()
			}
		}
	}()

	var query = `SELECT credit_limit, due_amount FROM ` + userTable + ` WHERE id = $1`
	var rows *sql.Rows
	rows, err = tx.QueryContext(ctx, query, txn.UserID)
	if err != nil {
		return fmt.Errorf("failed to retrieve user details: %w", err)
	}

	var creditLimit float64
	var dueAmount float64
	if rows.Next() {
		err = rows.Scan(&creditLimit, &dueAmount)
		if err != nil {
			return fmt.Errorf("failed to Unmarshal: %w", err)
		}
	}
	rows.Close()

	if creditLimit < txn.Amount {
		return errors.New("failed due to insufficient balance")
	}
	//  Update credit limit and due amount for user.
	var newCreditLimit = creditLimit - txn.Amount
	var newDueAmount = dueAmount + txn.Amount
	var now = time.Now().UTC()

	query = `UPDATE ` + userTable + `
			SET credit_limit= $1,
				due_amount = $2,
				updated_at = $3
			WHERE id = $4`

	_, err = tx.ExecContext(ctx, query,
		newCreditLimit, newDueAmount,
		now, txn.UserID)
	if err != nil {
		return fmt.Errorf("failed to update credit limt: %w", err)
	}

	var merchantQuery = `SELECT discount, total_payment FROM ` + merchantTable + ` WHERE id = $1`
	var merchantRows *sql.Rows
	merchantRows, err = tx.QueryContext(ctx, merchantQuery, txn.UserID)
	if err != nil {
		return fmt.Errorf("failed to retrieve merchant details: %w", err)
	}

	var discount float64
	var totalPayment float64
	if merchantRows.Next() {
		merchantRows.Scan(&discount, &totalPayment)
	}
	merchantRows.Close()

	var currentAmountToPay = txn.Amount - ((discount / 100) * txn.Amount)
	var newTotalPayment = totalPayment + currentAmountToPay
	query = `UPDATE ` + merchantTable + `
		SET total_payment=$1,
		updated_at =$2
		WHERE id = $3`

	_, err = tx.ExecContext(ctx, query, newTotalPayment, now, txn.UserID)
	if err != nil {
		return fmt.Errorf("failed to pay to merchant: %w", err)
	}

	var txnQuery = `INSERT INTO ` + transactionTable + `(
	user_id, merchant_id,
	amount, status,
	created_at, updated_at) 
	VALUES($1, $2, $3, $4, $5, $6)`

	_, err = tx.ExecContext(ctx, txnQuery,
		txn.UserID, txn.MerchantID,
		txn.Amount, txn.Status,
		txn.CreatedAt, txn.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create transaction: %w", err)
	}
	commit = true

	return nil
}

// NewTransactionManager creates TransactionManager's postgres implementation object.
func NewTransactionManager(ctx context.Context, uri string) (service.TransactionManager, error) {
	return newPostgresDB(ctx, uri)
}
