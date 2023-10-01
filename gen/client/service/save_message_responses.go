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

// SaveMessageReader is a Reader for the SaveMessage structure.
type SaveMessageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SaveMessageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSaveMessageCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSaveMessageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewSaveMessageConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewSaveMessageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSaveMessageCreated creates a SaveMessageCreated with default headers values
func NewSaveMessageCreated() *SaveMessageCreated {
	return &SaveMessageCreated{}
}

/*
SaveMessageCreated describes a response with status code 201, with default header values.

successfully save user object into database
*/
type SaveMessageCreated struct {
	Payload *models.SuccessResponse
}

// IsSuccess returns true when this save message created response has a 2xx status code
func (o *SaveMessageCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this save message created response has a 3xx status code
func (o *SaveMessageCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this save message created response has a 4xx status code
func (o *SaveMessageCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this save message created response has a 5xx status code
func (o *SaveMessageCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this save message created response a status code equal to that given
func (o *SaveMessageCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the save message created response
func (o *SaveMessageCreated) Code() int {
	return 201
}

func (o *SaveMessageCreated) Error() string {
	return fmt.Sprintf("[POST /chatrooms/{chatroom_id}/messages][%d] saveMessageCreated  %+v", 201, o.Payload)
}

func (o *SaveMessageCreated) String() string {
	return fmt.Sprintf("[POST /chatrooms/{chatroom_id}/messages][%d] saveMessageCreated  %+v", 201, o.Payload)
}

func (o *SaveMessageCreated) GetPayload() *models.SuccessResponse {
	return o.Payload
}

func (o *SaveMessageCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSaveMessageUnauthorized creates a SaveMessageUnauthorized with default headers values
func NewSaveMessageUnauthorized() *SaveMessageUnauthorized {
	return &SaveMessageUnauthorized{}
}

/*
SaveMessageUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SaveMessageUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this save message unauthorized response has a 2xx status code
func (o *SaveMessageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this save message unauthorized response has a 3xx status code
func (o *SaveMessageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this save message unauthorized response has a 4xx status code
func (o *SaveMessageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this save message unauthorized response has a 5xx status code
func (o *SaveMessageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this save message unauthorized response a status code equal to that given
func (o *SaveMessageUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the save message unauthorized response
func (o *SaveMessageUnauthorized) Code() int {
	return 401
}

func (o *SaveMessageUnauthorized) Error() string {
	return fmt.Sprintf("[POST /chatrooms/{chatroom_id}/messages][%d] saveMessageUnauthorized  %+v", 401, o.Payload)
}

func (o *SaveMessageUnauthorized) String() string {
	return fmt.Sprintf("[POST /chatrooms/{chatroom_id}/messages][%d] saveMessageUnauthorized  %+v", 401, o.Payload)
}

func (o *SaveMessageUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *SaveMessageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSaveMessageConflict creates a SaveMessageConflict with default headers values
func NewSaveMessageConflict() *SaveMessageConflict {
	return &SaveMessageConflict{}
}

/*
SaveMessageConflict describes a response with status code 409, with default header values.

Conflict
*/
type SaveMessageConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this save message conflict response has a 2xx status code
func (o *SaveMessageConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this save message conflict response has a 3xx status code
func (o *SaveMessageConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this save message conflict response has a 4xx status code
func (o *SaveMessageConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this save message conflict response has a 5xx status code
func (o *SaveMessageConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this save message conflict response a status code equal to that given
func (o *SaveMessageConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the save message conflict response
func (o *SaveMessageConflict) Code() int {
	return 409
}

func (o *SaveMessageConflict) Error() string {
	return fmt.Sprintf("[POST /chatrooms/{chatroom_id}/messages][%d] saveMessageConflict  %+v", 409, o.Payload)
}

func (o *SaveMessageConflict) String() string {
	return fmt.Sprintf("[POST /chatrooms/{chatroom_id}/messages][%d] saveMessageConflict  %+v", 409, o.Payload)
}

func (o *SaveMessageConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *SaveMessageConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSaveMessageDefault creates a SaveMessageDefault with default headers values
func NewSaveMessageDefault(code int) *SaveMessageDefault {
	return &SaveMessageDefault{
		_statusCode: code,
	}
}

/*
SaveMessageDefault describes a response with status code -1, with default header values.

Internal Server Error
*/
type SaveMessageDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this save message default response has a 2xx status code
func (o *SaveMessageDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this save message default response has a 3xx status code
func (o *SaveMessageDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this save message default response has a 4xx status code
func (o *SaveMessageDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this save message default response has a 5xx status code
func (o *SaveMessageDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this save message default response a status code equal to that given
func (o *SaveMessageDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the save message default response
func (o *SaveMessageDefault) Code() int {
	return o._statusCode
}

func (o *SaveMessageDefault) Error() string {
	return fmt.Sprintf("[POST /chatrooms/{chatroom_id}/messages][%d] SaveMessage default  %+v", o._statusCode, o.Payload)
}

func (o *SaveMessageDefault) String() string {
	return fmt.Sprintf("[POST /chatrooms/{chatroom_id}/messages][%d] SaveMessage default  %+v", o._statusCode, o.Payload)
}

func (o *SaveMessageDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *SaveMessageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
