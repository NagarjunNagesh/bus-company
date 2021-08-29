package add

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/NagarjunNagesh/bus-company/domain/models/city"
	add_a_trip "github.com/NagarjunNagesh/bus-company/domain/usecase/add-a-trip"
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

func Test_restHandler_AddATripHandler(t *testing.T) {
	setup()
	city_repo := city_data.New()
	trips_repo := trips.New(city_repo)
	t1OriginID, t1DestinationID, t1Dates, t1Price := int32(6), int32(5), "Mon Tue", 40.33
	t2OriginID, t2DestinationID, t2Dates, t2Price := int32(6), int32(7), "Mon Tue", 40.33
	t3OriginID, t3DestinationID, t3Dates, t3Price := int32(7), int32(5), "Mon Tue", 40.33
	t4OriginID, t4DestinationID, t4Dates, t4Price := int32(2), int32(5), "Mon Sat", 40.33
	apiResponse := models.APIResponse{
		Code:    201,
		Message: "Successfully created a new trip",
		Type:    "JSON",
	}
	t2ApiResponse := models.APIResponse{
		Code:    400,
		Message: "destination city invalid",
		Type:    "JSON",
	}
	t3ApiResponse := models.APIResponse{
		Code:    400,
		Message: "origin city invalid",
		Type:    "JSON",
	}
	t4ApiResponse := models.APIResponse{
		Code:    400,
		Message: "invalid date value. Should be one of 'Mon', 'Tue', 'Wed', 'Thu', 'Fri'",
		Type:    "JSON",
	}
	type fields struct {
		add_a_trip_uc add_a_trip.UseCase
	}
	type args struct {
		params trip.AddNewTripParams
	}
	type testCases struct {
		name   string
		fields fields
		args   args
		want   middleware.Responder
	}
	tests := []testCases{
		{
			name: "CreateTrip",
			fields: fields{
				add_a_trip_uc: add_a_trip.NewUseCase(trips_repo),
			},
			args: args{
				params: trip.AddNewTripParams{
					HTTPRequest: &http.Request{},
					Body: &models.AddTrip{
						Dates:         &t1Dates,
						DestinationID: &t1DestinationID,
						OriginID:      &t1OriginID,
						Price:         &t1Price,
					},
				},
			},
			want: trip.NewAddNewTripCreated().WithPayload(&apiResponse),
		},
		{
			name: "DestinationError",
			fields: fields{
				add_a_trip_uc: add_a_trip.NewUseCase(trips_repo),
			},
			args: args{
				params: trip.AddNewTripParams{
					HTTPRequest: &http.Request{},
					Body: &models.AddTrip{
						Dates:         &t2Dates,
						DestinationID: &t2DestinationID,
						OriginID:      &t2OriginID,
						Price:         &t2Price,
					},
				},
			},
			want: trip.NewAddNewTripBadRequest().WithPayload(&t2ApiResponse),
		},
		{
			name: "OriginError",
			fields: fields{
				add_a_trip_uc: add_a_trip.NewUseCase(trips_repo),
			},
			args: args{
				params: trip.AddNewTripParams{
					HTTPRequest: &http.Request{},
					Body: &models.AddTrip{
						Dates:         &t3Dates,
						DestinationID: &t3DestinationID,
						OriginID:      &t3OriginID,
						Price:         &t3Price,
					},
				},
			},
			want: trip.NewAddNewTripBadRequest().WithPayload(&t3ApiResponse),
		},
		{
			name: "DateError",
			fields: fields{
				add_a_trip_uc: add_a_trip.NewUseCase(trips_repo),
			},
			args: args{
				params: trip.AddNewTripParams{
					HTTPRequest: &http.Request{},
					Body: &models.AddTrip{
						Dates:         &t4Dates,
						DestinationID: &t4DestinationID,
						OriginID:      &t4OriginID,
						Price:         &t4Price,
					},
				},
			},
			want: trip.NewAddNewTripBadRequest().WithPayload(&t4ApiResponse),
		},
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
