package localdb

import (
	"bnpl/service"
	"sync"
)

const (
	envDatabaseURI = "BNPL_DB_URL"
)

// localDB stores user and merchant entities.
type localDB struct {
	lastUserID     int64
	lastMerchantID int64
	users          map[string]service.User
	merchants      map[string]service.Merchant
	m              *sync.RWMutex
}

// newlocalDB creates new localDB object.
func newlocalDB() (*localDB, error) {

	return &localDB{
		users:     make(map[string]service.User),
		merchants: make(map[string]service.Merchant),
		m:         &sync.RWMutex{},
	}, nil
}
