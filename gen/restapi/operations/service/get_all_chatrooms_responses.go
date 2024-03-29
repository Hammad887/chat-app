// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Hammad887/chat-app/gen/models"
)

// GetAllChatroomsOKCode is the HTTP code returned for type GetAllChatroomsOK
const GetAllChatroomsOKCode int = 200

/*
GetAllChatroomsOK successfully save user object into database

swagger:response getAllChatroomsOK
*/
type GetAllChatroomsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Chatroom `json:"body,omitempty"`
}

// NewGetAllChatroomsOK creates GetAllChatroomsOK with default headers values
func NewGetAllChatroomsOK() *GetAllChatroomsOK {

	return &GetAllChatroomsOK{}
}

// WithPayload adds the payload to the get all chatrooms o k response
func (o *GetAllChatroomsOK) WithPayload(payload []*models.Chatroom) *GetAllChatroomsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all chatrooms o k response
func (o *GetAllChatroomsOK) SetPayload(payload []*models.Chatroom) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllChatroomsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Chatroom, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetAllChatroomsNoContentCode is the HTTP code returned for type GetAllChatroomsNoContent
const GetAllChatroomsNoContentCode int = 204

/*
GetAllChatroomsNoContent no content in database

swagger:response getAllChatroomsNoContent
*/
type GetAllChatroomsNoContent struct {
}

// NewGetAllChatroomsNoContent creates GetAllChatroomsNoContent with default headers values
func NewGetAllChatroomsNoContent() *GetAllChatroomsNoContent {

	return &GetAllChatroomsNoContent{}
}

// WriteResponse to the client
func (o *GetAllChatroomsNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// GetAllChatroomsUnauthorizedCode is the HTTP code returned for type GetAllChatroomsUnauthorized
const GetAllChatroomsUnauthorizedCode int = 401

/*
GetAllChatroomsUnauthorized Unauthorized

swagger:response getAllChatroomsUnauthorized
*/
type GetAllChatroomsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAllChatroomsUnauthorized creates GetAllChatroomsUnauthorized with default headers values
func NewGetAllChatroomsUnauthorized() *GetAllChatroomsUnauthorized {

	return &GetAllChatroomsUnauthorized{}
}

// WithPayload adds the payload to the get all chatrooms unauthorized response
func (o *GetAllChatroomsUnauthorized) WithPayload(payload *models.Error) *GetAllChatroomsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all chatrooms unauthorized response
func (o *GetAllChatroomsUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllChatroomsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAllChatroomsNotFoundCode is the HTTP code returned for type GetAllChatroomsNotFound
const GetAllChatroomsNotFoundCode int = 404

/*
GetAllChatroomsNotFound Not Found

swagger:response getAllChatroomsNotFound
*/
type GetAllChatroomsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAllChatroomsNotFound creates GetAllChatroomsNotFound with default headers values
func NewGetAllChatroomsNotFound() *GetAllChatroomsNotFound {

	return &GetAllChatroomsNotFound{}
}

// WithPayload adds the payload to the get all chatrooms not found response
func (o *GetAllChatroomsNotFound) WithPayload(payload *models.Error) *GetAllChatroomsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all chatrooms not found response
func (o *GetAllChatroomsNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllChatroomsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
GetAllChatroomsDefault Internal Server Error

swagger:response getAllChatroomsDefault
*/
type GetAllChatroomsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAllChatroomsDefault creates GetAllChatroomsDefault with default headers values
func NewGetAllChatroomsDefault(code int) *GetAllChatroomsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetAllChatroomsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get all chatrooms default response
func (o *GetAllChatroomsDefault) WithStatusCode(code int) *GetAllChatroomsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get all chatrooms default response
func (o *GetAllChatroomsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get all chatrooms default response
func (o *GetAllChatroomsDefault) WithPayload(payload *models.Error) *GetAllChatroomsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all chatrooms default response
func (o *GetAllChatroomsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllChatroomsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
