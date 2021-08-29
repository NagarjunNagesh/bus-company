package addatrip

import (
	"errors"
	"fmt"
	"strings"

	trip_model "github.com/NagarjunNagesh/bus-company/domain/models/trip"
	"github.com/NagarjunNagesh/bus-company/domain/repository"
)

var dates = [...]string{"Mon", "Tue", "Wed", "Thu", "Fri"}

type usecase struct {
	trip_repo repository.TripRepository
}

func NewUseCase(trip_repo repository.TripRepository) UseCase {
	return &usecase{trip_repo: trip_repo}
}

func (uc *usecase) AddATrip(tripModel *trip_model.AddTrip) (bool, error) {
	isNotValidDate := hasInvalidDates(tripModel.Dates)

	if isNotValidDate {
		e := errors.New("invalid date value. Should be one of 'Mon', 'Tue', 'Wed', 'Thu', 'Fri'")
		return false, e
	}

	hasCreated, err := uc.trip_repo.Create(tripModel)
	return hasCreated, err
}

func hasInvalidDates(dateAsString string) bool {
	allDates := strings.Split(dateAsString, " ")
	exists := populatesDatesAsMap()

	for _, value := range allDates {
		if !exists[value] {
			fmt.Printf("The date value %s is invalid\n", value)
			return true
		}
	}
	return false
}

func populatesDatesAsMap() map[string]bool {
	exists := make(map[string]bool)
	for _, value := range dates {
		exists[value] = true
	}
	fmt.Println(exists)
	return exists
}
