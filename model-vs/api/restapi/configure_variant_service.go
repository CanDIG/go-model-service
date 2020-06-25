// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"net/http"

	"github.com/CanDIG/go-model-service/model-vs/api/restapi/handlers"
	"github.com/CanDIG/go-model-service/model-vs/api/restapi/operations"
	"github.com/CanDIG/go-model-service/utilities/log"
)

//go:generate swagger generate server --target .. --name model-vs --spec ../swagger.yml

func configureFlags(api *operations.VariantServiceAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.VariantServiceAPI) http.Handler {
	// Initialize custom logger. Configuration to the logger can be made here through this (or a similar) function
	log.Init()

	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.GetCallsHandler = operations.GetCallsHandlerFunc(func(params operations.GetCallsParams) middleware.Responder {
		return handlers.GetCalls(params)
	})
	api.GetIndividualsHandler = operations.GetIndividualsHandlerFunc(func(params operations.GetIndividualsParams) middleware.Responder {
		return handlers.GetIndividuals(params)
	})
	api.GetIndividualsByVariantHandler = operations.GetIndividualsByVariantHandlerFunc(func(params operations.GetIndividualsByVariantParams) middleware.Responder {
		return handlers.GetIndividualsByVariant(params)
	})
	api.GetOneCallHandler = operations.GetOneCallHandlerFunc(func(params operations.GetOneCallParams) middleware.Responder {
		return handlers.GetOneCall(params)
	})
	api.GetOneIndividualHandler = operations.GetOneIndividualHandlerFunc(func(params operations.GetOneIndividualParams) middleware.Responder {
		return handlers.GetOneIndividual(params)
	})
	api.GetOneVariantHandler = operations.GetOneVariantHandlerFunc(func(params operations.GetOneVariantParams) middleware.Responder {
		return handlers.GetOneVariant(params)
	})
	api.GetVariantsHandler = operations.GetVariantsHandlerFunc(func(params operations.GetVariantsParams) middleware.Responder {
		return handlers.GetVariants(params)
	})
	api.GetVariantsByIndividualHandler = operations.GetVariantsByIndividualHandlerFunc(func(params operations.GetVariantsByIndividualParams) middleware.Responder {
		return handlers.GetVariantsByIndividual(params)
	})
	api.PostCallHandler = operations.PostCallHandlerFunc(func(params operations.PostCallParams) middleware.Responder {
		return handlers.PostCall(params)
	})
	api.PostIndividualHandler = operations.PostIndividualHandlerFunc(func(params operations.PostIndividualParams) middleware.Responder {
		return handlers.PostIndividual(params)
	})
	api.PostVariantHandler = operations.PostVariantHandlerFunc(func(params operations.PostVariantParams) middleware.Responder {
		return handlers.PostVariant(params)
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
