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
)

// NewGetChatroomParams creates a new GetChatroomParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetChatroomParams() *GetChatroomParams {
	return &GetChatroomParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetChatroomParamsWithTimeout creates a new GetChatroomParams object
// with the ability to set a timeout on a request.
func NewGetChatroomParamsWithTimeout(timeout time.Duration) *GetChatroomParams {
	return &GetChatroomParams{
		timeout: timeout,
	}
}

// NewGetChatroomParamsWithContext creates a new GetChatroomParams object
// with the ability to set a context for a request.
func NewGetChatroomParamsWithContext(ctx context.Context) *GetChatroomParams {
	return &GetChatroomParams{
		Context: ctx,
	}
}

// NewGetChatroomParamsWithHTTPClient creates a new GetChatroomParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetChatroomParamsWithHTTPClient(client *http.Client) *GetChatroomParams {
	return &GetChatroomParams{
		HTTPClient: client,
	}
}

/*
GetChatroomParams contains all the parameters to send to the API endpoint

	for the get chatroom operation.

	Typically these are written to a http.Request.
*/
type GetChatroomParams struct {

	// ChatroomID.
	ChatroomID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get chatroom params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetChatroomParams) WithDefaults() *GetChatroomParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get chatroom params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetChatroomParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get chatroom params
func (o *GetChatroomParams) WithTimeout(timeout time.Duration) *GetChatroomParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get chatroom params
func (o *GetChatroomParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get chatroom params
func (o *GetChatroomParams) WithContext(ctx context.Context) *GetChatroomParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get chatroom params
func (o *GetChatroomParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get chatroom params
func (o *GetChatroomParams) WithHTTPClient(client *http.Client) *GetChatroomParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get chatroom params
func (o *GetChatroomParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChatroomID adds the chatroomID to the get chatroom params
func (o *GetChatroomParams) WithChatroomID(chatroomID string) *GetChatroomParams {
	o.SetChatroomID(chatroomID)
	return o
}

// SetChatroomID adds the chatroomId to the get chatroom params
func (o *GetChatroomParams) SetChatroomID(chatroomID string) {
	o.ChatroomID = chatroomID
}

// WriteToRequest writes these params to a swagger request
func (o *GetChatroomParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param chatroom_id
	if err := r.SetPathParam("chatroom_id", o.ChatroomID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
