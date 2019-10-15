// Code generated by go-swagger; DO NOT EDIT.

package outgoing_webhook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new outgoing webhook API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for outgoing webhook API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
OutgoingWebhookCreate creates an outgoing webhook for an app

Create an outgoing webhook for a specified Bitrise app: this can be used to send build events to a specified URL with custom headers. Currently, only build events can trigger outgoing webhooks.
*/
func (a *Client) OutgoingWebhookCreate(params *OutgoingWebhookCreateParams, authInfo runtime.ClientAuthInfoWriter) (*OutgoingWebhookCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOutgoingWebhookCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "outgoing-webhook-create",
		Method:             "POST",
		PathPattern:        "/apps/{app-slug}/outgoing-webhooks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &OutgoingWebhookCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*OutgoingWebhookCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for outgoing-webhook-create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OutgoingWebhookDelete deletes an outgoing webhook of an app

Delete an existing outgoing webhook for a specified Bitrise app.
*/
func (a *Client) OutgoingWebhookDelete(params *OutgoingWebhookDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*OutgoingWebhookDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOutgoingWebhookDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "outgoing-webhook-delete",
		Method:             "DELETE",
		PathPattern:        "/apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &OutgoingWebhookDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*OutgoingWebhookDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for outgoing-webhook-delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OutgoingWebhookList lists the outgoing webhooks of an app

List all the outgoing webhooks registered for a specified Bitrise app. This returns all the relevant data of the webhook, including the slug of the webhook and its URL.
*/
func (a *Client) OutgoingWebhookList(params *OutgoingWebhookListParams, authInfo runtime.ClientAuthInfoWriter) (*OutgoingWebhookListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOutgoingWebhookListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "outgoing-webhook-list",
		Method:             "GET",
		PathPattern:        "/apps/{app-slug}/outgoing-webhooks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &OutgoingWebhookListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*OutgoingWebhookListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for outgoing-webhook-list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OutgoingWebhookUpdate updates an outgoing webhook of an app

Update an existing outgoing webhook (URL, events, secrets and headers) for a specified Bitrise app. Even if you do not want to change one of the parameters, you still have to provide that parameter as well: simply use its existing value.
*/
func (a *Client) OutgoingWebhookUpdate(params *OutgoingWebhookUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*OutgoingWebhookUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOutgoingWebhookUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "outgoing-webhook-update",
		Method:             "PUT",
		PathPattern:        "/apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &OutgoingWebhookUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*OutgoingWebhookUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for outgoing-webhook-update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}