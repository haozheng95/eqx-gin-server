// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RequestBGPConfigReader is a Reader for the RequestBGPConfig structure.
type RequestBGPConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RequestBGPConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRequestBGPConfigNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRequestBGPConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRequestBGPConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRequestBGPConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewRequestBGPConfigUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRequestBGPConfigNoContent creates a RequestBGPConfigNoContent with default headers values
func NewRequestBGPConfigNoContent() *RequestBGPConfigNoContent {
	return &RequestBGPConfigNoContent{}
}

/* RequestBGPConfigNoContent describes a response with status code 204, with default header values.

no content
*/
type RequestBGPConfigNoContent struct {
}

func (o *RequestBGPConfigNoContent) Error() string {
	return fmt.Sprintf("[POST /projects/{id}/bgp-configs][%d] requestBgpConfigNoContent ", 204)
}

func (o *RequestBGPConfigNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRequestBGPConfigUnauthorized creates a RequestBGPConfigUnauthorized with default headers values
func NewRequestBGPConfigUnauthorized() *RequestBGPConfigUnauthorized {
	return &RequestBGPConfigUnauthorized{}
}

/* RequestBGPConfigUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type RequestBGPConfigUnauthorized struct {
}

func (o *RequestBGPConfigUnauthorized) Error() string {
	return fmt.Sprintf("[POST /projects/{id}/bgp-configs][%d] requestBgpConfigUnauthorized ", 401)
}

func (o *RequestBGPConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRequestBGPConfigForbidden creates a RequestBGPConfigForbidden with default headers values
func NewRequestBGPConfigForbidden() *RequestBGPConfigForbidden {
	return &RequestBGPConfigForbidden{}
}

/* RequestBGPConfigForbidden describes a response with status code 403, with default header values.

forbidden
*/
type RequestBGPConfigForbidden struct {
}

func (o *RequestBGPConfigForbidden) Error() string {
	return fmt.Sprintf("[POST /projects/{id}/bgp-configs][%d] requestBgpConfigForbidden ", 403)
}

func (o *RequestBGPConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRequestBGPConfigNotFound creates a RequestBGPConfigNotFound with default headers values
func NewRequestBGPConfigNotFound() *RequestBGPConfigNotFound {
	return &RequestBGPConfigNotFound{}
}

/* RequestBGPConfigNotFound describes a response with status code 404, with default header values.

not found
*/
type RequestBGPConfigNotFound struct {
}

func (o *RequestBGPConfigNotFound) Error() string {
	return fmt.Sprintf("[POST /projects/{id}/bgp-configs][%d] requestBgpConfigNotFound ", 404)
}

func (o *RequestBGPConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRequestBGPConfigUnprocessableEntity creates a RequestBGPConfigUnprocessableEntity with default headers values
func NewRequestBGPConfigUnprocessableEntity() *RequestBGPConfigUnprocessableEntity {
	return &RequestBGPConfigUnprocessableEntity{}
}

/* RequestBGPConfigUnprocessableEntity describes a response with status code 422, with default header values.

unprocessable entity
*/
type RequestBGPConfigUnprocessableEntity struct {
}

func (o *RequestBGPConfigUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /projects/{id}/bgp-configs][%d] requestBgpConfigUnprocessableEntity ", 422)
}

func (o *RequestBGPConfigUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
