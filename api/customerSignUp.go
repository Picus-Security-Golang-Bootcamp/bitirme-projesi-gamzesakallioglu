package api

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

type SignUpCustomer struct {
	Name     *string `json:"name"`
	Email    *string `gorm:"unique" json:"email"`
	Password *string `json:"password"`
	Address  string  `json:"address"`
	Phone    string  `json:"phone"`
}

func (c *SignUpCustomer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := c.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := c.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := c.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Validate name - not null
func (c *SignUpCustomer) validateName(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", c.Email); err != nil {
		return err
	}

	return nil
}

// Validate email - not null
func (c *SignUpCustomer) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", c.Email); err != nil {
		return err
	}

	return nil
}

// Validate password - not null
func (c *SignUpCustomer) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", c.Password); err != nil {
		return err
	}

	return nil
}
