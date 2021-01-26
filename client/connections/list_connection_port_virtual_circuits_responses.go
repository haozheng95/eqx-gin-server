// Code generated by go-swagger; DO NOT EDIT.

package connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/t0mk/gometal/types"
)

// ListConnectionPortVirtualCircuitsReader is a Reader for the ListConnectionPortVirtualCircuits structure.
type ListConnectionPortVirtualCircuitsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListConnectionPortVirtualCircuitsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListConnectionPortVirtualCircuitsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListConnectionPortVirtualCircuitsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListConnectionPortVirtualCircuitsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListConnectionPortVirtualCircuitsOK creates a ListConnectionPortVirtualCircuitsOK with default headers values
func NewListConnectionPortVirtualCircuitsOK() *ListConnectionPortVirtualCircuitsOK {
	return &ListConnectionPortVirtualCircuitsOK{}
}

/* ListConnectionPortVirtualCircuitsOK describes a response with status code 200, with default header values.

ok
*/
type ListConnectionPortVirtualCircuitsOK struct {
	Payload *types.VirtualCircuitList
}

func (o *ListConnectionPortVirtualCircuitsOK) Error() string {
	return fmt.Sprintf("[GET /connections/{connection_id}/ports/{port_id}/virtual-circuits][%d] listConnectionPortVirtualCircuitsOK  %+v", 200, o.Payload)
}
func (o *ListConnectionPortVirtualCircuitsOK) GetPayload() *types.VirtualCircuitList {
	return o.Payload
}

func (o *ListConnectionPortVirtualCircuitsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(types.VirtualCircuitList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListConnectionPortVirtualCircuitsForbidden creates a ListConnectionPortVirtualCircuitsForbidden with default headers values
func NewListConnectionPortVirtualCircuitsForbidden() *ListConnectionPortVirtualCircuitsForbidden {
	return &ListConnectionPortVirtualCircuitsForbidden{}
}

/* ListConnectionPortVirtualCircuitsForbidden describes a response with status code 403, with default header values.

forbidden
*/
type ListConnectionPortVirtualCircuitsForbidden struct {
}

func (o *ListConnectionPortVirtualCircuitsForbidden) Error() string {
	return fmt.Sprintf("[GET /connections/{connection_id}/ports/{port_id}/virtual-circuits][%d] listConnectionPortVirtualCircuitsForbidden ", 403)
}

func (o *ListConnectionPortVirtualCircuitsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListConnectionPortVirtualCircuitsNotFound creates a ListConnectionPortVirtualCircuitsNotFound with default headers values
func NewListConnectionPortVirtualCircuitsNotFound() *ListConnectionPortVirtualCircuitsNotFound {
	return &ListConnectionPortVirtualCircuitsNotFound{}
}

/* ListConnectionPortVirtualCircuitsNotFound describes a response with status code 404, with default header values.

not found
*/
type ListConnectionPortVirtualCircuitsNotFound struct {
}

func (o *ListConnectionPortVirtualCircuitsNotFound) Error() string {
	return fmt.Sprintf("[GET /connections/{connection_id}/ports/{port_id}/virtual-circuits][%d] listConnectionPortVirtualCircuitsNotFound ", 404)
}

func (o *ListConnectionPortVirtualCircuitsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
