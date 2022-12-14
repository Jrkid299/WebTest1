package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	//create a new httprouter router instances
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/user", app.listUsersHandler)
	router.HandlerFunc(http.MethodPost, "/v1/user", app.createUserHandler)
	router.HandlerFunc(http.MethodGet, "/v1/user/:id", app.showUserHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/user/:id", app.updateUserHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/user/:id", app.deleteUserHandler)

	return router
}
