package mocks

// This file contains all common error for package mocks.

import "errors"

var (
	errNotFound          = errors.New("not found")
	errDuplicateMerchant = errors.New("duplicate merchant")
	errTxnRejected       = errors.New("rejected! (reason: credit limit)")
)
