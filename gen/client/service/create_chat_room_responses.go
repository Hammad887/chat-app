// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Hammad887/chat-app/gen/models"
)

// CreateChatRoomReader is a Reader for the CreateChatRoom structure.
type CreateChatRoomReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateChatRoomReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateChatRoomOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateChatRoomUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateChatRoomConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateChatRoomDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateChatRoomOK creates a CreateChatRoomOK with default headers values
func NewCreateChatRoomOK() *CreateChatRoomOK {
	return &CreateChatRoomOK{}
}

/*
CreateChatRoomOK describes a response with status code 200, with default header values.

Successfully created a chatroom
*/
type CreateChatRoomOK struct {
	Payload *models.ChatroomSuccess
}

// IsSuccess returns true when this create chat room o k response has a 2xx status code
func (o *CreateChatRoomOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create chat room o k response has a 3xx status code
func (o *CreateChatRoomOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create chat room o k response has a 4xx status code
func (o *CreateChatRoomOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create chat room o k response has a 5xx status code
func (o *CreateChatRoomOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create chat room o k response a status code equal to that given
func (o *CreateChatRoomOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create chat room o k response
func (o *CreateChatRoomOK) Code() int {
	return 200
}

func (o *CreateChatRoomOK) Error() string {
	return fmt.Sprintf("[POST /create-chat-room][%d] createChatRoomOK  %+v", 200, o.Payload)
}

func (o *CreateChatRoomOK) String() string {
	return fmt.Sprintf("[POST /create-chat-room][%d] createChatRoomOK  %+v", 200, o.Payload)
}

func (o *CreateChatRoomOK) GetPayload() *models.ChatroomSuccess {
	return o.Payload
}

func (o *CreateChatRoomOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ChatroomSuccess)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateChatRoomUnauthorized creates a CreateChatRoomUnauthorized with default headers values
func NewCreateChatRoomUnauthorized() *CreateChatRoomUnauthorized {
	return &CreateChatRoomUnauthorized{}
}

/*
CreateChatRoomUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateChatRoomUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this create chat room unauthorized response has a 2xx status code
func (o *CreateChatRoomUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create chat room unauthorized response has a 3xx status code
func (o *CreateChatRoomUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create chat room unauthorized response has a 4xx status code
func (o *CreateChatRoomUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create chat room unauthorized response has a 5xx status code
func (o *CreateChatRoomUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create chat room unauthorized response a status code equal to that given
func (o *CreateChatRoomUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create chat room unauthorized response
func (o *CreateChatRoomUnauthorized) Code() int {
	return 401
}

func (o *CreateChatRoomUnauthorized) Error() string {
	return fmt.Sprintf("[POST /create-chat-room][%d] createChatRoomUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateChatRoomUnauthorized) String() string {
	return fmt.Sprintf("[POST /create-chat-room][%d] createChatRoomUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateChatRoomUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateChatRoomUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateChatRoomConflict creates a CreateChatRoomConflict with default headers values
func NewCreateChatRoomConflict() *CreateChatRoomConflict {
	return &CreateChatRoomConflict{}
}

/*
CreateChatRoomConflict describes a response with status code 409, with default header values.

Conflict
*/
type CreateChatRoomConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this create chat room conflict response has a 2xx status code
func (o *CreateChatRoomConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create chat room conflict response has a 3xx status code
func (o *CreateChatRoomConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create chat room conflict response has a 4xx status code
func (o *CreateChatRoomConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create chat room conflict response has a 5xx status code
func (o *CreateChatRoomConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create chat room conflict response a status code equal to that given
func (o *CreateChatRoomConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create chat room conflict response
func (o *CreateChatRoomConflict) Code() int {
	return 409
}

func (o *CreateChatRoomConflict) Error() string {
	return fmt.Sprintf("[POST /create-chat-room][%d] createChatRoomConflict  %+v", 409, o.Payload)
}

func (o *CreateChatRoomConflict) String() string {
	return fmt.Sprintf("[POST /create-chat-room][%d] createChatRoomConflict  %+v", 409, o.Payload)
}

func (o *CreateChatRoomConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateChatRoomConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateChatRoomDefault creates a CreateChatRoomDefault with default headers values
func NewCreateChatRoomDefault(code int) *CreateChatRoomDefault {
	return &CreateChatRoomDefault{
		_statusCode: code,
	}
}

/*
CreateChatRoomDefault describes a response with status code -1, with default header values.

Internal Server Error
*/
type CreateChatRoomDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this create chat room default response has a 2xx status code
func (o *CreateChatRoomDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create chat room default response has a 3xx status code
func (o *CreateChatRoomDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create chat room default response has a 4xx status code
func (o *CreateChatRoomDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create chat room default response has a 5xx status code
func (o *CreateChatRoomDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create chat room default response a status code equal to that given
func (o *CreateChatRoomDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create chat room default response
func (o *CreateChatRoomDefault) Code() int {
	return o._statusCode
}

func (o *CreateChatRoomDefault) Error() string {
	return fmt.Sprintf("[POST /create-chat-room][%d] CreateChatRoom default  %+v", o._statusCode, o.Payload)
}

func (o *CreateChatRoomDefault) String() string {
	return fmt.Sprintf("[POST /create-chat-room][%d] CreateChatRoom default  %+v", o._statusCode, o.Payload)
}

func (o *CreateChatRoomDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateChatRoomDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
