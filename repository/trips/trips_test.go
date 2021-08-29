package trips

import (
	"reflect"
	"testing"

	"github.com/NagarjunNagesh/bus-company/domain/models/city"
	trip_model "github.com/NagarjunNagesh/bus-company/domain/models/trip"
	irepository "github.com/NagarjunNagesh/bus-company/domain/repository"
	city_data "github.com/NagarjunNagesh/bus-company/repository/city"
)

func Test_repository_Get(t *testing.T) {
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
				city_repo: city_data.New(),
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
