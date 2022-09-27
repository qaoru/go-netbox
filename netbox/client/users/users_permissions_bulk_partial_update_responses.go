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

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// UsersPermissionsBulkPartialUpdateReader is a Reader for the UsersPermissionsBulkPartialUpdate structure.
type UsersPermissionsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersPermissionsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersPermissionsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUsersPermissionsBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUsersPermissionsBulkPartialUpdateOK creates a UsersPermissionsBulkPartialUpdateOK with default headers values
func NewUsersPermissionsBulkPartialUpdateOK() *UsersPermissionsBulkPartialUpdateOK {
	return &UsersPermissionsBulkPartialUpdateOK{}
}

/*
UsersPermissionsBulkPartialUpdateOK describes a response with status code 200, with default header values.

UsersPermissionsBulkPartialUpdateOK users permissions bulk partial update o k
*/
type UsersPermissionsBulkPartialUpdateOK struct {
	Payload *models.ObjectPermission
}

// IsSuccess returns true when this users permissions bulk partial update o k response has a 2xx status code
func (o *UsersPermissionsBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users permissions bulk partial update o k response has a 3xx status code
func (o *UsersPermissionsBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users permissions bulk partial update o k response has a 4xx status code
func (o *UsersPermissionsBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users permissions bulk partial update o k response has a 5xx status code
func (o *UsersPermissionsBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users permissions bulk partial update o k response a status code equal to that given
func (o *UsersPermissionsBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *UsersPermissionsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /users/permissions/][%d] usersPermissionsBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *UsersPermissionsBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /users/permissions/][%d] usersPermissionsBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *UsersPermissionsBulkPartialUpdateOK) GetPayload() *models.ObjectPermission {
	return o.Payload
}

func (o *UsersPermissionsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectPermission)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersPermissionsBulkPartialUpdateDefault creates a UsersPermissionsBulkPartialUpdateDefault with default headers values
func NewUsersPermissionsBulkPartialUpdateDefault(code int) *UsersPermissionsBulkPartialUpdateDefault {
	return &UsersPermissionsBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
UsersPermissionsBulkPartialUpdateDefault describes a response with status code -1, with default header values.

UsersPermissionsBulkPartialUpdateDefault users permissions bulk partial update default
*/
type UsersPermissionsBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the users permissions bulk partial update default response
func (o *UsersPermissionsBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this users permissions bulk partial update default response has a 2xx status code
func (o *UsersPermissionsBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this users permissions bulk partial update default response has a 3xx status code
func (o *UsersPermissionsBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this users permissions bulk partial update default response has a 4xx status code
func (o *UsersPermissionsBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this users permissions bulk partial update default response has a 5xx status code
func (o *UsersPermissionsBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this users permissions bulk partial update default response a status code equal to that given
func (o *UsersPermissionsBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UsersPermissionsBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /users/permissions/][%d] users_permissions_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *UsersPermissionsBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /users/permissions/][%d] users_permissions_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *UsersPermissionsBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersPermissionsBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
