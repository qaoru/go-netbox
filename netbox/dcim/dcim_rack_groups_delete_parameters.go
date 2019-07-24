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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDcimRackGroupsDeleteParams creates a new DcimRackGroupsDeleteParams object
// with the default values initialized.
func NewDcimRackGroupsDeleteParams() *DcimRackGroupsDeleteParams {
	var ()
	return &DcimRackGroupsDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRackGroupsDeleteParamsWithTimeout creates a new DcimRackGroupsDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimRackGroupsDeleteParamsWithTimeout(timeout time.Duration) *DcimRackGroupsDeleteParams {
	var ()
	return &DcimRackGroupsDeleteParams{

		timeout: timeout,
	}
}

// NewDcimRackGroupsDeleteParamsWithContext creates a new DcimRackGroupsDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimRackGroupsDeleteParamsWithContext(ctx context.Context) *DcimRackGroupsDeleteParams {
	var ()
	return &DcimRackGroupsDeleteParams{

		Context: ctx,
	}
}

// NewDcimRackGroupsDeleteParamsWithHTTPClient creates a new DcimRackGroupsDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimRackGroupsDeleteParamsWithHTTPClient(client *http.Client) *DcimRackGroupsDeleteParams {
	var ()
	return &DcimRackGroupsDeleteParams{
		HTTPClient: client,
	}
}

/*DcimRackGroupsDeleteParams contains all the parameters to send to the API endpoint
for the dcim rack groups delete operation typically these are written to a http.Request
*/
type DcimRackGroupsDeleteParams struct {

	/*ID
	  A unique integer value identifying this rack group.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim rack groups delete params
func (o *DcimRackGroupsDeleteParams) WithTimeout(timeout time.Duration) *DcimRackGroupsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rack groups delete params
func (o *DcimRackGroupsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rack groups delete params
func (o *DcimRackGroupsDeleteParams) WithContext(ctx context.Context) *DcimRackGroupsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rack groups delete params
func (o *DcimRackGroupsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rack groups delete params
func (o *DcimRackGroupsDeleteParams) WithHTTPClient(client *http.Client) *DcimRackGroupsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rack groups delete params
func (o *DcimRackGroupsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim rack groups delete params
func (o *DcimRackGroupsDeleteParams) WithID(id int64) *DcimRackGroupsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim rack groups delete params
func (o *DcimRackGroupsDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRackGroupsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
