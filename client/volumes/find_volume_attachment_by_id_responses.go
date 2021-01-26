// Code generated by go-swagger; DO NOT EDIT.

package volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/t0mk/gometal/types"
)

// FindVolumeAttachmentByIDReader is a Reader for the FindVolumeAttachmentByID structure.
type FindVolumeAttachmentByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindVolumeAttachmentByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindVolumeAttachmentByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewFindVolumeAttachmentByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFindVolumeAttachmentByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFindVolumeAttachmentByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFindVolumeAttachmentByIDOK creates a FindVolumeAttachmentByIDOK with default headers values
func NewFindVolumeAttachmentByIDOK() *FindVolumeAttachmentByIDOK {
	return &FindVolumeAttachmentByIDOK{}
}

/* FindVolumeAttachmentByIDOK describes a response with status code 200, with default header values.

ok
*/
type FindVolumeAttachmentByIDOK struct {
	Payload *types.VolumeAttachment
}

func (o *FindVolumeAttachmentByIDOK) Error() string {
	return fmt.Sprintf("[GET /storage/attachments/{id}][%d] findVolumeAttachmentByIdOK  %+v", 200, o.Payload)
}
func (o *FindVolumeAttachmentByIDOK) GetPayload() *types.VolumeAttachment {
	return o.Payload
}

func (o *FindVolumeAttachmentByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(types.VolumeAttachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindVolumeAttachmentByIDUnauthorized creates a FindVolumeAttachmentByIDUnauthorized with default headers values
func NewFindVolumeAttachmentByIDUnauthorized() *FindVolumeAttachmentByIDUnauthorized {
	return &FindVolumeAttachmentByIDUnauthorized{}
}

/* FindVolumeAttachmentByIDUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type FindVolumeAttachmentByIDUnauthorized struct {
}

func (o *FindVolumeAttachmentByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /storage/attachments/{id}][%d] findVolumeAttachmentByIdUnauthorized ", 401)
}

func (o *FindVolumeAttachmentByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindVolumeAttachmentByIDForbidden creates a FindVolumeAttachmentByIDForbidden with default headers values
func NewFindVolumeAttachmentByIDForbidden() *FindVolumeAttachmentByIDForbidden {
	return &FindVolumeAttachmentByIDForbidden{}
}

/* FindVolumeAttachmentByIDForbidden describes a response with status code 403, with default header values.

forbidden
*/
type FindVolumeAttachmentByIDForbidden struct {
}

func (o *FindVolumeAttachmentByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /storage/attachments/{id}][%d] findVolumeAttachmentByIdForbidden ", 403)
}

func (o *FindVolumeAttachmentByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindVolumeAttachmentByIDNotFound creates a FindVolumeAttachmentByIDNotFound with default headers values
func NewFindVolumeAttachmentByIDNotFound() *FindVolumeAttachmentByIDNotFound {
	return &FindVolumeAttachmentByIDNotFound{}
}

/* FindVolumeAttachmentByIDNotFound describes a response with status code 404, with default header values.

not found
*/
type FindVolumeAttachmentByIDNotFound struct {
}

func (o *FindVolumeAttachmentByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /storage/attachments/{id}][%d] findVolumeAttachmentByIdNotFound ", 404)
}

func (o *FindVolumeAttachmentByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
