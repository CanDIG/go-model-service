# Go Model Service

Based on Jonathan Dursi's [OpenAPI variant service demo](https://github.com/ljdursi/openapi_calls_example), this toy service demonstrates the go-swagger/pop stack with CanDIG API best practicNote that most auto-generated code, including binary files such as `main`, have been excluded from this repository. The exclusion of these files is considered to be best practice for repository maintenance. However, one-time auto-generated files that are *not* re-generated (and are therefore safe to edit) should be pushed to this repository. `model-vs/api/restapi/configure_variant_service.go` and the `model-vs/data/models` package are examples of such safe-to-edit auto-generated code.es.

[![Build Status](https://travis-ci.org/CanDIG/go-model-service.svg?branch=main)](https://travis-ci.org/CanDIG/go-model-service)
[![Go Report Card](https://goreportcard.com/badge/github.com/candig/go-model-service)](https://goreportcard.com/report/github.com/candig/go-model-service)

<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=6 orderedList=false} -->
<!-- code_chunk_output -->

- [Go Model Service](#go-model-service)
  - [Quick Start](#quick-start)
  - [Stack](#stack)
    - [Installing the stack](#installing-the-stack)
  - [Installing the service](#installing-the-service)
  - [Running The Service](#running-the-service)
    - [Request examples](#request-examples)
  - [For Developers](#for-developers)

<!-- /code_chunk_output -->

## Quick Start

1. Install [Docker v18.06.0+](https://docs.docker.com/get-docker/) and [Docker Compose v3.7+](https://docs.docker.com/compose/install/). You can check which version, if any, is already installed by running `docker --version` and `docker-compose --version`. Any versions greater than or equal to the ones stated here should do.
2. Clone this go-model-service repository:
  ```
  git clone https://github.com/CanDIG/go-model-service.git
  cd go-model-service
  ```
3. Set up the build environment by generating a `.env` file for Docker Compose to use:
  ```
  cp default.env .env
  ```
4. From the project root directory, run the server. It is presently configured to run on port `0.0.0.0:3000`.
  ```
  docker-compose up
  ```
5. In a new shell (or the same shell if `docker-compose up --detach` was run), run the migrations script to prepare the database for use. The script is located at the project root.
  ```
  ./migrate.sh
  ```
6. Send a request to the go-model-service from Postman (import `./tests/go-model-service.postman_collection.json` to run the collection against the `./tests/postman-data.csv` data file). Alternately, `curl` a request from the command line:
  ```
  curl -i http://0.0.0.0:3000/v1/individuals -d "{\"description\":\"Subject 17\"}" -H 'Content-Type: application/json'
  ```

## Stack

- [Docker](https://www.docker.com/) is used for containerizing the service and its dependencies.
- [PostgreSQL](https://www.postgresql.org/) database backend
- [Go](https://golang.org/) (Golang) backend
- Go [mod](https://blog.golang.org/using-go-modules) is used for dependency management, similarly to how `dep` was used in the past.
- [Swagger/OpenAPI 2.0](https://swagger.io/specification/v2/) is used to specify the API.
- [Postman](https://www.postman.com/) and the CLI runner [Newman](https://learning.postman.com/docs/postman/collection-runs/command-line-integration-with-newman/) are used for end-to-end API testing.
- [TravisCI](https://travis-ci.org/) is used for continuous integration, testing both the service build and, in conjunction with Newman, the API.
- [Go-swagger](https://goswagger.io/) auto-generates boilerplate Go code from a `swagger.yml` API definition. [Swagger](https://swagger.io/) tooling is based on the [OpenAPI](https://www.openapis.org/) specification.
- Gobuffalo [pop](https://github.com/gobuffalo/pop) is used as an orm-like for interfacing between the go code and the sqlite3 database. The `soda` CLI auto-generates boilerplate Go code for models and migrations, as well as performing migrations on the database. `fizz` files are used for defining database migrations in a Go-like syntax (see [syntax documentation](https://gobuffalo.io/en/docs/db/fizz/).)
- [genny](https://github.com/CanDIG/genny) is a code-generation solution to generics in Go.
- Gobuffalo [validate](https://github.com/gobuffalo/validate) is a framework used for writing custom validators. Some of their validators in the `validators` package are used as-is.

### Installing the stack

The `gms-webapp` image built from `./Dockerfile` depends upon the `gms-deps` image built from `./Dockerfile-gms-deps`, the image containing most of the stack and package dependencies for the app. These images have been split so that the `gms-webapp` build can be tested by TravisCI following every git push without time spent building the dependencies.

The stack and project dependencies image `gms-deps` can be pulled from `katpavlov/gms-deps`. Alternately, it can be built and run locally with the following commands:
  ```
  docker build -t <username>/gms-deps -f ./Dockerfile-gms-deps
  docker run -it --rm <username>/gms-deps
  ```

If you would like to push your altered image of `gms-deps`, consider using the `./push-image.sh` script provided to quickly semantically version the image. Run `./push-image.sh -h` for usage instructions.
  ```
  docker login && ./push-image.sh -f ./Dockerfile-gms-deps -u <username> gms-deps <patch>
  ```

## Installing the service

For containerized installation instructions, please see the [Quick Start](#quick-start).

If you are interested in attempting a non-containerized installation, the following files contain build instructions for Docker, Docker Compose, and Travis CI. They may equally be followed manually to accomplish a manual installation of the stack and service:
- `./Dockerfile-gms-deps` for the installation of the stack and of most other project dependencies
- `./Dockerfile` for the installation of the webapp, along with
- `docker-compose.yml` and `default.env` for the configuration of the service, database, and build-time environment
- `travis.yml` for running the service and its associated end-to-end API tests

Alternately, the [`v1.0` release](https://github.com/CanDIG/go-model-service/tree/v1.0) of this service contains instructions for building the stack and service without use of containers. However, the following parts of the stack are outdated for that version:
- `Sqlite3` database backend instead of `PostgreSQL`
- `dep` dependency management instead of go `mod`
- `Dredd` API testing instead of `Postman`/`Newman`
- Older versions of many current dependencies

## Running The Service

From the project root directory, run the server. It is presently configured to run on port `0.0.0.0:3000`.
  ```
  docker-compose up
  ```

### Request examples

If you have [Postman](https://www.postman.com/downloads/) installed, the quickest way to test the go-model-service API is to import `./tests/go-model-service.postman_collection.json` and run the collection against the `./tests/postman-data.csv` data file.

Alternately, you can `curl` requests to the service from the command line, ex.:
- POST an Individual:
  ```
  curl -i http://0.0.0.0:3000/v1/individuals -d "{\"description\":\"Subject 1\"}" -H 'Content-Type: application/json'
  ```
- GET all Individuals:
  ```
  curl -i "http://0.0.0.0:3000/v1/individuals"
  ```
- POST a Variant:
  ```
  curl -i http://0.0.0.0:3000/v1/variants -d "{\"name\":\"rs7054258\", \"chromosome\":\"chr1\", \"start\":5, \"ref\":\"A\", \"alt\":\"T\"}" -H 'Content-Type: application/json'
  ```
- POST a Call (change the `individual_id` and `variant_id` here to match your instance):
  ```
  curl -i http://0.0.0.0:3000/v1/calls -d "{\"individual_id\":\"0d583066-039a-4f61-832e-b0f8f5156f7d\", \"variant_id\":\"0e583067-039a-4f61-832e-b0f8f5156f7d\", \"genotype\":\"0/1\", \"format\":\"GQ:DP:HQ 48:1:51,51\"}" -H 'Content-Type: application/json'
  ```
- GET all Individuals with a specific Variant called (change the `variant_id` here to match your instance):
  ```
  curl -i "http://0.0.0.0:3000/v1/variants/0d583066-039a-4f61-832e-b0f8f5156f7d/individuals"
  ```

## For Developers

Note that most auto-generated code, including binary files such as `main`, are intentionally excluded from this repository.

If you would like to learn more about using this stack, especially with regards to its code-generation components, please take a look at the [DEVELOPER-GUIDE.md](https://github.com/CanDIG/go-model-service/blob/main/docs/DEVELOPER-GUIDE.md).
