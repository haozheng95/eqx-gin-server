// Code generated by go-swagger; DO NOT EDIT.

package user_verification_tokens

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

// NewCreateValidationRequestParams creates a new CreateValidationRequestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateValidationRequestParams() *CreateValidationRequestParams {
	return &CreateValidationRequestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateValidationRequestParamsWithTimeout creates a new CreateValidationRequestParams object
// with the ability to set a timeout on a request.
func NewCreateValidationRequestParamsWithTimeout(timeout time.Duration) *CreateValidationRequestParams {
	return &CreateValidationRequestParams{
		timeout: timeout,
	}
}

// NewCreateValidationRequestParamsWithContext creates a new CreateValidationRequestParams object
// with the ability to set a context for a request.
func NewCreateValidationRequestParamsWithContext(ctx context.Context) *CreateValidationRequestParams {
	return &CreateValidationRequestParams{
		Context: ctx,
	}
}

// NewCreateValidationRequestParamsWithHTTPClient creates a new CreateValidationRequestParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateValidationRequestParamsWithHTTPClient(client *http.Client) *CreateValidationRequestParams {
	return &CreateValidationRequestParams{
		HTTPClient: client,
	}
}

/* CreateValidationRequestParams contains all the parameters to send to the API endpoint
   for the create validation request operation.

   Typically these are written to a http.Request.
*/
type CreateValidationRequestParams struct {

	/* Login.

	   Email for verification request
	*/
	Login string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create validation request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateValidationRequestParams) WithDefaults() *CreateValidationRequestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create validation request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateValidationRequestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create validation request params
func (o *CreateValidationRequestParams) WithTimeout(timeout time.Duration) *CreateValidationRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create validation request params
func (o *CreateValidationRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create validation request params
func (o *CreateValidationRequestParams) WithContext(ctx context.Context) *CreateValidationRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create validation request params
func (o *CreateValidationRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create validation request params
func (o *CreateValidationRequestParams) WithHTTPClient(client *http.Client) *CreateValidationRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create validation request params
func (o *CreateValidationRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLogin adds the login to the create validation request params
func (o *CreateValidationRequestParams) WithLogin(login string) *CreateValidationRequestParams {
	o.SetLogin(login)
	return o
}

// SetLogin adds the login to the create validation request params
func (o *CreateValidationRequestParams) SetLogin(login string) {
	o.Login = login
}

// WriteToRequest writes these params to a swagger request
func (o *CreateValidationRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param login
	qrLogin := o.Login
	qLogin := qrLogin
	if qLogin != "" {

		if err := r.SetQueryParam("login", qLogin); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
