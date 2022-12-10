// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// RestWalletImportHandlerFunc turns a function with the right signature into a rest wallet import handler
type RestWalletImportHandlerFunc func(RestWalletImportParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RestWalletImportHandlerFunc) Handle(params RestWalletImportParams) middleware.Responder {
	return fn(params)
}

// RestWalletImportHandler interface for that can handle valid rest wallet import params
type RestWalletImportHandler interface {
	Handle(RestWalletImportParams) middleware.Responder
}

// NewRestWalletImport creates a new http.Handler for the rest wallet import operation
func NewRestWalletImport(ctx *middleware.Context, handler RestWalletImportHandler) *RestWalletImport {
	return &RestWalletImport{Context: ctx, Handler: handler}
}

/* RestWalletImport swagger:route PUT /rest/wallet restWalletImport

RestWalletImport rest wallet import API

*/
type RestWalletImport struct {
	Context *middleware.Context
	Handler RestWalletImportHandler
}

func (o *RestWalletImport) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewRestWalletImportParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}