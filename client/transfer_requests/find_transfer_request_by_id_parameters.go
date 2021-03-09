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
	"github.com/go-openapi/swag"
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

	/* Exclude.

	   Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects.
	*/
	Exclude []string

	/* ID.

	   Transfer request UUID

	   Format: uuid
	*/
	ID strfmt.UUID

	/* Include.

	   Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects.
	*/
	Include []string

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

// WithExclude adds the exclude to the find transfer request by Id params
func (o *FindTransferRequestByIDParams) WithExclude(exclude []string) *FindTransferRequestByIDParams {
	o.SetExclude(exclude)
	return o
}

// SetExclude adds the exclude to the find transfer request by Id params
func (o *FindTransferRequestByIDParams) SetExclude(exclude []string) {
	o.Exclude = exclude
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
func (o *FindTransferRequestByIDParams) WithInclude(include []string) *FindTransferRequestByIDParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the find transfer request by Id params
func (o *FindTransferRequestByIDParams) SetInclude(include []string) {
	o.Include = include
}

// WriteToRequest writes these params to a swagger request
func (o *FindTransferRequestByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Exclude != nil {

		// binding items for exclude
		joinedExclude := o.bindParamExclude(reg)

		// query array param exclude
		if err := r.SetQueryParam("exclude", joinedExclude...); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if o.Include != nil {

		// binding items for include
		joinedInclude := o.bindParamInclude(reg)

		// query array param include
		if err := r.SetQueryParam("include", joinedInclude...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamFindTransferRequestByID binds the parameter exclude
func (o *FindTransferRequestByIDParams) bindParamExclude(formats strfmt.Registry) []string {
	excludeIR := o.Exclude

	var excludeIC []string
	for _, excludeIIR := range excludeIR { // explode []string

		excludeIIV := excludeIIR // string as string
		excludeIC = append(excludeIC, excludeIIV)
	}

	// items.CollectionFormat: "csv"
	excludeIS := swag.JoinByFormat(excludeIC, "csv")

	return excludeIS
}

// bindParamFindTransferRequestByID binds the parameter include
func (o *FindTransferRequestByIDParams) bindParamInclude(formats strfmt.Registry) []string {
	includeIR := o.Include

	var includeIC []string
	for _, includeIIR := range includeIR { // explode []string

		includeIIV := includeIIR // string as string
		includeIC = append(includeIC, includeIIV)
	}

	// items.CollectionFormat: "csv"
	includeIS := swag.JoinByFormat(includeIC, "csv")

	return includeIS
}
