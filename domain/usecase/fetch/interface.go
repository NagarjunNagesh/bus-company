package fetch

import trip_model "github.com/NagarjunNagesh/bus-company/domain/models/trip"

type UseCase interface {
	FetchOneTrip(tripID int64) (*trip_model.Trip, error)
}
