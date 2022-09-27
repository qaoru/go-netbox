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

// IpamRirsUpdateReader is a Reader for the IpamRirsUpdate structure.
type IpamRirsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRirsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamRirsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamRirsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamRirsUpdateOK creates a IpamRirsUpdateOK with default headers values
func NewIpamRirsUpdateOK() *IpamRirsUpdateOK {
	return &IpamRirsUpdateOK{}
}

/*
IpamRirsUpdateOK describes a response with status code 200, with default header values.

IpamRirsUpdateOK ipam rirs update o k
*/
type IpamRirsUpdateOK struct {
	Payload *models.RIR
}

// IsSuccess returns true when this ipam rirs update o k response has a 2xx status code
func (o *IpamRirsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam rirs update o k response has a 3xx status code
func (o *IpamRirsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam rirs update o k response has a 4xx status code
func (o *IpamRirsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam rirs update o k response has a 5xx status code
func (o *IpamRirsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam rirs update o k response a status code equal to that given
func (o *IpamRirsUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *IpamRirsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/rirs/{id}/][%d] ipamRirsUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamRirsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ipam/rirs/{id}/][%d] ipamRirsUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamRirsUpdateOK) GetPayload() *models.RIR {
	return o.Payload
}

func (o *IpamRirsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RIR)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamRirsUpdateDefault creates a IpamRirsUpdateDefault with default headers values
func NewIpamRirsUpdateDefault(code int) *IpamRirsUpdateDefault {
	return &IpamRirsUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamRirsUpdateDefault describes a response with status code -1, with default header values.

IpamRirsUpdateDefault ipam rirs update default
*/
type IpamRirsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam rirs update default response
func (o *IpamRirsUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam rirs update default response has a 2xx status code
func (o *IpamRirsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam rirs update default response has a 3xx status code
func (o *IpamRirsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam rirs update default response has a 4xx status code
func (o *IpamRirsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam rirs update default response has a 5xx status code
func (o *IpamRirsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam rirs update default response a status code equal to that given
func (o *IpamRirsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamRirsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /ipam/rirs/{id}/][%d] ipam_rirs_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamRirsUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /ipam/rirs/{id}/][%d] ipam_rirs_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamRirsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamRirsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
