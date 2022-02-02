/*
 * Project X
 *
 * OpenAPI definition for project X endpoint and resources
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// TestsApiController binds http requests to an api service and writes the service results to the http response
type TestsApiController struct {
	service TestsApiServicer
	errorHandler ErrorHandler
}

// TestsApiOption for how the controller is set up.
type TestsApiOption func(*TestsApiController)

// WithTestsApiErrorHandler inject ErrorHandler into controller
func WithTestsApiErrorHandler(h ErrorHandler) TestsApiOption {
	return func(c *TestsApiController) {
		c.errorHandler = h
	}
}

// NewTestsApiController creates a default api controller
func NewTestsApiController(s TestsApiServicer, opts ...TestsApiOption) Router {
	controller := &TestsApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all of the api route for the TestsApiController
func (c *TestsApiController) Routes() Routes {
	return Routes{ 
		{
			"CreateTest",
			strings.ToUpper("Post"),
			"/tests",
			c.CreateTest,
		},
		{
			"GetTests",
			strings.ToUpper("Get"),
			"/tests",
			c.GetTests,
		},
	}
}

// CreateTest - Create new test
func (c *TestsApiController) CreateTest(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.CreateTest(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetTests - Create new test
func (c *TestsApiController) GetTests(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetTests(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
