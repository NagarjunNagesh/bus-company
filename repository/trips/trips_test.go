package trips

import (
	"reflect"
	"testing"

	"github.com/NagarjunNagesh/bus-company/domain/models/city"
	trip_model "github.com/NagarjunNagesh/bus-company/domain/models/trip"
	irepository "github.com/NagarjunNagesh/bus-company/domain/repository"
	city_data "github.com/NagarjunNagesh/bus-company/repository/city"
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

func Test_repository_Get(t *testing.T) {
	setup()
	city_repository := city_data.New()
	type fields struct {
		city_repo irepository.CityRepository
	}
	type args struct {
		id int32
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
			name: "simpleGETTest",
			fields: fields{
				city_repo: city_repository,
			},
			args: args{id: 1},
			want: &trip_model.Trip{
				Dates:       &t1Date,
				Destination: &t1Destination,
				Origin:      &t1Origin,
				Price:       &t1Price,
			},
			wantErr: false,
		},
		{
			name: "simpleGETTest2",
			fields: fields{
				city_repo: city_repository,
			},
			args:    args{id: 7},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				city_repo: tt.fields.city_repo,
			}
			got, err := r.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_GetAll(t *testing.T) {
	setup()
	city_repository := city_data.New()
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
		city_repo irepository.CityRepository
	}
	type testCases struct {
		name    string
		fields  fields
		want    []*trip_model.Trip
		wantErr bool
	}
	tests := []testCases{
		{
			name: "simpleGETALLTest",
			fields: fields{
				city_repo: city_repository,
			},
			want:    trips_response,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				city_repo: tt.fields.city_repo,
			}
			got, err := r.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_Create(t *testing.T) {
	setup()
	city_repository := city_data.New()
	type fields struct {
		city_repo irepository.CityRepository
	}
	type args struct {
		add_trip *trip_model.AddTrip
	}
	type testCases struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}
	tests := []testCases{
		{
			name: "simpleCreateTest",
			args: args{
				add_trip: &trip_model.AddTrip{
					ID:            1,
					OriginID:      2,
					DestinationID: 3,
					Dates:         "Mon",
					Price:         20.33,
				},
			},
			fields: fields{
				city_repo: city_repository,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "OriginError",
			args: args{
				add_trip: &trip_model.AddTrip{
					ID:            1,
					OriginID:      7,
					DestinationID: 3,
					Dates:         "Mon",
					Price:         20.33,
				},
			},
			fields: fields{
				city_repo: city_repository,
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "DestinationError",
			args: args{
				add_trip: &trip_model.AddTrip{
					ID:            1,
					OriginID:      2,
					DestinationID: 7,
					Dates:         "Mon",
					Price:         20.33,
				},
			},
			fields: fields{
				city_repo: city_repository,
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				city_repo: tt.fields.city_repo,
			}
			got, err := r.Create(tt.args.add_trip)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("repository.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
