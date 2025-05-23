// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewPatchSnapshotsSsIDMincoreParams creates a new PatchSnapshotsSsIDMincoreParams object
//
// There are no default values defined in the spec.
func NewPatchSnapshotsSsIDMincoreParams() PatchSnapshotsSsIDMincoreParams {

	return PatchSnapshotsSsIDMincoreParams{}
}

// PatchSnapshotsSsIDMincoreParams contains all the bound params for the patch snapshots ss ID mincore operation
// typically these are obtained from a http.Request
//
// swagger:parameters PatchSnapshotsSsIDMincore
type PatchSnapshotsSsIDMincoreParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	SsID string
	/*
	  Required: true
	  In: body
	*/
	State PatchSnapshotsSsIDMincoreBody
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPatchSnapshotsSsIDMincoreParams() beforehand.
func (o *PatchSnapshotsSsIDMincoreParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rSsID, rhkSsID, _ := route.Params.GetOK("ssId")
	if err := o.bindSsID(rSsID, rhkSsID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body PatchSnapshotsSsIDMincoreBody
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("state", "body", ""))
			} else {
				res = append(res, errors.NewParseError("state", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.State = body
			}
		}
	} else {
		res = append(res, errors.Required("state", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindSsID binds and validates parameter SsID from path.
func (o *PatchSnapshotsSsIDMincoreParams) bindSsID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.SsID = raw

	return nil
}
