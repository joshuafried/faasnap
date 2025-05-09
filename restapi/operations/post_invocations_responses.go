// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostInvocationsOKCode is the HTTP code returned for type PostInvocationsOK
const PostInvocationsOKCode int = 200

/*
PostInvocationsOK Success

swagger:response postInvocationsOK
*/
type PostInvocationsOK struct {

	/*
	  In: Body
	*/
	Payload *PostInvocationsOKBody `json:"body,omitempty"`
}

// NewPostInvocationsOK creates PostInvocationsOK with default headers values
func NewPostInvocationsOK() *PostInvocationsOK {

	return &PostInvocationsOK{}
}

// WithPayload adds the payload to the post invocations o k response
func (o *PostInvocationsOK) WithPayload(payload *PostInvocationsOKBody) *PostInvocationsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post invocations o k response
func (o *PostInvocationsOK) SetPayload(payload *PostInvocationsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostInvocationsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostInvocationsBadRequestCode is the HTTP code returned for type PostInvocationsBadRequest
const PostInvocationsBadRequestCode int = 400

/*
PostInvocationsBadRequest Invalid request

swagger:response postInvocationsBadRequest
*/
type PostInvocationsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *PostInvocationsBadRequestBody `json:"body,omitempty"`
}

// NewPostInvocationsBadRequest creates PostInvocationsBadRequest with default headers values
func NewPostInvocationsBadRequest() *PostInvocationsBadRequest {

	return &PostInvocationsBadRequest{}
}

// WithPayload adds the payload to the post invocations bad request response
func (o *PostInvocationsBadRequest) WithPayload(payload *PostInvocationsBadRequestBody) *PostInvocationsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post invocations bad request response
func (o *PostInvocationsBadRequest) SetPayload(payload *PostInvocationsBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostInvocationsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
