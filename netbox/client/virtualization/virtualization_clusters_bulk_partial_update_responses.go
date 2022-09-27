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

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// VirtualizationClustersBulkPartialUpdateReader is a Reader for the VirtualizationClustersBulkPartialUpdate structure.
type VirtualizationClustersBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClustersBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationClustersBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationClustersBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationClustersBulkPartialUpdateOK creates a VirtualizationClustersBulkPartialUpdateOK with default headers values
func NewVirtualizationClustersBulkPartialUpdateOK() *VirtualizationClustersBulkPartialUpdateOK {
	return &VirtualizationClustersBulkPartialUpdateOK{}
}

/*
VirtualizationClustersBulkPartialUpdateOK describes a response with status code 200, with default header values.

VirtualizationClustersBulkPartialUpdateOK virtualization clusters bulk partial update o k
*/
type VirtualizationClustersBulkPartialUpdateOK struct {
	Payload *models.Cluster
}

// IsSuccess returns true when this virtualization clusters bulk partial update o k response has a 2xx status code
func (o *VirtualizationClustersBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization clusters bulk partial update o k response has a 3xx status code
func (o *VirtualizationClustersBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization clusters bulk partial update o k response has a 4xx status code
func (o *VirtualizationClustersBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization clusters bulk partial update o k response has a 5xx status code
func (o *VirtualizationClustersBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization clusters bulk partial update o k response a status code equal to that given
func (o *VirtualizationClustersBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *VirtualizationClustersBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /virtualization/clusters/][%d] virtualizationClustersBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationClustersBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /virtualization/clusters/][%d] virtualizationClustersBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationClustersBulkPartialUpdateOK) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *VirtualizationClustersBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationClustersBulkPartialUpdateDefault creates a VirtualizationClustersBulkPartialUpdateDefault with default headers values
func NewVirtualizationClustersBulkPartialUpdateDefault(code int) *VirtualizationClustersBulkPartialUpdateDefault {
	return &VirtualizationClustersBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
VirtualizationClustersBulkPartialUpdateDefault describes a response with status code -1, with default header values.

VirtualizationClustersBulkPartialUpdateDefault virtualization clusters bulk partial update default
*/
type VirtualizationClustersBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the virtualization clusters bulk partial update default response
func (o *VirtualizationClustersBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this virtualization clusters bulk partial update default response has a 2xx status code
func (o *VirtualizationClustersBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this virtualization clusters bulk partial update default response has a 3xx status code
func (o *VirtualizationClustersBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this virtualization clusters bulk partial update default response has a 4xx status code
func (o *VirtualizationClustersBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this virtualization clusters bulk partial update default response has a 5xx status code
func (o *VirtualizationClustersBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this virtualization clusters bulk partial update default response a status code equal to that given
func (o *VirtualizationClustersBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *VirtualizationClustersBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /virtualization/clusters/][%d] virtualization_clusters_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationClustersBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /virtualization/clusters/][%d] virtualization_clusters_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationClustersBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationClustersBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
