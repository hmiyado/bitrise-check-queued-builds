// Code generated by go-swagger; DO NOT EDIT.

package addons

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new addons API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for addons API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddonListByApp gets list of the addons for apps

List all the provisioned addons for the authorized apps
*/
func (a *Client) AddonListByApp(params *AddonListByAppParams, authInfo runtime.ClientAuthInfoWriter) (*AddonListByAppOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddonListByAppParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addon-list-by-app",
		Method:             "GET",
		PathPattern:        "/apps/{app-slug}/addons",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddonListByAppReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddonListByAppOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addon-list-by-app: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AddonListByOrganization gets list of the addons for organization

List all the provisioned addons for organization
*/
func (a *Client) AddonListByOrganization(params *AddonListByOrganizationParams, authInfo runtime.ClientAuthInfoWriter) (*AddonListByOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddonListByOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addon-list-by-organization",
		Method:             "GET",
		PathPattern:        "/organizations/{organization-slug}/addons",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddonListByOrganizationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddonListByOrganizationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addon-list-by-organization: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AddonListByUser gets list of the addons for user

List all the provisioned addons for the authenticated user
*/
func (a *Client) AddonListByUser(params *AddonListByUserParams, authInfo runtime.ClientAuthInfoWriter) (*AddonListByUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddonListByUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addon-list-by-user",
		Method:             "GET",
		PathPattern:        "/users/{user-slug}/addons",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddonListByUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddonListByUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addon-list-by-user: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AddonsList gets list of available bitrise addons

List all the available Bitrise addons
*/
func (a *Client) AddonsList(params *AddonsListParams, authInfo runtime.ClientAuthInfoWriter) (*AddonsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddonsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addons-list",
		Method:             "GET",
		PathPattern:        "/addons",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddonsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddonsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addons-list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AddonsShow gets a specific bitrise addon

Show details of a specific Bitrise addon
*/
func (a *Client) AddonsShow(params *AddonsShowParams, authInfo runtime.ClientAuthInfoWriter) (*AddonsShowOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddonsShowParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addons-show",
		Method:             "GET",
		PathPattern:        "/addons/{addon-id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddonsShowReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddonsShowOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addons-show: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
