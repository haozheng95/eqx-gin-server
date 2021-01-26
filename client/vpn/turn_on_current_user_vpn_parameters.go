// Code generated by go-swagger; DO NOT EDIT.

package vpn

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

// NewTurnOnCurrentUserVPNParams creates a new TurnOnCurrentUserVPNParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTurnOnCurrentUserVPNParams() *TurnOnCurrentUserVPNParams {
	return &TurnOnCurrentUserVPNParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTurnOnCurrentUserVPNParamsWithTimeout creates a new TurnOnCurrentUserVPNParams object
// with the ability to set a timeout on a request.
func NewTurnOnCurrentUserVPNParamsWithTimeout(timeout time.Duration) *TurnOnCurrentUserVPNParams {
	return &TurnOnCurrentUserVPNParams{
		timeout: timeout,
	}
}

// NewTurnOnCurrentUserVPNParamsWithContext creates a new TurnOnCurrentUserVPNParams object
// with the ability to set a context for a request.
func NewTurnOnCurrentUserVPNParamsWithContext(ctx context.Context) *TurnOnCurrentUserVPNParams {
	return &TurnOnCurrentUserVPNParams{
		Context: ctx,
	}
}

// NewTurnOnCurrentUserVPNParamsWithHTTPClient creates a new TurnOnCurrentUserVPNParams object
// with the ability to set a custom HTTPClient for a request.
func NewTurnOnCurrentUserVPNParamsWithHTTPClient(client *http.Client) *TurnOnCurrentUserVPNParams {
	return &TurnOnCurrentUserVPNParams{
		HTTPClient: client,
	}
}

/* TurnOnCurrentUserVPNParams contains all the parameters to send to the API endpoint
   for the turn on current user Vpn operation.

   Typically these are written to a http.Request.
*/
type TurnOnCurrentUserVPNParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the turn on current user Vpn params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TurnOnCurrentUserVPNParams) WithDefaults() *TurnOnCurrentUserVPNParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the turn on current user Vpn params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TurnOnCurrentUserVPNParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the turn on current user Vpn params
func (o *TurnOnCurrentUserVPNParams) WithTimeout(timeout time.Duration) *TurnOnCurrentUserVPNParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the turn on current user Vpn params
func (o *TurnOnCurrentUserVPNParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the turn on current user Vpn params
func (o *TurnOnCurrentUserVPNParams) WithContext(ctx context.Context) *TurnOnCurrentUserVPNParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the turn on current user Vpn params
func (o *TurnOnCurrentUserVPNParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the turn on current user Vpn params
func (o *TurnOnCurrentUserVPNParams) WithHTTPClient(client *http.Client) *TurnOnCurrentUserVPNParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the turn on current user Vpn params
func (o *TurnOnCurrentUserVPNParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *TurnOnCurrentUserVPNParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
