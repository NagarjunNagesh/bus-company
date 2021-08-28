package fetch

import (
	trip_model "github.com/NagarjunNagesh/bus-company/domain/models/trip"
	"github.com/NagarjunNagesh/bus-company/domain/repository"
)

type usecase struct {
	trip_repo repository.TripRepository
}

func NewUseCase(trip_repo repository.TripRepository) UseCase {
	return &usecase{trip_repo: trip_repo}
}

func (uc *usecase) FetchOneTrip(tripID int64) *trip_model.Trip {
	trip, _ := uc.trip_repo.Get(int32(tripID))
	return trip
}
