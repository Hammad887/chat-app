// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/Hammad887/chat-app/gen/models"
)

// NewChatRoomParams creates a new ChatRoomParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewChatRoomParams() *ChatRoomParams {
	return &ChatRoomParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewChatRoomParamsWithTimeout creates a new ChatRoomParams object
// with the ability to set a timeout on a request.
func NewChatRoomParamsWithTimeout(timeout time.Duration) *ChatRoomParams {
	return &ChatRoomParams{
		timeout: timeout,
	}
}

// NewChatRoomParamsWithContext creates a new ChatRoomParams object
// with the ability to set a context for a request.
func NewChatRoomParamsWithContext(ctx context.Context) *ChatRoomParams {
	return &ChatRoomParams{
		Context: ctx,
	}
}

// NewChatRoomParamsWithHTTPClient creates a new ChatRoomParams object
// with the ability to set a custom HTTPClient for a request.
func NewChatRoomParamsWithHTTPClient(client *http.Client) *ChatRoomParams {
	return &ChatRoomParams{
		HTTPClient: client,
	}
}

/*
ChatRoomParams contains all the parameters to send to the API endpoint

	for the chat room operation.

	Typically these are written to a http.Request.
*/
type ChatRoomParams struct {

	/* Chatroom.

	   create chatroom Payload
	*/
	Chatroom *models.ChatroomInfo

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the chat room params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChatRoomParams) WithDefaults() *ChatRoomParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the chat room params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChatRoomParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the chat room params
func (o *ChatRoomParams) WithTimeout(timeout time.Duration) *ChatRoomParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the chat room params
func (o *ChatRoomParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the chat room params
func (o *ChatRoomParams) WithContext(ctx context.Context) *ChatRoomParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the chat room params
func (o *ChatRoomParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the chat room params
func (o *ChatRoomParams) WithHTTPClient(client *http.Client) *ChatRoomParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the chat room params
func (o *ChatRoomParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChatroom adds the chatroom to the chat room params
func (o *ChatRoomParams) WithChatroom(chatroom *models.ChatroomInfo) *ChatRoomParams {
	o.SetChatroom(chatroom)
	return o
}

// SetChatroom adds the chatroom to the chat room params
func (o *ChatRoomParams) SetChatroom(chatroom *models.ChatroomInfo) {
	o.Chatroom = chatroom
}

// WriteToRequest writes these params to a swagger request
func (o *ChatRoomParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Chatroom != nil {
		if err := r.SetBodyParam(o.Chatroom); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
