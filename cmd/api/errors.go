package main

import (
	"fmt"
	"net/http"
)

func (app *application) logError(r *http.Request, err error) {
	app.logger.Println(err)
}

// We want to send JSON-formatted error message
func (app *application) errorRepsonse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	//Create JSON response
	env := envelope{"error": message}
	err := app.writeJSON(w, status, env, nil)

	if err != nil {
		app.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// Server error response
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	// Welog the error
	app.logError(r, err)

	// Prepare a message with the error
	message := "the server encountered a problem and couldn't process the request"
	app.errorRepsonse(w, r, http.StatusInternalServerError, message)
}

// The not found response
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	// Create our message
	message := "The requested resource couldn't be found"
	app.errorRepsonse(w, r, http.StatusNotFound, message)
}

func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	//create msg
	message := fmt.Sprintf("The %s method is not supported for this resources", r.Method)
	app.errorRepsonse(w, r, http.StatusMethodNotAllowed, message)
}

// User provide bad request
func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {

	app.errorRepsonse(w, r, http.StatusBadRequest, err.Error())

}

// Validation error
func (app *application) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	app.errorRepsonse(w, r, http.StatusUnprocessableEntity, errors)
}
