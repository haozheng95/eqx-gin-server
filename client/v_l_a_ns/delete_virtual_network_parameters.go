// Code generated by go-swagger; DO NOT EDIT.

package v_l_a_ns

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
)

// NewDeleteVirtualNetworkParams creates a new DeleteVirtualNetworkParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteVirtualNetworkParams() *DeleteVirtualNetworkParams {
	return &DeleteVirtualNetworkParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVirtualNetworkParamsWithTimeout creates a new DeleteVirtualNetworkParams object
// with the ability to set a timeout on a request.
func NewDeleteVirtualNetworkParamsWithTimeout(timeout time.Duration) *DeleteVirtualNetworkParams {
	return &DeleteVirtualNetworkParams{
		timeout: timeout,
	}
}

// NewDeleteVirtualNetworkParamsWithContext creates a new DeleteVirtualNetworkParams object
// with the ability to set a context for a request.
func NewDeleteVirtualNetworkParamsWithContext(ctx context.Context) *DeleteVirtualNetworkParams {
	return &DeleteVirtualNetworkParams{
		Context: ctx,
	}
}

// NewDeleteVirtualNetworkParamsWithHTTPClient creates a new DeleteVirtualNetworkParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteVirtualNetworkParamsWithHTTPClient(client *http.Client) *DeleteVirtualNetworkParams {
	return &DeleteVirtualNetworkParams{
		HTTPClient: client,
	}
}

/* DeleteVirtualNetworkParams contains all the parameters to send to the API endpoint
   for the delete virtual network operation.

   Typically these are written to a http.Request.
*/
type DeleteVirtualNetworkParams struct {

	/* ID.

	   Virtual Network UUID

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete virtual network params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVirtualNetworkParams) WithDefaults() *DeleteVirtualNetworkParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete virtual network params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVirtualNetworkParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete virtual network params
func (o *DeleteVirtualNetworkParams) WithTimeout(timeout time.Duration) *DeleteVirtualNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete virtual network params
func (o *DeleteVirtualNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete virtual network params
func (o *DeleteVirtualNetworkParams) WithContext(ctx context.Context) *DeleteVirtualNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete virtual network params
func (o *DeleteVirtualNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete virtual network params
func (o *DeleteVirtualNetworkParams) WithHTTPClient(client *http.Client) *DeleteVirtualNetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete virtual network params
func (o *DeleteVirtualNetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete virtual network params
func (o *DeleteVirtualNetworkParams) WithID(id strfmt.UUID) *DeleteVirtualNetworkParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete virtual network params
func (o *DeleteVirtualNetworkParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVirtualNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
