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

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// DcimDeviceBaysPartialUpdateReader is a Reader for the DcimDeviceBaysPartialUpdate structure.
type DcimDeviceBaysPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceBaysPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceBaysPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimDeviceBaysPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDeviceBaysPartialUpdateOK creates a DcimDeviceBaysPartialUpdateOK with default headers values
func NewDcimDeviceBaysPartialUpdateOK() *DcimDeviceBaysPartialUpdateOK {
	return &DcimDeviceBaysPartialUpdateOK{}
}

/*
DcimDeviceBaysPartialUpdateOK describes a response with status code 200, with default header values.

DcimDeviceBaysPartialUpdateOK dcim device bays partial update o k
*/
type DcimDeviceBaysPartialUpdateOK struct {
	Payload *models.DeviceBay
}

// IsSuccess returns true when this dcim device bays partial update o k response has a 2xx status code
func (o *DcimDeviceBaysPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim device bays partial update o k response has a 3xx status code
func (o *DcimDeviceBaysPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim device bays partial update o k response has a 4xx status code
func (o *DcimDeviceBaysPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim device bays partial update o k response has a 5xx status code
func (o *DcimDeviceBaysPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim device bays partial update o k response a status code equal to that given
func (o *DcimDeviceBaysPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimDeviceBaysPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/device-bays/{id}/][%d] dcimDeviceBaysPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimDeviceBaysPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/device-bays/{id}/][%d] dcimDeviceBaysPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimDeviceBaysPartialUpdateOK) GetPayload() *models.DeviceBay {
	return o.Payload
}

func (o *DcimDeviceBaysPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceBay)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimDeviceBaysPartialUpdateDefault creates a DcimDeviceBaysPartialUpdateDefault with default headers values
func NewDcimDeviceBaysPartialUpdateDefault(code int) *DcimDeviceBaysPartialUpdateDefault {
	return &DcimDeviceBaysPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimDeviceBaysPartialUpdateDefault describes a response with status code -1, with default header values.

DcimDeviceBaysPartialUpdateDefault dcim device bays partial update default
*/
type DcimDeviceBaysPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim device bays partial update default response
func (o *DcimDeviceBaysPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim device bays partial update default response has a 2xx status code
func (o *DcimDeviceBaysPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim device bays partial update default response has a 3xx status code
func (o *DcimDeviceBaysPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim device bays partial update default response has a 4xx status code
func (o *DcimDeviceBaysPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim device bays partial update default response has a 5xx status code
func (o *DcimDeviceBaysPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim device bays partial update default response a status code equal to that given
func (o *DcimDeviceBaysPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimDeviceBaysPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/device-bays/{id}/][%d] dcim_device-bays_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceBaysPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/device-bays/{id}/][%d] dcim_device-bays_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceBaysPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDeviceBaysPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
