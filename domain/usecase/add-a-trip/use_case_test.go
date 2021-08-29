package addatrip

import (
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

func Test_usecase_AddATrip(t *testing.T) {
	setup()
	city_repo := city_data.New()
	trips_repo := trips.New(city_repo)
	type fields struct {
		trip_repo repository.TripRepository
	}
	type args struct {
		tripModel *trip_model.AddTrip
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
			name: "CreateTrip",
			fields: fields{
				trip_repo: trips_repo,
			},
			args: args{
				tripModel: &trip_model.AddTrip{
					ID:            1,
					OriginID:      2,
					DestinationID: 3,
					Dates:         "Mon",
					Price:         20.33,
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "SameOriginAndDestinationIDs",
			fields: fields{
				trip_repo: trips_repo,
			},
			args: args{
				tripModel: &trip_model.AddTrip{
					ID:            1,
					OriginID:      2,
					DestinationID: 2,
					Dates:         "Mon",
					Price:         20.33,
				},
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "InvalidDate",
			fields: fields{
				trip_repo: trips_repo,
			},
			args: args{
				tripModel: &trip_model.AddTrip{
					ID:            1,
					OriginID:      2,
					DestinationID: 2,
					Dates:         "Sun",
					Price:         20.33,
				},
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &usecase{
				trip_repo: tt.fields.trip_repo,
			}
			got, err := uc.AddATrip(tt.args.tripModel)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.AddATrip() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("usecase.AddATrip() = %v, want %v", got, tt.want)
			}
		})
	}
}
