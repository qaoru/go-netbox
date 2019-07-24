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

package secrets

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

// NewSecretsSecretsCreateParams creates a new SecretsSecretsCreateParams object
// with the default values initialized.
func NewSecretsSecretsCreateParams() *SecretsSecretsCreateParams {
	var ()
	return &SecretsSecretsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSecretsSecretsCreateParamsWithTimeout creates a new SecretsSecretsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSecretsSecretsCreateParamsWithTimeout(timeout time.Duration) *SecretsSecretsCreateParams {
	var ()
	return &SecretsSecretsCreateParams{

		timeout: timeout,
	}
}

// NewSecretsSecretsCreateParamsWithContext creates a new SecretsSecretsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewSecretsSecretsCreateParamsWithContext(ctx context.Context) *SecretsSecretsCreateParams {
	var ()
	return &SecretsSecretsCreateParams{

		Context: ctx,
	}
}

// NewSecretsSecretsCreateParamsWithHTTPClient creates a new SecretsSecretsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSecretsSecretsCreateParamsWithHTTPClient(client *http.Client) *SecretsSecretsCreateParams {
	var ()
	return &SecretsSecretsCreateParams{
		HTTPClient: client,
	}
}

/*SecretsSecretsCreateParams contains all the parameters to send to the API endpoint
for the secrets secrets create operation typically these are written to a http.Request
*/
type SecretsSecretsCreateParams struct {

	/*Data*/
	Data *models.WritableSecret

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the secrets secrets create params
func (o *SecretsSecretsCreateParams) WithTimeout(timeout time.Duration) *SecretsSecretsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the secrets secrets create params
func (o *SecretsSecretsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the secrets secrets create params
func (o *SecretsSecretsCreateParams) WithContext(ctx context.Context) *SecretsSecretsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the secrets secrets create params
func (o *SecretsSecretsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the secrets secrets create params
func (o *SecretsSecretsCreateParams) WithHTTPClient(client *http.Client) *SecretsSecretsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the secrets secrets create params
func (o *SecretsSecretsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the secrets secrets create params
func (o *SecretsSecretsCreateParams) WithData(data *models.WritableSecret) *SecretsSecretsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the secrets secrets create params
func (o *SecretsSecretsCreateParams) SetData(data *models.WritableSecret) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *SecretsSecretsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
