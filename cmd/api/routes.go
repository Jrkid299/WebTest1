package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	//create a new httprouter router instances
	router := httprouter.New()
	//router.NotFound = http.HandlerFunc(app.notFoundResponse)
	//router.MethodNotAllowed = http.HandlerFunc(app.routes().MethodNotAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	return router
}
