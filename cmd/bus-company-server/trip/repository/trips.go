package repository


type Trip struct {
	id            int32
	originId      int32
	destinationId int32
	dates         string
	price         float64
}

var trips = [] Trip {
	{id: 1, originId: 1, destinationId: 2, dates: "Mon Tue Wed Fri", price: 40.55},
	{id: 2, originId: 2, destinationId: 1, dates: "Sat Sun", price: 40.55},
	{id: 3, originId: 3, destinationId: 6, dates: "Mon Tue Wed Thu Fri", price: 32.10},
}