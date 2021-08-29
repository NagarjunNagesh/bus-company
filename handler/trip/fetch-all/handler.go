package fetchall

import (
	trip_entity "github.com/NagarjunNagesh/bus-company/domain/models/trip"
	fetch_all_trips "github.com/NagarjunNagesh/bus-company/domain/usecase/fetch-all"
	"github.com/NagarjunNagesh/bus-company/models"
	"github.com/NagarjunNagesh/bus-company/restapi/operations/trip"
	"github.com/go-openapi/runtime/middleware"
)

type restHandler struct {
	fetch_all_trips_uc fetch_all_trips.UseCase
}

func NewHandler(fetch_all_trips_uc fetch_all_trips.UseCase) RestHandler {
	return &restHandler{fetch_all_trips_uc: fetch_all_trips_uc}
}

func (h *restHandler) FetchAllTripHandler(params trip.GetAllTripsParams) middleware.Responder {
	trips, err := h.fetch_all_trips_uc.FetchAllTrips()
	if err != nil {
		return return400(err)
	}

	shouldReturn, returnValue := return404(trips)
	if shouldReturn {
		return returnValue
	}

	return return200(trips)
}

func return200(trips []*trip_entity.Trip) middleware.Responder {
	tripsModel := models.Trips{}
	for _, t := range trips {
		tripModel := models.Trip{
			Dates:       t.Dates,
			Destination: t.Destination,
			Origin:      t.Origin,
			Price:       t.Price,
		}
		tripsModel = append(tripsModel, &tripModel)
	}
	return trip.NewGetAllTripsOK().WithPayload(tripsModel)
}

func return404(trips []*trip_entity.Trip) (bool, middleware.Responder) {
	if len(trips) == 0 {
		apiResponse := models.APIResponse{
			Code:    404,
			Message: "Unable to fetch all the trips",
			Type:    "JSON",
		}
		return true, trip.NewGetAllTripsNotFound().WithPayload(&apiResponse)
	}
	return false, nil
}

func return400(err error) middleware.Responder {
	apiResponse := models.APIResponse{
		Code:    400,
		Message: err.Error(),
		Type:    "JSON",
	}
	return trip.NewGetAllTripsBadRequest().WithPayload(&apiResponse)
}
