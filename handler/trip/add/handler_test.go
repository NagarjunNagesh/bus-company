package add

import (
	"reflect"
	"testing"

	add_a_trip "github.com/NagarjunNagesh/bus-company/domain/usecase/add-a-trip"
	"github.com/NagarjunNagesh/bus-company/restapi/operations/trip"
	"github.com/go-openapi/runtime/middleware"
)

func Test_restHandler_AddATripHandler(t *testing.T) {
	type fields struct {
		add_a_trip_uc add_a_trip.UseCase
	}
	type args struct {
		params trip.AddNewTripParams
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
				add_a_trip_uc: tt.fields.add_a_trip_uc,
			}
			if got := h.AddATripHandler(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restHandler.AddATripHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
