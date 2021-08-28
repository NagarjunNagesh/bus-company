package fetchall

import (
	"reflect"
	"testing"

	fetch_all_trips "github.com/NagarjunNagesh/bus-company/domain/usecase/fetch-all"
	"github.com/NagarjunNagesh/bus-company/restapi/operations/trip"
	"github.com/go-openapi/runtime/middleware"
)

func Test_restHandler_FetchAllTripHandler(t *testing.T) {
	type fields struct {
		fetch_all_trips_uc fetch_all_trips.UseCase
	}
	type args struct {
		params trip.GetAllTripsParams
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   middleware.Responder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &restHandler{
				fetch_all_trips_uc: tt.fields.fetch_all_trips_uc,
			}
			if got := h.FetchAllTripHandler(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restHandler.FetchAllTripHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
