package trips

import (
	"reflect"
	"testing"

	trip_model "github.com/NagarjunNagesh/bus-company/domain/models/trip"
	irepository "github.com/NagarjunNagesh/bus-company/domain/repository"
)

func Test_repository_Get(t *testing.T) {
	type fields struct {
		city_repo irepository.CityRepository
	}
	type args struct {
		id int32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *trip_model.Trip
		wantErr bool
	}{
		// TODO: Add test cases.
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
	type fields struct {
		city_repo irepository.CityRepository
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*trip_model.Trip
		wantErr bool
	}{
		// TODO: Add test cases.
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
	type fields struct {
		city_repo irepository.CityRepository
	}
	type args struct {
		add_trip *trip_model.AddTrip
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
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
