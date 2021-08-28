package trips

import (
	trip_model "github.com/NagarjunNagesh/bus-company/domain/models/trip"
	irepository "github.com/NagarjunNagesh/bus-company/domain/repository"
)

type repository struct{}

func New() irepository.Trip {
	return &repository{}
}

func (r *repository) Get(id int) (*trip_model.Trip, error) {
	dates := "Mon Tue Wed Thurs"
	barcelona := "Barcelona"
	price := 40.22
	origin := "Seville"
	aTrip := trip_model.Trip{
		Dates:       &dates,
		Destination: &barcelona,
		Origin:      &origin,
		Price:       &price,
	}
	return &aTrip, nil
}

func (r *repository) GetAll() ([]*trip_model.Trip, error) {
	dates := "Mon Tue Wed Thurs"
	barcelona := "Barcelona"
	price := 40.22
	origin := "Seville"
	aTrip := trip_model.Trip{
		Dates:       &dates,
		Destination: &barcelona,
		Origin:      &origin,
		Price:       &price,
	}
	trips := []*trip_model.Trip{}
	trips = append(trips, &aTrip)
	return trips, nil
}

func (r *repository) Create(add_trip *trip_model.AddTrip) (bool, error) {
	return true, nil
}
