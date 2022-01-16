package service

import (
	"encoding/json"
	"time"
)

// User represent User Object.
type User struct {
	ID          int64
	Name        string
	Email       string
	CreditLimit float64
	DueAmount   float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type jsonUser struct {
	Name        string  `json:"name"`
	CreditLimit float64 `json:"limit"`
	DueAmount   float64 `json:"dueAmount"`
}

// MarshalJSON to encode User to JSON
func (u *User) MarshalJSON() ([]byte, error) {
	var temp = &jsonUser{
		Name:        u.Name,
		CreditLimit: u.CreditLimit,
		DueAmount:   u.CreditLimit,
	}

	return json.Marshal(temp)
}
