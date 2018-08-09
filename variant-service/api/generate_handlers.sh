#!/bin/bash

# TODO make smarter by only re-generating files if they've been modified (makefile?)

# Generate handler utilities for the following resources: Individual, Variant
cat ./generics/generic_resource_utilities.go.tmpl | genny gen "Resource=Individual,Variant" > ./restapi/handlers/resource_utilities.go

# Generate POST handlers for the following resources: Individual, Variant
cat ./generics/generic_post.go.tmpl | genny gen "Resource=Individual,Variant" > ./restapi/handlers/post.go

# Generate GET (many) handlers for the following resources: Individual, Variant
cat ./generics/generic_get_many.go.tmpl | genny gen "Resource=Individual,Variant" > ./restapi/handlers/get_many.go