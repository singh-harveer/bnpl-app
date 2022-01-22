package mocks

import "errors"

var (
	errNotFound          = errors.New("not found")
	errDuplicateMerchant = errors.New("duplicate merchant")
	errTxnRejected       = errors.New("rejected! (reason: credit limit)")
)
