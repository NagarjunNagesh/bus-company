package repository

import (
	trip_model "github.com/NagarjunNagesh/bus-company/domain/models/trip"
)

type TripRepository interface {
	Get(id int) (*trip_model.Trip, error)
	GetAll() ([]*trip_model.Trip, error)
	Create(add_trip *trip_model.AddTrip) (bool, error)
}
