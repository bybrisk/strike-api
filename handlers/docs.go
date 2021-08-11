// Package classification of Strike-Bots API
//
// Documentation for Strike-Bots API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta

package handlers
import "github.com/bybrisk/strike-api/data"

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handlers

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// Response with user room
// swagger:response roomPostResponse
type roomPostResponseWrapper struct {
	// Data structure with room info of a user
	// in: body
	Body data.Response_wrapper_structure
}

// swagger:parameters getMyRoom
type registerAUserParmsWrapper struct {
	// Data structure to accept the request from the app.
	// in: body
	// required: true
	Body data.Strike_Meta_Request_Structure
}