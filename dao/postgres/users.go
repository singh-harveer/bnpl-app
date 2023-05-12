package postgres

import (
	"bnpl/service"
	"context"
	"database/sql"
	"fmt"
	"time"
)

const (
	userTable = "users"
)

var (
	_ service.UserManager = (*postgresDB)(nil)
)

// Add create new Users.
func (db *postgresDB) AddUser(ctx context.Context, user *service.User) error {
	var query = `INSERT INTO ` + userTable + `(
		name, email,
		credit_limit,
		due_amount,
		created_at,
		updated_at
	) VALUES(
		$1, $2, $3, $4,
		$5, $6
	)
	`
	var now = time.Now().UTC()
	var _, err = db.conn.ExecContext(ctx, query,
		user.Name,
		user.Email,
		user.CreditLimit,
		user.DueAmount,
		now,
		now,
	)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	return nil
}

// DeleteUserByName delete User by name.
func (db *postgresDB) DeleteUserByName(ctx context.Context, name string) error {
	var query = `DELETE FROM ` + userTable + `
			WHERE name= $1
			`

	var result, err = db.conn.Exec(query, name)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	var affectedRow int64
	affectedRow, err = result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get affected row: %w", err)
	}

	if affectedRow == 0 {
		return fmt.Errorf("not record found")
	}

	return nil
}

// GetAllUsers retrieves all Users.
func (db *postgresDB) GetAllUsers(ctx context.Context) ([]service.User, error) {
	var query = `SELECT id, name, email,
	credit_limit, due_amount,
	created_at, updated_at
	FROM ` + userTable + `
	`

	var rows, err = db.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve user: %w", err)
	}

	var users []service.User
	if rows.Next() {
		var user service.User
		var err = rows.Scan(&user.ID, &user.Name,
			&user.Email, &user.CreditLimit,
			&user.DueAmount, &user.CreatedAt,
			&user.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to Unmarshal user details: %w", err)
		}

		users = append(users, user)
	}

	return users, nil
}

// GetUserByName retrieves user details by username.
func (db *postgresDB) GetUserByName(ctx context.Context, name string) (service.User, error) {
	var query = `SELECT id, name, email,
			credit_limit, due_amount,
			created_at, updated_at
			FROM ` + userTable + `
			WHERE name= $1
			`

	var rows, err = db.conn.QueryContext(ctx, query,
		name,
	)
	if err != nil {
		return service.User{}, fmt.Errorf("failed to retrieve user: %w", err)
	}

	var user service.User
	if rows.Next() {
		var err = rows.Scan(&user.ID, &user.Name,
			&user.Email, &user.CreditLimit,
			&user.DueAmount, &user.CreatedAt,
			&user.UpdatedAt)
		if err != nil {
			return service.User{}, fmt.Errorf("failed to Unmarshal user details: %w", err)
		}
	}

	return user, nil
}

// DuePayment deposit amount towards user due.
func (db *postgresDB) DuePayment(ctx context.Context, name string, amount float64) (service.User, error) {
	var tx, err = db.conn.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return service.User{}, fmt.Errorf("failed to begin transaction: %w", err)
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

	var query = `SELECT credit_limit, due_amount FROM ` + userTable + ` WHERE name = $1`

	var rows *sql.Rows
	rows, err = tx.QueryContext(ctx, query, name)
	if err != nil {
		return service.User{}, fmt.Errorf("failed to retrieve user details: %w", err)
	}

	var creditLimit float64
	var dueAmount float64
	if rows.Next() {
		err = rows.Scan(&creditLimit, &dueAmount)
		if err != nil {
			return service.User{}, fmt.Errorf("failed to Unmarshal credit and due amount details: %w", err)
		}
	}
	rows.Close()

	if amount > dueAmount {
		return service.User{}, fmt.Errorf("not allowed to pay more than due amount")
	}
	var newCreditLimit = creditLimit + amount
	var newDueAmount = dueAmount - amount

	var updateQuery = `UPDATE ` + userTable + `
					SET credit_limit= $1,
					due_amount = $2,
					updated_at = $3
					WHERE name = $4
					RETURNING name, credit_limit, due_amount`

	var user service.User
	err = tx.QueryRowContext(ctx, updateQuery,
		newCreditLimit, newDueAmount,
		time.Now().UTC(), name).Scan(&user.Name, &user.CreditLimit, &user.DueAmount)
	if err != nil {
		return service.User{}, fmt.Errorf("failed to update credit limt: %w", err)
	}

	commit = true

	return user, nil
}

// CreditLimit retrieves user credit limit by user'name.
func (db *postgresDB) CreditLimit(ctx context.Context, name string) (service.User, error) {
	var query = `SELECT name, credit_limit
			FROM ` + userTable + `
			WHERE name= $1
			`

	var rows, err = db.conn.QueryContext(ctx, query, name)
	if err != nil {
		return service.User{}, fmt.Errorf("failed to retrieve user credit limit: %w", err)
	}

	var user service.User
	if rows.Next() {
		var err = rows.Scan(&user.Name, &user.CreditLimit)
		if err != nil {
			return service.User{}, fmt.Errorf("failed to Unmarshal user credit details: %w", err)
		}
	}

	return user, nil
}

// NewUserManager creates UserManager's postgres implementation object.
func NewUserManager(ctx context.Context, uri string) (service.UserManager, error) {
	return newPostgresDB(ctx, uri)
}
