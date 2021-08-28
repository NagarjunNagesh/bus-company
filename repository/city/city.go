package city

import (
	"bufio"
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
	file, err := os.Open("repository/city/cities.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := int32(1)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		aCity := city.City{
			ID:   i,
			Name: scanner.Text(),
		}
		Cities = append(Cities, &aCity)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
