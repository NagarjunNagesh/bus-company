package trips

import (
	"errors"

	trip_model "github.com/NagarjunNagesh/bus-company/domain/models/trip"
	irepository "github.com/NagarjunNagesh/bus-company/domain/repository"
)

type repository struct {
	city_repo irepository.CityRepository
}

func New(city_repo irepository.CityRepository) irepository.TripRepository {
	return &repository{city_repo: city_repo}
}

func (r *repository) Get(id int32) (*trip_model.Trip, error) {
	var aTrip trip_model.Trip
	for _, t := range Trips {
		if t.ID == id {
			aTrip.Dates = &t.Dates
			aTrip.Destination, _ = r.city_repo.FindCity(t.DestinationID)
			aTrip.Origin, _ = r.city_repo.FindCity(t.OriginID)
			aTrip.Price = &t.Price
			return &aTrip, nil
		}
	}

	e := errors.New("no content")
	return nil, e
}

func (r *repository) GetAll() ([]*trip_model.Trip, error) {
	trips := []*trip_model.Trip{}
	for _, t := range Trips {
		var aTrip trip_model.Trip
		aTrip.Dates = &t.Dates
		aTrip.Destination, _ = r.city_repo.FindCity(t.DestinationID)
		aTrip.Origin, _ = r.city_repo.FindCity(t.OriginID)
		aTrip.Price = &t.Price
		trips = append(trips, &aTrip)
	}
	return trips, nil
}

func (r *repository) Create(add_trip *trip_model.AddTrip) (bool, error) {
	destinationCityFound, eDestination := r.city_repo.FindCity(add_trip.DestinationID)
	if len(*destinationCityFound) == 0 || eDestination != nil {
		e := errors.New("destination city invalid")
		return false, e
	} 
	
	originCityFound, eOrigin := r.city_repo.FindCity(add_trip.OriginID)
	if len(*originCityFound) == 0 || eOrigin != nil {
		e := errors.New("origin city invalid")
		return false, e
	}

	add_trip.ID = TripIDCounter
	Trips = append(Trips, add_trip)

	TripIDCounter++
	return true, nil
}
