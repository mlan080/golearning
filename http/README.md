# Training 4: HTTP and JSON

This session is about using HTTP and JSON to build a REST API.


## Documentation

The Go by Example website has good introductions of these topics:

* [Go by Example: JSON](https://gobyexample.com/json)
* [Go by Example: HTTP Clients](https://gobyexample.com/http-clients)
* [Go by Example: HTTP Servers](https://gobyexample.com/http-servers)


## Exercise

For the exercise we’ll build a service that lets clients look up cities by giving the first few
letters -- a bit like the fromAtoB clients do with the user’s input for departure and arrival
locations. We’re going to write both the server and a simple command-line client.

I’ve prepared some skeleton code and an implementation of a “city database” that the server will
use.


## API

The API we’re building will look have three endpoints:

    GET /cities
    
    Example response:
    {
        "count": 3,
        "last_update": "2020-09-28T09:54:43Z"
    }

This endpoint returns a JSON object with basic info: the number of cities and when the database was
last updated.

    POST /cities
    
    Example request:
    {
        "name": "Paris",
        "country_code": "FR",
        "latitude": 48.8589507,
        "longitude": 2.3426808
    }

This endpoint adds a new city to the database.

    GET /search?q=ber

    Example response:
    [
        {
            "name": "Berlin",
            "country_code": "DE",
            "latitude": 52.5203033,
            "longitude": 13.3977697
        },
        {
            "name": "Bern",
            "country_code": "CH",
            "latitude": 46.9490098,
            "longitude": 7.4450021
        }
    ]

This endpoint does a search (in this example looking for cities that start with “ber”) and returns a
JSON list of matches. The list will be empty if no results are found.


### Part 1: HTTP server

First, let’s implement the server side. The `http-server` directory has the “city database” in
`pkg/database`; your task is to add the API. For this exercise, we’ll just put the whole API code in
`cmd/server/main.go`.

A small part of main.go is already implemented to show how to work with HTTP methods (e.g. telling a
GET request apart from a POST request), sending HTTP status codes, and getting the request parameter
for the search query.

To keep things simple, we’ll always use port 8080 for the API, so you should be able to do

    go run cmd/server/main.go

and then the API will be available at `http://localhost:8080/`.

We’ll also skip writing unit tests for this exercise. It’s simple enough that you can just test it
with an HTTP client (e.g. curl or Postman). You can use the `sample_request.json` file to test the
POST endpoint.


### Part 2: HTTP client

The client is a command-line program that can be called in three ways. If you do

    go run main.go

it’ll print a help message that explains how to use it.

Your task is to implement the `info`, `add`, and `search` functions. To test it, you can run the
server from part 1 in one terminal, and run the client in another terminal (or directly from your
editor).

For error handling, you can use `log.Fatal` here. That’ll print an error and exit, which is fine for
a command-line program.
