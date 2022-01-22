package service

import "time"

type TransactionStatus int64

const (
	Successful = iota + 1
	Pending
	Failed
)

// Transaction status.
var (
	SuccessfulStatus = "successful"
	PendingStatus    = "pending"
	FailedStaus      = "failed"
)

var statusTransactionStatusToStrMap = map[TransactionStatus]string{
	Successful: SuccessfulStatus,
	Pending:    PendingStatus,
	Failed:     FailedStaus,
}

// Transactions represents an transaction entities.
type Transaction struct {
	ID         int64
	UserID     int64
	MerchantID int64
	Amount     float64
	Status     TransactionStatus
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
