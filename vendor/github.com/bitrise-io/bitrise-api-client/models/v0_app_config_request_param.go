// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V0AppConfigRequestParam v0 app config request param
// swagger:model v0.AppConfigRequestParam
type V0AppConfigRequestParam struct {

	// app config datastore yaml
	AppConfigDatastoreYaml string `json:"app_config_datastore_yaml,omitempty"`
}

// Validate validates this v0 app config request param
func (m *V0AppConfigRequestParam) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V0AppConfigRequestParam) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V0AppConfigRequestParam) UnmarshalBinary(b []byte) error {
	var res V0AppConfigRequestParam
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
