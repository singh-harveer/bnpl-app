package mocks

import (
	"bnpl/service"
	"context"
	"time"
)

var (
	_ service.UserManager = (*localDB)(nil)
)

// Add create new Users.
func (db *localDB) AddUser(ctx context.Context, user *service.User) error {

	db.m.Lock()
	defer db.m.Unlock()

	if _, ok := db.users[user.Name]; ok {
		return errDuplicateMerchant
	}

	var now = time.Now().UTC()
	user.ID = db.lastMerchantID + 1
	user.CreatedAt = &now
	user.UpdatedAt = &now
	db.users[user.Name] = *user

	db.lastUserID++

	return nil
}

// DeleteUserByName delete User by name.
func (db *localDB) DeleteUserByName(ctx context.Context, name string) error {
	db.m.Lock()
	defer db.m.Unlock()

	if _, ok := db.users[name]; !ok {
		return errNotFound
	}

	delete(db.users, name)

	return nil
}

// GetAllUsers retrieves all Users.
func (db *localDB) GetAllUsers(ctx context.Context) ([]service.User, error) {
	db.m.RLock()
	defer db.m.RUnlock()

	var result []service.User
	for _, v := range db.users {
		result = append(result, v)
	}

	return result, nil
}

// GetUserByName retrieve user by name.
func (db *localDB) GetUserByName(ctx context.Context, name string) (service.User, error) {
	db.m.Lock()
	defer db.m.Unlock()

	var user, ok = db.users[name]
	if !ok {
		return service.User{}, errNotFound
	}

	return user, nil
}

// CreditLimit retrieves user credit limit by user'name.
func (db *localDB) CreditLimit(ctx context.Context, name string) (service.User, error) {

	return service.User{}, nil
}

// NewLocalDBUserManager return UserManager's object for mock.
func NewLocalDBUserManager() (service.UserManager, error) {
	return newlocalDB()
}
