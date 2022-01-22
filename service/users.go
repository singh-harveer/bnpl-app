package service

import (
	"time"
)

// User represents User Object.
type User struct {
	ID          int64      `json:"id,omitempty"`
	Name        string     `json:"userName,omitempty"`
	Email       string     `json:"email,omitempty"`
	CreditLimit float64    `json:"creditLimit,omitempty"`
	DueAmount   float64    `json:"dueAmount,omitempty"`
	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
}
