// Code generated by go-swagger; DO NOT EDIT.

package projects

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

// NewFindProjectSSHKeysParams creates a new FindProjectSSHKeysParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindProjectSSHKeysParams() *FindProjectSSHKeysParams {
	return &FindProjectSSHKeysParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindProjectSSHKeysParamsWithTimeout creates a new FindProjectSSHKeysParams object
// with the ability to set a timeout on a request.
func NewFindProjectSSHKeysParamsWithTimeout(timeout time.Duration) *FindProjectSSHKeysParams {
	return &FindProjectSSHKeysParams{
		timeout: timeout,
	}
}

// NewFindProjectSSHKeysParamsWithContext creates a new FindProjectSSHKeysParams object
// with the ability to set a context for a request.
func NewFindProjectSSHKeysParamsWithContext(ctx context.Context) *FindProjectSSHKeysParams {
	return &FindProjectSSHKeysParams{
		Context: ctx,
	}
}

// NewFindProjectSSHKeysParamsWithHTTPClient creates a new FindProjectSSHKeysParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindProjectSSHKeysParamsWithHTTPClient(client *http.Client) *FindProjectSSHKeysParams {
	return &FindProjectSSHKeysParams{
		HTTPClient: client,
	}
}

/* FindProjectSSHKeysParams contains all the parameters to send to the API endpoint
   for the find project SSH keys operation.

   Typically these are written to a http.Request.
*/
type FindProjectSSHKeysParams struct {

	/* SearchString.

	   Search by key, label, or fingerprint
	*/
	SearchString *string

	/* ID.

	   Project UUID

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

// WithDefaults hydrates default values in the find project SSH keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindProjectSSHKeysParams) WithDefaults() *FindProjectSSHKeysParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find project SSH keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindProjectSSHKeysParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find project SSH keys params
func (o *FindProjectSSHKeysParams) WithTimeout(timeout time.Duration) *FindProjectSSHKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find project SSH keys params
func (o *FindProjectSSHKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find project SSH keys params
func (o *FindProjectSSHKeysParams) WithContext(ctx context.Context) *FindProjectSSHKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find project SSH keys params
func (o *FindProjectSSHKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find project SSH keys params
func (o *FindProjectSSHKeysParams) WithHTTPClient(client *http.Client) *FindProjectSSHKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find project SSH keys params
func (o *FindProjectSSHKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSearchString adds the searchString to the find project SSH keys params
func (o *FindProjectSSHKeysParams) WithSearchString(searchString *string) *FindProjectSSHKeysParams {
	o.SetSearchString(searchString)
	return o
}

// SetSearchString adds the searchString to the find project SSH keys params
func (o *FindProjectSSHKeysParams) SetSearchString(searchString *string) {
	o.SearchString = searchString
}

// WithID adds the id to the find project SSH keys params
func (o *FindProjectSSHKeysParams) WithID(id strfmt.UUID) *FindProjectSSHKeysParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find project SSH keys params
func (o *FindProjectSSHKeysParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithInclude adds the include to the find project SSH keys params
func (o *FindProjectSSHKeysParams) WithInclude(include *string) *FindProjectSSHKeysParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the find project SSH keys params
func (o *FindProjectSSHKeysParams) SetInclude(include *string) {
	o.Include = include
}

// WriteToRequest writes these params to a swagger request
func (o *FindProjectSSHKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SearchString != nil {

		// query param Search string
		var qrSearchString string

		if o.SearchString != nil {
			qrSearchString = *o.SearchString
		}
		qSearchString := qrSearchString
		if qSearchString != "" {

			if err := r.SetQueryParam("Search string", qSearchString); err != nil {
				return err
			}
		}
	}

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
