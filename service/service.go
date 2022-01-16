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
}

// TranscationManager manages transactions.
type TranscationManager interface {
}
