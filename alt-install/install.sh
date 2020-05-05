#!/bin/bash

# This script installs the service by auto-generating the boilerplate code necessary
# for running the service.
# It also creates a new development database to write into.

# Set up working directory for this script
export WORK_DIR=$GOPATH/src/github.com/CanDIG/go-model-service
cd $WORK_DIR

# Set up path to database configuration file. 
# This database config file must be kept within one of the following paths, relative to
# the file attempting to make the database connection:
# "", "./config", "/config", "../", "../config", "../..", "../../config", "APP_PATH", "POP_PATH"
# For more information: https://github.com/gobuffalo/pop/blob/master/config.go
export POP_PATH=$WORK_DIR

# Directory to API files, to be used by codegen
export API_PATH=$WORK_DIR/model-vs/api

# Use the dep tool to install all project import dependencies in a
# new vendor directory.
# It is run vendor-only to avoid modification of Gopkg.lock, which
# contains information about vital sub-packages for go-swagger that
# can not be explicitly constrained in Gopkg.toml.
# For example, package "github.com/go-openapi/runtime/flagext" is
# required by go-swagger but is *not* solved into Gopkg.lock if
# `dep ensure` is run prior to `swagger validate`. Therefore it is
# important to read the existing Gopkg.lock file in the initial
# installation, rather than solve for a new one.
# For more information, see: https://golang.github.io/dep/docs/ensure-mechanics.html
dep ensure -vendor-only

# Create a sqlite3 development database and migrate it to the schema
# defined in the model-vs/data directory, using the soda tool from pop
soda create -c ./database.yml -e development
soda migrate up -c ./database.yml -e development -p model-vs/data/migrations

# Swagger generate the boilerplate code necessary for handling API requests
# from the model-vs/api/swagger.yml template file.
# This will generate a server named variant-service. The name is important for 
# maintaining compatibility with the configure_variant_service.go middleware 
# configuration file.
cd $API_PATH
swagger generate server -A variant-service swagger.yml
cd $WORK_DIR

# Run a script to generate resource-specific request handlers for middleware,
# from the generic handlers defined in the model-vs/api/generics package,
# using the CanDIG-maintained CLI tool genny
$API_PATH/generate_handlers.sh

# Now that all the necessary boilerplate code has been auto-generated, compile the server
go build -tags sqlite -o ./main $API_PATH/cmd/variant-service-server/main.go

# Unset working directory since it is no longer needed
cd $WORK_DIR
unset WORK_DIR

# Now you can run ./main to launch the variant service
# ./main --port=3000
