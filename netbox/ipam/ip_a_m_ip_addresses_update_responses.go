// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
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

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/digitalocean/go-netbox/netbox/models"
)

// IPAMIPAddressesUpdateReader is a Reader for the IPAMIPAddressesUpdate structure.
type IPAMIPAddressesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPAMIPAddressesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIPAMIPAddressesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIPAMIPAddressesUpdateOK creates a IPAMIPAddressesUpdateOK with default headers values
func NewIPAMIPAddressesUpdateOK() *IPAMIPAddressesUpdateOK {
	return &IPAMIPAddressesUpdateOK{}
}

/*IPAMIPAddressesUpdateOK handles this case with default header values.

IPAMIPAddressesUpdateOK ipam Ip addresses update o k
*/
type IPAMIPAddressesUpdateOK struct {
	Payload *models.IPAddress
}

func (o *IPAMIPAddressesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/ip-addresses/{id}/][%d] ipamIpAddressesUpdateOK  %+v", 200, o.Payload)
}

func (o *IPAMIPAddressesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPAddress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
