package mocks

import (
	"bnpl/service"
	"sync"
)

// localDB stores user and merchant entities.
type localDB struct {
	lastUserID          int64
	lastMerchantID      int64
	lastTransactionID   int64
	userIDToNameMap     map[int64]string
	merchantIDToNameMap map[int64]string

	users     map[string]service.User
	usersTxn  map[string]service.Transaction
	merchants map[string]service.Merchant
	m         *sync.RWMutex
}

// newlocalDB creates new localDB object.
func newlocalDB() (*localDB, error) {

	return &localDB{
		users:     make(map[string]service.User),
		merchants: make(map[string]service.Merchant),
		m:         &sync.RWMutex{},
	}, nil
}
