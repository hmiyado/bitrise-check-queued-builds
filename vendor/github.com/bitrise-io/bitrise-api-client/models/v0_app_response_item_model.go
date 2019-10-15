// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V0AppResponseItemModel v0 app response item model
// swagger:model v0.AppResponseItemModel
type V0AppResponseItemModel struct {

	// avatar url
	AvatarURL string `json:"avatar_url,omitempty"`

	// is disabled
	IsDisabled bool `json:"is_disabled,omitempty"`

	// is public
	IsPublic bool `json:"is_public,omitempty"`

	// owner
	Owner *V0OwnerAccountResponseModel `json:"owner,omitempty"`

	// project type
	ProjectType string `json:"project_type,omitempty"`

	// provider
	Provider string `json:"provider,omitempty"`

	// repo owner
	RepoOwner string `json:"repo_owner,omitempty"`

	// repo slug
	RepoSlug string `json:"repo_slug,omitempty"`

	// repo url
	RepoURL string `json:"repo_url,omitempty"`

	// slug
	Slug string `json:"slug,omitempty"`

	// status
	Status int64 `json:"status,omitempty"`

	// title
	Title string `json:"title,omitempty"`
}

// Validate validates this v0 app response item model
func (m *V0AppResponseItemModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V0AppResponseItemModel) validateOwner(formats strfmt.Registry) error {

	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if m.Owner != nil {
		if err := m.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V0AppResponseItemModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V0AppResponseItemModel) UnmarshalBinary(b []byte) error {
	var res V0AppResponseItemModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
