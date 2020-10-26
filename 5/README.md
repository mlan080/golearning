# Training 5: Docker Compose, Production Logs

This session is about two ops-related tasks: using Docker and Docker Compose for local development,
and accessing logs for a service in production.


## Docker and Docker Compose

For our backend services, we use Docker to build them, Docker Compose to run them in development and
Kubernetes to run them in staging and production. Since we did a Kubernetes workshop quite recently,
we’ll look at the other two here.

To get started with writing Docker files, read the Quickstart guide on the Docker website:
* [Part 1: Orientation and setup](https://docs.docker.com/get-started/)
* [Part 2: Build and run your image](https://docs.docker.com/get-started/part2/)

There’s also a part 3 that’s about Docker Hub, but we’re using Google Cloud Registry instead of
Docker Hub so that part isn’t relevant for us.

For Docker Compose there’s a good tutorial here: [Overview of Docker
Compose](https://docs.docker.com/compose/).


### Exercise

This exercise is based on the “http-server” exercise from the previous session. In the http-server
directory you’ll find a service like the one from session 4, but it uses Redis as the database. It
also reads some configuration options from environment variables.

Your task is to

1. Write a `Dockerfile` so you can build a Docker image for the service. You can use the
   `golang:1.15` image as the base.
2. Write a `docker-compose.yml` so you can run the http-server together with Redis. You’ll have to
   set all the environment variables that the http-server needs; take a look at the `main.go` to see
   which ones those are. For Redis you can use the `redis:5` image.

The documentation linked above should help you get started. You can use the [Dockerfile
reference](https://docs.docker.com/engine/reference/builder/) and [Docker Compose file
reference](https://docs.docker.com/compose/compose-file/) to look up specific options.

The goal is that you can run it something like this (I’m using
[lwp-request](https://metacpan.org/pod/distribution/libwww-perl/bin/lwp-request) and
[jq](https://stedolan.github.io/jq/) here, but you can also use Postman or some other client):

    http-server$ docker-compose build
    ...
    http-server$ docker-compose up -d
    ...
    http-server$ GET http://localhost:8080/cities | jq .
    {
      "count": 4
    }
    http-server$ cat sample_request.json | POST http://localhost:8080/cities
    http-server$ GET http://localhost:8080/search?q=par | jq .
    [
      {
        "name": "Paris",
        "country_code": "FR",
        "latitude": 48.8589507,
        "longitude": 2.3426808
      }
    ]

You don’t need to change the Go code.


## Logs on Google Cloud

One important tool to understand what’s going on with services in production (and staging) is
viewing logs. With our setup, the way this works is:

1. We (the developers) add logging statements to our code. We’re been using the [logrus
   package](https://github.com/sirupsen/logrus) in our Go code.
2. The log messages get printed to standard output -- that means if you’re running a service
   locally, it just prints to the terminal.
3. Kubernetes collects those log messages from all services and sends them to Google Cloud Logs.
4. Google Cloud Logs stores them in a database of some sort and makes them available to view in the
   Google Cloud console.

For this exercise, we’ll use the Logs Viewer to take a look at the logs for some of our services.
The idea is for you to get a bit of experience using it so you’ll be ready when you need it for your
own services.

The documentation for the Logs Viewer is here: [Viewing logs
(Classic)](https://cloud.google.com/logging/docs/view/overview). Most of it is pretty
self-explaining, though.

To see the production logs for a service, go to Google Cloud Console, select the a2b-master project,
go to Legacy Logs Viewer, and in the drop-down for the resource, go to Kubernetes Container >
k8n-cluster-2 > default namespace > *service name*. For staging it’s the same except with a2b-exp
and k8n-cluster-1. For some projects we also have a link to the logs in the README so you don’t have
to navigate all those menus...

Your task is to answer these questions:

1. Take a look at the logs for provider-hitchhiker on October 2nd. Are there any logs with level
   Error?
   1. You should have found exactly one log message. Write down the exact time and message.
   2. Go to the source code and find where exactly that message came from. Write down filename and
      line number.
   3. If you look at the source, the code is logging a bit more information than just the time and
      error message. Find that additional information in the Logs Viewer.

2. Take a look at the logs for provider-sbb between October 7th and October 11th. Are there any logs
   with level Error?
   1. You should have found exactly six log messages, all of which have the same message. Write down
      that message.
   2. Go to the source code and find where exactly that message came from. Write down filename and
      line number.
   3. If you look at the source, the code is logging a bit more information than just the time and
      error message. Find that additional information in the Logs Viewer.
   4. Based on the additional information that you found, write down the names of all unique
      translations that caused this error.

3. Take a look at the logs for provider-sbb on October 8th. Are there any logs with level Warning?
   1. You should have found exactly two log messages, both of which have the same message. Write
      down that message.
   2. Go to the source code and find where exactly that message came from. Write down filename and
      line number.
   3. If you look at the source, the code is logging a bit more information than just the time and
      error message. Find that additional information in the Logs Viewer.
   4. The additional information that you found is an error returned by a function near the line
      number that you noted before. Go to the definition of this function (hint: this function is
      defined in a separate repository, and used here via a dependent module) and note down the
      filename and line number of the origin of this error.
