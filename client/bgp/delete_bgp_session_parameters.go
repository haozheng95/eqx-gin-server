// Code generated by go-swagger; DO NOT EDIT.

package bgp

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

// NewDeleteBGPSessionParams creates a new DeleteBGPSessionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteBGPSessionParams() *DeleteBGPSessionParams {
	return &DeleteBGPSessionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBGPSessionParamsWithTimeout creates a new DeleteBGPSessionParams object
// with the ability to set a timeout on a request.
func NewDeleteBGPSessionParamsWithTimeout(timeout time.Duration) *DeleteBGPSessionParams {
	return &DeleteBGPSessionParams{
		timeout: timeout,
	}
}

// NewDeleteBGPSessionParamsWithContext creates a new DeleteBGPSessionParams object
// with the ability to set a context for a request.
func NewDeleteBGPSessionParamsWithContext(ctx context.Context) *DeleteBGPSessionParams {
	return &DeleteBGPSessionParams{
		Context: ctx,
	}
}

// NewDeleteBGPSessionParamsWithHTTPClient creates a new DeleteBGPSessionParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteBGPSessionParamsWithHTTPClient(client *http.Client) *DeleteBGPSessionParams {
	return &DeleteBGPSessionParams{
		HTTPClient: client,
	}
}

/* DeleteBGPSessionParams contains all the parameters to send to the API endpoint
   for the delete Bgp session operation.

   Typically these are written to a http.Request.
*/
type DeleteBGPSessionParams struct {

	/* ID.

	   BGP session UUID

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete Bgp session params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBGPSessionParams) WithDefaults() *DeleteBGPSessionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete Bgp session params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBGPSessionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete Bgp session params
func (o *DeleteBGPSessionParams) WithTimeout(timeout time.Duration) *DeleteBGPSessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete Bgp session params
func (o *DeleteBGPSessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete Bgp session params
func (o *DeleteBGPSessionParams) WithContext(ctx context.Context) *DeleteBGPSessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete Bgp session params
func (o *DeleteBGPSessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete Bgp session params
func (o *DeleteBGPSessionParams) WithHTTPClient(client *http.Client) *DeleteBGPSessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete Bgp session params
func (o *DeleteBGPSessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete Bgp session params
func (o *DeleteBGPSessionParams) WithID(id strfmt.UUID) *DeleteBGPSessionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete Bgp session params
func (o *DeleteBGPSessionParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBGPSessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
