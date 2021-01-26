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

// UpdateVolumeReader is a Reader for the UpdateVolume structure.
type UpdateVolumeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateVolumeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateVolumeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateVolumeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateVolumeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateVolumeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateVolumeUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateVolumeOK creates a UpdateVolumeOK with default headers values
func NewUpdateVolumeOK() *UpdateVolumeOK {
	return &UpdateVolumeOK{}
}

/* UpdateVolumeOK describes a response with status code 200, with default header values.

ok
*/
type UpdateVolumeOK struct {
	Payload *types.Volume
}

func (o *UpdateVolumeOK) Error() string {
	return fmt.Sprintf("[PUT /storage/{id}][%d] updateVolumeOK  %+v", 200, o.Payload)
}
func (o *UpdateVolumeOK) GetPayload() *types.Volume {
	return o.Payload
}

func (o *UpdateVolumeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(types.Volume)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVolumeUnauthorized creates a UpdateVolumeUnauthorized with default headers values
func NewUpdateVolumeUnauthorized() *UpdateVolumeUnauthorized {
	return &UpdateVolumeUnauthorized{}
}

/* UpdateVolumeUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type UpdateVolumeUnauthorized struct {
}

func (o *UpdateVolumeUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /storage/{id}][%d] updateVolumeUnauthorized ", 401)
}

func (o *UpdateVolumeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateVolumeForbidden creates a UpdateVolumeForbidden with default headers values
func NewUpdateVolumeForbidden() *UpdateVolumeForbidden {
	return &UpdateVolumeForbidden{}
}

/* UpdateVolumeForbidden describes a response with status code 403, with default header values.

forbidden
*/
type UpdateVolumeForbidden struct {
}

func (o *UpdateVolumeForbidden) Error() string {
	return fmt.Sprintf("[PUT /storage/{id}][%d] updateVolumeForbidden ", 403)
}

func (o *UpdateVolumeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateVolumeNotFound creates a UpdateVolumeNotFound with default headers values
func NewUpdateVolumeNotFound() *UpdateVolumeNotFound {
	return &UpdateVolumeNotFound{}
}

/* UpdateVolumeNotFound describes a response with status code 404, with default header values.

not found
*/
type UpdateVolumeNotFound struct {
}

func (o *UpdateVolumeNotFound) Error() string {
	return fmt.Sprintf("[PUT /storage/{id}][%d] updateVolumeNotFound ", 404)
}

func (o *UpdateVolumeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateVolumeUnprocessableEntity creates a UpdateVolumeUnprocessableEntity with default headers values
func NewUpdateVolumeUnprocessableEntity() *UpdateVolumeUnprocessableEntity {
	return &UpdateVolumeUnprocessableEntity{}
}

/* UpdateVolumeUnprocessableEntity describes a response with status code 422, with default header values.

unprocessable entity
*/
type UpdateVolumeUnprocessableEntity struct {
}

func (o *UpdateVolumeUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /storage/{id}][%d] updateVolumeUnprocessableEntity ", 422)
}

func (o *UpdateVolumeUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
