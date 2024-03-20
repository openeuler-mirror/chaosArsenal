/*
Copyright 2023 Sangfor Technologies Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"arsenal/restapi/operations"
)

//go:generate swagger generate server --target ../../swagger --name ChaosArsenalFaultInjection --spec ../arsenal-server-1.0.0.yaml --principal interface{}

func configureFlags(api *operations.ChaosArsenalFaultInjectionAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ChaosArsenalFaultInjectionAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.DeleteFaultsIDHandler == nil {
		api.DeleteFaultsIDHandler = operations.DeleteFaultsIDHandlerFunc(func(params operations.DeleteFaultsIDParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeleteFaultsID has not yet been implemented")
		})
	}
	if api.GetFaultsHandler == nil {
		api.GetFaultsHandler = operations.GetFaultsHandlerFunc(func(params operations.GetFaultsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetFaults has not yet been implemented")
		})
	}
	if api.PostFaultsHandler == nil {
		api.PostFaultsHandler = operations.PostFaultsHandlerFunc(func(params operations.PostFaultsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostFaults has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

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
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
