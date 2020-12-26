package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// CompanyLevel ...
type CompanyLevel struct {
	ID         int64    `json:"id"`
	Company    *Company `json:"-"`
	Experience int64    `json:"experience"`
	Level      int      `json:"level"`
}

// Validate ...
func (cl *CompanyLevel) Validate() error {
	return validation.ValidateStruct(
		cl,
		validation.Field(&cl.Experience, validation.Required),
		validation.Field(&cl.Level, validation.Required))
}
