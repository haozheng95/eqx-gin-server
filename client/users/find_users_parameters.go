// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewFindUsersParams creates a new FindUsersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindUsersParams() *FindUsersParams {
	return &FindUsersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindUsersParamsWithTimeout creates a new FindUsersParams object
// with the ability to set a timeout on a request.
func NewFindUsersParamsWithTimeout(timeout time.Duration) *FindUsersParams {
	return &FindUsersParams{
		timeout: timeout,
	}
}

// NewFindUsersParamsWithContext creates a new FindUsersParams object
// with the ability to set a context for a request.
func NewFindUsersParamsWithContext(ctx context.Context) *FindUsersParams {
	return &FindUsersParams{
		Context: ctx,
	}
}

// NewFindUsersParamsWithHTTPClient creates a new FindUsersParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindUsersParamsWithHTTPClient(client *http.Client) *FindUsersParams {
	return &FindUsersParams{
		HTTPClient: client,
	}
}

/* FindUsersParams contains all the parameters to send to the API endpoint
   for the find users operation.

   Typically these are written to a http.Request.
*/
type FindUsersParams struct {

	/* Exclude.

	   Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects.
	*/
	Exclude []string

	/* Include.

	   Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects.
	*/
	Include []string

	/* Page.

	   Page to return

	   Format: int32
	   Default: 1
	*/
	Page *int32

	/* PerPage.

	   Items returned per page

	   Format: int32
	   Default: 10
	*/
	PerPage *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindUsersParams) WithDefaults() *FindUsersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindUsersParams) SetDefaults() {
	var (
		pageDefault = int32(1)

		perPageDefault = int32(10)
	)

	val := FindUsersParams{
		Page:    &pageDefault,
		PerPage: &perPageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the find users params
func (o *FindUsersParams) WithTimeout(timeout time.Duration) *FindUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find users params
func (o *FindUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find users params
func (o *FindUsersParams) WithContext(ctx context.Context) *FindUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find users params
func (o *FindUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find users params
func (o *FindUsersParams) WithHTTPClient(client *http.Client) *FindUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find users params
func (o *FindUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExclude adds the exclude to the find users params
func (o *FindUsersParams) WithExclude(exclude []string) *FindUsersParams {
	o.SetExclude(exclude)
	return o
}

// SetExclude adds the exclude to the find users params
func (o *FindUsersParams) SetExclude(exclude []string) {
	o.Exclude = exclude
}

// WithInclude adds the include to the find users params
func (o *FindUsersParams) WithInclude(include []string) *FindUsersParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the find users params
func (o *FindUsersParams) SetInclude(include []string) {
	o.Include = include
}

// WithPage adds the page to the find users params
func (o *FindUsersParams) WithPage(page *int32) *FindUsersParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the find users params
func (o *FindUsersParams) SetPage(page *int32) {
	o.Page = page
}

// WithPerPage adds the perPage to the find users params
func (o *FindUsersParams) WithPerPage(perPage *int32) *FindUsersParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the find users params
func (o *FindUsersParams) SetPerPage(perPage *int32) {
	o.PerPage = perPage
}

// WriteToRequest writes these params to a swagger request
func (o *FindUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Include != nil {

		// binding items for include
		joinedInclude := o.bindParamInclude(reg)

		// query array param include
		if err := r.SetQueryParam("include", joinedInclude...); err != nil {
			return err
		}
	}

	if o.Page != nil {

		// query param page
		var qrPage int32

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage int32

		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt32(qrPerPage)
		if qPerPage != "" {

			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamFindUsers binds the parameter exclude
func (o *FindUsersParams) bindParamExclude(formats strfmt.Registry) []string {
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

// bindParamFindUsers binds the parameter include
func (o *FindUsersParams) bindParamInclude(formats strfmt.Registry) []string {
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
