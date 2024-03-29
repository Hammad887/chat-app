// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// LogoutHandlerFunc turns a function with the right signature into a logout handler
type LogoutHandlerFunc func(LogoutParams) middleware.Responder

// Handle executing the request and returning a response
func (fn LogoutHandlerFunc) Handle(params LogoutParams) middleware.Responder {
	return fn(params)
}

// LogoutHandler interface for that can handle valid logout params
type LogoutHandler interface {
	Handle(LogoutParams) middleware.Responder
}

// NewLogout creates a new http.Handler for the logout operation
func NewLogout(ctx *middleware.Context, handler LogoutHandler) *Logout {
	return &Logout{Context: ctx, Handler: handler}
}

/*
	Logout swagger:route POST /logout service logout

logout a user
*/
type Logout struct {
	Context *middleware.Context
	Handler LogoutHandler
}

func (o *Logout) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewLogoutParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
