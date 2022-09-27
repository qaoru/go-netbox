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

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IpamVlanGroupsBulkDeleteReader is a Reader for the IpamVlanGroupsBulkDelete structure.
type IpamVlanGroupsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlanGroupsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamVlanGroupsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamVlanGroupsBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamVlanGroupsBulkDeleteNoContent creates a IpamVlanGroupsBulkDeleteNoContent with default headers values
func NewIpamVlanGroupsBulkDeleteNoContent() *IpamVlanGroupsBulkDeleteNoContent {
	return &IpamVlanGroupsBulkDeleteNoContent{}
}

/*
IpamVlanGroupsBulkDeleteNoContent describes a response with status code 204, with default header values.

IpamVlanGroupsBulkDeleteNoContent ipam vlan groups bulk delete no content
*/
type IpamVlanGroupsBulkDeleteNoContent struct {
}

// IsSuccess returns true when this ipam vlan groups bulk delete no content response has a 2xx status code
func (o *IpamVlanGroupsBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam vlan groups bulk delete no content response has a 3xx status code
func (o *IpamVlanGroupsBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam vlan groups bulk delete no content response has a 4xx status code
func (o *IpamVlanGroupsBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam vlan groups bulk delete no content response has a 5xx status code
func (o *IpamVlanGroupsBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam vlan groups bulk delete no content response a status code equal to that given
func (o *IpamVlanGroupsBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *IpamVlanGroupsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/vlan-groups/][%d] ipamVlanGroupsBulkDeleteNoContent ", 204)
}

func (o *IpamVlanGroupsBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ipam/vlan-groups/][%d] ipamVlanGroupsBulkDeleteNoContent ", 204)
}

func (o *IpamVlanGroupsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIpamVlanGroupsBulkDeleteDefault creates a IpamVlanGroupsBulkDeleteDefault with default headers values
func NewIpamVlanGroupsBulkDeleteDefault(code int) *IpamVlanGroupsBulkDeleteDefault {
	return &IpamVlanGroupsBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
IpamVlanGroupsBulkDeleteDefault describes a response with status code -1, with default header values.

IpamVlanGroupsBulkDeleteDefault ipam vlan groups bulk delete default
*/
type IpamVlanGroupsBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam vlan groups bulk delete default response
func (o *IpamVlanGroupsBulkDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam vlan groups bulk delete default response has a 2xx status code
func (o *IpamVlanGroupsBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam vlan groups bulk delete default response has a 3xx status code
func (o *IpamVlanGroupsBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam vlan groups bulk delete default response has a 4xx status code
func (o *IpamVlanGroupsBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam vlan groups bulk delete default response has a 5xx status code
func (o *IpamVlanGroupsBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam vlan groups bulk delete default response a status code equal to that given
func (o *IpamVlanGroupsBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamVlanGroupsBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /ipam/vlan-groups/][%d] ipam_vlan-groups_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *IpamVlanGroupsBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /ipam/vlan-groups/][%d] ipam_vlan-groups_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *IpamVlanGroupsBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamVlanGroupsBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
