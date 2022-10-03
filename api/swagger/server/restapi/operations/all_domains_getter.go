// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AllDomainsGetterHandlerFunc turns a function with the right signature into a all domains getter handler
type AllDomainsGetterHandlerFunc func(AllDomainsGetterParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AllDomainsGetterHandlerFunc) Handle(params AllDomainsGetterParams) middleware.Responder {
	return fn(params)
}

// AllDomainsGetterHandler interface for that can handle valid all domains getter params
type AllDomainsGetterHandler interface {
	Handle(AllDomainsGetterParams) middleware.Responder
}

// NewAllDomainsGetter creates a new http.Handler for the all domains getter operation
func NewAllDomainsGetter(ctx *middleware.Context, handler AllDomainsGetterHandler) *AllDomainsGetter {
	return &AllDomainsGetter{Context: ctx, Handler: handler}
}

/* AllDomainsGetter swagger:route GET /all/domains allDomainsGetter

AllDomainsGetter all domains getter API

*/
type AllDomainsGetter struct {
	Context *middleware.Context
	Handler AllDomainsGetterHandler
}

func (o *AllDomainsGetter) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAllDomainsGetterParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}