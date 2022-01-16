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
	_ service.UserManager = (*PostgresDB)(nil)
)

// Add create new Users.
func (db *PostgresDB) AddUser(ctx context.Context, user *service.User) error {
	var statment, err = db.conn.PrepareContext(ctx, `INSERT INTO users(
			name,
			email,
			credit_limit,
			due_amount,
			created_at,
			updated_at
		) VALUES(?,?,?,?,?,?)`)

	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}

	var now = time.Now().UTC()

	var result sql.Result
	result, err = statment.Exec(user.Name, user.Email,
		user.CreditLimit, user.DueAmount,
		now, now)
	if err != nil {
		return fmt.Errorf("failed to add user: %w", err)
	}

	if row, _ := result.RowsAffected(); row > 0 {
		var lastInsertedID int64
		lastInsertedID, err = result.LastInsertId()
		if err != nil {
			return fmt.Errorf("failed to get last inserted id: %v", err)
		}

		user.ID = service.ID(lastInsertedID)
	}

	return nil
}

// DeleteUser delete User by userID.
func (db *PostgresDB) DeleteUser(ctx context.Context, userID service.ID) error {
	var statement, err = db.conn.PrepareContext(ctx, `
	DELETE FROM users
	WHERE id=?
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare delete statement: %w\n", err)
	}

	_, err = statement.Exec(userID)
	if err != nil {
		return fmt.Errorf("failed to delete: %w\n", err)
	}

	return nil
}

// GetAllUsers retrieves all Users.
func (db *PostgresDB) GetAllUsers(ctx context.Context) ([]service.User, error) {
	var rows, err = db.conn.QueryContext(ctx, `
	SELECT id, name, email,
	credit_limit, due_amount,
	 created_at, updated_at
	 FROM `+userTable+`
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieves users: %w\n", err)
	}
	defer rows.Close()

	var result []service.User
	for rows.Next() {
		var user service.User
		rows.Scan(&user.ID, &user.Name, &user.Email,
			&user.CreditLimit, &user.DueAmount,
			&user.CreatedAt, &user.UpdatedAt)
		result = append(result, user)
	}

	return result, nil
}

// GetUserByID retrieve user by userID.
func (db *PostgresDB) GetUserByID(ctx context.Context, id service.ID) (service.User, error) {
	var rows, err = db.conn.QueryContext(ctx, `
	SELECT id, name, email,
	credit_limit, due_amount,
	created_at, updated_at
	FROM `+userTable+` 
	WHERE id= ?
	`, id)
	if err != nil {
		return service.User{}, fmt.Errorf("failed to retreive user: %w\n", err)
	}
	defer rows.Close()

	var user service.User
	if rows.Next() {
		rows.Scan(&user.ID, &user.Name, &user.Email,
			&user.CreditLimit, &user.DueAmount,
			&user.CreatedAt, &user.UpdatedAt)
	}

	return user, nil
}

func NewLocalDBUserManager() (service.UserManager, error) {
	return newPostgresDB()
}
