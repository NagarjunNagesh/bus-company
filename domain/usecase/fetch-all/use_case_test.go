package fetchall

import (
	"reflect"
	"testing"

	trip_model "github.com/NagarjunNagesh/bus-company/domain/models/trip"
	"github.com/NagarjunNagesh/bus-company/domain/repository"
)

func Test_usecase_FetchAllTrips(t *testing.T) {
	type fields struct {
		trip_repo repository.TripRepository
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
