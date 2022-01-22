package service

import "time"

// Merchant represents Merchants object
type Merchant struct {
	ID           int64      `json:"id,omitempty"`
	Name         string     `json:"merchantName,omitempty"`
	Email        string     `json:"email,omitempty"`
	Discount     float64    `json:"discount,omitempty"`
	TotalPayment float64    `json:"totalPayment,omitempty"`
	CreatedAt    *time.Time `json:"createdAt,omitempty"`
	UpdatedAt    *time.Time `json:"updatedAt,omitempty"`
}
