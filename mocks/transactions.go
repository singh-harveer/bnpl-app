package mocks

import (
	"bnpl/service"
	"context"
	"time"
)

var (
	_ service.TransactionManager = (*localDB)(nil)
)

// Create creates new transaction in localDB.
func (db *localDB) Create(ctx context.Context, txn service.Transaction) error {
	db.m.Lock()
	defer db.m.Unlock()

	var now = time.Now().UTC()

	var userName, ok = db.userIDToNameMap[txn.UserID]
	if !ok {
		txn.Status = service.Failed

		return errNotFound
	}
	var userDetails service.User
	userDetails, ok = db.users[userName]
	if !ok {
		txn.Status = service.Failed

		return errNotFound
	}

	if userDetails.CreditLimit < txn.Amount {
		txn.Status = service.Failed
		return errTxnRejected
	}

	userDetails.CreditLimit = userDetails.CreditLimit - txn.Amount
	userDetails.DueAmount = userDetails.DueAmount + txn.Amount
	userDetails.UpdatedAt = &now

	var merchantName string
	merchantName, ok = db.userIDToNameMap[txn.MerchantID]
	if !ok {
		txn.Status = service.Failed

		return errNotFound
	}

	var merchantDetails service.Merchant
	merchantDetails, ok = db.merchants[merchantName]
	if !ok {
		txn.Status = service.Failed

		return errNotFound
	}

	var amountToPayToMerchants = txn.Amount - (txn.Amount * merchantDetails.Discount / 100)
	merchantDetails.TotalPayment = merchantDetails.TotalPayment + amountToPayToMerchants
	merchantDetails.UpdatedAt = &now

	db.lastTransactionID++
	txn.ID = db.lastTransactionID
	txn.CreatedAt = now
	txn.UpdatedAt = now

	txn.Status = service.Successful

	db.users[userName] = userDetails
	db.merchants[merchantName] = merchantDetails
	db.usersTxn[userName] = txn

	return nil
}

// DuePayment deposit amount towards user due.
func (db *localDB) DuePayment(ctx context.Context, name string, amount float64) (service.User, error) {

	return service.User{}, nil
}
