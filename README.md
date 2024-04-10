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

The goal of this project is to create a simple caching api for a webserver and then use different testing methods to test the system. E.g. Unit Tests, Integration Tests, E2E Tests, etc.

This Repository is only the backend the Frontend can be found [here](https://codeberg.org/Birkenfunk/SQS-Frontend)

## How to run the project
1. Clone the repository
2. copy the `.env.example` file to `.env` and fill in the values
3. run `docker-compose up --build`
4. The server should now be running on `localhost:4000`

