package add

import (
	"github.com/NagarjunNagesh/bus-company/restapi/operations/trip"
	"github.com/go-openapi/runtime/middleware"
)

type RestHandler interface {
	AddATripHandler(params trip.AddNewTripParams) middleware.Responder
}
