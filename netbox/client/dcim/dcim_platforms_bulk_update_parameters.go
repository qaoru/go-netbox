// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package dcim

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

	"github.com/fbreckle/go-netbox/netbox/models"
)

// NewDcimPlatformsBulkUpdateParams creates a new DcimPlatformsBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPlatformsBulkUpdateParams() *DcimPlatformsBulkUpdateParams {
	return &DcimPlatformsBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPlatformsBulkUpdateParamsWithTimeout creates a new DcimPlatformsBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimPlatformsBulkUpdateParamsWithTimeout(timeout time.Duration) *DcimPlatformsBulkUpdateParams {
	return &DcimPlatformsBulkUpdateParams{
		timeout: timeout,
	}
}

// NewDcimPlatformsBulkUpdateParamsWithContext creates a new DcimPlatformsBulkUpdateParams object
// with the ability to set a context for a request.
func NewDcimPlatformsBulkUpdateParamsWithContext(ctx context.Context) *DcimPlatformsBulkUpdateParams {
	return &DcimPlatformsBulkUpdateParams{
		Context: ctx,
	}
}

// NewDcimPlatformsBulkUpdateParamsWithHTTPClient creates a new DcimPlatformsBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPlatformsBulkUpdateParamsWithHTTPClient(client *http.Client) *DcimPlatformsBulkUpdateParams {
	return &DcimPlatformsBulkUpdateParams{
		HTTPClient: client,
	}
}

/*
DcimPlatformsBulkUpdateParams contains all the parameters to send to the API endpoint

	for the dcim platforms bulk update operation.

	Typically these are written to a http.Request.
*/
type DcimPlatformsBulkUpdateParams struct {

	// Data.
	Data *models.WritablePlatform

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim platforms bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPlatformsBulkUpdateParams) WithDefaults() *DcimPlatformsBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim platforms bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPlatformsBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim platforms bulk update params
func (o *DcimPlatformsBulkUpdateParams) WithTimeout(timeout time.Duration) *DcimPlatformsBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim platforms bulk update params
func (o *DcimPlatformsBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim platforms bulk update params
func (o *DcimPlatformsBulkUpdateParams) WithContext(ctx context.Context) *DcimPlatformsBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim platforms bulk update params
func (o *DcimPlatformsBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim platforms bulk update params
func (o *DcimPlatformsBulkUpdateParams) WithHTTPClient(client *http.Client) *DcimPlatformsBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim platforms bulk update params
func (o *DcimPlatformsBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim platforms bulk update params
func (o *DcimPlatformsBulkUpdateParams) WithData(data *models.WritablePlatform) *DcimPlatformsBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim platforms bulk update params
func (o *DcimPlatformsBulkUpdateParams) SetData(data *models.WritablePlatform) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPlatformsBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
