// Code generated by go-swagger; DO NOT EDIT.

package ports

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

// NewDeleteNativeVLANParams creates a new DeleteNativeVLANParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteNativeVLANParams() *DeleteNativeVLANParams {
	return &DeleteNativeVLANParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNativeVLANParamsWithTimeout creates a new DeleteNativeVLANParams object
// with the ability to set a timeout on a request.
func NewDeleteNativeVLANParamsWithTimeout(timeout time.Duration) *DeleteNativeVLANParams {
	return &DeleteNativeVLANParams{
		timeout: timeout,
	}
}

// NewDeleteNativeVLANParamsWithContext creates a new DeleteNativeVLANParams object
// with the ability to set a context for a request.
func NewDeleteNativeVLANParamsWithContext(ctx context.Context) *DeleteNativeVLANParams {
	return &DeleteNativeVLANParams{
		Context: ctx,
	}
}

// NewDeleteNativeVLANParamsWithHTTPClient creates a new DeleteNativeVLANParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteNativeVLANParamsWithHTTPClient(client *http.Client) *DeleteNativeVLANParams {
	return &DeleteNativeVLANParams{
		HTTPClient: client,
	}
}

/* DeleteNativeVLANParams contains all the parameters to send to the API endpoint
   for the delete native Vlan operation.

   Typically these are written to a http.Request.
*/
type DeleteNativeVLANParams struct {

	/* ID.

	   Port UUID

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete native Vlan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNativeVLANParams) WithDefaults() *DeleteNativeVLANParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete native Vlan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNativeVLANParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete native Vlan params
func (o *DeleteNativeVLANParams) WithTimeout(timeout time.Duration) *DeleteNativeVLANParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete native Vlan params
func (o *DeleteNativeVLANParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete native Vlan params
func (o *DeleteNativeVLANParams) WithContext(ctx context.Context) *DeleteNativeVLANParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete native Vlan params
func (o *DeleteNativeVLANParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete native Vlan params
func (o *DeleteNativeVLANParams) WithHTTPClient(client *http.Client) *DeleteNativeVLANParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete native Vlan params
func (o *DeleteNativeVLANParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete native Vlan params
func (o *DeleteNativeVLANParams) WithID(id strfmt.UUID) *DeleteNativeVLANParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete native Vlan params
func (o *DeleteNativeVLANParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNativeVLANParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
