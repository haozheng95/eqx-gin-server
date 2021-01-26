// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CapacityPerFacility capacity per facility
//
// swagger:model CapacityPerFacility
type CapacityPerFacility struct {

	// baremetal 0
	Baremetal0 *CapacityLevelPerBaremetal `json:"baremetal_0,omitempty"`

	// baremetal 1
	Baremetal1 *CapacityLevelPerBaremetal `json:"baremetal_1,omitempty"`

	// baremetal 2
	Baremetal2 *CapacityLevelPerBaremetal `json:"baremetal_2,omitempty"`

	// baremetal 2a
	Baremetal2a *CapacityLevelPerBaremetal `json:"baremetal_2a,omitempty"`

	// baremetal 2a2
	Baremetal2a2 *CapacityLevelPerBaremetal `json:"baremetal_2a2,omitempty"`

	// baremetal 3
	Baremetal3 *CapacityLevelPerBaremetal `json:"baremetal_3,omitempty"`

	// baremetal s
	Baremetals *CapacityLevelPerBaremetal `json:"baremetal_s,omitempty"`

	// c2 medium x86
	C2MediumX86 *CapacityLevelPerBaremetal `json:"c2.medium.x86,omitempty"`

	// m2 xlarge x86
	M2XlargeX86 *CapacityLevelPerBaremetal `json:"m2.xlarge.x86,omitempty"`
}

// Validate validates this capacity per facility
func (m *CapacityPerFacility) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBaremetal0(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBaremetal1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBaremetal2(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBaremetal2a(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBaremetal2a2(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBaremetal3(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBaremetals(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateC2MediumX86(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateM2XlargeX86(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CapacityPerFacility) validateBaremetal0(formats strfmt.Registry) error {
	if swag.IsZero(m.Baremetal0) { // not required
		return nil
	}

	if m.Baremetal0 != nil {
		if err := m.Baremetal0.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("baremetal_0")
			}
			return err
		}
	}

	return nil
}

func (m *CapacityPerFacility) validateBaremetal1(formats strfmt.Registry) error {
	if swag.IsZero(m.Baremetal1) { // not required
		return nil
	}

	if m.Baremetal1 != nil {
		if err := m.Baremetal1.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("baremetal_1")
			}
			return err
		}
	}

	return nil
}

func (m *CapacityPerFacility) validateBaremetal2(formats strfmt.Registry) error {
	if swag.IsZero(m.Baremetal2) { // not required
		return nil
	}

	if m.Baremetal2 != nil {
		if err := m.Baremetal2.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("baremetal_2")
			}
			return err
		}
	}

	return nil
}

func (m *CapacityPerFacility) validateBaremetal2a(formats strfmt.Registry) error {
	if swag.IsZero(m.Baremetal2a) { // not required
		return nil
	}

	if m.Baremetal2a != nil {
		if err := m.Baremetal2a.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("baremetal_2a")
			}
			return err
		}
	}

	return nil
}

func (m *CapacityPerFacility) validateBaremetal2a2(formats strfmt.Registry) error {
	if swag.IsZero(m.Baremetal2a2) { // not required
		return nil
	}

	if m.Baremetal2a2 != nil {
		if err := m.Baremetal2a2.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("baremetal_2a2")
			}
			return err
		}
	}

	return nil
}

func (m *CapacityPerFacility) validateBaremetal3(formats strfmt.Registry) error {
	if swag.IsZero(m.Baremetal3) { // not required
		return nil
	}

	if m.Baremetal3 != nil {
		if err := m.Baremetal3.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("baremetal_3")
			}
			return err
		}
	}

	return nil
}

func (m *CapacityPerFacility) validateBaremetals(formats strfmt.Registry) error {
	if swag.IsZero(m.Baremetals) { // not required
		return nil
	}

	if m.Baremetals != nil {
		if err := m.Baremetals.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("baremetal_s")
			}
			return err
		}
	}

	return nil
}

func (m *CapacityPerFacility) validateC2MediumX86(formats strfmt.Registry) error {
	if swag.IsZero(m.C2MediumX86) { // not required
		return nil
	}

	if m.C2MediumX86 != nil {
		if err := m.C2MediumX86.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("c2.medium.x86")
			}
			return err
		}
	}

	return nil
}

func (m *CapacityPerFacility) validateM2XlargeX86(formats strfmt.Registry) error {
	if swag.IsZero(m.M2XlargeX86) { // not required
		return nil
	}

	if m.M2XlargeX86 != nil {
		if err := m.M2XlargeX86.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("m2.xlarge.x86")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this capacity per facility based on the context it is used
func (m *CapacityPerFacility) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBaremetal0(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBaremetal1(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBaremetal2(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBaremetal2a(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBaremetal2a2(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBaremetal3(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBaremetals(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateC2MediumX86(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateM2XlargeX86(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CapacityPerFacility) contextValidateBaremetal0(ctx context.Context, formats strfmt.Registry) error {

	if m.Baremetal0 != nil {
		if err := m.Baremetal0.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("baremetal_0")
			}
			return err
		}
	}

	return nil
}

func (m *CapacityPerFacility) contextValidateBaremetal1(ctx context.Context, formats strfmt.Registry) error {

	if m.Baremetal1 != nil {
		if err := m.Baremetal1.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("baremetal_1")
			}
			return err
		}
	}

	return nil
}

func (m *CapacityPerFacility) contextValidateBaremetal2(ctx context.Context, formats strfmt.Registry) error {

	if m.Baremetal2 != nil {
		if err := m.Baremetal2.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("baremetal_2")
			}
			return err
		}
	}

	return nil
}

func (m *CapacityPerFacility) contextValidateBaremetal2a(ctx context.Context, formats strfmt.Registry) error {

	if m.Baremetal2a != nil {
		if err := m.Baremetal2a.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("baremetal_2a")
			}
			return err
		}
	}

	return nil
}

func (m *CapacityPerFacility) contextValidateBaremetal2a2(ctx context.Context, formats strfmt.Registry) error {

	if m.Baremetal2a2 != nil {
		if err := m.Baremetal2a2.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("baremetal_2a2")
			}
			return err
		}
	}

	return nil
}

func (m *CapacityPerFacility) contextValidateBaremetal3(ctx context.Context, formats strfmt.Registry) error {

	if m.Baremetal3 != nil {
		if err := m.Baremetal3.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("baremetal_3")
			}
			return err
		}
	}

	return nil
}

func (m *CapacityPerFacility) contextValidateBaremetals(ctx context.Context, formats strfmt.Registry) error {

	if m.Baremetals != nil {
		if err := m.Baremetals.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("baremetal_s")
			}
			return err
		}
	}

	return nil
}

func (m *CapacityPerFacility) contextValidateC2MediumX86(ctx context.Context, formats strfmt.Registry) error {

	if m.C2MediumX86 != nil {
		if err := m.C2MediumX86.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("c2.medium.x86")
			}
			return err
		}
	}

	return nil
}

func (m *CapacityPerFacility) contextValidateM2XlargeX86(ctx context.Context, formats strfmt.Registry) error {

	if m.M2XlargeX86 != nil {
		if err := m.M2XlargeX86.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("m2.xlarge.x86")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CapacityPerFacility) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CapacityPerFacility) UnmarshalBinary(b []byte) error {
	var res CapacityPerFacility
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
