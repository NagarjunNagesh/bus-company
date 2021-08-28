package add

import (
	trip_model "github.com/NagarjunNagesh/bus-company/domain/models/trip"
	add_a_trip "github.com/NagarjunNagesh/bus-company/domain/usecase/add-a-trip"
	"github.com/NagarjunNagesh/bus-company/models"
	"github.com/NagarjunNagesh/bus-company/restapi/operations/trip"
	"github.com/go-openapi/runtime/middleware"
)

type restHandler struct {
	add_a_trip_uc add_a_trip.UseCase
}

func NewHandler(add_a_trip_uc add_a_trip.UseCase) RestHandler {
	return &restHandler{add_a_trip_uc: add_a_trip_uc}
}

func (h *restHandler) AddATripHandler(params trip.AddNewTripParams) middleware.Responder {
	body := params.Body
	tripModel := trip_model.AddTrip{
		OriginID:      *body.OriginID,
		DestinationID: *body.DestinationID,
		Dates:         *body.Dates,
		Price:         *body.Price,
	}
	apiResponse := models.APIResponse{
		Message: "Successfully created a new trip",
		Type:    "JSON",
	}

	isAddedSuccessfully := h.add_a_trip_uc.AddATrip(&tripModel)
	if isAddedSuccessfully {
		apiResponse.Code = 201
		return trip.NewAddNewTripCreated().WithPayload(&apiResponse)
	}

	return return404(apiResponse)
}

func return404(apiResponse models.APIResponse) middleware.Responder {
	apiResponse.Code = 404
	apiResponse.Message = "Unable to create a new trip"
	return trip.NewGetAllTripsNotFound().WithPayload(&apiResponse)
}
