// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostSnapshotsSsIDMincoreOKCode is the HTTP code returned for type PostSnapshotsSsIDMincoreOK
const PostSnapshotsSsIDMincoreOKCode int = 200

/*
PostSnapshotsSsIDMincoreOK OK

swagger:response postSnapshotsSsIdMincoreOK
*/
type PostSnapshotsSsIDMincoreOK struct {
}

// NewPostSnapshotsSsIDMincoreOK creates PostSnapshotsSsIDMincoreOK with default headers values
func NewPostSnapshotsSsIDMincoreOK() *PostSnapshotsSsIDMincoreOK {

	return &PostSnapshotsSsIDMincoreOK{}
}

// WriteResponse to the client
func (o *PostSnapshotsSsIDMincoreOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PostSnapshotsSsIDMincoreBadRequestCode is the HTTP code returned for type PostSnapshotsSsIDMincoreBadRequest
const PostSnapshotsSsIDMincoreBadRequestCode int = 400

/*
PostSnapshotsSsIDMincoreBadRequest Invalid request

swagger:response postSnapshotsSsIdMincoreBadRequest
*/
type PostSnapshotsSsIDMincoreBadRequest struct {

	/*
	  In: Body
	*/
	Payload *PostSnapshotsSsIDMincoreBadRequestBody `json:"body,omitempty"`
}

// NewPostSnapshotsSsIDMincoreBadRequest creates PostSnapshotsSsIDMincoreBadRequest with default headers values
func NewPostSnapshotsSsIDMincoreBadRequest() *PostSnapshotsSsIDMincoreBadRequest {

	return &PostSnapshotsSsIDMincoreBadRequest{}
}

// WithPayload adds the payload to the post snapshots ss Id mincore bad request response
func (o *PostSnapshotsSsIDMincoreBadRequest) WithPayload(payload *PostSnapshotsSsIDMincoreBadRequestBody) *PostSnapshotsSsIDMincoreBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post snapshots ss Id mincore bad request response
func (o *PostSnapshotsSsIDMincoreBadRequest) SetPayload(payload *PostSnapshotsSsIDMincoreBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostSnapshotsSsIDMincoreBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
