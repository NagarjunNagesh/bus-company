// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Trip trip
//
// swagger:model Trip
type Trip struct {

	// dates
	// Required: true
	Dates *string `json:"dates"`

	// destination
	// Required: true
	Destination *string `json:"destination"`

	// origin
	// Required: true
	Origin *string `json:"origin"`

	// price
	// Required: true
	Price *float64 `json:"price"`
}

// Validate validates this trip
func (m *Trip) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrigin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrice(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Trip) validateDates(formats strfmt.Registry) error {

	if err := validate.Required("dates", "body", m.Dates); err != nil {
		return err
	}

	return nil
}

func (m *Trip) validateDestination(formats strfmt.Registry) error {

	if err := validate.Required("destination", "body", m.Destination); err != nil {
		return err
	}

	return nil
}

func (m *Trip) validateOrigin(formats strfmt.Registry) error {

	if err := validate.Required("origin", "body", m.Origin); err != nil {
		return err
	}

	return nil
}

func (m *Trip) validatePrice(formats strfmt.Registry) error {

	if err := validate.Required("price", "body", m.Price); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this trip based on context it is used
func (m *Trip) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Trip) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Trip) UnmarshalBinary(b []byte) error {
	var res Trip
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
