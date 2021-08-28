package city

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/NagarjunNagesh/bus-company/domain/models/city"
	irepository "github.com/NagarjunNagesh/bus-company/domain/repository"
)

type repository struct{}

func New() irepository.CityRepository {
	return &repository{}
}

func (r *repository) PopulateCities() {
	file, err := os.Open("resources/cities.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := int32(1)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		aCity := city.City{
			ID:   i,
			Name: scanner.Text(),
		}
		Cities = append(Cities, &aCity)
		fmt.Println(aCity)
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (r *repository) FindCity(id int32) (*string, error) {
	var city string
	for _, c := range Cities {
		if c.ID == id {
			city = c.Name
		}
	}

	if len(city) == 0 {
		e := errors.New("no content")
		return nil, e
	}

	return &city, nil
}
