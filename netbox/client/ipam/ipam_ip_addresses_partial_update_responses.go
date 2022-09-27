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

// IpamIPAddressesPartialUpdateReader is a Reader for the IpamIPAddressesPartialUpdate structure.
type IpamIPAddressesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamIPAddressesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamIPAddressesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamIPAddressesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamIPAddressesPartialUpdateOK creates a IpamIPAddressesPartialUpdateOK with default headers values
func NewIpamIPAddressesPartialUpdateOK() *IpamIPAddressesPartialUpdateOK {
	return &IpamIPAddressesPartialUpdateOK{}
}

/*
IpamIPAddressesPartialUpdateOK describes a response with status code 200, with default header values.

IpamIPAddressesPartialUpdateOK ipam Ip addresses partial update o k
*/
type IpamIPAddressesPartialUpdateOK struct {
	Payload *models.IPAddress
}

// IsSuccess returns true when this ipam Ip addresses partial update o k response has a 2xx status code
func (o *IpamIPAddressesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam Ip addresses partial update o k response has a 3xx status code
func (o *IpamIPAddressesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam Ip addresses partial update o k response has a 4xx status code
func (o *IpamIPAddressesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam Ip addresses partial update o k response has a 5xx status code
func (o *IpamIPAddressesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam Ip addresses partial update o k response a status code equal to that given
func (o *IpamIPAddressesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *IpamIPAddressesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/ip-addresses/{id}/][%d] ipamIpAddressesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamIPAddressesPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /ipam/ip-addresses/{id}/][%d] ipamIpAddressesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamIPAddressesPartialUpdateOK) GetPayload() *models.IPAddress {
	return o.Payload
}

func (o *IpamIPAddressesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPAddress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamIPAddressesPartialUpdateDefault creates a IpamIPAddressesPartialUpdateDefault with default headers values
func NewIpamIPAddressesPartialUpdateDefault(code int) *IpamIPAddressesPartialUpdateDefault {
	return &IpamIPAddressesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamIPAddressesPartialUpdateDefault describes a response with status code -1, with default header values.

IpamIPAddressesPartialUpdateDefault ipam ip addresses partial update default
*/
type IpamIPAddressesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam ip addresses partial update default response
func (o *IpamIPAddressesPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam ip addresses partial update default response has a 2xx status code
func (o *IpamIPAddressesPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam ip addresses partial update default response has a 3xx status code
func (o *IpamIPAddressesPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam ip addresses partial update default response has a 4xx status code
func (o *IpamIPAddressesPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam ip addresses partial update default response has a 5xx status code
func (o *IpamIPAddressesPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam ip addresses partial update default response a status code equal to that given
func (o *IpamIPAddressesPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamIPAddressesPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /ipam/ip-addresses/{id}/][%d] ipam_ip-addresses_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamIPAddressesPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /ipam/ip-addresses/{id}/][%d] ipam_ip-addresses_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamIPAddressesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamIPAddressesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
