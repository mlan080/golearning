package database

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/go-redis/redis/v8"
)

type City struct {
	Name        string  `json:"name"`
	CountryCode string  `json:"country_code"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}

type Database struct {
	client *redis.Client
	cities []City
}

func New(redisHost string, redisPort int) (*Database, error) {
	ctx := context.Background()

	var db Database
	db.client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", redisHost, redisPort),
		DB:   0,
	})
	_, err := db.client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("error connecting to Redis: %w", err)
	}

	if db.Count() == 0 {
		db.Add(City{"Berlin", "DE", 52.5203033, 13.3977697})
		db.Add(City{"Vienna", "AT", 48.2086788, 16.3681014})
		db.Add(City{"Zurich", "CH", 47.3754557, 8.5385552})
		db.Add(City{"Bern", "CH", 46.9490098, 7.4450021})
	}

	return &db, nil
}

func (d *Database) Count() int {
	ctx := context.Background()
	count, err := d.client.DBSize(ctx).Result()
	if err != nil {
		log.Printf("error reading from Redis: %v", err)
	}
	return int(count)
}

func (d *Database) Add(city City) {
	ctx := context.Background()
	key := strings.ToLower(city.Name)
	_, err := d.client.HSet(ctx, key, map[string]interface{}{
		"name":         city.Name,
		"country_code": city.CountryCode,
		"latitude":     city.Latitude,
		"longitude":    city.Longitude,
	}).Result()
	if err != nil {
		log.Printf("error adding entry to Redis: %v", err)
	}
}

func (d *Database) Search(prefix string) []City {
	cities := []City{}
	ctx := context.Background()
	keys, err := d.client.Keys(ctx, fmt.Sprintf("%s*", prefix)).Result()
	if err != nil {
		log.Printf("error looking up keys in Redis: %v", err)
		return cities
	}
	for _, key := range keys {
		fields, err := d.client.HGetAll(ctx, key).Result()
		if err != nil {
			log.Printf("error looking up city for %+v in Redis: %v", key, err)
			continue
		}
		var city City
		city.Name = fields["name"]
		city.CountryCode = fields["country_code"]
		city.Latitude, err = strconv.ParseFloat(fields["latitude"], 64)
		if err != nil {
			log.Printf("invalid latitude in Redis for %+v: %v", key, err)
			continue
		}
		city.Longitude, err = strconv.ParseFloat(fields["longitude"], 64)
		if err != nil {
			log.Printf("invalid longitude in Redis for %+v: %v", key, err)
			continue
		}
		cities = append(cities, city)
	}
	return cities
}
