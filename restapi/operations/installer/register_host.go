// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// RegisterHostHandlerFunc turns a function with the right signature into a register host handler
type RegisterHostHandlerFunc func(RegisterHostParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn RegisterHostHandlerFunc) Handle(params RegisterHostParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// RegisterHostHandler interface for that can handle valid register host params
type RegisterHostHandler interface {
	Handle(RegisterHostParams, interface{}) middleware.Responder
}

// NewRegisterHost creates a new http.Handler for the register host operation
func NewRegisterHost(ctx *middleware.Context, handler RegisterHostHandler) *RegisterHost {
	return &RegisterHost{Context: ctx, Handler: handler}
}

/*RegisterHost swagger:route POST /clusters/{cluster_id}/hosts installer registerHost

Registers a new OpenShift host.

*/
type RegisterHost struct {
	Context *middleware.Context
	Handler RegisterHostHandler
}

func (o *RegisterHost) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRegisterHostParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
