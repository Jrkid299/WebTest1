package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"AWtest1.jalenlamb.net/internals/validator/data"
)

// createSchoolHandler for the "POST /v1/entries" endpoint
func (app *application) createEntryHandler(w http.ResponseWriter, r *http.Request) {
	//Our target decode destination
	var input struct {
		Id       int    `json:"id"`
		Username string `json:"age"`
		Email    string `json:"email"`
	}
	// Initialize new JSON.Decoder instance
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		app.errorRepsonse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	//Display the request
	fmt.Fprintf(w, "%+v\n", input)
}

// showSchoolHandler for the "GET /v1/entries/:id" endpoint
func (app *application) showEntryHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	// Create a new instance of the school struct containing the ID we extracted
	// From our Url and some sample data
	user := data.User{
		Id:       id,
		Username: "Apple Tree",
		Email:    "2018118881@ub.edu.bz",
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"school": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
