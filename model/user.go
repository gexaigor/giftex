package model

import (
	"time"

	"github.com/gexaigor/MyRestAPI/service"
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// User ...
type User struct {
	ID        int64     `json:"id"`
	Login     string    `json:"login"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Role      Role      `json:"role"`
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
