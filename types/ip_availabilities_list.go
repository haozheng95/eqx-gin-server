// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IPAvailabilitiesList IP availabilities list
//
// swagger:model IPAvailabilitiesList
type IPAvailabilitiesList struct {

	// available
	Available []string `json:"available"`
}

// Validate validates this IP availabilities list
func (m *IPAvailabilitiesList) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this IP availabilities list based on context it is used
func (m *IPAvailabilitiesList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IPAvailabilitiesList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPAvailabilitiesList) UnmarshalBinary(b []byte) error {
	var res IPAvailabilitiesList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
