package localdb

import "errors"

var (
	errNotFound          = errors.New("not found")
	errDuplicateMerchant = errors.New("duplicate merchant")
)