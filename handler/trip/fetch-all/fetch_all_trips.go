package fetch

import (
	"github.com/NagarjunNagesh/bus-company/models"
	"github.com/NagarjunNagesh/bus-company/restapi/operations/trip"
	"github.com/go-openapi/runtime/middleware"
)

func FetchAllTripHandler(params trip.GetAllTripsParams) middleware.Responder {
	dates := "Mon Tue Wed Thurs"
	barcelona := "Barcelona"
	price := float32(40.22)
	origin := "Seville"
	aTrip := models.Trip{
		Dates:       &dates,
		Destination: &barcelona,
		Origin:      &origin,
		Price:       &price,
	}
	trips := []*models.Trip{}
	trips = append(trips, &aTrip)
	return trip.NewGetAllTripsOK().WithPayload(trips)
}
