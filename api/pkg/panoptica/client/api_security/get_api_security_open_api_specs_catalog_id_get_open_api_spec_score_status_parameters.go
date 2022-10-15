// Code generated by go-swagger; DO NOT EDIT.

package api_security

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

// NewGetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParams creates a new GetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParams() *GetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParams {
	return &GetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParamsWithTimeout creates a new GetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParams object
// with the ability to set a timeout on a request.
func NewGetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParamsWithTimeout(timeout time.Duration) *GetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParams {
	return &GetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParams{
		timeout: timeout,
	}
}

// NewGetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParamsWithContext creates a new GetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParams object
// with the ability to set a context for a request.
func NewGetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParamsWithContext(ctx context.Context) *GetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParams {
	return &GetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParams{
		Context: ctx,
	}
}

// NewGetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParamsWithHTTPClient creates a new GetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParamsWithHTTPClient(client *http.Client) *GetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParams {
	return &GetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParams{
		HTTPClient: client,
	}
}

/* GetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParams contains all the parameters to send to the API endpoint
   for the get API security open API specs catalog ID get open API spec score status operation.

   Typically these are written to a http.Request.
*/
type GetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParams struct {

	// CatalogID.
	//
	// Format: uuid
	CatalogID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API security open API specs catalog ID get open API spec score status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParams) WithDefaults() *GetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API security open API specs catalog ID get open API spec score status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API security open API specs catalog ID get open API spec score status params
func (o *GetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParams) WithTimeout(timeout time.Duration) *GetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API security open API specs catalog ID get open API spec score status params
func (o *GetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API security open API specs catalog ID get open API spec score status params
func (o *GetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParams) WithContext(ctx context.Context) *GetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API security open API specs catalog ID get open API spec score status params
func (o *GetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API security open API specs catalog ID get open API spec score status params
func (o *GetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParams) WithHTTPClient(client *http.Client) *GetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API security open API specs catalog ID get open API spec score status params
func (o *GetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCatalogID adds the catalogID to the get API security open API specs catalog ID get open API spec score status params
func (o *GetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParams) WithCatalogID(catalogID strfmt.UUID) *GetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParams {
	o.SetCatalogID(catalogID)
	return o
}

// SetCatalogID adds the catalogId to the get API security open API specs catalog ID get open API spec score status params
func (o *GetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParams) SetCatalogID(catalogID strfmt.UUID) {
	o.CatalogID = catalogID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPISecurityOpenAPISpecsCatalogIDGetOpenAPISpecScoreStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param catalogId
	if err := r.SetPathParam("catalogId", o.CatalogID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}