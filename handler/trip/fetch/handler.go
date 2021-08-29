package fetch

import (
	trip_entity "github.com/NagarjunNagesh/bus-company/domain/models/trip"
	fetch_one_trip "github.com/NagarjunNagesh/bus-company/domain/usecase/fetch"
	"github.com/NagarjunNagesh/bus-company/models"
	"github.com/NagarjunNagesh/bus-company/restapi/operations/trip"
	"github.com/go-openapi/runtime/middleware"
)

type restHandler struct {
	fetch_one_trip_uc fetch_one_trip.UseCase
}

func NewHandler(fetch_one_trip_uc fetch_one_trip.UseCase) RestHandler {
	return &restHandler{fetch_one_trip_uc: fetch_one_trip_uc}
}

func (h *restHandler) FetchATripHandler(params trip.GetTripByIDParams) middleware.Responder {
	tripEntity, err := h.fetch_one_trip_uc.FetchOneTrip(params.TripID)
	if err != nil {
		return return400()
	}

	shouldReturn, returnValue := return200(tripEntity)
	if shouldReturn {
		return returnValue
	}

	return return404()
}

func return200(tripEntity *trip_entity.Trip) (bool, middleware.Responder) {
	if tripEntity != nil {
		tripModel := models.Trip{
			Dates:       tripEntity.Dates,
			Destination: tripEntity.Destination,
			Origin:      tripEntity.Origin,
			Price:       tripEntity.Price,
		}
		return true, trip.NewGetTripByIDOK().WithPayload(&tripModel)
	}
	return false, nil
}

func return404() middleware.Responder {
	apiResponse := models.APIResponse{
		Code:    404,
		Message: "trip does not exist",
		Type:    "JSON",
	}
	return trip.NewGetAllTripsNotFound().WithPayload(&apiResponse)
}

func return400() middleware.Responder {
	apiResponse := models.APIResponse{
		Code:    404,
		Message: "Unable to fetch the trip",
		Type:    "JSON",
	}
	return trip.NewGetAllTripsBadRequest().WithPayload(&apiResponse)
}
