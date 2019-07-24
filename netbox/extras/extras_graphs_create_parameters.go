// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/digitalocean/go-netbox/netbox/models"
)

// NewExtrasGraphsCreateParams creates a new ExtrasGraphsCreateParams object
// with the default values initialized.
func NewExtrasGraphsCreateParams() *ExtrasGraphsCreateParams {
	var ()
	return &ExtrasGraphsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasGraphsCreateParamsWithTimeout creates a new ExtrasGraphsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExtrasGraphsCreateParamsWithTimeout(timeout time.Duration) *ExtrasGraphsCreateParams {
	var ()
	return &ExtrasGraphsCreateParams{

		timeout: timeout,
	}
}

// NewExtrasGraphsCreateParamsWithContext creates a new ExtrasGraphsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewExtrasGraphsCreateParamsWithContext(ctx context.Context) *ExtrasGraphsCreateParams {
	var ()
	return &ExtrasGraphsCreateParams{

		Context: ctx,
	}
}

// NewExtrasGraphsCreateParamsWithHTTPClient creates a new ExtrasGraphsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExtrasGraphsCreateParamsWithHTTPClient(client *http.Client) *ExtrasGraphsCreateParams {
	var ()
	return &ExtrasGraphsCreateParams{
		HTTPClient: client,
	}
}

/*ExtrasGraphsCreateParams contains all the parameters to send to the API endpoint
for the extras graphs create operation typically these are written to a http.Request
*/
type ExtrasGraphsCreateParams struct {

	/*Data*/
	Data *models.WritableGraph

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the extras graphs create params
func (o *ExtrasGraphsCreateParams) WithTimeout(timeout time.Duration) *ExtrasGraphsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras graphs create params
func (o *ExtrasGraphsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras graphs create params
func (o *ExtrasGraphsCreateParams) WithContext(ctx context.Context) *ExtrasGraphsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras graphs create params
func (o *ExtrasGraphsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras graphs create params
func (o *ExtrasGraphsCreateParams) WithHTTPClient(client *http.Client) *ExtrasGraphsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras graphs create params
func (o *ExtrasGraphsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras graphs create params
func (o *ExtrasGraphsCreateParams) WithData(data *models.WritableGraph) *ExtrasGraphsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras graphs create params
func (o *ExtrasGraphsCreateParams) SetData(data *models.WritableGraph) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasGraphsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
