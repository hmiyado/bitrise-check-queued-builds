// Code generated by go-swagger; DO NOT EDIT.

package outgoing_webhook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bitrise-io/bitrise-api-client/models"
)

// OutgoingWebhookUpdateReader is a Reader for the OutgoingWebhookUpdate structure.
type OutgoingWebhookUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OutgoingWebhookUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOutgoingWebhookUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOutgoingWebhookUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOutgoingWebhookUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOutgoingWebhookUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOutgoingWebhookUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOutgoingWebhookUpdateOK creates a OutgoingWebhookUpdateOK with default headers values
func NewOutgoingWebhookUpdateOK() *OutgoingWebhookUpdateOK {
	return &OutgoingWebhookUpdateOK{}
}

/*OutgoingWebhookUpdateOK handles this case with default header values.

OK
*/
type OutgoingWebhookUpdateOK struct {
	Payload *models.V0AppWebhookResponseModel
}

func (o *OutgoingWebhookUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}][%d] outgoingWebhookUpdateOK  %+v", 200, o.Payload)
}

func (o *OutgoingWebhookUpdateOK) GetPayload() *models.V0AppWebhookResponseModel {
	return o.Payload
}

func (o *OutgoingWebhookUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0AppWebhookResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOutgoingWebhookUpdateBadRequest creates a OutgoingWebhookUpdateBadRequest with default headers values
func NewOutgoingWebhookUpdateBadRequest() *OutgoingWebhookUpdateBadRequest {
	return &OutgoingWebhookUpdateBadRequest{}
}

/*OutgoingWebhookUpdateBadRequest handles this case with default header values.

Bad Request
*/
type OutgoingWebhookUpdateBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *OutgoingWebhookUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}][%d] outgoingWebhookUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *OutgoingWebhookUpdateBadRequest) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *OutgoingWebhookUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOutgoingWebhookUpdateUnauthorized creates a OutgoingWebhookUpdateUnauthorized with default headers values
func NewOutgoingWebhookUpdateUnauthorized() *OutgoingWebhookUpdateUnauthorized {
	return &OutgoingWebhookUpdateUnauthorized{}
}

/*OutgoingWebhookUpdateUnauthorized handles this case with default header values.

Unauthorized
*/
type OutgoingWebhookUpdateUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *OutgoingWebhookUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}][%d] outgoingWebhookUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *OutgoingWebhookUpdateUnauthorized) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *OutgoingWebhookUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOutgoingWebhookUpdateNotFound creates a OutgoingWebhookUpdateNotFound with default headers values
func NewOutgoingWebhookUpdateNotFound() *OutgoingWebhookUpdateNotFound {
	return &OutgoingWebhookUpdateNotFound{}
}

/*OutgoingWebhookUpdateNotFound handles this case with default header values.

Not Found
*/
type OutgoingWebhookUpdateNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *OutgoingWebhookUpdateNotFound) Error() string {
	return fmt.Sprintf("[PUT /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}][%d] outgoingWebhookUpdateNotFound  %+v", 404, o.Payload)
}

func (o *OutgoingWebhookUpdateNotFound) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *OutgoingWebhookUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOutgoingWebhookUpdateInternalServerError creates a OutgoingWebhookUpdateInternalServerError with default headers values
func NewOutgoingWebhookUpdateInternalServerError() *OutgoingWebhookUpdateInternalServerError {
	return &OutgoingWebhookUpdateInternalServerError{}
}

/*OutgoingWebhookUpdateInternalServerError handles this case with default header values.

Internal Server Error
*/
type OutgoingWebhookUpdateInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *OutgoingWebhookUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /apps/{app-slug}/outgoing-webhooks/{app-webhook-slug}][%d] outgoingWebhookUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *OutgoingWebhookUpdateInternalServerError) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *OutgoingWebhookUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}