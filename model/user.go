package model

import (
	"time"

	"github.com/gexaigor/MyRestAPI/service"
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// User represents the user for this application
//
// swagger:model User
type User struct {
	// the id for this user
	//
	// required: false
	// min: 1
	ID int64 `json:"id"`

	// the login for this user
	// required: true
	// min length: 5
	// max length: 50
	Login string `json:"login"`

	// the email for this user
	// required: true
	Email string `json:"email"`

	// the password for this user
	// required: false
	// min length: 6
	// max length: 100
	Password string `json:"-"`

	// swagger:allOf
	Role Role `json:"role"`

	// the date of user created
	// required: false
	CreatedOn time.Time `json:"created_on"`
}

// Validate ...
func (a *User) Validate() error {
	return validation.ValidateStruct(
		a,
		validation.Field(&a.Login, validation.Required, validation.Length(5, 50)),
		validation.Field(&a.Email, validation.Required, is.Email),
		validation.Field(&a.Password, validation.Required, validation.Length(6, 100)),
		validation.Field(&a.Role, validation.By(isRole())))
}

// BeforeSave ...
func (a *User) BeforeSave() error {
	enc, err := service.EncryptString(a.Password)
	if err != nil {
		return err
	}

	a.Password = enc
	return nil
}
