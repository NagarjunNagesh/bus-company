package add

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/NagarjunNagesh/bus-company/restapi/operations/trip"
	"github.com/NagarjunNagesh/bus-company/models"
)

func AddATripHandler(params trip.AddNewTripParams) middleware.Responder {
	apiResponse := models.APIResponse {
		Code: 201,
		Message: "Successfully created a new trip",
		Type: "JSON",
	}
	return trip.NewAddNewTripCreated().WithPayload(&apiResponse)
}