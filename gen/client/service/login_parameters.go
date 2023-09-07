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

// NewLoginParams creates a new LoginParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLoginParams() *LoginParams {
	return &LoginParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLoginParamsWithTimeout creates a new LoginParams object
// with the ability to set a timeout on a request.
func NewLoginParamsWithTimeout(timeout time.Duration) *LoginParams {
	return &LoginParams{
		timeout: timeout,
	}
}

// NewLoginParamsWithContext creates a new LoginParams object
// with the ability to set a context for a request.
func NewLoginParamsWithContext(ctx context.Context) *LoginParams {
	return &LoginParams{
		Context: ctx,
	}
}

// NewLoginParamsWithHTTPClient creates a new LoginParams object
// with the ability to set a custom HTTPClient for a request.
func NewLoginParamsWithHTTPClient(client *http.Client) *LoginParams {
	return &LoginParams{
		HTTPClient: client,
	}
}

/*
LoginParams contains all the parameters to send to the API endpoint

	for the login operation.

	Typically these are written to a http.Request.
*/
type LoginParams struct {

	/* Login.

	   Login Payload
	*/
	Login *models.LoginInfo

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the login params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LoginParams) WithDefaults() *LoginParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the login params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LoginParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the login params
func (o *LoginParams) WithTimeout(timeout time.Duration) *LoginParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the login params
func (o *LoginParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the login params
func (o *LoginParams) WithContext(ctx context.Context) *LoginParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the login params
func (o *LoginParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the login params
func (o *LoginParams) WithHTTPClient(client *http.Client) *LoginParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the login params
func (o *LoginParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLogin adds the login to the login params
func (o *LoginParams) WithLogin(login *models.LoginInfo) *LoginParams {
	o.SetLogin(login)
	return o
}

// SetLogin adds the login to the login params
func (o *LoginParams) SetLogin(login *models.LoginInfo) {
	o.Login = login
}

// WriteToRequest writes these params to a swagger request
func (o *LoginParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Login != nil {
		if err := r.SetBodyParam(o.Login); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
