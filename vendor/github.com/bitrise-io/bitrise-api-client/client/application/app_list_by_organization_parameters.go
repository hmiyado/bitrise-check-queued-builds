// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewAppListByOrganizationParams creates a new AppListByOrganizationParams object
// with the default values initialized.
func NewAppListByOrganizationParams() *AppListByOrganizationParams {
	var ()
	return &AppListByOrganizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAppListByOrganizationParamsWithTimeout creates a new AppListByOrganizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAppListByOrganizationParamsWithTimeout(timeout time.Duration) *AppListByOrganizationParams {
	var ()
	return &AppListByOrganizationParams{

		timeout: timeout,
	}
}

// NewAppListByOrganizationParamsWithContext creates a new AppListByOrganizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewAppListByOrganizationParamsWithContext(ctx context.Context) *AppListByOrganizationParams {
	var ()
	return &AppListByOrganizationParams{

		Context: ctx,
	}
}

// NewAppListByOrganizationParamsWithHTTPClient creates a new AppListByOrganizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAppListByOrganizationParamsWithHTTPClient(client *http.Client) *AppListByOrganizationParams {
	var ()
	return &AppListByOrganizationParams{
		HTTPClient: client,
	}
}

/*AppListByOrganizationParams contains all the parameters to send to the API endpoint
for the app list by organization operation typically these are written to a http.Request
*/
type AppListByOrganizationParams struct {

	/*Limit
	  Max number of elements per page (default: 50)

	*/
	Limit *int64
	/*Next
	  Slug of the first app in the response

	*/
	Next *string
	/*OrgSlug
	  Organization slug

	*/
	OrgSlug string
	/*SortBy
	  Order of applications: sort them based on when they were created or the time of their last build

	*/
	SortBy *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the app list by organization params
func (o *AppListByOrganizationParams) WithTimeout(timeout time.Duration) *AppListByOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the app list by organization params
func (o *AppListByOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the app list by organization params
func (o *AppListByOrganizationParams) WithContext(ctx context.Context) *AppListByOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the app list by organization params
func (o *AppListByOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the app list by organization params
func (o *AppListByOrganizationParams) WithHTTPClient(client *http.Client) *AppListByOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the app list by organization params
func (o *AppListByOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the app list by organization params
func (o *AppListByOrganizationParams) WithLimit(limit *int64) *AppListByOrganizationParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the app list by organization params
func (o *AppListByOrganizationParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithNext adds the next to the app list by organization params
func (o *AppListByOrganizationParams) WithNext(next *string) *AppListByOrganizationParams {
	o.SetNext(next)
	return o
}

// SetNext adds the next to the app list by organization params
func (o *AppListByOrganizationParams) SetNext(next *string) {
	o.Next = next
}

// WithOrgSlug adds the orgSlug to the app list by organization params
func (o *AppListByOrganizationParams) WithOrgSlug(orgSlug string) *AppListByOrganizationParams {
	o.SetOrgSlug(orgSlug)
	return o
}

// SetOrgSlug adds the orgSlug to the app list by organization params
func (o *AppListByOrganizationParams) SetOrgSlug(orgSlug string) {
	o.OrgSlug = orgSlug
}

// WithSortBy adds the sortBy to the app list by organization params
func (o *AppListByOrganizationParams) WithSortBy(sortBy *string) *AppListByOrganizationParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the app list by organization params
func (o *AppListByOrganizationParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WriteToRequest writes these params to a swagger request
func (o *AppListByOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Next != nil {

		// query param next
		var qrNext string
		if o.Next != nil {
			qrNext = *o.Next
		}
		qNext := qrNext
		if qNext != "" {
			if err := r.SetQueryParam("next", qNext); err != nil {
				return err
			}
		}

	}

	// path param org-slug
	if err := r.SetPathParam("org-slug", o.OrgSlug); err != nil {
		return err
	}

	if o.SortBy != nil {

		// query param sort_by
		var qrSortBy string
		if o.SortBy != nil {
			qrSortBy = *o.SortBy
		}
		qSortBy := qrSortBy
		if qSortBy != "" {
			if err := r.SetQueryParam("sort_by", qSortBy); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}