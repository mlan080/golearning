package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/fromatob/engineering/training/backend/5/http-server/pkg/database"
	"github.com/lfritz/env"
)

var config struct {
	Port      int
	RedisHost string
	RedisPort int
}

func main() {
	// load configuration from environment variables
	e := env.New()
	e.Int("PORT", &config.Port, "the port to listen on")
	e.String("REDIS_HOST", &config.RedisHost, "Redis host")
	e.Int("REDIS_PORT", &config.RedisPort, "Redis port")
	err := e.Load()
	if err != nil {
		log.Fatal(err)
	}

	// set up database
	db, err := database.New(config.RedisHost, config.RedisPort)
	if err != nil {
		log.Fatal(err)
	}

	// set up HTTP handlers
	http.HandleFunc("/cities", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "GET":
			citiesInfo(db, w, req)
		case "POST":
			postCity(db, w, req)
		default:
			http.Error(w, "allowed methods for /cities: POST", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/search", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "GET":
			search(db, w, req)
		default:
			http.Error(w, "allowed methods for /search: GET", http.StatusMethodNotAllowed)
		}
	})

	// run!
	addr := fmt.Sprintf(":%d", config.Port)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func citiesInfo(db *database.Database, w http.ResponseWriter, req *http.Request) {
	response := struct {
		Count int `json:"count"`
	}{
		Count: db.Count(),
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		log.Printf("error marshaling JSON: %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(bytes)
	if err != nil {
		log.Printf("error writing JSON response: %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
}

func search(db *database.Database, w http.ResponseWriter, req *http.Request) {
	q := req.URL.Query().Get("q")
	if q == "" {
		http.Error(w, "parameter 'q' is required", http.StatusBadRequest)
		return
	}
	cities := db.Search(q)

	bytes, err := json.Marshal(cities)
	if err != nil {
		log.Printf("error marshaling JSON: %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(bytes)
	if err != nil {
		log.Printf("error writing JSON response: %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
}

func postCity(db *database.Database, w http.ResponseWriter, req *http.Request) {
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("error reading request body: %v", err)
		return
	}

	var city database.City
	err = json.Unmarshal(data, &city)
	if err != nil {
		log.Printf("error parsing request: %v", err)
		http.Error(w, "parameter 'q' is required", http.StatusBadRequest)
		return
	}

	db.Add(city)
	w.WriteHeader(http.StatusCreated)
}
