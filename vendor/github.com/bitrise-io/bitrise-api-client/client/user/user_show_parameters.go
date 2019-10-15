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

// NewUserShowParams creates a new UserShowParams object
// with the default values initialized.
func NewUserShowParams() *UserShowParams {
	var ()
	return &UserShowParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserShowParamsWithTimeout creates a new UserShowParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserShowParamsWithTimeout(timeout time.Duration) *UserShowParams {
	var ()
	return &UserShowParams{

		timeout: timeout,
	}
}

// NewUserShowParamsWithContext creates a new UserShowParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserShowParamsWithContext(ctx context.Context) *UserShowParams {
	var ()
	return &UserShowParams{

		Context: ctx,
	}
}

// NewUserShowParamsWithHTTPClient creates a new UserShowParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserShowParamsWithHTTPClient(client *http.Client) *UserShowParams {
	var ()
	return &UserShowParams{
		HTTPClient: client,
	}
}

/*UserShowParams contains all the parameters to send to the API endpoint
for the user show operation typically these are written to a http.Request
*/
type UserShowParams struct {

	/*UserSlug
	  User slug

	*/
	UserSlug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user show params
func (o *UserShowParams) WithTimeout(timeout time.Duration) *UserShowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user show params
func (o *UserShowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user show params
func (o *UserShowParams) WithContext(ctx context.Context) *UserShowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user show params
func (o *UserShowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user show params
func (o *UserShowParams) WithHTTPClient(client *http.Client) *UserShowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user show params
func (o *UserShowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserSlug adds the userSlug to the user show params
func (o *UserShowParams) WithUserSlug(userSlug string) *UserShowParams {
	o.SetUserSlug(userSlug)
	return o
}

// SetUserSlug adds the userSlug to the user show params
func (o *UserShowParams) SetUserSlug(userSlug string) {
	o.UserSlug = userSlug
}

// WriteToRequest writes these params to a swagger request
func (o *UserShowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param user-slug
	if err := r.SetPathParam("user-slug", o.UserSlug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}