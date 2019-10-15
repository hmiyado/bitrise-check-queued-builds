// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V0AvatarCandidateCreateResponseItem v0 avatar candidate create response item
// swagger:model v0.AvatarCandidateCreateResponseItem
type V0AvatarCandidateCreateResponseItem struct {

	// filename
	Filename string `json:"filename,omitempty"`

	// filesize
	Filesize int64 `json:"filesize,omitempty"`

	// slug
	Slug string `json:"slug,omitempty"`

	// upload url
	UploadURL string `json:"upload_url,omitempty"`
}

// Validate validates this v0 avatar candidate create response item
func (m *V0AvatarCandidateCreateResponseItem) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V0AvatarCandidateCreateResponseItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V0AvatarCandidateCreateResponseItem) UnmarshalBinary(b []byte) error {
	var res V0AvatarCandidateCreateResponseItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
