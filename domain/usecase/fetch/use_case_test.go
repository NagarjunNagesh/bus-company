package fetch

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

func Test_usecase_FetchOneTrip(t *testing.T) {
	setup()
	city_repo := city_data.New()
	trips_repo := trips.New(city_repo)
	type fields struct {
		trip_repo repository.TripRepository
	}
	type args struct {
		tripID int64
	}
	type testCases struct {
		name    string
		fields  fields
		args    args
		want    *trip_model.Trip
		wantErr bool
	}
	t1Date, t1Destination, t1Origin, t1Price := "Mon Tue Wed Fri", "Seville", "Barcelona", 40.55
	tests := []testCases{
		{
			name: "Fetch",
			fields: fields{
				trip_repo: trips_repo,
			},
			args: args{
				tripID: 1,
			},
			want: &trip_model.Trip{
				Dates:       &t1Date,
				Destination: &t1Destination,
				Origin:      &t1Origin,
				Price:       &t1Price,
			},
			wantErr: false,
		},
		{
			name: "FetchInvalidID",
			fields: fields{
				trip_repo: trips_repo,
			},
			args: args{
				tripID: 0,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &usecase{
				trip_repo: tt.fields.trip_repo,
			}
			got, err := uc.FetchOneTrip(tt.args.tripID)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.FetchOneTrip() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usecase.FetchOneTrip() = %v, want %v", got, tt.want)
			}
		})
	}
}
