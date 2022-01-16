package service

import "time"

type Merchant struct {
	// Merchant unique ID
	ID   int64 // Merchant name
	Name string
	// Merchant Email address.
	Email string
	// Discount offer by Merchant.
	Discount float64
	// Total payment from BNPL service to merchant.
	TotalPayment float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
