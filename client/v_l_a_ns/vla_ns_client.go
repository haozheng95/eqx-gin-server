// Code generated by go-swagger; DO NOT EDIT.

package v_l_a_ns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new v l a ns API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v l a ns API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteVirtualNetwork(params *DeleteVirtualNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteVirtualNetworkOK, error)

	GetVirtualNetwork(params *GetVirtualNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVirtualNetworkOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteVirtualNetwork deletes a virtual network

  Deletes a virtual network.
*/
func (a *Client) DeleteVirtualNetwork(params *DeleteVirtualNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteVirtualNetworkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteVirtualNetworkParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteVirtualNetwork",
		Method:             "DELETE",
		PathPattern:        "/virtual-networks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteVirtualNetworkReader{formats: a.formats},
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
	success, ok := result.(*DeleteVirtualNetworkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteVirtualNetwork: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetVirtualNetwork gets a virtual network

  Get a virtual network.
*/
func (a *Client) GetVirtualNetwork(params *GetVirtualNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVirtualNetworkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVirtualNetworkParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getVirtualNetwork",
		Method:             "GET",
		PathPattern:        "/virtual-networks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVirtualNetworkReader{formats: a.formats},
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
	success, ok := result.(*GetVirtualNetworkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getVirtualNetwork: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
