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

	"github.com/fbreckle/go-netbox/netbox/models"
)

// IpamFhrpGroupsCreateReader is a Reader for the IpamFhrpGroupsCreate structure.
type IpamFhrpGroupsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamFhrpGroupsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamFhrpGroupsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamFhrpGroupsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamFhrpGroupsCreateCreated creates a IpamFhrpGroupsCreateCreated with default headers values
func NewIpamFhrpGroupsCreateCreated() *IpamFhrpGroupsCreateCreated {
	return &IpamFhrpGroupsCreateCreated{}
}

/*
IpamFhrpGroupsCreateCreated describes a response with status code 201, with default header values.

IpamFhrpGroupsCreateCreated ipam fhrp groups create created
*/
type IpamFhrpGroupsCreateCreated struct {
	Payload *models.FHRPGroup
}

// IsSuccess returns true when this ipam fhrp groups create created response has a 2xx status code
func (o *IpamFhrpGroupsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam fhrp groups create created response has a 3xx status code
func (o *IpamFhrpGroupsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam fhrp groups create created response has a 4xx status code
func (o *IpamFhrpGroupsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam fhrp groups create created response has a 5xx status code
func (o *IpamFhrpGroupsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam fhrp groups create created response a status code equal to that given
func (o *IpamFhrpGroupsCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *IpamFhrpGroupsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/fhrp-groups/][%d] ipamFhrpGroupsCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamFhrpGroupsCreateCreated) String() string {
	return fmt.Sprintf("[POST /ipam/fhrp-groups/][%d] ipamFhrpGroupsCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamFhrpGroupsCreateCreated) GetPayload() *models.FHRPGroup {
	return o.Payload
}

func (o *IpamFhrpGroupsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FHRPGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamFhrpGroupsCreateDefault creates a IpamFhrpGroupsCreateDefault with default headers values
func NewIpamFhrpGroupsCreateDefault(code int) *IpamFhrpGroupsCreateDefault {
	return &IpamFhrpGroupsCreateDefault{
		_statusCode: code,
	}
}

/*
IpamFhrpGroupsCreateDefault describes a response with status code -1, with default header values.

IpamFhrpGroupsCreateDefault ipam fhrp groups create default
*/
type IpamFhrpGroupsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam fhrp groups create default response
func (o *IpamFhrpGroupsCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam fhrp groups create default response has a 2xx status code
func (o *IpamFhrpGroupsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam fhrp groups create default response has a 3xx status code
func (o *IpamFhrpGroupsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam fhrp groups create default response has a 4xx status code
func (o *IpamFhrpGroupsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam fhrp groups create default response has a 5xx status code
func (o *IpamFhrpGroupsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam fhrp groups create default response a status code equal to that given
func (o *IpamFhrpGroupsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamFhrpGroupsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /ipam/fhrp-groups/][%d] ipam_fhrp-groups_create default  %+v", o._statusCode, o.Payload)
}

func (o *IpamFhrpGroupsCreateDefault) String() string {
	return fmt.Sprintf("[POST /ipam/fhrp-groups/][%d] ipam_fhrp-groups_create default  %+v", o._statusCode, o.Payload)
}

func (o *IpamFhrpGroupsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamFhrpGroupsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
