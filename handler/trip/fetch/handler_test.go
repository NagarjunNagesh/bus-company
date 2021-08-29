package fetch

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/NagarjunNagesh/bus-company/domain/models/city"
	fetch_one_trip "github.com/NagarjunNagesh/bus-company/domain/usecase/fetch"
	"github.com/NagarjunNagesh/bus-company/models"
	city_data "github.com/NagarjunNagesh/bus-company/repository/city"
	"github.com/NagarjunNagesh/bus-company/repository/trips"
	"github.com/NagarjunNagesh/bus-company/restapi/operations/trip"
	"github.com/go-openapi/runtime/middleware"
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

func Test_restHandler_FetchATripHandler(t *testing.T) {
	setup()
	city_repo := city_data.New()
	trips_repo := trips.New(city_repo)
	t1Date, t1Destination, t1Origin, t1Price := "Mon Tue Wed Fri", "Seville", "Barcelona", 40.55
	tripFound := models.Trip{
		Dates:       &t1Date,
		Destination: &t1Destination,
		Origin:      &t1Origin,
		Price:       &t1Price,
	}
	apiResponse := models.APIResponse{
		Code:    400,
		Message: "no content",
		Type:    "JSON",
	}
	type fields struct {
		fetch_one_trip_uc fetch_one_trip.UseCase
	}
	type args struct {
		params trip.GetTripByIDParams
	}
	type testCases struct {
		name   string
		fields fields
		args   args
		want   middleware.Responder
	}
	tests := []testCases{
		{
			name: "FetchOneTrip",
			fields: fields{
				fetch_one_trip_uc: fetch_one_trip.NewUseCase(trips_repo),
			},
			args: args{
				params: trip.GetTripByIDParams{
					HTTPRequest: &http.Request{},
					TripID:      1,
				},
			},
			want: trip.NewGetTripByIDOK().WithPayload(&tripFound),
		},
		{
			name: "FetchOneTripError",
			fields: fields{
				fetch_one_trip_uc: fetch_one_trip.NewUseCase(trips_repo),
			},
			args: args{
				params: trip.GetTripByIDParams{
					HTTPRequest: &http.Request{},
					TripID:      7,
				},
			},
			want: trip.NewGetAllTripsBadRequest().WithPayload(&apiResponse),
		},
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
