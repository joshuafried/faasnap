// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PatchSnapshotsSsIDReapOKCode is the HTTP code returned for type PatchSnapshotsSsIDReapOK
const PatchSnapshotsSsIDReapOKCode int = 200

/*
PatchSnapshotsSsIDReapOK OK

swagger:response patchSnapshotsSsIdReapOK
*/
type PatchSnapshotsSsIDReapOK struct {
}

// NewPatchSnapshotsSsIDReapOK creates PatchSnapshotsSsIDReapOK with default headers values
func NewPatchSnapshotsSsIDReapOK() *PatchSnapshotsSsIDReapOK {

	return &PatchSnapshotsSsIDReapOK{}
}

// WriteResponse to the client
func (o *PatchSnapshotsSsIDReapOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PatchSnapshotsSsIDReapBadRequestCode is the HTTP code returned for type PatchSnapshotsSsIDReapBadRequest
const PatchSnapshotsSsIDReapBadRequestCode int = 400

/*
PatchSnapshotsSsIDReapBadRequest Invalid request

swagger:response patchSnapshotsSsIdReapBadRequest
*/
type PatchSnapshotsSsIDReapBadRequest struct {

	/*
	  In: Body
	*/
	Payload *PatchSnapshotsSsIDReapBadRequestBody `json:"body,omitempty"`
}

// NewPatchSnapshotsSsIDReapBadRequest creates PatchSnapshotsSsIDReapBadRequest with default headers values
func NewPatchSnapshotsSsIDReapBadRequest() *PatchSnapshotsSsIDReapBadRequest {

	return &PatchSnapshotsSsIDReapBadRequest{}
}

// WithPayload adds the payload to the patch snapshots ss Id reap bad request response
func (o *PatchSnapshotsSsIDReapBadRequest) WithPayload(payload *PatchSnapshotsSsIDReapBadRequestBody) *PatchSnapshotsSsIDReapBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch snapshots ss Id reap bad request response
func (o *PatchSnapshotsSsIDReapBadRequest) SetPayload(payload *PatchSnapshotsSsIDReapBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchSnapshotsSsIDReapBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
