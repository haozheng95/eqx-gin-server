// Code generated by go-swagger; DO NOT EDIT.

package plans

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

// NewFindPlansByProjectParams creates a new FindPlansByProjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindPlansByProjectParams() *FindPlansByProjectParams {
	return &FindPlansByProjectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindPlansByProjectParamsWithTimeout creates a new FindPlansByProjectParams object
// with the ability to set a timeout on a request.
func NewFindPlansByProjectParamsWithTimeout(timeout time.Duration) *FindPlansByProjectParams {
	return &FindPlansByProjectParams{
		timeout: timeout,
	}
}

// NewFindPlansByProjectParamsWithContext creates a new FindPlansByProjectParams object
// with the ability to set a context for a request.
func NewFindPlansByProjectParamsWithContext(ctx context.Context) *FindPlansByProjectParams {
	return &FindPlansByProjectParams{
		Context: ctx,
	}
}

// NewFindPlansByProjectParamsWithHTTPClient creates a new FindPlansByProjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindPlansByProjectParamsWithHTTPClient(client *http.Client) *FindPlansByProjectParams {
	return &FindPlansByProjectParams{
		HTTPClient: client,
	}
}

/* FindPlansByProjectParams contains all the parameters to send to the API endpoint
   for the find plans by project operation.

   Typically these are written to a http.Request.
*/
type FindPlansByProjectParams struct {

	/* Exclude.

	   Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects.
	*/
	Exclude []string

	/* ID.

	   Project UUID

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

// WithDefaults hydrates default values in the find plans by project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindPlansByProjectParams) WithDefaults() *FindPlansByProjectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find plans by project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindPlansByProjectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find plans by project params
func (o *FindPlansByProjectParams) WithTimeout(timeout time.Duration) *FindPlansByProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find plans by project params
func (o *FindPlansByProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find plans by project params
func (o *FindPlansByProjectParams) WithContext(ctx context.Context) *FindPlansByProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find plans by project params
func (o *FindPlansByProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find plans by project params
func (o *FindPlansByProjectParams) WithHTTPClient(client *http.Client) *FindPlansByProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find plans by project params
func (o *FindPlansByProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExclude adds the exclude to the find plans by project params
func (o *FindPlansByProjectParams) WithExclude(exclude []string) *FindPlansByProjectParams {
	o.SetExclude(exclude)
	return o
}

// SetExclude adds the exclude to the find plans by project params
func (o *FindPlansByProjectParams) SetExclude(exclude []string) {
	o.Exclude = exclude
}

// WithID adds the id to the find plans by project params
func (o *FindPlansByProjectParams) WithID(id strfmt.UUID) *FindPlansByProjectParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find plans by project params
func (o *FindPlansByProjectParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithInclude adds the include to the find plans by project params
func (o *FindPlansByProjectParams) WithInclude(include []string) *FindPlansByProjectParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the find plans by project params
func (o *FindPlansByProjectParams) SetInclude(include []string) {
	o.Include = include
}

// WriteToRequest writes these params to a swagger request
func (o *FindPlansByProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamFindPlansByProject binds the parameter exclude
func (o *FindPlansByProjectParams) bindParamExclude(formats strfmt.Registry) []string {
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

// bindParamFindPlansByProject binds the parameter include
func (o *FindPlansByProjectParams) bindParamInclude(formats strfmt.Registry) []string {
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
