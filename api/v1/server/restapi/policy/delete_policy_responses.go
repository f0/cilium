package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/cilium/api/v1/models"
)

// HTTP code for type DeletePolicyNoContent
const DeletePolicyNoContentCode int = 204

/*DeletePolicyNoContent Success

swagger:response deletePolicyNoContent
*/
type DeletePolicyNoContent struct {
}

// NewDeletePolicyNoContent creates DeletePolicyNoContent with default headers values
func NewDeletePolicyNoContent() *DeletePolicyNoContent {
	return &DeletePolicyNoContent{}
}

// WriteResponse to the client
func (o *DeletePolicyNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

// HTTP code for type DeletePolicyInvalid
const DeletePolicyInvalidCode int = 400

/*DeletePolicyInvalid Invalid request

swagger:response deletePolicyInvalid
*/
type DeletePolicyInvalid struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewDeletePolicyInvalid creates DeletePolicyInvalid with default headers values
func NewDeletePolicyInvalid() *DeletePolicyInvalid {
	return &DeletePolicyInvalid{}
}

// WithPayload adds the payload to the delete policy invalid response
func (o *DeletePolicyInvalid) WithPayload(payload models.Error) *DeletePolicyInvalid {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete policy invalid response
func (o *DeletePolicyInvalid) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeletePolicyInvalid) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// HTTP code for type DeletePolicyNotFound
const DeletePolicyNotFoundCode int = 404

/*DeletePolicyNotFound Policy tree not found

swagger:response deletePolicyNotFound
*/
type DeletePolicyNotFound struct {
}

// NewDeletePolicyNotFound creates DeletePolicyNotFound with default headers values
func NewDeletePolicyNotFound() *DeletePolicyNotFound {
	return &DeletePolicyNotFound{}
}

// WriteResponse to the client
func (o *DeletePolicyNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

// HTTP code for type DeletePolicyFailure
const DeletePolicyFailureCode int = 500

/*DeletePolicyFailure Error while deleting policy

swagger:response deletePolicyFailure
*/
type DeletePolicyFailure struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewDeletePolicyFailure creates DeletePolicyFailure with default headers values
func NewDeletePolicyFailure() *DeletePolicyFailure {
	return &DeletePolicyFailure{}
}

// WithPayload adds the payload to the delete policy failure response
func (o *DeletePolicyFailure) WithPayload(payload models.Error) *DeletePolicyFailure {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete policy failure response
func (o *DeletePolicyFailure) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeletePolicyFailure) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}