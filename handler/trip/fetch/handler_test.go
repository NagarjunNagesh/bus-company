package fetch

import (
	"reflect"
	"testing"

	fetch_one_trip "github.com/NagarjunNagesh/bus-company/domain/usecase/fetch"
	"github.com/NagarjunNagesh/bus-company/restapi/operations/trip"
	"github.com/go-openapi/runtime/middleware"
)

func Test_restHandler_FetchATripHandler(t *testing.T) {
	type fields struct {
		fetch_one_trip_uc fetch_one_trip.UseCase
	}
	type args struct {
		params trip.GetTripByIDParams
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
				fetch_one_trip_uc: tt.fields.fetch_one_trip_uc,
			}
			if got := h.FetchATripHandler(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restHandler.FetchATripHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
