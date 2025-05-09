// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewPutSnapshotsParams creates a new PutSnapshotsParams object
//
// There are no default values defined in the spec.
func NewPutSnapshotsParams() PutSnapshotsParams {

	return PutSnapshotsParams{}
}

// PutSnapshotsParams contains all the bound params for the put snapshots operation
// typically these are obtained from a http.Request
//
// swagger:parameters PutSnapshots
type PutSnapshotsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: query
	*/
	FromSnapshot string
	/*
	  Required: true
	  In: query
	*/
	MemFilePath string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPutSnapshotsParams() beforehand.
func (o *PutSnapshotsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qFromSnapshot, qhkFromSnapshot, _ := qs.GetOK("from_snapshot")
	if err := o.bindFromSnapshot(qFromSnapshot, qhkFromSnapshot, route.Formats); err != nil {
		res = append(res, err)
	}

	qMemFilePath, qhkMemFilePath, _ := qs.GetOK("mem_file_path")
	if err := o.bindMemFilePath(qMemFilePath, qhkMemFilePath, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindFromSnapshot binds and validates parameter FromSnapshot from query.
func (o *PutSnapshotsParams) bindFromSnapshot(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("from_snapshot", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("from_snapshot", "query", raw); err != nil {
		return err
	}
	o.FromSnapshot = raw

	return nil
}

// bindMemFilePath binds and validates parameter MemFilePath from query.
func (o *PutSnapshotsParams) bindMemFilePath(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("mem_file_path", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("mem_file_path", "query", raw); err != nil {
		return err
	}
	o.MemFilePath = raw

	return nil
}
