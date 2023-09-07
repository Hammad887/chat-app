// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Hammad887/chat-app/gen/models"
)

// GetChatroomOKCode is the HTTP code returned for type GetChatroomOK
const GetChatroomOKCode int = 200

/*
GetChatroomOK successfully save user object into database

swagger:response getChatroomOK
*/
type GetChatroomOK struct {

	/*
	  In: Body
	*/
	Payload *models.Chatroom `json:"body,omitempty"`
}

// NewGetChatroomOK creates GetChatroomOK with default headers values
func NewGetChatroomOK() *GetChatroomOK {

	return &GetChatroomOK{}
}

// WithPayload adds the payload to the get chatroom o k response
func (o *GetChatroomOK) WithPayload(payload *models.Chatroom) *GetChatroomOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get chatroom o k response
func (o *GetChatroomOK) SetPayload(payload *models.Chatroom) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChatroomOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetChatroomUnauthorizedCode is the HTTP code returned for type GetChatroomUnauthorized
const GetChatroomUnauthorizedCode int = 401

/*
GetChatroomUnauthorized Unauthorized

swagger:response getChatroomUnauthorized
*/
type GetChatroomUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetChatroomUnauthorized creates GetChatroomUnauthorized with default headers values
func NewGetChatroomUnauthorized() *GetChatroomUnauthorized {

	return &GetChatroomUnauthorized{}
}

// WithPayload adds the payload to the get chatroom unauthorized response
func (o *GetChatroomUnauthorized) WithPayload(payload *models.Error) *GetChatroomUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get chatroom unauthorized response
func (o *GetChatroomUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChatroomUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetChatroomNotFoundCode is the HTTP code returned for type GetChatroomNotFound
const GetChatroomNotFoundCode int = 404

/*
GetChatroomNotFound Not Found

swagger:response getChatroomNotFound
*/
type GetChatroomNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetChatroomNotFound creates GetChatroomNotFound with default headers values
func NewGetChatroomNotFound() *GetChatroomNotFound {

	return &GetChatroomNotFound{}
}

// WithPayload adds the payload to the get chatroom not found response
func (o *GetChatroomNotFound) WithPayload(payload *models.Error) *GetChatroomNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get chatroom not found response
func (o *GetChatroomNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChatroomNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
GetChatroomDefault Internal Server Error

swagger:response getChatroomDefault
*/
type GetChatroomDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetChatroomDefault creates GetChatroomDefault with default headers values
func NewGetChatroomDefault(code int) *GetChatroomDefault {
	if code <= 0 {
		code = 500
	}

	return &GetChatroomDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get chatroom default response
func (o *GetChatroomDefault) WithStatusCode(code int) *GetChatroomDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get chatroom default response
func (o *GetChatroomDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get chatroom default response
func (o *GetChatroomDefault) WithPayload(payload *models.Error) *GetChatroomDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get chatroom default response
func (o *GetChatroomDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChatroomDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
