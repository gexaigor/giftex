package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Transaction ...
type Transaction struct {
	ID         int64    `json:"id"`
	User       *User    `json:"-"`
	Company    *Company `json:"-"`
	Experience int64    `json:"experience"`
}

// Validate ...
func (t *Transaction) Validate() error {
	return validation.ValidateStruct(
		t,
		validation.Field(&t.Experience, validation.Required, validation.When(t.Experience > 0)))
}
