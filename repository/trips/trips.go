package trips

import (
	"errors"
	"fmt"

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
	var eDestination, eOrigin error

	for _, t := range Trips {
		if t.ID == id {
			aTrip.Dates = &t.Dates
			aTrip.Destination, eDestination = r.city_repo.FindCity(t.DestinationID)
			aTrip.Origin, eOrigin = r.city_repo.FindCity(t.OriginID)
			aTrip.Price = &t.Price

			if eOrigin != nil {
				e := fmt.Errorf("cannot find origin city for the trip id %d", id)
				return nil, e
			} else if eDestination != nil {
				e := fmt.Errorf("cannot find destination city for the trip id %d", id)
				return nil, e
			}

			return &aTrip, nil
		}
	}

	e := errors.New("no content")
	return nil, e
}

func (r *repository) GetAll() ([]*trip_model.Trip, error) {
	trips := []*trip_model.Trip{}
	var eDestination, eOrigin error

	for _, t := range Trips {
		var aTrip trip_model.Trip
		aTrip.Dates = &t.Dates
		aTrip.Destination, eDestination = r.city_repo.FindCity(t.DestinationID)
		aTrip.Origin, eOrigin = r.city_repo.FindCity(t.OriginID)
		aTrip.Price = &t.Price

		if eOrigin != nil {
			e := fmt.Errorf("cannot find origin city for the trip with id %d", t.ID)
			return nil, e
		} else if eDestination != nil {
			e := fmt.Errorf("cannot find destination city for the trip with id %d", t.ID)
			return nil, e
		}

		trips = append(trips, &aTrip)
	}
	return trips, nil
}

func (r *repository) Create(add_trip *trip_model.AddTrip) (bool, error) {
	destinationCityFound, eDestination := r.city_repo.FindCity(add_trip.DestinationID)
	if eDestination != nil || len(*destinationCityFound) == 0 {
		e := errors.New("destination city invalid")
		return false, e
	}

	originCityFound, eOrigin := r.city_repo.FindCity(add_trip.OriginID)
	if eOrigin != nil || len(*originCityFound) == 0 {
		e := errors.New("origin city invalid")
		return false, e
	}

	add_trip.ID = TripIDCounter
	Trips = append(Trips, add_trip)

	TripIDCounter++
	return true, nil
}
