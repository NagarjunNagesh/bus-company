// Code generated by go-swagger; DO NOT EDIT.

package trip

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/NagarjunNagesh/bus-company/models"
)

// GetAllTripsOKCode is the HTTP code returned for type GetAllTripsOK
const GetAllTripsOKCode int = 200

/*GetAllTripsOK successful operation

swagger:response getAllTripsOK
*/
type GetAllTripsOK struct {

	/*
	  In: Body
	*/
	Payload models.Trips `json:"body,omitempty"`
}

// NewGetAllTripsOK creates GetAllTripsOK with default headers values
func NewGetAllTripsOK() *GetAllTripsOK {

	return &GetAllTripsOK{}
}

// WithPayload adds the payload to the get all trips o k response
func (o *GetAllTripsOK) WithPayload(payload models.Trips) *GetAllTripsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all trips o k response
func (o *GetAllTripsOK) SetPayload(payload models.Trips) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllTripsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.Trips{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetAllTripsBadRequestCode is the HTTP code returned for type GetAllTripsBadRequest
const GetAllTripsBadRequestCode int = 400

/*GetAllTripsBadRequest Bad Request

swagger:response getAllTripsBadRequest
*/
type GetAllTripsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetAllTripsBadRequest creates GetAllTripsBadRequest with default headers values
func NewGetAllTripsBadRequest() *GetAllTripsBadRequest {

	return &GetAllTripsBadRequest{}
}

// WithPayload adds the payload to the get all trips bad request response
func (o *GetAllTripsBadRequest) WithPayload(payload *models.APIResponse) *GetAllTripsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all trips bad request response
func (o *GetAllTripsBadRequest) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllTripsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAllTripsNotFoundCode is the HTTP code returned for type GetAllTripsNotFound
const GetAllTripsNotFoundCode int = 404

/*GetAllTripsNotFound No trips found

swagger:response getAllTripsNotFound
*/
type GetAllTripsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetAllTripsNotFound creates GetAllTripsNotFound with default headers values
func NewGetAllTripsNotFound() *GetAllTripsNotFound {

	return &GetAllTripsNotFound{}
}

// WithPayload adds the payload to the get all trips not found response
func (o *GetAllTripsNotFound) WithPayload(payload *models.APIResponse) *GetAllTripsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all trips not found response
func (o *GetAllTripsNotFound) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllTripsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
