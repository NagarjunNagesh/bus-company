package fetchall

import (
	"github.com/NagarjunNagesh/bus-company/restapi/operations/trip"
	"github.com/go-openapi/runtime/middleware"
)

type RestHandler interface {
	FetchAllTripHandler(params trip.GetAllTripsParams) middleware.Responder
}
