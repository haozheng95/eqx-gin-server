// Code generated by go-swagger; DO NOT EDIT.

package spot_market_request

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteSpotMarketRequestReader is a Reader for the DeleteSpotMarketRequest structure.
type DeleteSpotMarketRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSpotMarketRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteSpotMarketRequestNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteSpotMarketRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteSpotMarketRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteSpotMarketRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteSpotMarketRequestNoContent creates a DeleteSpotMarketRequestNoContent with default headers values
func NewDeleteSpotMarketRequestNoContent() *DeleteSpotMarketRequestNoContent {
	return &DeleteSpotMarketRequestNoContent{}
}

/* DeleteSpotMarketRequestNoContent describes a response with status code 204, with default header values.

no content
*/
type DeleteSpotMarketRequestNoContent struct {
}

func (o *DeleteSpotMarketRequestNoContent) Error() string {
	return fmt.Sprintf("[DELETE /spot-market-requests/{id}][%d] deleteSpotMarketRequestNoContent ", 204)
}

func (o *DeleteSpotMarketRequestNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSpotMarketRequestUnauthorized creates a DeleteSpotMarketRequestUnauthorized with default headers values
func NewDeleteSpotMarketRequestUnauthorized() *DeleteSpotMarketRequestUnauthorized {
	return &DeleteSpotMarketRequestUnauthorized{}
}

/* DeleteSpotMarketRequestUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type DeleteSpotMarketRequestUnauthorized struct {
}

func (o *DeleteSpotMarketRequestUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /spot-market-requests/{id}][%d] deleteSpotMarketRequestUnauthorized ", 401)
}

func (o *DeleteSpotMarketRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSpotMarketRequestForbidden creates a DeleteSpotMarketRequestForbidden with default headers values
func NewDeleteSpotMarketRequestForbidden() *DeleteSpotMarketRequestForbidden {
	return &DeleteSpotMarketRequestForbidden{}
}

/* DeleteSpotMarketRequestForbidden describes a response with status code 403, with default header values.

forbidden
*/
type DeleteSpotMarketRequestForbidden struct {
}

func (o *DeleteSpotMarketRequestForbidden) Error() string {
	return fmt.Sprintf("[DELETE /spot-market-requests/{id}][%d] deleteSpotMarketRequestForbidden ", 403)
}

func (o *DeleteSpotMarketRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSpotMarketRequestNotFound creates a DeleteSpotMarketRequestNotFound with default headers values
func NewDeleteSpotMarketRequestNotFound() *DeleteSpotMarketRequestNotFound {
	return &DeleteSpotMarketRequestNotFound{}
}

/* DeleteSpotMarketRequestNotFound describes a response with status code 404, with default header values.

not found
*/
type DeleteSpotMarketRequestNotFound struct {
}

func (o *DeleteSpotMarketRequestNotFound) Error() string {
	return fmt.Sprintf("[DELETE /spot-market-requests/{id}][%d] deleteSpotMarketRequestNotFound ", 404)
}

func (o *DeleteSpotMarketRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
