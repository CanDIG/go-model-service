### This is the main Dockerfile for the go-model-service app,
### excluding the database which is built seperately.

# TODO is GOPATH going to cause us issues?

# Modify this line if you want to use a different stack-image
FROM katpavlov/gms-stack-v1 AS webapp

WORKDIR /go/src/github.com/CanDIG/go-model-service
COPY . .

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
# For more information, see: 
# https://golang.github.io/dep/docs/ensure-mechanics.html
RUN dep ensure -vendor-only

# TODO fix this to suit a persistent postgres database
# Create a sqlite3 development database and migrate it to the schema
# defined in the model-vs/data directory, using the soda tool from pop
RUN soda create -c ./database.yml -e development
RUN soda migrate up -c ./database.yml -e development -p model-vs/data/migrations

# Swagger generate the boilerplate code necessary for handling API requests
# from the model-vs/api/swagger.yml template file.
# This will generate a server named variant-service. The name is important for 
# maintaining compatibility with the configure_variant_service.go middleware 
# configuration file.
RUN cd $API_PATH 
###&& swagger generate server -A variant-service swagger.yml

# Run a script to generate resource-specific request handlers for middleware,
# from the generic handlers defined in the model-vs/api/generics package,
# using the CanDIG-maintained CLI tool genny
###RUN $API_PATH/generate_handlers.sh

# Now that all the necessary boilerplate code has been auto-generated, compile 
# the server
###RUN go build -tags sqlite -o ./main $API_PATH/cmd/variant-service-server/main.go

# Run the variant service
###EXPOSE 3000
###ENTRYPOINT "./main" --port=3000

# TODO write dockerignore file
