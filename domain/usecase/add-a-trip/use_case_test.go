package addatrip

import (
	"testing"

	trip_model "github.com/NagarjunNagesh/bus-company/domain/models/trip"
	"github.com/NagarjunNagesh/bus-company/domain/repository"
)

func Test_usecase_AddATrip(t *testing.T) {
	type fields struct {
		trip_repo repository.TripRepository
	}
	type args struct {
		tripModel *trip_model.AddTrip
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &usecase{
				trip_repo: tt.fields.trip_repo,
			}
			if got := uc.AddATrip(tt.args.tripModel); got != tt.want {
				t.Errorf("usecase.AddATrip() = %v, want %v", got, tt.want)
			}
		})
	}
}
