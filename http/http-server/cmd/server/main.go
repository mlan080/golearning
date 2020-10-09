package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	database "github.com/fromatob/engineering/training/backend/4/http-server/pkg/database"
)

func main() {
	// set up database
	db := database.New()

	// set up handlers for the different HTTP methods and paths
	http.HandleFunc("/cities", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "GET":
			citiesInfo(db, w, req)
		case "POST":
			postCity(db, w, req)
		default:
			// send status "405 Method not allowed" to the client
			http.Error(w, "allowed methods for /cities: POST", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/search", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "GET":
			search(db, w, req)
		default:
			http.Error(w, "allowed methods for /search:GET", http.StatusMethodNotAllowed)
		}
		// TODO call search, but only if it's a GET request
		// you can use a switch statment similar to the one above
	})

	// listen port 8080 for requests
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// citiesInfo handles the "GET /cities" endpoint
func citiesInfo(db *database.Database, w http.ResponseWriter, req *http.Request) {
	//create a struct
	type CitiesInfo struct {
		Count   int
		Updated time.Time
	}
	citiesInfo := CitiesInfo{
		Count:   db.Count(),
		Updated: db.LastUpdate(), //format time stamp
	}
	//fmt.Printf("count %v, last update: %v", a, b)
	//bytearray, _ := json.Marshal(citiesInfo)
	err := json.NewEncoder(w).Encode(citiesInfo)
	if err != nil {
		http.Error(w, "Unable to convert citiesInfo to json for /cities:GET", 400) //q: can you do this httpstatus thing or the number directly?
	}
	//fmt.Println(string(bytearray)) //excercise figure out bytearray printing format
}

// postCity handles the "POST /cities" endpoint
func postCity(db *database.Database, w http.ResponseWriter, req *http.Request) {
	city := database.City{}
	err := json.NewDecoder(req.Body).Decode(&city)
	if err != nil {
		http.Error(w, "Unable to convert city request to struct for /cities:POST", 400)
	}
	//byteValue, _ := ioutil.ReadAll(req)
	//json.Unmarshal(req.Body, &city)
	db.Add(city)
	//Q: add badreuqest statauscode before Add method or successstatus code after?
	fmt.Fprintf(w, "City added sucessfully %v", http.StatusCreated)

}

// search handles the "GET /search" endpoint
//localhost:8080//search?q=ber
func search(db *database.Database, w http.ResponseWriter, req *http.Request) {
	// get the parameter from the request
	q := req.URL.Query().Get("q")
	if q == "" {
		http.Error(w, "400  Bad reuqest for /search: GET", http.StatusBadRequest)
		// TODO use the http.Error function to send a "400 Bad request" to the client
		return
	}
	cities := db.Search(q)
	//json.NewEncoder(w).Encode(cities)
	byteArray, err := json.MarshalIndent(cities, "", "")
	if err != nil {
		http.Error(w, "Unable to convert cities response to json for /search: GET", 400)
	}
	fmt.Fprintf(w, string(byteArray))
}

//_ = q Q: what is this for?
