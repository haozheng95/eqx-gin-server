// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SSHKeyInput SSH key input
//
// swagger:model SSHKeyInput
type SSHKeyInput struct {

	// key
	Key string `json:"key,omitempty"`

	// label
	Label string `json:"label,omitempty"`
}

// Validate validates this SSH key input
func (m *SSHKeyInput) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this SSH key input based on context it is used
func (m *SSHKeyInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SSHKeyInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SSHKeyInput) UnmarshalBinary(b []byte) error {
	var res SSHKeyInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
