package trip

type AddTrip struct {
	ID            int32
	OriginID      int32
	DestinationID int32
	Dates         string
	Price         float64
}
