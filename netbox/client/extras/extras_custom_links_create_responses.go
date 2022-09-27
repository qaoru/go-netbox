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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// ExtrasCustomLinksCreateReader is a Reader for the ExtrasCustomLinksCreate structure.
type ExtrasCustomLinksCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomLinksCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewExtrasCustomLinksCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasCustomLinksCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasCustomLinksCreateCreated creates a ExtrasCustomLinksCreateCreated with default headers values
func NewExtrasCustomLinksCreateCreated() *ExtrasCustomLinksCreateCreated {
	return &ExtrasCustomLinksCreateCreated{}
}

/*
ExtrasCustomLinksCreateCreated describes a response with status code 201, with default header values.

ExtrasCustomLinksCreateCreated extras custom links create created
*/
type ExtrasCustomLinksCreateCreated struct {
	Payload *models.CustomLink
}

// IsSuccess returns true when this extras custom links create created response has a 2xx status code
func (o *ExtrasCustomLinksCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras custom links create created response has a 3xx status code
func (o *ExtrasCustomLinksCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras custom links create created response has a 4xx status code
func (o *ExtrasCustomLinksCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras custom links create created response has a 5xx status code
func (o *ExtrasCustomLinksCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this extras custom links create created response a status code equal to that given
func (o *ExtrasCustomLinksCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *ExtrasCustomLinksCreateCreated) Error() string {
	return fmt.Sprintf("[POST /extras/custom-links/][%d] extrasCustomLinksCreateCreated  %+v", 201, o.Payload)
}

func (o *ExtrasCustomLinksCreateCreated) String() string {
	return fmt.Sprintf("[POST /extras/custom-links/][%d] extrasCustomLinksCreateCreated  %+v", 201, o.Payload)
}

func (o *ExtrasCustomLinksCreateCreated) GetPayload() *models.CustomLink {
	return o.Payload
}

func (o *ExtrasCustomLinksCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomLink)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasCustomLinksCreateDefault creates a ExtrasCustomLinksCreateDefault with default headers values
func NewExtrasCustomLinksCreateDefault(code int) *ExtrasCustomLinksCreateDefault {
	return &ExtrasCustomLinksCreateDefault{
		_statusCode: code,
	}
}

/*
ExtrasCustomLinksCreateDefault describes a response with status code -1, with default header values.

ExtrasCustomLinksCreateDefault extras custom links create default
*/
type ExtrasCustomLinksCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras custom links create default response
func (o *ExtrasCustomLinksCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this extras custom links create default response has a 2xx status code
func (o *ExtrasCustomLinksCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras custom links create default response has a 3xx status code
func (o *ExtrasCustomLinksCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras custom links create default response has a 4xx status code
func (o *ExtrasCustomLinksCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras custom links create default response has a 5xx status code
func (o *ExtrasCustomLinksCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras custom links create default response a status code equal to that given
func (o *ExtrasCustomLinksCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ExtrasCustomLinksCreateDefault) Error() string {
	return fmt.Sprintf("[POST /extras/custom-links/][%d] extras_custom-links_create default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasCustomLinksCreateDefault) String() string {
	return fmt.Sprintf("[POST /extras/custom-links/][%d] extras_custom-links_create default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasCustomLinksCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasCustomLinksCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
