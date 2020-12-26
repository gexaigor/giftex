package model

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func isRole() validation.RuleFunc {
	return func(value interface{}) error {
		role, ok := value.(Role)

		if !ok || role.String() == "" {
			return errors.New("field is not role")
		}

		return nil
	}
}
