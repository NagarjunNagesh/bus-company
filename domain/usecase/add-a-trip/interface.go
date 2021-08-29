package addatrip

import trip_model "github.com/NagarjunNagesh/bus-company/domain/models/trip"

type UseCase interface {
	AddATrip(tripModel *trip_model.AddTrip) (bool, error)
}
