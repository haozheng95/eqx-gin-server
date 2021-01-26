// Code generated by go-swagger; DO NOT EDIT.

package connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/t0mk/gometal/types"
)

// NewCreateConnectionPortVirtualCircuitParams creates a new CreateConnectionPortVirtualCircuitParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateConnectionPortVirtualCircuitParams() *CreateConnectionPortVirtualCircuitParams {
	return &CreateConnectionPortVirtualCircuitParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateConnectionPortVirtualCircuitParamsWithTimeout creates a new CreateConnectionPortVirtualCircuitParams object
// with the ability to set a timeout on a request.
func NewCreateConnectionPortVirtualCircuitParamsWithTimeout(timeout time.Duration) *CreateConnectionPortVirtualCircuitParams {
	return &CreateConnectionPortVirtualCircuitParams{
		timeout: timeout,
	}
}

// NewCreateConnectionPortVirtualCircuitParamsWithContext creates a new CreateConnectionPortVirtualCircuitParams object
// with the ability to set a context for a request.
func NewCreateConnectionPortVirtualCircuitParamsWithContext(ctx context.Context) *CreateConnectionPortVirtualCircuitParams {
	return &CreateConnectionPortVirtualCircuitParams{
		Context: ctx,
	}
}

// NewCreateConnectionPortVirtualCircuitParamsWithHTTPClient creates a new CreateConnectionPortVirtualCircuitParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateConnectionPortVirtualCircuitParamsWithHTTPClient(client *http.Client) *CreateConnectionPortVirtualCircuitParams {
	return &CreateConnectionPortVirtualCircuitParams{
		HTTPClient: client,
	}
}

/* CreateConnectionPortVirtualCircuitParams contains all the parameters to send to the API endpoint
   for the create connection port virtual circuit operation.

   Typically these are written to a http.Request.
*/
type CreateConnectionPortVirtualCircuitParams struct {

	/* ConnectionID.

	   UUID of the connection

	   Format: uuid
	*/
	ConnectionID strfmt.UUID

	/* PortID.

	   UUID of the connection port

	   Format: uuid
	*/
	PortID strfmt.UUID

	/* VirtualCircuit.

	   Virtual Circuit details
	*/
	VirtualCircuit *types.VirtualCircuitCreateInput

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create connection port virtual circuit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateConnectionPortVirtualCircuitParams) WithDefaults() *CreateConnectionPortVirtualCircuitParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create connection port virtual circuit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateConnectionPortVirtualCircuitParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create connection port virtual circuit params
func (o *CreateConnectionPortVirtualCircuitParams) WithTimeout(timeout time.Duration) *CreateConnectionPortVirtualCircuitParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create connection port virtual circuit params
func (o *CreateConnectionPortVirtualCircuitParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create connection port virtual circuit params
func (o *CreateConnectionPortVirtualCircuitParams) WithContext(ctx context.Context) *CreateConnectionPortVirtualCircuitParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create connection port virtual circuit params
func (o *CreateConnectionPortVirtualCircuitParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create connection port virtual circuit params
func (o *CreateConnectionPortVirtualCircuitParams) WithHTTPClient(client *http.Client) *CreateConnectionPortVirtualCircuitParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create connection port virtual circuit params
func (o *CreateConnectionPortVirtualCircuitParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnectionID adds the connectionID to the create connection port virtual circuit params
func (o *CreateConnectionPortVirtualCircuitParams) WithConnectionID(connectionID strfmt.UUID) *CreateConnectionPortVirtualCircuitParams {
	o.SetConnectionID(connectionID)
	return o
}

// SetConnectionID adds the connectionId to the create connection port virtual circuit params
func (o *CreateConnectionPortVirtualCircuitParams) SetConnectionID(connectionID strfmt.UUID) {
	o.ConnectionID = connectionID
}

// WithPortID adds the portID to the create connection port virtual circuit params
func (o *CreateConnectionPortVirtualCircuitParams) WithPortID(portID strfmt.UUID) *CreateConnectionPortVirtualCircuitParams {
	o.SetPortID(portID)
	return o
}

// SetPortID adds the portId to the create connection port virtual circuit params
func (o *CreateConnectionPortVirtualCircuitParams) SetPortID(portID strfmt.UUID) {
	o.PortID = portID
}

// WithVirtualCircuit adds the virtualCircuit to the create connection port virtual circuit params
func (o *CreateConnectionPortVirtualCircuitParams) WithVirtualCircuit(virtualCircuit *types.VirtualCircuitCreateInput) *CreateConnectionPortVirtualCircuitParams {
	o.SetVirtualCircuit(virtualCircuit)
	return o
}

// SetVirtualCircuit adds the virtualCircuit to the create connection port virtual circuit params
func (o *CreateConnectionPortVirtualCircuitParams) SetVirtualCircuit(virtualCircuit *types.VirtualCircuitCreateInput) {
	o.VirtualCircuit = virtualCircuit
}

// WriteToRequest writes these params to a swagger request
func (o *CreateConnectionPortVirtualCircuitParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param connection_id
	if err := r.SetPathParam("connection_id", o.ConnectionID.String()); err != nil {
		return err
	}

	// path param port_id
	if err := r.SetPathParam("port_id", o.PortID.String()); err != nil {
		return err
	}
	if o.VirtualCircuit != nil {
		if err := r.SetBodyParam(o.VirtualCircuit); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
