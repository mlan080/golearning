package database

import (
	"strings"
	"time"
)

type City struct {
	Name        string  `json:"name"`
	CountryCode string  `json:"country_code"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}

type Database struct {
	cities  []City
	updated time.Time
}

func New() *Database {
	return &Database{
		cities: []City{
			{"Berlin", "DE", 52.5203033, 13.3977697},
			{"Vienna", "AT", 48.2086788, 16.3681014},
			{"Zurich", "CH", 47.3754557, 8.5385552},
			{"Bern", "CH", 46.9490098, 7.4450021},
		},
		updated: now(),
	}
}

func (d *Database) Count() int {
	return len(d.cities)
}

func (d *Database) LastUpdate() time.Time {
	return now()
}

func (d *Database) Add(city City) {
	d.cities = append(d.cities, city)
	d.updated = time.Now()
}

func (d *Database) Search(prefix string) []City {
	prefix = strings.ToLower(prefix)
	result := []City{}
	for _, city := range d.cities {
		if strings.HasPrefix(strings.ToLower(city.Name), prefix) {
			result = append(result, city)
		}
	}
	return result
}

func now() time.Time {
	return time.Now().UTC()
}
