package fetch

import (
	"github.com/NagarjunNagesh/bus-company/restapi/operations/trip"
	"github.com/go-openapi/runtime/middleware"
)

type RestHandler interface {
	FetchATripHandler(params trip.GetTripByIDParams) middleware.Responder
}
