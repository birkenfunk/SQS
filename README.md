# SQS
[![License: Apache 2](https://img.shields.io/badge/License-Apache2-blue.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)
[![status-badge](https://woodpecker.birkenfunk.de/api/badges/4/status.svg)](https://woodpecker.birkenfunk.de/repos/4)
[![Coverage](https://sonarqube.birkenfunk.de/api/project_badges/measure?project=SQS&metric=coverage&token=sqb_f7ecb4947be271a3f2e849f9bc1b28ddd0dc4a1c)](https://sonarqube.birkenfunk.de/dashboard?id=SQS)
[![Quality Gate Status](https://sonarqube.birkenfunk.de/api/project_badges/measure?project=SQS&metric=alert_status&token=sqb_f7ecb4947be271a3f2e849f9bc1b28ddd0dc4a1c)](https://sonarqube.birkenfunk.de/dashboard?id=SQS)

<a href="https://codeberg.org/Birkenfunk/SQS">
 <img alt="Get it on Codeberg" src="https://codeberg.org/Codeberg/GetItOnCodeberg/media/branch/main/get-it-on-neon-blue.png" height="60">
</a>

## What is SQS?
SQS is a subject at the University of Applied Sciences in Rosenheim.

A full documentation of the quality specifications can be found [here](./docs/Documentation.adoc).

This Repository is only the backend the Frontend can be found [here](https://codeberg.org/Birkenfunk/SQS-Frontend)

## How to get started

This has to be done once to set up the project

1. Clone the repository
2. copy the `.env.example` file to `.env` and fill in the values

## How to run the project
1. run `docker-compose up --build`
2. The server should now be running on `localhost:4000`

## How to configure the project
The project can be configured by changing the `.env` file. The following values can be set:
- `PORT` The port the server should run on
- `WEATHER_SERVICE_API_URL` The URL of the weather service
- `REDIS_URL` The URL of the redis server

## How to run the tests
1. run `docker-compose -f integration.docker-compose.yml up`
2. run make test

## How to run the linter
1. install [golangci-lint](https://golangci-lint.run/welcome/install/)
2. run `golanci-lint run ./...`

## How to run the coverage
1. run `make coverage`
2. open the file `coverage.html` in the build folder

## Used Libraries
- [Go-Chi](https://go-chi.io/#/) This is the router used in the project
- [Redigo](https://github.com/gomodule/redigo) This is the redis client used in the project
- [Godotenv](https://github.com/joho/godotenv) This is the library used to read the .env file
- [Zerolog](https://github.com/rs/zerolog) This is the logging library used in the project
