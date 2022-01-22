package service

import "context"

// MerchantManager manages merchants.
type MerchantManager interface {
	// Add create new merchants.
	AddMerchant(ctx context.Context, merchant *Merchant) error

	// DeleteMerchantByName delete merchant by name.
	DeleteMerchantByName(ctx context.Context, name string) error

	// GetAllMerchants retrieves all merchants.
	GetAllMerchants(ctx context.Context) ([]Merchant, error)

	// GetMerchantByName
	GetMerchantByName(ctx context.Context, name string) (Merchant, error)

	// Discount retrieves marchant's discount by name.
	Discount(ctx context.Context, merchantName string) (Merchant, error)
}

// UserManager manages users.
type UserManager interface {
	// Add create new Users.
	AddUser(ctx context.Context, user *User) error

	// DeleteUserByName delete User by name.
	DeleteUserByName(ctx context.Context, name string) error

	// GetAllUsers retrieves all Users.
	GetAllUsers(ctx context.Context) ([]User, error)

	// GetUserByName
	GetUserByName(ctx context.Context, name string) (User, error)

	// DuePayment deposit amount towards user due.
	DuePayment(ctx context.Context, name string, amount float64) (User, error)

	// UserCreditLimit retrieves user credit limit by user'name.
	CreditLimit(ctx context.Context, name string) (User, error)
}

// TransactionManager manages transactions.
type TransactionManager interface {
	// Create creates new transactions.
	Create(ctx context.Context, txn Transaction) error
}

// Reporter manage all consolidating reports.
type Reporter interface {
	// TotalDue retrieves all due amount for all users.
	TotalDue(ctx context.Context) ([]User, error)
	// AllUserAtCreditLimit retrieves all users which reached credit limits.
	AllUserAtCreditLimit(ctx context.Context) ([]User, error)
}
