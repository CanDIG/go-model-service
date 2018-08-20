#!/bin/bash

# This script installs the service by auto-generating the boilerplate code necessary
# for running the service.
# It also creates a new development database to write into.

# Use the dep tool to install all project import dependencies in a
# new vendor directory
dep ensure

# Create a sqlite3 development database and migrate it to the schema
# defined in the model-vs/data directory, using the soda tool from pop
cd model-vs/data
soda create -e development
soda migrate up -e development
cd ../..

# Swagger generate the boilerplate code necessary for handling API requests
# from the model-vs/api/swagger.yml template file
cd model-vs/api
swagger generate server -A variant-service swagger.yml # This will generate a server named variant-service, which is important for maintaining compatibility with the configure_variant_service.go middleware configuration file.
cd ../..

# Run a script to generate resource-specific request handlers for middleware,
# from the generic handlers defined in the model-vs/api/generics package,
# using the CanDIG-maintained CLI tool genny
./model-vs/api/generate_handlers.sh

# Now that all the necessary boilerplate code has been auto-generated, compile the server
go build -tags sqlite3 -o ./main model-vs/api/cmd/variant-service-server/main.go
