// Code generated by go-swagger; DO NOT EDIT.

package transfer_requests

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

// NewFindTransferRequestByIDParams creates a new FindTransferRequestByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindTransferRequestByIDParams() *FindTransferRequestByIDParams {
	return &FindTransferRequestByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindTransferRequestByIDParamsWithTimeout creates a new FindTransferRequestByIDParams object
// with the ability to set a timeout on a request.
func NewFindTransferRequestByIDParamsWithTimeout(timeout time.Duration) *FindTransferRequestByIDParams {
	return &FindTransferRequestByIDParams{
		timeout: timeout,
	}
}

// NewFindTransferRequestByIDParamsWithContext creates a new FindTransferRequestByIDParams object
// with the ability to set a context for a request.
func NewFindTransferRequestByIDParamsWithContext(ctx context.Context) *FindTransferRequestByIDParams {
	return &FindTransferRequestByIDParams{
		Context: ctx,
	}
}

// NewFindTransferRequestByIDParamsWithHTTPClient creates a new FindTransferRequestByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindTransferRequestByIDParamsWithHTTPClient(client *http.Client) *FindTransferRequestByIDParams {
	return &FindTransferRequestByIDParams{
		HTTPClient: client,
	}
}

/* FindTransferRequestByIDParams contains all the parameters to send to the API endpoint
   for the find transfer request by Id operation.

   Typically these are written to a http.Request.
*/
type FindTransferRequestByIDParams struct {

	/* ID.

	   Transfer request UUID

	   Format: uuid
	*/
	ID strfmt.UUID

	/* Include.

	   related attributes to include
	*/
	Include *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find transfer request by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindTransferRequestByIDParams) WithDefaults() *FindTransferRequestByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find transfer request by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindTransferRequestByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find transfer request by Id params
func (o *FindTransferRequestByIDParams) WithTimeout(timeout time.Duration) *FindTransferRequestByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find transfer request by Id params
func (o *FindTransferRequestByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find transfer request by Id params
func (o *FindTransferRequestByIDParams) WithContext(ctx context.Context) *FindTransferRequestByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find transfer request by Id params
func (o *FindTransferRequestByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find transfer request by Id params
func (o *FindTransferRequestByIDParams) WithHTTPClient(client *http.Client) *FindTransferRequestByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find transfer request by Id params
func (o *FindTransferRequestByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find transfer request by Id params
func (o *FindTransferRequestByIDParams) WithID(id strfmt.UUID) *FindTransferRequestByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find transfer request by Id params
func (o *FindTransferRequestByIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithInclude adds the include to the find transfer request by Id params
func (o *FindTransferRequestByIDParams) WithInclude(include *string) *FindTransferRequestByIDParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the find transfer request by Id params
func (o *FindTransferRequestByIDParams) SetInclude(include *string) {
	o.Include = include
}

// WriteToRequest writes these params to a swagger request
func (o *FindTransferRequestByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if o.Include != nil {

		// query param include
		var qrInclude string

		if o.Include != nil {
			qrInclude = *o.Include
		}
		qInclude := qrInclude
		if qInclude != "" {

			if err := r.SetQueryParam("include", qInclude); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
