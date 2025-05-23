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

// PatchSnapshotsSsIDReapHandlerFunc turns a function with the right signature into a patch snapshots ss ID reap handler
type PatchSnapshotsSsIDReapHandlerFunc func(PatchSnapshotsSsIDReapParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchSnapshotsSsIDReapHandlerFunc) Handle(params PatchSnapshotsSsIDReapParams) middleware.Responder {
	return fn(params)
}

// PatchSnapshotsSsIDReapHandler interface for that can handle valid patch snapshots ss ID reap params
type PatchSnapshotsSsIDReapHandler interface {
	Handle(PatchSnapshotsSsIDReapParams) middleware.Responder
}

// NewPatchSnapshotsSsIDReap creates a new http.Handler for the patch snapshots ss ID reap operation
func NewPatchSnapshotsSsIDReap(ctx *middleware.Context, handler PatchSnapshotsSsIDReapHandler) *PatchSnapshotsSsIDReap {
	return &PatchSnapshotsSsIDReap{Context: ctx, Handler: handler}
}

/*
	PatchSnapshotsSsIDReap swagger:route PATCH /snapshots/{ssId}/reap patchSnapshotsSsIdReap

Change reap state
*/
type PatchSnapshotsSsIDReap struct {
	Context *middleware.Context
	Handler PatchSnapshotsSsIDReapHandler
}

func (o *PatchSnapshotsSsIDReap) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPatchSnapshotsSsIDReapParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PatchSnapshotsSsIDReapBadRequestBody patch snapshots ss ID reap bad request body
//
// swagger:model PatchSnapshotsSsIDReapBadRequestBody
type PatchSnapshotsSsIDReapBadRequestBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this patch snapshots ss ID reap bad request body
func (o *PatchSnapshotsSsIDReapBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this patch snapshots ss ID reap bad request body based on context it is used
func (o *PatchSnapshotsSsIDReapBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PatchSnapshotsSsIDReapBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchSnapshotsSsIDReapBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PatchSnapshotsSsIDReapBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
