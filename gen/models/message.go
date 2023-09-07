// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Message message
//
// swagger:model message
type Message struct {

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// room id
	// Required: true
	RoomID *string `json:"room_id"`

	// sender id
	// Required: true
	SenderID *string `json:"sender_id"`

	// text
	// Required: true
	Text *string `json:"text"`
}

// Validate validates this message
func (m *Message) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRoomID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSenderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateText(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Message) validateRoomID(formats strfmt.Registry) error {

	if err := validate.Required("room_id", "body", m.RoomID); err != nil {
		return err
	}

	return nil
}

func (m *Message) validateSenderID(formats strfmt.Registry) error {

	if err := validate.Required("sender_id", "body", m.SenderID); err != nil {
		return err
	}

	return nil
}

func (m *Message) validateText(formats strfmt.Registry) error {

	if err := validate.Required("text", "body", m.Text); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this message based on context it is used
func (m *Message) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Message) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Message) UnmarshalBinary(b []byte) error {
	var res Message
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
