package fetchall

import trip_model "github.com/NagarjunNagesh/bus-company/domain/models/trip"

type UseCase interface {
	FetchAllTrips() []*trip_model.Trip
}
