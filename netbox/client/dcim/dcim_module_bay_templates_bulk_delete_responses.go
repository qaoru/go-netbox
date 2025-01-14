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

// DcimModuleBayTemplatesBulkDeleteReader is a Reader for the DcimModuleBayTemplatesBulkDelete structure.
type DcimModuleBayTemplatesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimModuleBayTemplatesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimModuleBayTemplatesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimModuleBayTemplatesBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimModuleBayTemplatesBulkDeleteNoContent creates a DcimModuleBayTemplatesBulkDeleteNoContent with default headers values
func NewDcimModuleBayTemplatesBulkDeleteNoContent() *DcimModuleBayTemplatesBulkDeleteNoContent {
	return &DcimModuleBayTemplatesBulkDeleteNoContent{}
}

/*
DcimModuleBayTemplatesBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimModuleBayTemplatesBulkDeleteNoContent dcim module bay templates bulk delete no content
*/
type DcimModuleBayTemplatesBulkDeleteNoContent struct {
}

// IsSuccess returns true when this dcim module bay templates bulk delete no content response has a 2xx status code
func (o *DcimModuleBayTemplatesBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim module bay templates bulk delete no content response has a 3xx status code
func (o *DcimModuleBayTemplatesBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim module bay templates bulk delete no content response has a 4xx status code
func (o *DcimModuleBayTemplatesBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim module bay templates bulk delete no content response has a 5xx status code
func (o *DcimModuleBayTemplatesBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim module bay templates bulk delete no content response a status code equal to that given
func (o *DcimModuleBayTemplatesBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DcimModuleBayTemplatesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/module-bay-templates/][%d] dcimModuleBayTemplatesBulkDeleteNoContent ", 204)
}

func (o *DcimModuleBayTemplatesBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/module-bay-templates/][%d] dcimModuleBayTemplatesBulkDeleteNoContent ", 204)
}

func (o *DcimModuleBayTemplatesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimModuleBayTemplatesBulkDeleteDefault creates a DcimModuleBayTemplatesBulkDeleteDefault with default headers values
func NewDcimModuleBayTemplatesBulkDeleteDefault(code int) *DcimModuleBayTemplatesBulkDeleteDefault {
	return &DcimModuleBayTemplatesBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimModuleBayTemplatesBulkDeleteDefault describes a response with status code -1, with default header values.

DcimModuleBayTemplatesBulkDeleteDefault dcim module bay templates bulk delete default
*/
type DcimModuleBayTemplatesBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim module bay templates bulk delete default response
func (o *DcimModuleBayTemplatesBulkDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim module bay templates bulk delete default response has a 2xx status code
func (o *DcimModuleBayTemplatesBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim module bay templates bulk delete default response has a 3xx status code
func (o *DcimModuleBayTemplatesBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim module bay templates bulk delete default response has a 4xx status code
func (o *DcimModuleBayTemplatesBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim module bay templates bulk delete default response has a 5xx status code
func (o *DcimModuleBayTemplatesBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim module bay templates bulk delete default response a status code equal to that given
func (o *DcimModuleBayTemplatesBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimModuleBayTemplatesBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/module-bay-templates/][%d] dcim_module-bay-templates_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModuleBayTemplatesBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/module-bay-templates/][%d] dcim_module-bay-templates_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModuleBayTemplatesBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimModuleBayTemplatesBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
