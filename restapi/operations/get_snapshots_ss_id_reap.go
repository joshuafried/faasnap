// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetSnapshotsSsIDReapHandlerFunc turns a function with the right signature into a get snapshots ss ID reap handler
type GetSnapshotsSsIDReapHandlerFunc func(GetSnapshotsSsIDReapParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSnapshotsSsIDReapHandlerFunc) Handle(params GetSnapshotsSsIDReapParams) middleware.Responder {
	return fn(params)
}

// GetSnapshotsSsIDReapHandler interface for that can handle valid get snapshots ss ID reap params
type GetSnapshotsSsIDReapHandler interface {
	Handle(GetSnapshotsSsIDReapParams) middleware.Responder
}

// NewGetSnapshotsSsIDReap creates a new http.Handler for the get snapshots ss ID reap operation
func NewGetSnapshotsSsIDReap(ctx *middleware.Context, handler GetSnapshotsSsIDReapHandler) *GetSnapshotsSsIDReap {
	return &GetSnapshotsSsIDReap{Context: ctx, Handler: handler}
}

/*
	GetSnapshotsSsIDReap swagger:route GET /snapshots/{ssId}/reap getSnapshotsSsIdReap

get reap state
*/
type GetSnapshotsSsIDReap struct {
	Context *middleware.Context
	Handler GetSnapshotsSsIDReapHandler
}

func (o *GetSnapshotsSsIDReap) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetSnapshotsSsIDReapParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetSnapshotsSsIDReapBadRequestBody get snapshots ss ID reap bad request body
//
// swagger:model GetSnapshotsSsIDReapBadRequestBody
type GetSnapshotsSsIDReapBadRequestBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get snapshots ss ID reap bad request body
func (o *GetSnapshotsSsIDReapBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get snapshots ss ID reap bad request body based on context it is used
func (o *GetSnapshotsSsIDReapBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetSnapshotsSsIDReapBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSnapshotsSsIDReapBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetSnapshotsSsIDReapBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
