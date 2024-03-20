/*
Copyright 2023 Sangfor Technologies Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteFaultsIDHandlerFunc turns a function with the right signature into a delete faults ID handler
type DeleteFaultsIDHandlerFunc func(DeleteFaultsIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteFaultsIDHandlerFunc) Handle(params DeleteFaultsIDParams) middleware.Responder {
	return fn(params)
}

// DeleteFaultsIDHandler interface for that can handle valid delete faults ID params
type DeleteFaultsIDHandler interface {
	Handle(DeleteFaultsIDParams) middleware.Responder
}

// NewDeleteFaultsID creates a new http.Handler for the delete faults ID operation
func NewDeleteFaultsID(ctx *middleware.Context, handler DeleteFaultsIDHandler) *DeleteFaultsID {
	return &DeleteFaultsID{Context: ctx, Handler: handler}
}

/*
	DeleteFaultsID swagger:route DELETE /faults/{id} deleteFaultsId

Delete the specific fault.
*/
type DeleteFaultsID struct {
	Context *middleware.Context
	Handler DeleteFaultsIDHandler
}

func (o *DeleteFaultsID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteFaultsIDParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}