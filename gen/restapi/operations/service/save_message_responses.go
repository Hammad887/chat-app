// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Hammad887/chat-app/gen/models"
)

// SaveMessageCreatedCode is the HTTP code returned for type SaveMessageCreated
const SaveMessageCreatedCode int = 201

/*
SaveMessageCreated successfully save user object into database

swagger:response saveMessageCreated
*/
type SaveMessageCreated struct {

	/*
	  In: Body
	*/
	Payload *models.SuccessResponse `json:"body,omitempty"`
}

// NewSaveMessageCreated creates SaveMessageCreated with default headers values
func NewSaveMessageCreated() *SaveMessageCreated {

	return &SaveMessageCreated{}
}

// WithPayload adds the payload to the save message created response
func (o *SaveMessageCreated) WithPayload(payload *models.SuccessResponse) *SaveMessageCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the save message created response
func (o *SaveMessageCreated) SetPayload(payload *models.SuccessResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SaveMessageCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SaveMessageUnauthorizedCode is the HTTP code returned for type SaveMessageUnauthorized
const SaveMessageUnauthorizedCode int = 401

/*
SaveMessageUnauthorized Unauthorized

swagger:response saveMessageUnauthorized
*/
type SaveMessageUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSaveMessageUnauthorized creates SaveMessageUnauthorized with default headers values
func NewSaveMessageUnauthorized() *SaveMessageUnauthorized {

	return &SaveMessageUnauthorized{}
}

// WithPayload adds the payload to the save message unauthorized response
func (o *SaveMessageUnauthorized) WithPayload(payload *models.Error) *SaveMessageUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the save message unauthorized response
func (o *SaveMessageUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SaveMessageUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SaveMessageConflictCode is the HTTP code returned for type SaveMessageConflict
const SaveMessageConflictCode int = 409

/*
SaveMessageConflict Conflict

swagger:response saveMessageConflict
*/
type SaveMessageConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSaveMessageConflict creates SaveMessageConflict with default headers values
func NewSaveMessageConflict() *SaveMessageConflict {

	return &SaveMessageConflict{}
}

// WithPayload adds the payload to the save message conflict response
func (o *SaveMessageConflict) WithPayload(payload *models.Error) *SaveMessageConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the save message conflict response
func (o *SaveMessageConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SaveMessageConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
SaveMessageDefault Internal Server Error

swagger:response saveMessageDefault
*/
type SaveMessageDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSaveMessageDefault creates SaveMessageDefault with default headers values
func NewSaveMessageDefault(code int) *SaveMessageDefault {
	if code <= 0 {
		code = 500
	}

	return &SaveMessageDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the save message default response
func (o *SaveMessageDefault) WithStatusCode(code int) *SaveMessageDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the save message default response
func (o *SaveMessageDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the save message default response
func (o *SaveMessageDefault) WithPayload(payload *models.Error) *SaveMessageDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the save message default response
func (o *SaveMessageDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SaveMessageDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
