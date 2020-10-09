package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		usageAndExit()
	}
	switch args[0] {
	case "info":
		info()
	case "add":
		add(args[1:])
	case "search":
		search(args[1:])
	default:
		usageAndExit()
	}
}

func usageAndExit() {
	fmt.Print(`Usage:

go run main.go info
    Print some info on the service

go run main.go add <name> <country> <lat> <lon>
    Add a city

go run main.go search <query>
    Search for cities
`)
	os.Exit(1)
}

func info() {
	resp, err := http.Get("http://localhost:3000/cities")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(body))

	// TODO get info from server and print it
}

func add(args []string) {
	if len(args) != 4 {
		usageAndExit()
	}
	name := args[0]

	country := args[1]
	lat, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		usageAndExit()
	}
	lon, err := strconv.ParseFloat(args[3], 64)
	if err != nil {
		usageAndExit()
	}

	requestBody := strings.NewReader(`
	{
        "name": "Paris",
        "country_code": "FR",
        "latitude": 48.8589507,
        "longitude": 2.3426808
	}
	`)
	//If you have a JSON compatible data structure like struct or map, then you can use json.Marshal function to encode Go data structure to JSON data and bytes.NewReader function to convert JSON data to a io.Reader object.

	resp, err := http.Post("http://localhost:3000/cities", "application/json", requestBody)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	resp.Body.Close()

	fmt.Printf("%s\n", body)

	// 	_ = name
	// 	_ = country
	// 	_ = lat
	// 	_ = lon
}

func search(args []string) {
	if len(args) != 1 {
		usageAndExit()
	}
	query := args[0]
	resp, err := http.Get("http://localhost:3000/cities")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(body))
	// TODO make a request to the server and print search results
	_ = query
}
