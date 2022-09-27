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
)

// DcimModuleTypesDeleteReader is a Reader for the DcimModuleTypesDelete structure.
type DcimModuleTypesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimModuleTypesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimModuleTypesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimModuleTypesDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimModuleTypesDeleteNoContent creates a DcimModuleTypesDeleteNoContent with default headers values
func NewDcimModuleTypesDeleteNoContent() *DcimModuleTypesDeleteNoContent {
	return &DcimModuleTypesDeleteNoContent{}
}

/*
DcimModuleTypesDeleteNoContent describes a response with status code 204, with default header values.

DcimModuleTypesDeleteNoContent dcim module types delete no content
*/
type DcimModuleTypesDeleteNoContent struct {
}

// IsSuccess returns true when this dcim module types delete no content response has a 2xx status code
func (o *DcimModuleTypesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim module types delete no content response has a 3xx status code
func (o *DcimModuleTypesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim module types delete no content response has a 4xx status code
func (o *DcimModuleTypesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim module types delete no content response has a 5xx status code
func (o *DcimModuleTypesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim module types delete no content response a status code equal to that given
func (o *DcimModuleTypesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DcimModuleTypesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/module-types/{id}/][%d] dcimModuleTypesDeleteNoContent ", 204)
}

func (o *DcimModuleTypesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/module-types/{id}/][%d] dcimModuleTypesDeleteNoContent ", 204)
}

func (o *DcimModuleTypesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimModuleTypesDeleteDefault creates a DcimModuleTypesDeleteDefault with default headers values
func NewDcimModuleTypesDeleteDefault(code int) *DcimModuleTypesDeleteDefault {
	return &DcimModuleTypesDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimModuleTypesDeleteDefault describes a response with status code -1, with default header values.

DcimModuleTypesDeleteDefault dcim module types delete default
*/
type DcimModuleTypesDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim module types delete default response
func (o *DcimModuleTypesDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim module types delete default response has a 2xx status code
func (o *DcimModuleTypesDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim module types delete default response has a 3xx status code
func (o *DcimModuleTypesDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim module types delete default response has a 4xx status code
func (o *DcimModuleTypesDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim module types delete default response has a 5xx status code
func (o *DcimModuleTypesDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim module types delete default response a status code equal to that given
func (o *DcimModuleTypesDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimModuleTypesDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/module-types/{id}/][%d] dcim_module-types_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModuleTypesDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/module-types/{id}/][%d] dcim_module-types_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModuleTypesDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimModuleTypesDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
