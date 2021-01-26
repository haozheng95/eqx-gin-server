// Code generated by go-swagger; DO NOT EDIT.

package invitations

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

// NewAcceptInvitationParams creates a new AcceptInvitationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAcceptInvitationParams() *AcceptInvitationParams {
	return &AcceptInvitationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAcceptInvitationParamsWithTimeout creates a new AcceptInvitationParams object
// with the ability to set a timeout on a request.
func NewAcceptInvitationParamsWithTimeout(timeout time.Duration) *AcceptInvitationParams {
	return &AcceptInvitationParams{
		timeout: timeout,
	}
}

// NewAcceptInvitationParamsWithContext creates a new AcceptInvitationParams object
// with the ability to set a context for a request.
func NewAcceptInvitationParamsWithContext(ctx context.Context) *AcceptInvitationParams {
	return &AcceptInvitationParams{
		Context: ctx,
	}
}

// NewAcceptInvitationParamsWithHTTPClient creates a new AcceptInvitationParams object
// with the ability to set a custom HTTPClient for a request.
func NewAcceptInvitationParamsWithHTTPClient(client *http.Client) *AcceptInvitationParams {
	return &AcceptInvitationParams{
		HTTPClient: client,
	}
}

/* AcceptInvitationParams contains all the parameters to send to the API endpoint
   for the accept invitation operation.

   Typically these are written to a http.Request.
*/
type AcceptInvitationParams struct {

	/* ID.

	   Invitation UUID

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the accept invitation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AcceptInvitationParams) WithDefaults() *AcceptInvitationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the accept invitation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AcceptInvitationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the accept invitation params
func (o *AcceptInvitationParams) WithTimeout(timeout time.Duration) *AcceptInvitationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the accept invitation params
func (o *AcceptInvitationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the accept invitation params
func (o *AcceptInvitationParams) WithContext(ctx context.Context) *AcceptInvitationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the accept invitation params
func (o *AcceptInvitationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the accept invitation params
func (o *AcceptInvitationParams) WithHTTPClient(client *http.Client) *AcceptInvitationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the accept invitation params
func (o *AcceptInvitationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the accept invitation params
func (o *AcceptInvitationParams) WithID(id strfmt.UUID) *AcceptInvitationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the accept invitation params
func (o *AcceptInvitationParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AcceptInvitationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
