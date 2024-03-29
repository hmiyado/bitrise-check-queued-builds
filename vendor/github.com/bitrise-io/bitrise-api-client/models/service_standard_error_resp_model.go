// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ServiceStandardErrorRespModel service standard error resp model
// swagger:model service.StandardErrorRespModel
type ServiceStandardErrorRespModel struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this service standard error resp model
func (m *ServiceStandardErrorRespModel) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceStandardErrorRespModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceStandardErrorRespModel) UnmarshalBinary(b []byte) error {
	var res ServiceStandardErrorRespModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
