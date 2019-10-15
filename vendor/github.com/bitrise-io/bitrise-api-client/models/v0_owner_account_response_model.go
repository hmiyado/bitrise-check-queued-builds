// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V0OwnerAccountResponseModel v0 owner account response model
// swagger:model v0.OwnerAccountResponseModel
type V0OwnerAccountResponseModel struct {

	// account type
	AccountType string `json:"account_type,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// slug
	Slug string `json:"slug,omitempty"`
}

// Validate validates this v0 owner account response model
func (m *V0OwnerAccountResponseModel) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V0OwnerAccountResponseModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V0OwnerAccountResponseModel) UnmarshalBinary(b []byte) error {
	var res V0OwnerAccountResponseModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}