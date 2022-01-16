package service

import "context"

// ID represent objectID.
type ID int64

// MerchantManager manages merchants.
type MerchantManager interface {
	// Add create new merchants.
	AddMerchant(ctx context.Context, merchant *Merchant) error

	// DeleteMerchant delete merchant by merchantID.
	DeleteMerchant(ctx context.Context, merchantID ID) error

	// GetAllMerchants retrieves all merchants.
	GetAllMerchants(ctx context.Context) ([]Merchant, error)

	// GetMerchantByID
	GetMerchantByID(ctx context.Context, id ID) (Merchant, error)
}

// UserManager manages users.
type UserManager interface {
	// Add create new Users.
	AddUser(ctx context.Context, user *User) error

	// DeleteUser delete User by userID.
	DeleteUser(ctx context.Context, userID ID) error

	// GetAllUsers retrieves all Users.
	GetAllUsers(ctx context.Context) ([]User, error)

	// GetUserByID
	GetUserByID(ctx context.Context, id ID) (User, error)
}

// TranscationManager manages transactions.
type TranscationManager interface {
}
