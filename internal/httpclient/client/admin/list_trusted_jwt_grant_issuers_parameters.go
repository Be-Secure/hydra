// Code generated by go-swagger; DO NOT EDIT.

package admin

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

// NewListTrustedJwtGrantIssuersParams creates a new ListTrustedJwtGrantIssuersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListTrustedJwtGrantIssuersParams() *ListTrustedJwtGrantIssuersParams {
	return &ListTrustedJwtGrantIssuersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListTrustedJwtGrantIssuersParamsWithTimeout creates a new ListTrustedJwtGrantIssuersParams object
// with the ability to set a timeout on a request.
func NewListTrustedJwtGrantIssuersParamsWithTimeout(timeout time.Duration) *ListTrustedJwtGrantIssuersParams {
	return &ListTrustedJwtGrantIssuersParams{
		timeout: timeout,
	}
}

// NewListTrustedJwtGrantIssuersParamsWithContext creates a new ListTrustedJwtGrantIssuersParams object
// with the ability to set a context for a request.
func NewListTrustedJwtGrantIssuersParamsWithContext(ctx context.Context) *ListTrustedJwtGrantIssuersParams {
	return &ListTrustedJwtGrantIssuersParams{
		Context: ctx,
	}
}

// NewListTrustedJwtGrantIssuersParamsWithHTTPClient creates a new ListTrustedJwtGrantIssuersParams object
// with the ability to set a custom HTTPClient for a request.
func NewListTrustedJwtGrantIssuersParamsWithHTTPClient(client *http.Client) *ListTrustedJwtGrantIssuersParams {
	return &ListTrustedJwtGrantIssuersParams{
		HTTPClient: client,
	}
}

/* ListTrustedJwtGrantIssuersParams contains all the parameters to send to the API endpoint
   for the list trusted jwt grant issuers operation.

   Typically these are written to a http.Request.
*/
type ListTrustedJwtGrantIssuersParams struct {

	/* Issuer.

	   If optional "issuer" is supplied, only jwt-bearer grants with this issuer will be returned.
	*/
	Issuer *string

	/* Limit.

	   The maximum amount of policies returned, upper bound is 500 policies

	   Format: int64
	*/
	Limit *int64

	/* Offset.

	   The offset from where to start looking.

	   Format: int64
	*/
	Offset *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list trusted jwt grant issuers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListTrustedJwtGrantIssuersParams) WithDefaults() *ListTrustedJwtGrantIssuersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list trusted jwt grant issuers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListTrustedJwtGrantIssuersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list trusted jwt grant issuers params
func (o *ListTrustedJwtGrantIssuersParams) WithTimeout(timeout time.Duration) *ListTrustedJwtGrantIssuersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list trusted jwt grant issuers params
func (o *ListTrustedJwtGrantIssuersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list trusted jwt grant issuers params
func (o *ListTrustedJwtGrantIssuersParams) WithContext(ctx context.Context) *ListTrustedJwtGrantIssuersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list trusted jwt grant issuers params
func (o *ListTrustedJwtGrantIssuersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list trusted jwt grant issuers params
func (o *ListTrustedJwtGrantIssuersParams) WithHTTPClient(client *http.Client) *ListTrustedJwtGrantIssuersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list trusted jwt grant issuers params
func (o *ListTrustedJwtGrantIssuersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIssuer adds the issuer to the list trusted jwt grant issuers params
func (o *ListTrustedJwtGrantIssuersParams) WithIssuer(issuer *string) *ListTrustedJwtGrantIssuersParams {
	o.SetIssuer(issuer)
	return o
}

// SetIssuer adds the issuer to the list trusted jwt grant issuers params
func (o *ListTrustedJwtGrantIssuersParams) SetIssuer(issuer *string) {
	o.Issuer = issuer
}

// WithLimit adds the limit to the list trusted jwt grant issuers params
func (o *ListTrustedJwtGrantIssuersParams) WithLimit(limit *int64) *ListTrustedJwtGrantIssuersParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list trusted jwt grant issuers params
func (o *ListTrustedJwtGrantIssuersParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the list trusted jwt grant issuers params
func (o *ListTrustedJwtGrantIssuersParams) WithOffset(offset *int64) *ListTrustedJwtGrantIssuersParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the list trusted jwt grant issuers params
func (o *ListTrustedJwtGrantIssuersParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *ListTrustedJwtGrantIssuersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Issuer != nil {

		// query param issuer
		var qrIssuer string

		if o.Issuer != nil {
			qrIssuer = *o.Issuer
		}
		qIssuer := qrIssuer
		if qIssuer != "" {

			if err := r.SetQueryParam("issuer", qIssuer); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
