// Code generated by go-swagger; DO NOT EDIT.

package events

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

// NewFindEventByIDParams creates a new FindEventByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindEventByIDParams() *FindEventByIDParams {
	return &FindEventByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindEventByIDParamsWithTimeout creates a new FindEventByIDParams object
// with the ability to set a timeout on a request.
func NewFindEventByIDParamsWithTimeout(timeout time.Duration) *FindEventByIDParams {
	return &FindEventByIDParams{
		timeout: timeout,
	}
}

// NewFindEventByIDParamsWithContext creates a new FindEventByIDParams object
// with the ability to set a context for a request.
func NewFindEventByIDParamsWithContext(ctx context.Context) *FindEventByIDParams {
	return &FindEventByIDParams{
		Context: ctx,
	}
}

// NewFindEventByIDParamsWithHTTPClient creates a new FindEventByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindEventByIDParamsWithHTTPClient(client *http.Client) *FindEventByIDParams {
	return &FindEventByIDParams{
		HTTPClient: client,
	}
}

/* FindEventByIDParams contains all the parameters to send to the API endpoint
   for the find event by Id operation.

   Typically these are written to a http.Request.
*/
type FindEventByIDParams struct {

	/* ID.

	   Event UUID

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

// WithDefaults hydrates default values in the find event by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindEventByIDParams) WithDefaults() *FindEventByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find event by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindEventByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find event by Id params
func (o *FindEventByIDParams) WithTimeout(timeout time.Duration) *FindEventByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find event by Id params
func (o *FindEventByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find event by Id params
func (o *FindEventByIDParams) WithContext(ctx context.Context) *FindEventByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find event by Id params
func (o *FindEventByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find event by Id params
func (o *FindEventByIDParams) WithHTTPClient(client *http.Client) *FindEventByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find event by Id params
func (o *FindEventByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find event by Id params
func (o *FindEventByIDParams) WithID(id strfmt.UUID) *FindEventByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find event by Id params
func (o *FindEventByIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithInclude adds the include to the find event by Id params
func (o *FindEventByIDParams) WithInclude(include *string) *FindEventByIDParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the find event by Id params
func (o *FindEventByIDParams) SetInclude(include *string) {
	o.Include = include
}

// WriteToRequest writes these params to a swagger request
func (o *FindEventByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
