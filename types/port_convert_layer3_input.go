// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PortConvertLayer3Input port convert layer3 input
//
// swagger:model PortConvertLayer3Input
type PortConvertLayer3Input struct {

	// request ips
	RequestIps []*PortConvertLayer3InputRequestIpsItems0 `json:"request_ips"`
}

// Validate validates this port convert layer3 input
func (m *PortConvertLayer3Input) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRequestIps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortConvertLayer3Input) validateRequestIps(formats strfmt.Registry) error {
	if swag.IsZero(m.RequestIps) { // not required
		return nil
	}

	for i := 0; i < len(m.RequestIps); i++ {
		if swag.IsZero(m.RequestIps[i]) { // not required
			continue
		}

		if m.RequestIps[i] != nil {
			if err := m.RequestIps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("request_ips" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this port convert layer3 input based on the context it is used
func (m *PortConvertLayer3Input) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRequestIps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortConvertLayer3Input) contextValidateRequestIps(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RequestIps); i++ {

		if m.RequestIps[i] != nil {
			if err := m.RequestIps[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("request_ips" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PortConvertLayer3Input) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortConvertLayer3Input) UnmarshalBinary(b []byte) error {
	var res PortConvertLayer3Input
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PortConvertLayer3InputRequestIpsItems0 port convert layer3 input request ips items0
//
// swagger:model PortConvertLayer3InputRequestIpsItems0
type PortConvertLayer3InputRequestIpsItems0 struct {

	// address family
	AddressFamily int64 `json:"address_family,omitempty"`

	// public
	Public bool `json:"public,omitempty"`
}

// Validate validates this port convert layer3 input request ips items0
func (m *PortConvertLayer3InputRequestIpsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this port convert layer3 input request ips items0 based on context it is used
func (m *PortConvertLayer3InputRequestIpsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortConvertLayer3InputRequestIpsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortConvertLayer3InputRequestIpsItems0) UnmarshalBinary(b []byte) error {
	var res PortConvertLayer3InputRequestIpsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
