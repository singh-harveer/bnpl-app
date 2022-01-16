package service

import "time"

type Merchant struct {
	// Merchant unique ID
	ID ID `json:"id"`
	// Merchant name
	Name string `json:"name"`
	// Merchant Email address.
	Email string `json:"email"`
	// Discount offer by Merchant.
	Discount float64 `json:"discount"`
	// Total payment from BNPL service to merchant.
	TotalPayment float64   `json:"total_payment"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
