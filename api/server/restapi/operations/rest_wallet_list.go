// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// RestWalletListHandlerFunc turns a function with the right signature into a rest wallet list handler
type RestWalletListHandlerFunc func(RestWalletListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RestWalletListHandlerFunc) Handle(params RestWalletListParams) middleware.Responder {
	return fn(params)
}

// RestWalletListHandler interface for that can handle valid rest wallet list params
type RestWalletListHandler interface {
	Handle(RestWalletListParams) middleware.Responder
}

// NewRestWalletList creates a new http.Handler for the rest wallet list operation
func NewRestWalletList(ctx *middleware.Context, handler RestWalletListHandler) *RestWalletList {
	return &RestWalletList{Context: ctx, Handler: handler}
}

/* RestWalletList swagger:route GET /rest/wallet restWalletList

RestWalletList rest wallet list API

*/
type RestWalletList struct {
	Context *middleware.Context
	Handler RestWalletListHandler
}

func (o *RestWalletList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewRestWalletListParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}