package fetchall

import (
	"reflect"
	"testing"

	"github.com/NagarjunNagesh/bus-company/domain/models/city"
	fetch_all_trips "github.com/NagarjunNagesh/bus-company/domain/usecase/fetch-all"
	fetchall "github.com/NagarjunNagesh/bus-company/domain/usecase/fetch-all"
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

func Test_restHandler_FetchAllTripHandler(t *testing.T) {
	setup()
	city_repo := city_data.New()
	trips_repo := trips.New(city_repo)
	t1Date, t1Destination, t1Origin, t1Price := "Mon Tue Wed Fri", "Seville", "Barcelona", 40.55
	t2Date, t2Destination, t2Origin, t2Price := "Sat Sun", "Barcelona", "Seville", 40.55
	t3Date, t3Destination, t3Origin, t3Price := "Mon Tue Wed Thu Fri", "Malaga", "Madrid", 32.1
	trips_response := []*models.Trip{
		{
			Dates:       &t1Date,
			Destination: &t1Destination,
			Origin:      &t1Origin,
			Price:       &t1Price,
		},
		{
			Dates:       &t2Date,
			Destination: &t2Destination,
			Origin:      &t2Origin,
			Price:       &t2Price,
		},
		{
			Dates:       &t3Date,
			Destination: &t3Destination,
			Origin:      &t3Origin,
			Price:       &t3Price,
		},
	}
	type fields struct {
		fetch_all_trips_uc fetch_all_trips.UseCase
	}
	type args struct {
		params trip.GetAllTripsParams
	}
	type testCases struct {
		name   string
		fields fields
		args   args
		want   middleware.Responder
	}
	tests := []testCases{
		{
			name: "fetchAll",
			fields: fields{
				fetch_all_trips_uc: fetchall.NewUseCase(trips_repo),
			},
			args: args{
				params: trip.GetAllTripsParams{},
			},
			want: trip.NewGetAllTripsOK().WithPayload(trips_response),
		},
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
