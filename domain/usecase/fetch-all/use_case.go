package fetchall

import (
	trip_model "github.com/NagarjunNagesh/bus-company/domain/models/trip"
	"github.com/NagarjunNagesh/bus-company/domain/repository"
)

type usecase struct {
	trip_repo repository.Trip
}

func NewUseCase(trip_repo repository.Trip) UseCase {
	return &usecase{trip_repo: trip_repo}
}

func (uc *usecase) FetchAllTrips() []*trip_model.Trip {
	trips, _ := uc.trip_repo.GetAll()
	return trips
}
