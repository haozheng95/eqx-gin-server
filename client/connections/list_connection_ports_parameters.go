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
)

// NewListConnectionPortsParams creates a new ListConnectionPortsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListConnectionPortsParams() *ListConnectionPortsParams {
	return &ListConnectionPortsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListConnectionPortsParamsWithTimeout creates a new ListConnectionPortsParams object
// with the ability to set a timeout on a request.
func NewListConnectionPortsParamsWithTimeout(timeout time.Duration) *ListConnectionPortsParams {
	return &ListConnectionPortsParams{
		timeout: timeout,
	}
}

// NewListConnectionPortsParamsWithContext creates a new ListConnectionPortsParams object
// with the ability to set a context for a request.
func NewListConnectionPortsParamsWithContext(ctx context.Context) *ListConnectionPortsParams {
	return &ListConnectionPortsParams{
		Context: ctx,
	}
}

// NewListConnectionPortsParamsWithHTTPClient creates a new ListConnectionPortsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListConnectionPortsParamsWithHTTPClient(client *http.Client) *ListConnectionPortsParams {
	return &ListConnectionPortsParams{
		HTTPClient: client,
	}
}

/* ListConnectionPortsParams contains all the parameters to send to the API endpoint
   for the list connection ports operation.

   Typically these are written to a http.Request.
*/
type ListConnectionPortsParams struct {

	/* ConnectionID.

	   UUID of the connection

	   Format: uuid
	*/
	ConnectionID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list connection ports params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListConnectionPortsParams) WithDefaults() *ListConnectionPortsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list connection ports params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListConnectionPortsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list connection ports params
func (o *ListConnectionPortsParams) WithTimeout(timeout time.Duration) *ListConnectionPortsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list connection ports params
func (o *ListConnectionPortsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list connection ports params
func (o *ListConnectionPortsParams) WithContext(ctx context.Context) *ListConnectionPortsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list connection ports params
func (o *ListConnectionPortsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list connection ports params
func (o *ListConnectionPortsParams) WithHTTPClient(client *http.Client) *ListConnectionPortsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list connection ports params
func (o *ListConnectionPortsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnectionID adds the connectionID to the list connection ports params
func (o *ListConnectionPortsParams) WithConnectionID(connectionID strfmt.UUID) *ListConnectionPortsParams {
	o.SetConnectionID(connectionID)
	return o
}

// SetConnectionID adds the connectionId to the list connection ports params
func (o *ListConnectionPortsParams) SetConnectionID(connectionID strfmt.UUID) {
	o.ConnectionID = connectionID
}

// WriteToRequest writes these params to a swagger request
func (o *ListConnectionPortsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param connection_id
	if err := r.SetPathParam("connection_id", o.ConnectionID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
