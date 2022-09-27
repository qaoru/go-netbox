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

// UsersUsersBulkPartialUpdateReader is a Reader for the UsersUsersBulkPartialUpdate structure.
type UsersUsersBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersUsersBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersUsersBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUsersUsersBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUsersUsersBulkPartialUpdateOK creates a UsersUsersBulkPartialUpdateOK with default headers values
func NewUsersUsersBulkPartialUpdateOK() *UsersUsersBulkPartialUpdateOK {
	return &UsersUsersBulkPartialUpdateOK{}
}

/*
UsersUsersBulkPartialUpdateOK describes a response with status code 200, with default header values.

UsersUsersBulkPartialUpdateOK users users bulk partial update o k
*/
type UsersUsersBulkPartialUpdateOK struct {
	Payload *models.User
}

// IsSuccess returns true when this users users bulk partial update o k response has a 2xx status code
func (o *UsersUsersBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users users bulk partial update o k response has a 3xx status code
func (o *UsersUsersBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users users bulk partial update o k response has a 4xx status code
func (o *UsersUsersBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users users bulk partial update o k response has a 5xx status code
func (o *UsersUsersBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users users bulk partial update o k response a status code equal to that given
func (o *UsersUsersBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *UsersUsersBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /users/users/][%d] usersUsersBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *UsersUsersBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /users/users/][%d] usersUsersBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *UsersUsersBulkPartialUpdateOK) GetPayload() *models.User {
	return o.Payload
}

func (o *UsersUsersBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersUsersBulkPartialUpdateDefault creates a UsersUsersBulkPartialUpdateDefault with default headers values
func NewUsersUsersBulkPartialUpdateDefault(code int) *UsersUsersBulkPartialUpdateDefault {
	return &UsersUsersBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
UsersUsersBulkPartialUpdateDefault describes a response with status code -1, with default header values.

UsersUsersBulkPartialUpdateDefault users users bulk partial update default
*/
type UsersUsersBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the users users bulk partial update default response
func (o *UsersUsersBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this users users bulk partial update default response has a 2xx status code
func (o *UsersUsersBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this users users bulk partial update default response has a 3xx status code
func (o *UsersUsersBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this users users bulk partial update default response has a 4xx status code
func (o *UsersUsersBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this users users bulk partial update default response has a 5xx status code
func (o *UsersUsersBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this users users bulk partial update default response a status code equal to that given
func (o *UsersUsersBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UsersUsersBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /users/users/][%d] users_users_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *UsersUsersBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /users/users/][%d] users_users_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *UsersUsersBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersUsersBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
