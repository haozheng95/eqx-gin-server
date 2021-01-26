// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProjectUpdateInput project update input
//
// swagger:model ProjectUpdateInput
type ProjectUpdateInput struct {

	// backend transfer enabled
	BackendTransferEnabled bool `json:"backend_transfer_enabled,omitempty"`

	// customdata
	Customdata interface{} `json:"customdata,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// payment method id
	// Format: uuid
	PaymentMethodID strfmt.UUID `json:"payment_method_id,omitempty"`
}

// Validate validates this project update input
func (m *ProjectUpdateInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePaymentMethodID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectUpdateInput) validatePaymentMethodID(formats strfmt.Registry) error {
	if swag.IsZero(m.PaymentMethodID) { // not required
		return nil
	}

	if err := validate.FormatOf("payment_method_id", "body", "uuid", m.PaymentMethodID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this project update input based on context it is used
func (m *ProjectUpdateInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProjectUpdateInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectUpdateInput) UnmarshalBinary(b []byte) error {
	var res ProjectUpdateInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
