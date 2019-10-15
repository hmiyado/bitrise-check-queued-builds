// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bitrise-io/bitrise-api-client/models"
)

// UserShowReader is a Reader for the UserShow structure.
type UserShowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserShowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserShowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserShowBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUserShowUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUserShowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUserShowInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserShowOK creates a UserShowOK with default headers values
func NewUserShowOK() *UserShowOK {
	return &UserShowOK{}
}

/*UserShowOK handles this case with default header values.

OK
*/
type UserShowOK struct {
	Payload *models.V0UserProfileRespModel
}

func (o *UserShowOK) Error() string {
	return fmt.Sprintf("[GET /users/{user-slug}][%d] userShowOK  %+v", 200, o.Payload)
}

func (o *UserShowOK) GetPayload() *models.V0UserProfileRespModel {
	return o.Payload
}

func (o *UserShowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0UserProfileRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserShowBadRequest creates a UserShowBadRequest with default headers values
func NewUserShowBadRequest() *UserShowBadRequest {
	return &UserShowBadRequest{}
}

/*UserShowBadRequest handles this case with default header values.

Bad Request
*/
type UserShowBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *UserShowBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{user-slug}][%d] userShowBadRequest  %+v", 400, o.Payload)
}

func (o *UserShowBadRequest) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *UserShowBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserShowUnauthorized creates a UserShowUnauthorized with default headers values
func NewUserShowUnauthorized() *UserShowUnauthorized {
	return &UserShowUnauthorized{}
}

/*UserShowUnauthorized handles this case with default header values.

Unauthorized
*/
type UserShowUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *UserShowUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/{user-slug}][%d] userShowUnauthorized  %+v", 401, o.Payload)
}

func (o *UserShowUnauthorized) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *UserShowUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserShowNotFound creates a UserShowNotFound with default headers values
func NewUserShowNotFound() *UserShowNotFound {
	return &UserShowNotFound{}
}

/*UserShowNotFound handles this case with default header values.

Not Found
*/
type UserShowNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *UserShowNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{user-slug}][%d] userShowNotFound  %+v", 404, o.Payload)
}

func (o *UserShowNotFound) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *UserShowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserShowInternalServerError creates a UserShowInternalServerError with default headers values
func NewUserShowInternalServerError() *UserShowInternalServerError {
	return &UserShowInternalServerError{}
}

/*UserShowInternalServerError handles this case with default header values.

Internal Server Error
*/
type UserShowInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *UserShowInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/{user-slug}][%d] userShowInternalServerError  %+v", 500, o.Payload)
}

func (o *UserShowInternalServerError) GetPayload() *models.ServiceStandardErrorRespModel {
	return o.Payload
}

func (o *UserShowInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}