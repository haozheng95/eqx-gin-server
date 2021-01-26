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

// SpotPricesReport spot prices report
//
// swagger:model SpotPricesReport
type SpotPricesReport struct {

	// ams1
	Ams1 *SpotPricesPerFacility `json:"ams1,omitempty"`

	// atl1
	Atl1 *SpotPricesPerNewFacility `json:"atl1,omitempty"`

	// dfw1
	Dfw1 *SpotPricesPerNewFacility `json:"dfw1,omitempty"`

	// ewr1
	Ewr1 *SpotPricesPerFacility `json:"ewr1,omitempty"`

	// fra1
	Fra1 *SpotPricesPerNewFacility `json:"fra1,omitempty"`

	// iad1
	Iad1 *SpotPricesPerNewFacility `json:"iad1,omitempty"`

	// lax1
	Lax1 *SpotPricesPerNewFacility `json:"lax1,omitempty"`

	// nrt1
	Nrt1 *SpotPricesPerFacility `json:"nrt1,omitempty"`

	// ord1
	Ord1 *SpotPricesPerNewFacility `json:"ord1,omitempty"`

	// sea1
	Sea1 *SpotPricesPerNewFacility `json:"sea1,omitempty"`

	// sin1
	Sin1 *SpotPricesPerNewFacility `json:"sin1,omitempty"`

	// sjc1
	Sjc1 *SpotPricesPerFacility `json:"sjc1,omitempty"`

	// syd1
	Syd1 *SpotPricesPerNewFacility `json:"syd1,omitempty"`

	// yyz1
	Yyz1 *SpotPricesPerNewFacility `json:"yyz1,omitempty"`
}

// Validate validates this spot prices report
func (m *SpotPricesReport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAms1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAtl1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDfw1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEwr1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFra1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIad1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLax1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNrt1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrd1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSea1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSin1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSjc1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSyd1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateYyz1(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SpotPricesReport) validateAms1(formats strfmt.Registry) error {
	if swag.IsZero(m.Ams1) { // not required
		return nil
	}

	if m.Ams1 != nil {
		if err := m.Ams1.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ams1")
			}
			return err
		}
	}

	return nil
}

func (m *SpotPricesReport) validateAtl1(formats strfmt.Registry) error {
	if swag.IsZero(m.Atl1) { // not required
		return nil
	}

	if m.Atl1 != nil {
		if err := m.Atl1.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("atl1")
			}
			return err
		}
	}

	return nil
}

func (m *SpotPricesReport) validateDfw1(formats strfmt.Registry) error {
	if swag.IsZero(m.Dfw1) { // not required
		return nil
	}

	if m.Dfw1 != nil {
		if err := m.Dfw1.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dfw1")
			}
			return err
		}
	}

	return nil
}

func (m *SpotPricesReport) validateEwr1(formats strfmt.Registry) error {
	if swag.IsZero(m.Ewr1) { // not required
		return nil
	}

	if m.Ewr1 != nil {
		if err := m.Ewr1.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ewr1")
			}
			return err
		}
	}

	return nil
}

func (m *SpotPricesReport) validateFra1(formats strfmt.Registry) error {
	if swag.IsZero(m.Fra1) { // not required
		return nil
	}

	if m.Fra1 != nil {
		if err := m.Fra1.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fra1")
			}
			return err
		}
	}

	return nil
}

func (m *SpotPricesReport) validateIad1(formats strfmt.Registry) error {
	if swag.IsZero(m.Iad1) { // not required
		return nil
	}

	if m.Iad1 != nil {
		if err := m.Iad1.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iad1")
			}
			return err
		}
	}

	return nil
}

func (m *SpotPricesReport) validateLax1(formats strfmt.Registry) error {
	if swag.IsZero(m.Lax1) { // not required
		return nil
	}

	if m.Lax1 != nil {
		if err := m.Lax1.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lax1")
			}
			return err
		}
	}

	return nil
}

func (m *SpotPricesReport) validateNrt1(formats strfmt.Registry) error {
	if swag.IsZero(m.Nrt1) { // not required
		return nil
	}

	if m.Nrt1 != nil {
		if err := m.Nrt1.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nrt1")
			}
			return err
		}
	}

	return nil
}

func (m *SpotPricesReport) validateOrd1(formats strfmt.Registry) error {
	if swag.IsZero(m.Ord1) { // not required
		return nil
	}

	if m.Ord1 != nil {
		if err := m.Ord1.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ord1")
			}
			return err
		}
	}

	return nil
}

func (m *SpotPricesReport) validateSea1(formats strfmt.Registry) error {
	if swag.IsZero(m.Sea1) { // not required
		return nil
	}

	if m.Sea1 != nil {
		if err := m.Sea1.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sea1")
			}
			return err
		}
	}

	return nil
}

func (m *SpotPricesReport) validateSin1(formats strfmt.Registry) error {
	if swag.IsZero(m.Sin1) { // not required
		return nil
	}

	if m.Sin1 != nil {
		if err := m.Sin1.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sin1")
			}
			return err
		}
	}

	return nil
}

func (m *SpotPricesReport) validateSjc1(formats strfmt.Registry) error {
	if swag.IsZero(m.Sjc1) { // not required
		return nil
	}

	if m.Sjc1 != nil {
		if err := m.Sjc1.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sjc1")
			}
			return err
		}
	}

	return nil
}

func (m *SpotPricesReport) validateSyd1(formats strfmt.Registry) error {
	if swag.IsZero(m.Syd1) { // not required
		return nil
	}

	if m.Syd1 != nil {
		if err := m.Syd1.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("syd1")
			}
			return err
		}
	}

	return nil
}

func (m *SpotPricesReport) validateYyz1(formats strfmt.Registry) error {
	if swag.IsZero(m.Yyz1) { // not required
		return nil
	}

	if m.Yyz1 != nil {
		if err := m.Yyz1.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("yyz1")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this spot prices report based on the context it is used
func (m *SpotPricesReport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAms1(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAtl1(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDfw1(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEwr1(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFra1(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIad1(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLax1(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNrt1(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOrd1(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSea1(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSin1(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSjc1(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSyd1(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateYyz1(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SpotPricesReport) contextValidateAms1(ctx context.Context, formats strfmt.Registry) error {

	if m.Ams1 != nil {
		if err := m.Ams1.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ams1")
			}
			return err
		}
	}

	return nil
}

func (m *SpotPricesReport) contextValidateAtl1(ctx context.Context, formats strfmt.Registry) error {

	if m.Atl1 != nil {
		if err := m.Atl1.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("atl1")
			}
			return err
		}
	}

	return nil
}

func (m *SpotPricesReport) contextValidateDfw1(ctx context.Context, formats strfmt.Registry) error {

	if m.Dfw1 != nil {
		if err := m.Dfw1.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dfw1")
			}
			return err
		}
	}

	return nil
}

func (m *SpotPricesReport) contextValidateEwr1(ctx context.Context, formats strfmt.Registry) error {

	if m.Ewr1 != nil {
		if err := m.Ewr1.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ewr1")
			}
			return err
		}
	}

	return nil
}

func (m *SpotPricesReport) contextValidateFra1(ctx context.Context, formats strfmt.Registry) error {

	if m.Fra1 != nil {
		if err := m.Fra1.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fra1")
			}
			return err
		}
	}

	return nil
}

func (m *SpotPricesReport) contextValidateIad1(ctx context.Context, formats strfmt.Registry) error {

	if m.Iad1 != nil {
		if err := m.Iad1.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iad1")
			}
			return err
		}
	}

	return nil
}

func (m *SpotPricesReport) contextValidateLax1(ctx context.Context, formats strfmt.Registry) error {

	if m.Lax1 != nil {
		if err := m.Lax1.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lax1")
			}
			return err
		}
	}

	return nil
}

func (m *SpotPricesReport) contextValidateNrt1(ctx context.Context, formats strfmt.Registry) error {

	if m.Nrt1 != nil {
		if err := m.Nrt1.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nrt1")
			}
			return err
		}
	}

	return nil
}

func (m *SpotPricesReport) contextValidateOrd1(ctx context.Context, formats strfmt.Registry) error {

	if m.Ord1 != nil {
		if err := m.Ord1.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ord1")
			}
			return err
		}
	}

	return nil
}

func (m *SpotPricesReport) contextValidateSea1(ctx context.Context, formats strfmt.Registry) error {

	if m.Sea1 != nil {
		if err := m.Sea1.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sea1")
			}
			return err
		}
	}

	return nil
}

func (m *SpotPricesReport) contextValidateSin1(ctx context.Context, formats strfmt.Registry) error {

	if m.Sin1 != nil {
		if err := m.Sin1.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sin1")
			}
			return err
		}
	}

	return nil
}

func (m *SpotPricesReport) contextValidateSjc1(ctx context.Context, formats strfmt.Registry) error {

	if m.Sjc1 != nil {
		if err := m.Sjc1.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sjc1")
			}
			return err
		}
	}

	return nil
}

func (m *SpotPricesReport) contextValidateSyd1(ctx context.Context, formats strfmt.Registry) error {

	if m.Syd1 != nil {
		if err := m.Syd1.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("syd1")
			}
			return err
		}
	}

	return nil
}

func (m *SpotPricesReport) contextValidateYyz1(ctx context.Context, formats strfmt.Registry) error {

	if m.Yyz1 != nil {
		if err := m.Yyz1.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("yyz1")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SpotPricesReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SpotPricesReport) UnmarshalBinary(b []byte) error {
	var res SpotPricesReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
