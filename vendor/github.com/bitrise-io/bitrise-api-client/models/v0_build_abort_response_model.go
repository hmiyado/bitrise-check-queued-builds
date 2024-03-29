// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V0BuildAbortResponseModel v0 build abort response model
// swagger:model v0.BuildAbortResponseModel
type V0BuildAbortResponseModel struct {

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this v0 build abort response model
func (m *V0BuildAbortResponseModel) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V0BuildAbortResponseModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V0BuildAbortResponseModel) UnmarshalBinary(b []byte) error {
	var res V0BuildAbortResponseModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
