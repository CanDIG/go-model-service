### This is the main Dockerfile for the go-model-service app,
### excluding the database which is built seperately.

# TODO is GOPATH going to cause us issues?

# Modify this line if you want to use a different stack/dependencies-image
FROM katpavlov/gms-deps-v1 AS webapp

ARG API_PATH

WORKDIR /go/src/github.com/CanDIG/go-model-service
COPY . .

# Check that all dependencies are present.
# See `./Dockerfile-gms-deps-v1` for more information.
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
RUN cd "$API_PATH" && swagger generate server -A variant-service swagger.yml

# Run a script to generate resource-specific request handlers for middleware,
# from the generic handlers defined in the model-vs/api/generics package,
# using the CanDIG-maintained CLI tool genny
###RUN $API_PATH/generate_handlers.sh

# Now that all the necessary boilerplate code has been auto-generated, compile 
# the server
###RUN go build -tags sqlite -o ./main $API_PATH/cmd/variant-service-server/main.go

#TODO rm
RUN ls

# Run the variant service
###EXPOSE 3000
###ENTRYPOINT "./main" --port=3000

# TODO write dockerignore file
