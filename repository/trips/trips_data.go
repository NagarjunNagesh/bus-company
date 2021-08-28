package trips

import (
	"github.com/NagarjunNagesh/bus-company/domain/models/trip"
)

var Trips = []*trip.AddTrip{
	{
		ID:            1,
		OriginID:      1,
		DestinationID: 2,
		Dates:         "Mon Tue Wed Fri",
		Price:         40.55,
	},
	{ID: 2, OriginID: 2, DestinationID: 1, Dates: "Sat Sun", Price: 40.55},
	{ID: 3, OriginID: 3, DestinationID: 6, Dates: "Mon Tue Wed Thu Fri", Price: 32.10},
}

var TripIDCounter = 4
