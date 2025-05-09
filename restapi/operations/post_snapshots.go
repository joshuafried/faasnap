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

// PostSnapshotsHandlerFunc turns a function with the right signature into a post snapshots handler
type PostSnapshotsHandlerFunc func(PostSnapshotsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostSnapshotsHandlerFunc) Handle(params PostSnapshotsParams) middleware.Responder {
	return fn(params)
}

// PostSnapshotsHandler interface for that can handle valid post snapshots params
type PostSnapshotsHandler interface {
	Handle(PostSnapshotsParams) middleware.Responder
}

// NewPostSnapshots creates a new http.Handler for the post snapshots operation
func NewPostSnapshots(ctx *middleware.Context, handler PostSnapshotsHandler) *PostSnapshots {
	return &PostSnapshots{Context: ctx, Handler: handler}
}

/*
	PostSnapshots swagger:route POST /snapshots postSnapshots

Take a snapshot
*/
type PostSnapshots struct {
	Context *middleware.Context
	Handler PostSnapshotsHandler
}

func (o *PostSnapshots) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostSnapshotsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostSnapshotsBadRequestBody post snapshots bad request body
//
// swagger:model PostSnapshotsBadRequestBody
type PostSnapshotsBadRequestBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this post snapshots bad request body
func (o *PostSnapshotsBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post snapshots bad request body based on context it is used
func (o *PostSnapshotsBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostSnapshotsBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostSnapshotsBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PostSnapshotsBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
