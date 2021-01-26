// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// NewFindOrganizationsParams creates a new FindOrganizationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindOrganizationsParams() *FindOrganizationsParams {
	return &FindOrganizationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindOrganizationsParamsWithTimeout creates a new FindOrganizationsParams object
// with the ability to set a timeout on a request.
func NewFindOrganizationsParamsWithTimeout(timeout time.Duration) *FindOrganizationsParams {
	return &FindOrganizationsParams{
		timeout: timeout,
	}
}

// NewFindOrganizationsParamsWithContext creates a new FindOrganizationsParams object
// with the ability to set a context for a request.
func NewFindOrganizationsParamsWithContext(ctx context.Context) *FindOrganizationsParams {
	return &FindOrganizationsParams{
		Context: ctx,
	}
}

// NewFindOrganizationsParamsWithHTTPClient creates a new FindOrganizationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindOrganizationsParamsWithHTTPClient(client *http.Client) *FindOrganizationsParams {
	return &FindOrganizationsParams{
		HTTPClient: client,
	}
}

/* FindOrganizationsParams contains all the parameters to send to the API endpoint
   for the find organizations operation.

   Typically these are written to a http.Request.
*/
type FindOrganizationsParams struct {

	/* Include.

	   related attributes to include
	*/
	Include *string

	/* Page.

	   page to display, default to 1, max 100_000
	*/
	Page *int64

	/* PerPage.

	   items per page, default to 10, max 1_000
	*/
	PerPage *int64

	/* Personal.

	   Include, exclude or show only personal organizations.
	*/
	Personal *string

	/* WithoutProjects.

	   Include, exclude or show only organizations that have no projects.
	*/
	WithoutProjects *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find organizations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindOrganizationsParams) WithDefaults() *FindOrganizationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find organizations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindOrganizationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find organizations params
func (o *FindOrganizationsParams) WithTimeout(timeout time.Duration) *FindOrganizationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find organizations params
func (o *FindOrganizationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find organizations params
func (o *FindOrganizationsParams) WithContext(ctx context.Context) *FindOrganizationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find organizations params
func (o *FindOrganizationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find organizations params
func (o *FindOrganizationsParams) WithHTTPClient(client *http.Client) *FindOrganizationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find organizations params
func (o *FindOrganizationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInclude adds the include to the find organizations params
func (o *FindOrganizationsParams) WithInclude(include *string) *FindOrganizationsParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the find organizations params
func (o *FindOrganizationsParams) SetInclude(include *string) {
	o.Include = include
}

// WithPage adds the page to the find organizations params
func (o *FindOrganizationsParams) WithPage(page *int64) *FindOrganizationsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the find organizations params
func (o *FindOrganizationsParams) SetPage(page *int64) {
	o.Page = page
}

// WithPerPage adds the perPage to the find organizations params
func (o *FindOrganizationsParams) WithPerPage(perPage *int64) *FindOrganizationsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the find organizations params
func (o *FindOrganizationsParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WithPersonal adds the personal to the find organizations params
func (o *FindOrganizationsParams) WithPersonal(personal *string) *FindOrganizationsParams {
	o.SetPersonal(personal)
	return o
}

// SetPersonal adds the personal to the find organizations params
func (o *FindOrganizationsParams) SetPersonal(personal *string) {
	o.Personal = personal
}

// WithWithoutProjects adds the withoutProjects to the find organizations params
func (o *FindOrganizationsParams) WithWithoutProjects(withoutProjects *string) *FindOrganizationsParams {
	o.SetWithoutProjects(withoutProjects)
	return o
}

// SetWithoutProjects adds the withoutProjects to the find organizations params
func (o *FindOrganizationsParams) SetWithoutProjects(withoutProjects *string) {
	o.WithoutProjects = withoutProjects
}

// WriteToRequest writes these params to a swagger request
func (o *FindOrganizationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage int64

		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt64(qrPerPage)
		if qPerPage != "" {

			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
				return err
			}
		}
	}

	if o.Personal != nil {

		// query param personal
		var qrPersonal string

		if o.Personal != nil {
			qrPersonal = *o.Personal
		}
		qPersonal := qrPersonal
		if qPersonal != "" {

			if err := r.SetQueryParam("personal", qPersonal); err != nil {
				return err
			}
		}
	}

	if o.WithoutProjects != nil {

		// query param without_projects
		var qrWithoutProjects string

		if o.WithoutProjects != nil {
			qrWithoutProjects = *o.WithoutProjects
		}
		qWithoutProjects := qrWithoutProjects
		if qWithoutProjects != "" {

			if err := r.SetQueryParam("without_projects", qWithoutProjects); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
