// Code generated by go-swagger; DO NOT EDIT.

package hardware_reservations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new hardware reservations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for hardware reservations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	PostHardwareReservationsIDMove(params *PostHardwareReservationsIDMoveParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostHardwareReservationsIDMoveCreated, error)

	FindHardwareReservationByID(params *FindHardwareReservationByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindHardwareReservationByIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  PostHardwareReservationsIDMove moves a hardware reservation

  Move a hardware reservation to another project
*/
func (a *Client) PostHardwareReservationsIDMove(params *PostHardwareReservationsIDMoveParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostHardwareReservationsIDMoveCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostHardwareReservationsIDMoveParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostHardwareReservationsIDMove",
		Method:             "POST",
		PathPattern:        "/hardware-reservations/{id}/move",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostHardwareReservationsIDMoveReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostHardwareReservationsIDMoveCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostHardwareReservationsIDMove: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  FindHardwareReservationByID retrieves a hardware reservation

  Returns a single hardware reservation
*/
func (a *Client) FindHardwareReservationByID(params *FindHardwareReservationByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindHardwareReservationByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindHardwareReservationByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findHardwareReservationById",
		Method:             "GET",
		PathPattern:        "/hardware-reservations/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindHardwareReservationByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FindHardwareReservationByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findHardwareReservationById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
