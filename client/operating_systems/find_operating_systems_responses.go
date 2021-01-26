// Code generated by go-swagger; DO NOT EDIT.

package operating_systems

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/t0mk/gometal/types"
)

// FindOperatingSystemsReader is a Reader for the FindOperatingSystems structure.
type FindOperatingSystemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindOperatingSystemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindOperatingSystemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewFindOperatingSystemsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFindOperatingSystemsOK creates a FindOperatingSystemsOK with default headers values
func NewFindOperatingSystemsOK() *FindOperatingSystemsOK {
	return &FindOperatingSystemsOK{}
}

/* FindOperatingSystemsOK describes a response with status code 200, with default header values.

ok
*/
type FindOperatingSystemsOK struct {
	Payload types.OperatingSystemList
}

func (o *FindOperatingSystemsOK) Error() string {
	return fmt.Sprintf("[GET /operating-systems][%d] findOperatingSystemsOK  %+v", 200, o.Payload)
}
func (o *FindOperatingSystemsOK) GetPayload() types.OperatingSystemList {
	return o.Payload
}

func (o *FindOperatingSystemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindOperatingSystemsUnauthorized creates a FindOperatingSystemsUnauthorized with default headers values
func NewFindOperatingSystemsUnauthorized() *FindOperatingSystemsUnauthorized {
	return &FindOperatingSystemsUnauthorized{}
}

/* FindOperatingSystemsUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type FindOperatingSystemsUnauthorized struct {
}

func (o *FindOperatingSystemsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /operating-systems][%d] findOperatingSystemsUnauthorized ", 401)
}

func (o *FindOperatingSystemsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
