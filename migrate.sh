#!/bin/bash

#TODO make this a makefile instead?

# Create a Postgres development database and migrate it to the schema
# defined in the model-vs/data directory, using the soda tool from pop.
# wait-for-it waits for the Postgres DB to be ready to accept connections.

docker-compose exec gms-webapp sh -c \
	"until pg_isready --timeout=60 --username=$POSTGRES_USER --host=database; \
	do sleep 1; \
	done &&\
	soda migrate up -c ./database.yml -e development -p model-vs/data/migrations"
