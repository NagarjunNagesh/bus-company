package fetch

import (
	"github.com/NagarjunNagesh/bus-company/models"
	"github.com/NagarjunNagesh/bus-company/restapi/operations/trip"
	"github.com/go-openapi/runtime/middleware"
)

func FetchATripHandler(params trip.GetTripByIDParams) middleware.Responder {
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
	return trip.NewGetTripByIDOK().WithPayload(&aTrip)
}
