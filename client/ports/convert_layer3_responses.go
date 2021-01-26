// Code generated by go-swagger; DO NOT EDIT.

package ports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/t0mk/gometal/types"
)

// ConvertLayer3Reader is a Reader for the ConvertLayer3 structure.
type ConvertLayer3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConvertLayer3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConvertLayer3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewConvertLayer3Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewConvertLayer3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewConvertLayer3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewConvertLayer3UnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewConvertLayer3OK creates a ConvertLayer3OK with default headers values
func NewConvertLayer3OK() *ConvertLayer3OK {
	return &ConvertLayer3OK{}
}

/* ConvertLayer3OK describes a response with status code 200, with default header values.

ok
*/
type ConvertLayer3OK struct {
	Payload *types.Port
}

func (o *ConvertLayer3OK) Error() string {
	return fmt.Sprintf("[POST /ports/{id}/convert/layer-3][%d] convertLayer3OK  %+v", 200, o.Payload)
}
func (o *ConvertLayer3OK) GetPayload() *types.Port {
	return o.Payload
}

func (o *ConvertLayer3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(types.Port)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConvertLayer3Unauthorized creates a ConvertLayer3Unauthorized with default headers values
func NewConvertLayer3Unauthorized() *ConvertLayer3Unauthorized {
	return &ConvertLayer3Unauthorized{}
}

/* ConvertLayer3Unauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type ConvertLayer3Unauthorized struct {
}

func (o *ConvertLayer3Unauthorized) Error() string {
	return fmt.Sprintf("[POST /ports/{id}/convert/layer-3][%d] convertLayer3Unauthorized ", 401)
}

func (o *ConvertLayer3Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewConvertLayer3Forbidden creates a ConvertLayer3Forbidden with default headers values
func NewConvertLayer3Forbidden() *ConvertLayer3Forbidden {
	return &ConvertLayer3Forbidden{}
}

/* ConvertLayer3Forbidden describes a response with status code 403, with default header values.

forbidden
*/
type ConvertLayer3Forbidden struct {
}

func (o *ConvertLayer3Forbidden) Error() string {
	return fmt.Sprintf("[POST /ports/{id}/convert/layer-3][%d] convertLayer3Forbidden ", 403)
}

func (o *ConvertLayer3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewConvertLayer3NotFound creates a ConvertLayer3NotFound with default headers values
func NewConvertLayer3NotFound() *ConvertLayer3NotFound {
	return &ConvertLayer3NotFound{}
}

/* ConvertLayer3NotFound describes a response with status code 404, with default header values.

not found
*/
type ConvertLayer3NotFound struct {
}

func (o *ConvertLayer3NotFound) Error() string {
	return fmt.Sprintf("[POST /ports/{id}/convert/layer-3][%d] convertLayer3NotFound ", 404)
}

func (o *ConvertLayer3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewConvertLayer3UnprocessableEntity creates a ConvertLayer3UnprocessableEntity with default headers values
func NewConvertLayer3UnprocessableEntity() *ConvertLayer3UnprocessableEntity {
	return &ConvertLayer3UnprocessableEntity{}
}

/* ConvertLayer3UnprocessableEntity describes a response with status code 422, with default header values.

unprocessable entity
*/
type ConvertLayer3UnprocessableEntity struct {
}

func (o *ConvertLayer3UnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /ports/{id}/convert/layer-3][%d] convertLayer3UnprocessableEntity ", 422)
}

func (o *ConvertLayer3UnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
