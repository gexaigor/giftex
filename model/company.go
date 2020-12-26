package model

import validation "github.com/go-ozzo/ozzo-validation/v4"

// Company represents the company for this application
//
// swagger:model Company
type Company struct {
	// the id for this company
	//
	// required: false
	// min: 1
	ID int64 `json:"id"`

	// swagger:allOf
	User *User `json:"user"`

	// the bin for this company
	// required: true
	// min length: 12
	// max length: 12
	BIN string `json:"bin"`

	// the name for this company
	// required: true
	// min length: 3
	// max length: 255
	Name string `json:"name"`

	// the address for this company
	// required: false
	// min length: 6
	// max length: 255
	Address string `json:"address"`
}

// Validate ...
func (c *Company) Validate() error {
	return validation.ValidateStruct(
		c,
		validation.Field(&c.BIN, validation.Length(12, 12)),
		validation.Field(&c.Name, validation.Length(3, 255)),
		validation.Field(&c.Address, validation.Length(6, 255)))
}
