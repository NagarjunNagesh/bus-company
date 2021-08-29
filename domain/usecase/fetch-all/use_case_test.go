package fetchall

import (
	"reflect"
	"testing"

	"github.com/NagarjunNagesh/bus-company/domain/models/city"
	trip_model "github.com/NagarjunNagesh/bus-company/domain/models/trip"
	"github.com/NagarjunNagesh/bus-company/domain/repository"
	city_data "github.com/NagarjunNagesh/bus-company/repository/city"
	"github.com/NagarjunNagesh/bus-company/repository/trips"
)

func setup() {
	city_data.Cities = []*city.City{
		{
			ID:   1,
			Name: "Barcelona",
		},
		{
			ID:   2,
			Name: "Seville",
		},
		{
			ID:   3,
			Name: "Madrid",
		},
		{
			ID:   4,
			Name: "Valencia",
		},
		{
			ID:   5,
			Name: "Andorra la Vella",
		},
		{
			ID:   6,
			Name: "Malaga",
		},
	}
}

func Test_usecase_FetchAllTrips(t *testing.T) {
	setup()
	city_repo := city_data.New()
	trips_repo := trips.New(city_repo)
	t1Date, t1Destination, t1Origin, t1Price := "Mon Tue Wed Fri", "Seville", "Barcelona", 40.55
	t2Date, t2Destination, t2Origin, t2Price := "Sat Sun", "Barcelona", "Seville", 40.55
	t3Date, t3Destination, t3Origin, t3Price := "Mon Tue Wed Thu Fri", "Malaga", "Madrid", 32.1
	trips_response := []*trip_model.Trip{
		{
			Dates:       &t1Date,
			Destination: &t1Destination,
			Origin:      &t1Origin,
			Price:       &t1Price,
		},
		{
			Dates:       &t2Date,
			Destination: &t2Destination,
			Origin:      &t2Origin,
			Price:       &t2Price,
		},
		{
			Dates:       &t3Date,
			Destination: &t3Destination,
			Origin:      &t3Origin,
			Price:       &t3Price,
		},
	}
	type fields struct {
		trip_repo repository.TripRepository
	}
	type testCases struct {
		name    string
		fields  fields
		want    []*trip_model.Trip
		wantErr bool
	}
	tests := []testCases{
		{
			name: "FetchAllTrips",
			fields: fields{
				trip_repo: trips_repo,
			},
			want:    trips_response,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &usecase{
				trip_repo: tt.fields.trip_repo,
			}
			got, err := uc.FetchAllTrips()
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.FetchAllTrips() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usecase.FetchAllTrips() = %v, want %v", got, tt.want)
			}
		})
	}
}
