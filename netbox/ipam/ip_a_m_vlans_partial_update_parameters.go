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

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/digitalocean/go-netbox/netbox/models"
)

// NewIPAMVlansPartialUpdateParams creates a new IPAMVlansPartialUpdateParams object
// with the default values initialized.
func NewIPAMVlansPartialUpdateParams() *IPAMVlansPartialUpdateParams {
	var ()
	return &IPAMVlansPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIPAMVlansPartialUpdateParamsWithTimeout creates a new IPAMVlansPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIPAMVlansPartialUpdateParamsWithTimeout(timeout time.Duration) *IPAMVlansPartialUpdateParams {
	var ()
	return &IPAMVlansPartialUpdateParams{

		timeout: timeout,
	}
}

// NewIPAMVlansPartialUpdateParamsWithContext creates a new IPAMVlansPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewIPAMVlansPartialUpdateParamsWithContext(ctx context.Context) *IPAMVlansPartialUpdateParams {
	var ()
	return &IPAMVlansPartialUpdateParams{

		Context: ctx,
	}
}

// NewIPAMVlansPartialUpdateParamsWithHTTPClient creates a new IPAMVlansPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIPAMVlansPartialUpdateParamsWithHTTPClient(client *http.Client) *IPAMVlansPartialUpdateParams {
	var ()
	return &IPAMVlansPartialUpdateParams{
		HTTPClient: client,
	}
}

/*IPAMVlansPartialUpdateParams contains all the parameters to send to the API endpoint
for the ipam vlans partial update operation typically these are written to a http.Request
*/
type IPAMVlansPartialUpdateParams struct {

	/*Data*/
	Data *models.WritableVLAN
	/*ID
	  A unique integer value identifying this VLAN.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam vlans partial update params
func (o *IPAMVlansPartialUpdateParams) WithTimeout(timeout time.Duration) *IPAMVlansPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam vlans partial update params
func (o *IPAMVlansPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam vlans partial update params
func (o *IPAMVlansPartialUpdateParams) WithContext(ctx context.Context) *IPAMVlansPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam vlans partial update params
func (o *IPAMVlansPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam vlans partial update params
func (o *IPAMVlansPartialUpdateParams) WithHTTPClient(client *http.Client) *IPAMVlansPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam vlans partial update params
func (o *IPAMVlansPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam vlans partial update params
func (o *IPAMVlansPartialUpdateParams) WithData(data *models.WritableVLAN) *IPAMVlansPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam vlans partial update params
func (o *IPAMVlansPartialUpdateParams) SetData(data *models.WritableVLAN) {
	o.Data = data
}

// WithID adds the id to the ipam vlans partial update params
func (o *IPAMVlansPartialUpdateParams) WithID(id int64) *IPAMVlansPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam vlans partial update params
func (o *IPAMVlansPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IPAMVlansPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
