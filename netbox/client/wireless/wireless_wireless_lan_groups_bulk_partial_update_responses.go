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

package wireless

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// WirelessWirelessLanGroupsBulkPartialUpdateReader is a Reader for the WirelessWirelessLanGroupsBulkPartialUpdate structure.
type WirelessWirelessLanGroupsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WirelessWirelessLanGroupsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWirelessWirelessLanGroupsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWirelessWirelessLanGroupsBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWirelessWirelessLanGroupsBulkPartialUpdateOK creates a WirelessWirelessLanGroupsBulkPartialUpdateOK with default headers values
func NewWirelessWirelessLanGroupsBulkPartialUpdateOK() *WirelessWirelessLanGroupsBulkPartialUpdateOK {
	return &WirelessWirelessLanGroupsBulkPartialUpdateOK{}
}

/* WirelessWirelessLanGroupsBulkPartialUpdateOK describes a response with status code 200, with default header values.

WirelessWirelessLanGroupsBulkPartialUpdateOK wireless wireless lan groups bulk partial update o k
*/
type WirelessWirelessLanGroupsBulkPartialUpdateOK struct {
	Payload *models.WirelessLANGroup
}

func (o *WirelessWirelessLanGroupsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /wireless/wireless-lan-groups/][%d] wirelessWirelessLanGroupsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *WirelessWirelessLanGroupsBulkPartialUpdateOK) GetPayload() *models.WirelessLANGroup {
	return o.Payload
}

func (o *WirelessWirelessLanGroupsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WirelessLANGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWirelessWirelessLanGroupsBulkPartialUpdateDefault creates a WirelessWirelessLanGroupsBulkPartialUpdateDefault with default headers values
func NewWirelessWirelessLanGroupsBulkPartialUpdateDefault(code int) *WirelessWirelessLanGroupsBulkPartialUpdateDefault {
	return &WirelessWirelessLanGroupsBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/* WirelessWirelessLanGroupsBulkPartialUpdateDefault describes a response with status code -1, with default header values.

WirelessWirelessLanGroupsBulkPartialUpdateDefault wireless wireless lan groups bulk partial update default
*/
type WirelessWirelessLanGroupsBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the wireless wireless lan groups bulk partial update default response
func (o *WirelessWirelessLanGroupsBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *WirelessWirelessLanGroupsBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /wireless/wireless-lan-groups/][%d] wireless_wireless-lan-groups_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *WirelessWirelessLanGroupsBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *WirelessWirelessLanGroupsBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
