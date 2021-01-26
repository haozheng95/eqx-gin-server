// Code generated by go-swagger; DO NOT EDIT.

package invitations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/t0mk/gometal/types"
)

// FindInvitationByIDReader is a Reader for the FindInvitationByID structure.
type FindInvitationByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindInvitationByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindInvitationByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewFindInvitationByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFindInvitationByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFindInvitationByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFindInvitationByIDOK creates a FindInvitationByIDOK with default headers values
func NewFindInvitationByIDOK() *FindInvitationByIDOK {
	return &FindInvitationByIDOK{}
}

/* FindInvitationByIDOK describes a response with status code 200, with default header values.

ok
*/
type FindInvitationByIDOK struct {
	Payload *types.Invitation
}

func (o *FindInvitationByIDOK) Error() string {
	return fmt.Sprintf("[GET /invitations/{id}][%d] findInvitationByIdOK  %+v", 200, o.Payload)
}
func (o *FindInvitationByIDOK) GetPayload() *types.Invitation {
	return o.Payload
}

func (o *FindInvitationByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(types.Invitation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindInvitationByIDUnauthorized creates a FindInvitationByIDUnauthorized with default headers values
func NewFindInvitationByIDUnauthorized() *FindInvitationByIDUnauthorized {
	return &FindInvitationByIDUnauthorized{}
}

/* FindInvitationByIDUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type FindInvitationByIDUnauthorized struct {
}

func (o *FindInvitationByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /invitations/{id}][%d] findInvitationByIdUnauthorized ", 401)
}

func (o *FindInvitationByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindInvitationByIDForbidden creates a FindInvitationByIDForbidden with default headers values
func NewFindInvitationByIDForbidden() *FindInvitationByIDForbidden {
	return &FindInvitationByIDForbidden{}
}

/* FindInvitationByIDForbidden describes a response with status code 403, with default header values.

forbidden
*/
type FindInvitationByIDForbidden struct {
}

func (o *FindInvitationByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /invitations/{id}][%d] findInvitationByIdForbidden ", 403)
}

func (o *FindInvitationByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindInvitationByIDNotFound creates a FindInvitationByIDNotFound with default headers values
func NewFindInvitationByIDNotFound() *FindInvitationByIDNotFound {
	return &FindInvitationByIDNotFound{}
}

/* FindInvitationByIDNotFound describes a response with status code 404, with default header values.

not found
*/
type FindInvitationByIDNotFound struct {
}

func (o *FindInvitationByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /invitations/{id}][%d] findInvitationByIdNotFound ", 404)
}

func (o *FindInvitationByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
