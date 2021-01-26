// Code generated by go-swagger; DO NOT EDIT.

package ssh_keys

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

// NewFindSSHKeysParams creates a new FindSSHKeysParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindSSHKeysParams() *FindSSHKeysParams {
	return &FindSSHKeysParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindSSHKeysParamsWithTimeout creates a new FindSSHKeysParams object
// with the ability to set a timeout on a request.
func NewFindSSHKeysParamsWithTimeout(timeout time.Duration) *FindSSHKeysParams {
	return &FindSSHKeysParams{
		timeout: timeout,
	}
}

// NewFindSSHKeysParamsWithContext creates a new FindSSHKeysParams object
// with the ability to set a context for a request.
func NewFindSSHKeysParamsWithContext(ctx context.Context) *FindSSHKeysParams {
	return &FindSSHKeysParams{
		Context: ctx,
	}
}

// NewFindSSHKeysParamsWithHTTPClient creates a new FindSSHKeysParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindSSHKeysParamsWithHTTPClient(client *http.Client) *FindSSHKeysParams {
	return &FindSSHKeysParams{
		HTTPClient: client,
	}
}

/* FindSSHKeysParams contains all the parameters to send to the API endpoint
   for the find SSH keys operation.

   Typically these are written to a http.Request.
*/
type FindSSHKeysParams struct {

	/* SearchString.

	   Search by key, label, or fingerprint
	*/
	SearchString *string

	/* Include.

	   related attributes to include
	*/
	Include *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find SSH keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindSSHKeysParams) WithDefaults() *FindSSHKeysParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find SSH keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindSSHKeysParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find SSH keys params
func (o *FindSSHKeysParams) WithTimeout(timeout time.Duration) *FindSSHKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find SSH keys params
func (o *FindSSHKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find SSH keys params
func (o *FindSSHKeysParams) WithContext(ctx context.Context) *FindSSHKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find SSH keys params
func (o *FindSSHKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find SSH keys params
func (o *FindSSHKeysParams) WithHTTPClient(client *http.Client) *FindSSHKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find SSH keys params
func (o *FindSSHKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSearchString adds the searchString to the find SSH keys params
func (o *FindSSHKeysParams) WithSearchString(searchString *string) *FindSSHKeysParams {
	o.SetSearchString(searchString)
	return o
}

// SetSearchString adds the searchString to the find SSH keys params
func (o *FindSSHKeysParams) SetSearchString(searchString *string) {
	o.SearchString = searchString
}

// WithInclude adds the include to the find SSH keys params
func (o *FindSSHKeysParams) WithInclude(include *string) *FindSSHKeysParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the find SSH keys params
func (o *FindSSHKeysParams) SetInclude(include *string) {
	o.Include = include
}

// WriteToRequest writes these params to a swagger request
func (o *FindSSHKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
