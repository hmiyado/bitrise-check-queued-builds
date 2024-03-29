// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUserProfileParams creates a new UserProfileParams object
// with the default values initialized.
func NewUserProfileParams() *UserProfileParams {

	return &UserProfileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserProfileParamsWithTimeout creates a new UserProfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserProfileParamsWithTimeout(timeout time.Duration) *UserProfileParams {

	return &UserProfileParams{

		timeout: timeout,
	}
}

// NewUserProfileParamsWithContext creates a new UserProfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserProfileParamsWithContext(ctx context.Context) *UserProfileParams {

	return &UserProfileParams{

		Context: ctx,
	}
}

// NewUserProfileParamsWithHTTPClient creates a new UserProfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserProfileParamsWithHTTPClient(client *http.Client) *UserProfileParams {

	return &UserProfileParams{
		HTTPClient: client,
	}
}

/*UserProfileParams contains all the parameters to send to the API endpoint
for the user profile operation typically these are written to a http.Request
*/
type UserProfileParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user profile params
func (o *UserProfileParams) WithTimeout(timeout time.Duration) *UserProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user profile params
func (o *UserProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user profile params
func (o *UserProfileParams) WithContext(ctx context.Context) *UserProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user profile params
func (o *UserProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user profile params
func (o *UserProfileParams) WithHTTPClient(client *http.Client) *UserProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user profile params
func (o *UserProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *UserProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
