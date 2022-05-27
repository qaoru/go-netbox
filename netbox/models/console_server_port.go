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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConsoleServerPort console server port
//
// swagger:model ConsoleServerPort
type ConsoleServerPort struct {

	// occupied
	// Read Only: true
	Occupied *bool `json:"_occupied,omitempty"`

	// cable
	Cable *NestedCable `json:"cable,omitempty"`

	// Connected endpoint
	//
	//
	// Return the appropriate serializer for the type of connected object.
	//
	// Read Only: true
	ConnectedEndpoint map[string]*string `json:"connected_endpoint,omitempty"`

	// Connected endpoint reachable
	// Read Only: true
	ConnectedEndpointReachable *bool `json:"connected_endpoint_reachable,omitempty"`

	// Connected endpoint type
	// Read Only: true
	ConnectedEndpointType string `json:"connected_endpoint_type,omitempty"`

	// Created
	// Read Only: true
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// device
	// Required: true
	Device *NestedDevice `json:"device"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Label
	//
	// Physical label
	// Max Length: 64
	Label string `json:"label,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Link peer
	//
	//
	// Return the appropriate serializer for the link termination model.
	//
	// Read Only: true
	LinkPeer map[string]*string `json:"link_peer,omitempty"`

	// Link peer type
	// Read Only: true
	LinkPeerType string `json:"link_peer_type,omitempty"`

	// Mark connected
	//
	// Treat as if a cable is connected
	MarkConnected bool `json:"mark_connected,omitempty"`

	// module
	Module *ComponentNestedModule `json:"module,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// speed
	Speed *ConsoleServerPortSpeed `json:"speed,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// type
	Type *ConsoleServerPortType `json:"type,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this console server port
func (m *ConsoleServerPort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpeed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsoleServerPort) validateCable(formats strfmt.Registry) error {
	if swag.IsZero(m.Cable) { // not required
		return nil
	}

	if m.Cable != nil {
		if err := m.Cable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *ConsoleServerPort) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConsoleServerPort) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *ConsoleServerPort) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	if m.Device != nil {
		if err := m.Device.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *ConsoleServerPort) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", m.Label, 64); err != nil {
		return err
	}

	return nil
}

func (m *ConsoleServerPort) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConsoleServerPort) validateModule(formats strfmt.Registry) error {
	if swag.IsZero(m.Module) { // not required
		return nil
	}

	if m.Module != nil {
		if err := m.Module.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("module")
			}
			return err
		}
	}

	return nil
}

func (m *ConsoleServerPort) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 64); err != nil {
		return err
	}

	return nil
}

func (m *ConsoleServerPort) validateSpeed(formats strfmt.Registry) error {
	if swag.IsZero(m.Speed) { // not required
		return nil
	}

	if m.Speed != nil {
		if err := m.Speed.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("speed")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("speed")
			}
			return err
		}
	}

	return nil
}

func (m *ConsoleServerPort) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConsoleServerPort) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *ConsoleServerPort) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this console server port based on the context it is used
func (m *ConsoleServerPort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOccupied(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectedEndpoint(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectedEndpointReachable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectedEndpointType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinkPeer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinkPeerType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSpeed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsoleServerPort) contextValidateOccupied(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "_occupied", "body", m.Occupied); err != nil {
		return err
	}

	return nil
}

func (m *ConsoleServerPort) contextValidateCable(ctx context.Context, formats strfmt.Registry) error {

	if m.Cable != nil {
		if err := m.Cable.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *ConsoleServerPort) contextValidateConnectedEndpoint(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ConsoleServerPort) contextValidateConnectedEndpointReachable(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connected_endpoint_reachable", "body", m.ConnectedEndpointReachable); err != nil {
		return err
	}

	return nil
}

func (m *ConsoleServerPort) contextValidateConnectedEndpointType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connected_endpoint_type", "body", string(m.ConnectedEndpointType)); err != nil {
		return err
	}

	return nil
}

func (m *ConsoleServerPort) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.DateTime(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *ConsoleServerPort) contextValidateDevice(ctx context.Context, formats strfmt.Registry) error {

	if m.Device != nil {
		if err := m.Device.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *ConsoleServerPort) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *ConsoleServerPort) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *ConsoleServerPort) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", strfmt.DateTime(m.LastUpdated)); err != nil {
		return err
	}

	return nil
}

func (m *ConsoleServerPort) contextValidateLinkPeer(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ConsoleServerPort) contextValidateLinkPeerType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "link_peer_type", "body", string(m.LinkPeerType)); err != nil {
		return err
	}

	return nil
}

func (m *ConsoleServerPort) contextValidateModule(ctx context.Context, formats strfmt.Registry) error {

	if m.Module != nil {
		if err := m.Module.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("module")
			}
			return err
		}
	}

	return nil
}

func (m *ConsoleServerPort) contextValidateSpeed(ctx context.Context, formats strfmt.Registry) error {

	if m.Speed != nil {
		if err := m.Speed.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("speed")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("speed")
			}
			return err
		}
	}

	return nil
}

func (m *ConsoleServerPort) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConsoleServerPort) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *ConsoleServerPort) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsoleServerPort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsoleServerPort) UnmarshalBinary(b []byte) error {
	var res ConsoleServerPort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ConsoleServerPortSpeed Speed
//
// swagger:model ConsoleServerPortSpeed
type ConsoleServerPortSpeed struct {

	// label
	// Required: true
	// Enum: [1200 bps 2400 bps 4800 bps 9600 bps 19.2 kbps 38.4 kbps 57.6 kbps 115.2 kbps]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [1200 2400 4800 9600 19200 38400 57600 115200]
	Value *int64 `json:"value"`
}

// Validate validates this console server port speed
func (m *ConsoleServerPortSpeed) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var consoleServerPortSpeedTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["1200 bps","2400 bps","4800 bps","9600 bps","19.2 kbps","38.4 kbps","57.6 kbps","115.2 kbps"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		consoleServerPortSpeedTypeLabelPropEnum = append(consoleServerPortSpeedTypeLabelPropEnum, v)
	}
}

const (

	// ConsoleServerPortSpeedLabelNr1200Bps captures enum value "1200 bps"
	ConsoleServerPortSpeedLabelNr1200Bps string = "1200 bps"

	// ConsoleServerPortSpeedLabelNr2400Bps captures enum value "2400 bps"
	ConsoleServerPortSpeedLabelNr2400Bps string = "2400 bps"

	// ConsoleServerPortSpeedLabelNr4800Bps captures enum value "4800 bps"
	ConsoleServerPortSpeedLabelNr4800Bps string = "4800 bps"

	// ConsoleServerPortSpeedLabelNr9600Bps captures enum value "9600 bps"
	ConsoleServerPortSpeedLabelNr9600Bps string = "9600 bps"

	// ConsoleServerPortSpeedLabelNr19Dot2Kbps captures enum value "19.2 kbps"
	ConsoleServerPortSpeedLabelNr19Dot2Kbps string = "19.2 kbps"

	// ConsoleServerPortSpeedLabelNr38Dot4Kbps captures enum value "38.4 kbps"
	ConsoleServerPortSpeedLabelNr38Dot4Kbps string = "38.4 kbps"

	// ConsoleServerPortSpeedLabelNr57Dot6Kbps captures enum value "57.6 kbps"
	ConsoleServerPortSpeedLabelNr57Dot6Kbps string = "57.6 kbps"

	// ConsoleServerPortSpeedLabelNr115Dot2Kbps captures enum value "115.2 kbps"
	ConsoleServerPortSpeedLabelNr115Dot2Kbps string = "115.2 kbps"
)

// prop value enum
func (m *ConsoleServerPortSpeed) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, consoleServerPortSpeedTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConsoleServerPortSpeed) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("speed"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("speed"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var consoleServerPortSpeedTypeValuePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1200,2400,4800,9600,19200,38400,57600,115200]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		consoleServerPortSpeedTypeValuePropEnum = append(consoleServerPortSpeedTypeValuePropEnum, v)
	}
}

// prop value enum
func (m *ConsoleServerPortSpeed) validateValueEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, consoleServerPortSpeedTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConsoleServerPortSpeed) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("speed"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("speed"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this console server port speed based on context it is used
func (m *ConsoleServerPortSpeed) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConsoleServerPortSpeed) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsoleServerPortSpeed) UnmarshalBinary(b []byte) error {
	var res ConsoleServerPortSpeed
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ConsoleServerPortType Type
//
// swagger:model ConsoleServerPortType
type ConsoleServerPortType struct {

	// label
	// Required: true
	// Enum: [DE-9 DB-25 RJ-11 RJ-12 RJ-45 Mini-DIN 8 USB Type A USB Type B USB Type C USB Mini A USB Mini B USB Micro A USB Micro B USB Micro AB Other]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [de-9 db-25 rj-11 rj-12 rj-45 mini-din-8 usb-a usb-b usb-c usb-mini-a usb-mini-b usb-micro-a usb-micro-b usb-micro-ab other]
	Value *string `json:"value"`
}

// Validate validates this console server port type
func (m *ConsoleServerPortType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var consoleServerPortTypeTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DE-9","DB-25","RJ-11","RJ-12","RJ-45","Mini-DIN 8","USB Type A","USB Type B","USB Type C","USB Mini A","USB Mini B","USB Micro A","USB Micro B","USB Micro AB","Other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		consoleServerPortTypeTypeLabelPropEnum = append(consoleServerPortTypeTypeLabelPropEnum, v)
	}
}

const (

	// ConsoleServerPortTypeLabelDEDash9 captures enum value "DE-9"
	ConsoleServerPortTypeLabelDEDash9 string = "DE-9"

	// ConsoleServerPortTypeLabelDBDash25 captures enum value "DB-25"
	ConsoleServerPortTypeLabelDBDash25 string = "DB-25"

	// ConsoleServerPortTypeLabelRJDash11 captures enum value "RJ-11"
	ConsoleServerPortTypeLabelRJDash11 string = "RJ-11"

	// ConsoleServerPortTypeLabelRJDash12 captures enum value "RJ-12"
	ConsoleServerPortTypeLabelRJDash12 string = "RJ-12"

	// ConsoleServerPortTypeLabelRJDash45 captures enum value "RJ-45"
	ConsoleServerPortTypeLabelRJDash45 string = "RJ-45"

	// ConsoleServerPortTypeLabelMiniDashDIN8 captures enum value "Mini-DIN 8"
	ConsoleServerPortTypeLabelMiniDashDIN8 string = "Mini-DIN 8"

	// ConsoleServerPortTypeLabelUSBTypeA captures enum value "USB Type A"
	ConsoleServerPortTypeLabelUSBTypeA string = "USB Type A"

	// ConsoleServerPortTypeLabelUSBTypeB captures enum value "USB Type B"
	ConsoleServerPortTypeLabelUSBTypeB string = "USB Type B"

	// ConsoleServerPortTypeLabelUSBTypeC captures enum value "USB Type C"
	ConsoleServerPortTypeLabelUSBTypeC string = "USB Type C"

	// ConsoleServerPortTypeLabelUSBMiniA captures enum value "USB Mini A"
	ConsoleServerPortTypeLabelUSBMiniA string = "USB Mini A"

	// ConsoleServerPortTypeLabelUSBMiniB captures enum value "USB Mini B"
	ConsoleServerPortTypeLabelUSBMiniB string = "USB Mini B"

	// ConsoleServerPortTypeLabelUSBMicroA captures enum value "USB Micro A"
	ConsoleServerPortTypeLabelUSBMicroA string = "USB Micro A"

	// ConsoleServerPortTypeLabelUSBMicroB captures enum value "USB Micro B"
	ConsoleServerPortTypeLabelUSBMicroB string = "USB Micro B"

	// ConsoleServerPortTypeLabelUSBMicroAB captures enum value "USB Micro AB"
	ConsoleServerPortTypeLabelUSBMicroAB string = "USB Micro AB"

	// ConsoleServerPortTypeLabelOther captures enum value "Other"
	ConsoleServerPortTypeLabelOther string = "Other"
)

// prop value enum
func (m *ConsoleServerPortType) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, consoleServerPortTypeTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConsoleServerPortType) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("type"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var consoleServerPortTypeTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["de-9","db-25","rj-11","rj-12","rj-45","mini-din-8","usb-a","usb-b","usb-c","usb-mini-a","usb-mini-b","usb-micro-a","usb-micro-b","usb-micro-ab","other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		consoleServerPortTypeTypeValuePropEnum = append(consoleServerPortTypeTypeValuePropEnum, v)
	}
}

const (

	// ConsoleServerPortTypeValueDeDash9 captures enum value "de-9"
	ConsoleServerPortTypeValueDeDash9 string = "de-9"

	// ConsoleServerPortTypeValueDbDash25 captures enum value "db-25"
	ConsoleServerPortTypeValueDbDash25 string = "db-25"

	// ConsoleServerPortTypeValueRjDash11 captures enum value "rj-11"
	ConsoleServerPortTypeValueRjDash11 string = "rj-11"

	// ConsoleServerPortTypeValueRjDash12 captures enum value "rj-12"
	ConsoleServerPortTypeValueRjDash12 string = "rj-12"

	// ConsoleServerPortTypeValueRjDash45 captures enum value "rj-45"
	ConsoleServerPortTypeValueRjDash45 string = "rj-45"

	// ConsoleServerPortTypeValueMiniDashDinDash8 captures enum value "mini-din-8"
	ConsoleServerPortTypeValueMiniDashDinDash8 string = "mini-din-8"

	// ConsoleServerPortTypeValueUsbDasha captures enum value "usb-a"
	ConsoleServerPortTypeValueUsbDasha string = "usb-a"

	// ConsoleServerPortTypeValueUsbDashb captures enum value "usb-b"
	ConsoleServerPortTypeValueUsbDashb string = "usb-b"

	// ConsoleServerPortTypeValueUsbDashc captures enum value "usb-c"
	ConsoleServerPortTypeValueUsbDashc string = "usb-c"

	// ConsoleServerPortTypeValueUsbDashMiniDasha captures enum value "usb-mini-a"
	ConsoleServerPortTypeValueUsbDashMiniDasha string = "usb-mini-a"

	// ConsoleServerPortTypeValueUsbDashMiniDashb captures enum value "usb-mini-b"
	ConsoleServerPortTypeValueUsbDashMiniDashb string = "usb-mini-b"

	// ConsoleServerPortTypeValueUsbDashMicroDasha captures enum value "usb-micro-a"
	ConsoleServerPortTypeValueUsbDashMicroDasha string = "usb-micro-a"

	// ConsoleServerPortTypeValueUsbDashMicroDashb captures enum value "usb-micro-b"
	ConsoleServerPortTypeValueUsbDashMicroDashb string = "usb-micro-b"

	// ConsoleServerPortTypeValueUsbDashMicroDashAb captures enum value "usb-micro-ab"
	ConsoleServerPortTypeValueUsbDashMicroDashAb string = "usb-micro-ab"

	// ConsoleServerPortTypeValueOther captures enum value "other"
	ConsoleServerPortTypeValueOther string = "other"
)

// prop value enum
func (m *ConsoleServerPortType) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, consoleServerPortTypeTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConsoleServerPortType) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("type"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this console server port type based on context it is used
func (m *ConsoleServerPortType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConsoleServerPortType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsoleServerPortType) UnmarshalBinary(b []byte) error {
	var res ConsoleServerPortType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
