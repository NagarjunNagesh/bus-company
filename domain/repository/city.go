package repository

type CityRepository interface {
	PopulateCities()
	FindCity(id int32) (*string, error)
}
