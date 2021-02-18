// Code generated by go-swagger; DO NOT EDIT.

package plans

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new plans API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for plans API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	FindPlans(params *FindPlansParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindPlansOK, error)

	FindPlansByOrganization(params *FindPlansByOrganizationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindPlansByOrganizationOK, error)

	FindPlansByProject(params *FindPlansByProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindPlansByProjectOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  FindPlans retrieves all plans

  Provides a listing of available plans to provision your device on.
*/
func (a *Client) FindPlans(params *FindPlansParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindPlansOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindPlansParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findPlans",
		Method:             "GET",
		PathPattern:        "/plans",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindPlansReader{formats: a.formats},
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
	success, ok := result.(*FindPlansOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findPlans: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  FindPlansByOrganization retrieves all plans visible by the organization

  Returns a listing of available plans for the given organization
*/
func (a *Client) FindPlansByOrganization(params *FindPlansByOrganizationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindPlansByOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindPlansByOrganizationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findPlansByOrganization",
		Method:             "GET",
		PathPattern:        "/organizations/{id}/plans",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindPlansByOrganizationReader{formats: a.formats},
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
	success, ok := result.(*FindPlansByOrganizationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findPlansByOrganization: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  FindPlansByProject retrieves all plans visible by the project

  Returns a listing of available plans for the given project
*/
func (a *Client) FindPlansByProject(params *FindPlansByProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindPlansByProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindPlansByProjectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findPlansByProject",
		Method:             "GET",
		PathPattern:        "/projects/{id}/plans",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindPlansByProjectReader{formats: a.formats},
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
	success, ok := result.(*FindPlansByProjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findPlansByProject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
