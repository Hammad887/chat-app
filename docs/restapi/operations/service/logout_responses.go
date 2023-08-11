// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Hammad887/chat-app/docs/models"
)

// LogoutOKCode is the HTTP code returned for type LogoutOK
const LogoutOKCode int = 200

/*
LogoutOK Successful logout

swagger:response logoutOK
*/
type LogoutOK struct {

	/*
	  In: Body
	*/
	Payload *models.LogoutSuccess `json:"body,omitempty"`
}

// NewLogoutOK creates LogoutOK with default headers values
func NewLogoutOK() *LogoutOK {

	return &LogoutOK{}
}

// WithPayload adds the payload to the logout o k response
func (o *LogoutOK) WithPayload(payload *models.LogoutSuccess) *LogoutOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the logout o k response
func (o *LogoutOK) SetPayload(payload *models.LogoutSuccess) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LogoutOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// LogoutBadRequestCode is the HTTP code returned for type LogoutBadRequest
const LogoutBadRequestCode int = 400

/*
LogoutBadRequest Bad Request

swagger:response logoutBadRequest
*/
type LogoutBadRequest struct {
}

// NewLogoutBadRequest creates LogoutBadRequest with default headers values
func NewLogoutBadRequest() *LogoutBadRequest {

	return &LogoutBadRequest{}
}

// WriteResponse to the client
func (o *LogoutBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// LogoutNotFoundCode is the HTTP code returned for type LogoutNotFound
const LogoutNotFoundCode int = 404

/*
LogoutNotFound User not found

swagger:response logoutNotFound
*/
type LogoutNotFound struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewLogoutNotFound creates LogoutNotFound with default headers values
func NewLogoutNotFound() *LogoutNotFound {

	return &LogoutNotFound{}
}

// WithPayload adds the payload to the logout not found response
func (o *LogoutNotFound) WithPayload(payload string) *LogoutNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the logout not found response
func (o *LogoutNotFound) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LogoutNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// LogoutInternalServerErrorCode is the HTTP code returned for type LogoutInternalServerError
const LogoutInternalServerErrorCode int = 500

/*
LogoutInternalServerError Server error

swagger:response logoutInternalServerError
*/
type LogoutInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewLogoutInternalServerError creates LogoutInternalServerError with default headers values
func NewLogoutInternalServerError() *LogoutInternalServerError {

	return &LogoutInternalServerError{}
}

// WithPayload adds the payload to the logout internal server error response
func (o *LogoutInternalServerError) WithPayload(payload string) *LogoutInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the logout internal server error response
func (o *LogoutInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LogoutInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
