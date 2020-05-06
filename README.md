# Go Model Service

Based on Jonathan Dursi's [OpenAPI variant service demo](https://github.com/ljdursi/openapi_calls_example), this toy service demonstrates the go-swagger/pop stack with CanDIG API best practices.

[![Build Status](https://travis-ci.org/CanDIG/go-model-service.svg?branch=master)](https://travis-ci.org/CanDIG/go-model-service)
[![Go Report Card](https://goreportcard.com/badge/github.com/candig/go-model-service)](https://goreportcard.com/report/github.com/candig/go-model-service)

## Quick Start

Once you have [installed the stack](#installing-the-stack), run the following commands (as described in [Scripted Installation](#scripted-installation-of-the-go-model-service) and [Running the Service](#running-the-service)):

1. Checkout this go-model-service into `$GOPATH/src/github.com/CanDIG` via:
  ```
  $ cd $GOPATH/src/github.com/CanDIG
  $ git checkout https://github.com/CanDIG/go-model-service.git
  $ cd go-model-service
  ```
2. Run the installation script from the project root directory. It is important that you run it from within the active shell, so that pertinent environment variables are set in the same shell as the one that will be running the program via `$ ./main`.
  ```
  $ cd $GOPATH/src/github.com/CanDIG/go-model-service
  $ . ./install.sh
  ```
3. From the project root directory, run the server on a port of your choosing (eg. port 3000):
  ```
  $ cd $GOPATH/src/github.com/CanDIG/go-model-service
  $ ./main --port=3000
  ```

## Stack

- [Sqlite3](https://www.sqlite.org/index.html) database backend
- [Go](https://golang.org/) (Golang) backend
- Gobuffalo [pop](https://github.com/gobuffalo/pop) is used as an orm-like for interfacing between the go code and the sqlite3 database. The `soda` CLI auto-generates boilerplate Go code for models and migrations, as well as performing migrations on the database. `fizz` files are used for defining database migrations in a Go-like syntax (see [syntax documentation](https://gobuffalo.io/en/docs/db/fizz/).)
- [Go-swagger](https://goswagger.io/) auto-generates boilerplate Go code from a `swagger.yml` API definition. [Swagger](https://swagger.io/) tooling is based on the [OpenAPI](https://www.openapis.org/) specification.
- [dep](https://golang.github.io/dep/docs/introduction.html) is used for dependency management, specifically for libraries imported by Go code in the project.
- [genny](https://github.com/CanDIG/genny) is a code-generation solution to generics in Go.
- Gobuffalo [validate](https://github.com/gobuffalo/validate) is a framework used for writing custom validators. Some of their validators in the `validators` package are used as-is.

## Installation

Following the [installation of the stack](#installing-the-stack), There are two sets of installation instructions provided for this project.

1. The [Descriptive Installation](#descriptive-installation-of-the-go-model-service) of this project is verbose so as to offer a tutorial on which aspects of the server backend are auto-generated by tools in the stack, and which files act as configuration to this auto-generation step. You can learn more about the code-generating tools employed here in the [For Developers](#for-developers) section of this README.

2. The [Scripted Installation](#scripted-installation-of-the-go-model-service) of this project hides the individual installation steps in a script and allows for a quick-start approach to running the server.

### Installing the Stack

Prior to installing new programs, run `$ which <program-name>` to check if it is already installed on your machine. If there is a blank output rather than a path to the program binary, it needs to be installed.

See `install_dep.sh` for an example of the installation of steps 3-7. It s not recommended that you run this script locally, as some of these programs may already be installed on your system and the version of some tools may matter (eg. Go-swagger v0.16.0).

1. [Install Go](https://golang.org/doc/install). Make sure to set up the `$PATH` and `$GOPATH` environment variables as described in bash below, and to understand the expected contents of the three `$GOPATH` subdirectories: `$GOPATH/src`, `$GOPATH/pkg`, and `$GOPATH/bin`.
  ```
  export GOROOT=/usr/local/go # Set $GOROOT to whatever your installation directory for go is, eg. /usr/local/go
  export PATH=$PATH:$GOROOT/bin # Append go bin to path
  export GOPATH=$HOME/go # Set $GOPATH to where you want your go source (src), binaries (bin), and packaged (pkg) to lie, eg. $HOME/go
  ```
2. [Install gcc](https://gcc.gnu.org/install/).
3. [Install sqlite3](https://www.tutorialspoint.com/sqlite/sqlite_installation.htm).
4. [Install dep](https://golang.github.io/dep/docs/installation.html).
5. [Install go-swagger](https://goswagger.io/install.html) (release 0.16.0 strongly recommended.)
6. [Install pop](https://github.com/gobuffalo/pop). See the [Unnoficial Pop Book](https://andrew-sledge.gitbooks.io/the-unofficial-pop-book/content/installation.html) for instructions. Make sure to include sqlite3 support with `tags sqlite` in your installation commands, as follows:
  ```
  $ go get -u -v -tags sqlite github.com/gobuffalo/pop/...
  $ go install -tags sqlite github.com/gobuffalo/pop/soda
  ```
7. [Install genny](https://github.com/CanDIG/genny#genny---generics-for-go).

### Descriptive Installation of the Go-Model-Service

1. Checkout this go-model-service into `$GOPATH/src/github.com/CanDIG` via:
  ```
  $ cd $GOPATH/src/github.com/CanDIG
  $ git checkout https://github.com/CanDIG/go-model-service.git
  $ cd go-model-service
  ```
2. In the root directory of this project (ie. the directory where `Gopkg.lock` and `Gopkg.toml` are found) use the `dep` CLI tool to install all project import dependencies in a new `vendor` directory. See this README's [developers' notes for dep](#dep) for an explanation of the `-vendor-only` option used here.
  ```
  $ cd $GOPATH/src/github.com/CanDIG/go-model-service
  $ dep ensure -vendor-only
  ```
3. Set the path for for the database and its configuration file. Read more about the `POP_PATH` variable [here](https://github.com/CanDIG/go-model-service/tree/upgrade-stack#pop_path).
  ```
  $ export POP_PATH=$GOPATH/src/github.com/CanDIG/go-model-service/database
  ```
4. Create a sqlite3 development database and migrate it to the schema defined in the `model-vs/data` directory, using the pop CLI tool `soda`:
  ```
  $ cd $GOPATH/src/github.com/CanDIG/go-model-service
  $ soda create -c ./database.yml -e development
  $ soda migrate up -c ./database.yml -e development -p model-vs/data/migrations
  ```
5. Generate the boilerplate code necessary for handling API requests, from the `model-vs/api/swagger.yml` template file, with the Go-swagger CLI tool `swagger`. The following commands will generate a server named `variant-service`. This name is important for maintaining compatibility with the `configure_variant_service.go` middleware configuration file.
  ```
  $ cd $GOPATH/src/github.com/CanDIG/go-model-service/model-vs/api
  $ swagger generate server -A variant-service swagger.yml
  ```
6. Run a script to generate resource-specific request handlers for middleware, from the generic handlers defined in the `model-vs/api/generics` package, using the CanDIG-maintained CLI tool `genny`:
  ```
  $ cd $GOPATH/src/github.com/CanDIG/go-model-service/model-vs/api
  $ ./generate_handlers.sh
  ```
7. Now that all the necessary boilerplate code has been auto-generated, compile the server by running:
  ```
  $ cd $GOPATH/src/github.com/CanDIG/go-model-service
  $ go build -tags sqlite model-vs/api/cmd/variant-service-server/main.go
  ```

### Scripted Installation of the Go-Model-Service

1. Checkout this go-model-service into `$GOPATH/src/github.com/CanDIG` via:
  ```
  $ cd $GOPATH/src/github.com/CanDIG
  $ git checkout https://github.com/CanDIG/go-model-service.git
  $ cd go-model-service
  ```

2. Run the installation script from the project root directory:
  ```
  $ cd $GOPATH/src/github.com/CanDIG/go-model-service
  $ ./install.sh
  ```

## Running The Service

From the project root directory, run the server on a port of your choosing (eg. port 3000):
```
$ cd $GOPATH/src/github.com/CanDIG/go-model-service
$ ./main --port=3000
```

### Request examples

Some examples for correctly-formatted CURL requests to the server:

#### GET

##### GET Individuals

by ID

`$ curl -i "localhost:3000/individuals/0d583066-039a-4f61-832e-b0f8f5156f7d"`

all

`$ curl -i "localhost:3000/individuals"`

by Variant

`$ curl -i "localhost:3000/variants/0d583066-039a-4f61-832e-b0f8f5156f7d/individuals"`

##### GET Variants

by ID

`$ curl -i "localhost:3000/variants/0d583066-039a-4f61-832e-b0f8f5156f7d"`

by parameter

`$ curl -i "localhost:3000/variants?chromosome=chr1&start=3&end=105"`

by Individual

`$ curl -i "localhost:3000/individuals/0d583066-039a-4f61-832e-b0f8f5156f7d/variants"`

##### GET Calls

by ID

`$ curl -i "localhost:3000/calls/0d583066-039a-4f61-832e-b0f8f5156f7d"`

all

`$ curl -i "localhost:3000/calls"`

#### POST

##### POST Individuals

`$ curl -i localhost:3000/individuals -d "{\"description\":\"Subject 17\"}" -H 'Content-Type: application/json'`

##### POST Variants

`$ curl -i localhost:3000/variants -d "{\"name\":\"rs7054258\", \"chromosome\":\"chr1\", \"start\":5, \"ref\":\"A\", \"alt\":\"T\"}" -H 'Content-Type: application/json'`

##### POST Calls

`$ curl -i localhost:3000/calls -d "{\"individual_id\":\"0d583066-039a-4f61-832e-b0f8f5156f7d\", \"variant_id\":\"0e583067-039a-4f61-832e-b0f8f5156f7d\", \"genotype\":\"0/1\", \"format\":\"GQ:DP:HQ 48:1:51,51\"}" -H 'Content-Type: application/json'`

## For Developers

This section contains tips and tricks for using the code-generation tools presented in this model service.

Note that most auto-generated code has been excluded from this repository, and that instead there are instructions for generating this code locally provided in the [Descriptive Installation](#descriptive-installation-of-the-go-model-service) section of this README. Binary files such as main and development.sqlite have also been excluded from this repository.
The exclusion of these files is considered to be best practice for version control repository maintenance. However, one-time auto-generated files that are *not* re-generated (and are therefore safe to edit) should be pushed to this repository. `model-vs/api/restapi/configure_variant_service.go` and the `model-vs/data/models` package are examples of such safe-to-edit auto-generated code.

### Dep

In the initial installation of the service, the vendor-building step is run with `$ dep ensure -vendor-only`. This is to avoid modification of Gopkg.lock, which contains information about vital sub-packages for go-swagger that can not be explicitly constrained in Gopkg.toml.

For example, package "github.com/go-openapi/runtime/flagext" is required by go-swagger but is *not* solved into Gopkg.lock if `dep ensure` is run (without the `vendor-only` tag) prior to `swagger validate`. Therefore it is important to read the existing Gopkg.lock file in the initial installation, rather than solve for a new one.

For more information, see the [dep documentation](https://golang.github.io/dep/docs/ensure-mechanics.html).

### Go-Swagger

Go-Swagger reads the API definitions for a service and automatically generates the boilerplate Go code needed to build the server.

See [goswagger.io](https://goswagger.io/) for installation instructions, tutorials, use-cases, etc. If you find yourself having trouble with the installation, check the [prerequisites](https://goswagger.io/generate/requirements.html). The [Todo List Tutorial](https://goswagger.io/tutorial/todo-list.html) (Simple Server) is a good place to start if you've never used the go-swagger before. 

Go-Swagger uses Swagger 2.0, which is based on the OpenAPI specification. "Swagger" and "OpenAPI" are often used interchangeably. See [this post](https://swagger.io/blog/api-strategy/difference-between-swagger-and-openapi/) for an explanation of the relationship between the two.

#### Generating The Server

The API definitions are written by the developer in a `swagger.yml` file. To validate that the `swagger.yml` file follows the specification, run
`$ swagger validate <path-to-target-swagger.yml>`.

To auto-generate a server based on the entities and endpoints described in the `swagger.yml` file, run
`$ swagger generate server -A <server-name> <path-to-target-swagger.yml>.`
For example, for this service, from the `go-model-service/model-vs/api` directory, you would run the following to re-generate the server:
`$ swagger generate server -A model-vs swagger.yml`

The backend can now be implemented by modifying the endpoint handlers in `restapi/configure_<server-name>.go`. The connection to the data backend is made in these handlers. Other configuration such as middleware setup is also written in this file, in its respective methods.

##### Adding New Paths

To prevent overwrite of the backend implementation, the `restapi/configure_<server-name>.go` file is not re-generated if it already exists. Therefore, if new paths are added in the `swagger.yml` file, new handlers for those paths will not be automatically generated into the existing `restapi/configure_<server-name>.go` file. Moving the existing file to a different directory will allow swagger to generate the configuration file with the up-to-date set of handlers upon the next call of `swagger generate server`. The two copies of the file can then be reconciled to include both the new handlers and the previously-implemented ones. 

#### Boilerplate Code and Directory Structure

All files in the api directory are auto-generated (and auto-replaced upon calling
`$ swagger generate server <path-to-target-swagger.yml>`) except for the following:
- swagger.yml: The swagger definition.
- configure_variant_service.go: Auto-generated but safe to edit.
- main: Generated by calling `$ go build cmd/model-vs-server/main.go`

The auto-generated boilerplate code includes:
- models
The API-facing models for entities are generated into the `models` package.
Models are generated from the `definitions` defined in the `swagger.yml` file.
- endpoints
The endpoints for the API are defined in `paths` in the `swagger.yml` file, and from their definition the `operations` package is populated with endpoint parameters, validation, responses, URL building, etc.
However the backend handlers for these endpoints, ie. what is done with the received request, must be written manually in confifure_variant_service.go, in the configureAPI method. By default, the handlers return 501: Not Implemented responses.
- server
The go server files and main.go are auto-generated.
- configuration
The `configure_variant_service.go` file is auto-generated but safe to edit. This is where the manually written backend goes, where requests are handled following their automatic transformation into go stucts, and where responses/payloads are assigned.
Middleware can be plugged in here.
The connection to the data backend/memory store (ie. the ORM and/or database) should be made here.

### GoBuffalo Pop

Pop is an ORM-like that is used to interface between a go backend and a database.

See the [pop README](https://github.com/gobuffalo/pop#pop--) for installation and use instructions. Most documentation is now maintained at [gobuffalo.io](https://gobuffalo.io/en/docs/db/getting-started/) in the `Database` section. There is also an [Unofficial pop Book](https://andrew-sledge.gitbooks.io/the-unofficial-pop-book/content/) with tutorials, [Quick Start](https://andrew-sledge.gitbooks.io/the-unofficial-pop-book/content/installation.html) being a good place to begin.

Since the database used in this project is `sqlite3`, there are slight modifications that must be made to some commands in the form of a `-tags sqlite` option. These are detailed in the [Installing CLI Support](https://github.com/gobuffalo/pop#installing-cli-support) section of the Pop documentation.

#### Pop Migrations

Soda is a CLI tool for generating pop migration files and models, as well as for running up- and down-migrations. Migrations are described in `.fizz` or `.sql` files. Files for simple migrations such as adding/dropping columns are auto-generated by soda. For more complicated migrations, the migration files must be manually populated with explicit instructions.

Fizz provides a Go-like syntax for writing migrations, but [you may instead opt to write SQL migrations](https://github.com/gobuffalo/pop#generating-migrations). The fizz syntax is described [here](https://gobuffalo.io/en/docs/db/fizz/).

##### Migrating Pop Models

*Note: Pop models are distinct from go-swagger models! Go-swagger models represent API calls in golang, while pop models represent database entities in golang. See the [go-swagger docs](https://goswagger.io/use/models/schemas.html) to read about go-swagger models.*

Pop models are go files that describe database entities in terms of the go language. Each pop model corresponds to a table in the database. Soda can generate models from command-line input. When a migration modifies the database table that a model corresponds to, the associated model file must be manually edited.

For example, if you add a `province` column to the `individual` table in a migration, the `individual.go` model must have that field added to its `type Individual struct`. You may also want to add validations for this new field in the `Validate` method of the `individual.go` model.

#### Validating Pop Models

The `Validate` method contained in each model Go file, called upon each `ValidateAndSave` (or similar) call, checks that the data being pushed to the database meets desired constraints.

The `validators` package from Gobuffalo [validate](https://github.com/gobuffalo/validate) is a set of validators automatically imported by Pop. The `validate` package also allows for the creation of custom validators. See the `tools/validators` directory for a simple example, or  the Unofficial pop Book's tutorial on [Writin Validations](https://andrew-sledge.gitbooks.io/the-unofficial-pop-book/content/advanced-topics/writing-validations.html) for a more complex example.

#### Handling Nulls With Pop

There is some subtlety to representing database tables with Go structs. Since Go only allows `nil` values for pointers, a work-around is needed to handle nulls retrieved from the database, or to validate on nulls in data being pushed to the database.

By default, for a db datatype that is pop-converted into a non-nillable type in Go, null-values from the database are transformed into zero-values. For example, the `Chromosome` field of the `Variant` model (in `variant.go`) is of type string, and if a value for `chromosome` is not supplied in an entry, the value of the `Chromosome` field is `""`. The validators that check *required* fields provided in the `validators` package only check for these fields having a non-zero value. Thus, no distinction is made between zero-valued entries such as `0` or `""` and `nil`. These validators are insufficient for cases where when zero-valued entries are acceptable but `nil` entries are not.

This project uses the `pop/nulls` package to handle non-nullable fields that should be permitted to have zero values, such as the `Start` field of the `Variant` model. This field is of type nulls.Int, which is able to differentiate between null values and zero values. A custom `int_is_not_null.go` validator is needed to validate this datatype.

#### POP_PATH

The database configuration file used by pop must be kept within one of the following paths, relative to the file attempting to connect with the database: 
  ```
  "", "./config", "/config", "../", "../config", "../..", "../../config", "APP_PATH", "POP_PATH"
  ```
Therefore the `POP_PATH` variable may be set to point to the folder containing the `database.yml` config file. See the [pop configuration code](https://github.com/gobuffalo/pop/blob/master/config.go) for more information.

### Genny

Genny is a code-generation solution to the lack of generics in Go. We use it handle the myriad of similarly-named auto-generated go files created by the pop and Go-Swagger tools.

See the [genny README](https://github.com/CanDIG/genny#genny---generics-for-go) for usage instructions and examples. It operates essentially like a copy-then-search-and-replace for generating typed functions out of functions written for generic types. This reduces code duplication by allowing for the development of type-agnostic code.

Run `$ model-vs/api/generate_handlers.sh` to re-generate the handlers in the `github.com/CanDIG/go-model-service/model-vs/api/restapi/handlers` package.

There remain several issues with the genny tool that block the complete integration of generic code-gen into our project. Please contribute to the resolution of these issues in our [genny](https://github.com/CanDIG/genny/issues) repository.

### Dredd

Dredd can use the `swagger.yml` api definitions to run automated testing on the API. It uses `example` headings defined in the swagger file to generate its input bodies and parameters.

It can be helpful to only populate the examples with the required fields, to check that all non-required fields are handled properly for possible nil pointers.

### Docker
Docker containerizes processes so that they can be run in predictable environments, and minimize interference with the host environment. See the [Docker tutorial](https://docker-curriculum.com/) and [Dockerfile reference](https://docs.docker.com/engine/reference/builder/#add) for more information.

Docker Compose automates the building of these containers. See the [tutorial](https://docs.docker.com/compose/gettingstarted/) and [reference](https://docs.docker.com/compose/compose-file/) documentation.

It is common to run `docker` and `docker-compose` commands as `sudo`, because the socket used by the Docker daemon is owned by the `root` user. There is a noteworthy downside to running  commands as `sudo`: environment variables are not kept by default. This especially causes problems when building go services, which often depend upon the `$GOPATH` environment variable.

There are two solutions to this issue.
1) Always run sudo with the `-E` option (ie. `sudo -E <command>`) to explicitly preserve environment variables.
2) Create a user group to act as `root` when docker is run, as detailed [here](https://docs.docker.com/engine/install/linux-postinstall/#manage-docker-as-a-non-root-user). Be sure to read about the [security consequences](https://docs.docker.com/engine/security/security/#docker-daemon-attack-surface) in advance of doing this!
