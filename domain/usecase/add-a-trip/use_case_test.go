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
