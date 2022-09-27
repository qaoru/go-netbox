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

// IpamServicesPartialUpdateReader is a Reader for the IpamServicesPartialUpdate structure.
type IpamServicesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamServicesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamServicesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamServicesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamServicesPartialUpdateOK creates a IpamServicesPartialUpdateOK with default headers values
func NewIpamServicesPartialUpdateOK() *IpamServicesPartialUpdateOK {
	return &IpamServicesPartialUpdateOK{}
}

/*
IpamServicesPartialUpdateOK describes a response with status code 200, with default header values.

IpamServicesPartialUpdateOK ipam services partial update o k
*/
type IpamServicesPartialUpdateOK struct {
	Payload *models.Service
}

// IsSuccess returns true when this ipam services partial update o k response has a 2xx status code
func (o *IpamServicesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam services partial update o k response has a 3xx status code
func (o *IpamServicesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam services partial update o k response has a 4xx status code
func (o *IpamServicesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam services partial update o k response has a 5xx status code
func (o *IpamServicesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam services partial update o k response a status code equal to that given
func (o *IpamServicesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *IpamServicesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/services/{id}/][%d] ipamServicesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamServicesPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /ipam/services/{id}/][%d] ipamServicesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamServicesPartialUpdateOK) GetPayload() *models.Service {
	return o.Payload
}

func (o *IpamServicesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Service)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamServicesPartialUpdateDefault creates a IpamServicesPartialUpdateDefault with default headers values
func NewIpamServicesPartialUpdateDefault(code int) *IpamServicesPartialUpdateDefault {
	return &IpamServicesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamServicesPartialUpdateDefault describes a response with status code -1, with default header values.

IpamServicesPartialUpdateDefault ipam services partial update default
*/
type IpamServicesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam services partial update default response
func (o *IpamServicesPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam services partial update default response has a 2xx status code
func (o *IpamServicesPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam services partial update default response has a 3xx status code
func (o *IpamServicesPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam services partial update default response has a 4xx status code
func (o *IpamServicesPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam services partial update default response has a 5xx status code
func (o *IpamServicesPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam services partial update default response a status code equal to that given
func (o *IpamServicesPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamServicesPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /ipam/services/{id}/][%d] ipam_services_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamServicesPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /ipam/services/{id}/][%d] ipam_services_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamServicesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamServicesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
