package localdb

import (
	"bnpl/service"
	"context"
	"time"
)

var (
	_ service.MerchantManager = (*localDB)(nil)
)

// AddMerchant add new merchant to localDB.
func (db *localDB) AddMerchant(ctx context.Context, merchant *service.Merchant) error {
	db.m.Lock()
	defer db.m.Unlock()

	if _, ok := db.merchants[merchant.Name]; ok {
		return errDuplicateMerchant
	}

	var now = time.Now().UTC()
	merchant.ID = db.lastMerchantID + 1
	merchant.CreatedAt = now
	merchant.UpdatedAt = now
	db.merchants[merchant.Name] = *merchant

	db.lastMerchantID++

	return nil
}

// GetAllMerchants retrieves all merchants from localDB.
func (db *localDB) GetAllMerchants(ctx context.Context) ([]service.Merchant, error) {
	db.m.RLock()
	defer db.m.RUnlock()

	var result []service.Merchant
	for _, v := range db.merchants {
		result = append(result, v)
	}

	return result, nil
}

// GetMerchantByName retrieve merchant details by name.
func (db *localDB) GetMerchantByName(ctx context.Context, name string) (service.Merchant, error) {
	db.m.Lock()
	defer db.m.Unlock()

	var merchant, ok = db.merchants[name]
	if !ok {
		return service.Merchant{}, errNotFound
	}

	return merchant, nil
}

// DeleteMerchantByName remove merchant from localDB by name
func (db *localDB) DeleteMerchantByName(ctx context.Context, name string) error {
	db.m.Lock()
	defer db.m.Unlock()

	if _, ok := db.merchants[name]; !ok {
		return errNotFound
	}

	delete(db.merchants, name)

	return nil
}

func NewLocalDBMerchantManager() (service.MerchantManager, error) {
	return newlocalDB()
}
