package fetchall

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

func (uc *usecase) FetchAllTrips() ([]*trip_model.Trip, error) {
	trips, err := uc.trip_repo.GetAll()
	return trips, err
}
