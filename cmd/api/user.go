package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"AWtest1.jalenlamb.net/internals/validator"
	"AWtest1.jalenlamb.net/internals/validator/data"
)

// createSchoolHandler for the "POST /v1/entries" endpoint
func (app *application) createUserHandler(w http.ResponseWriter, r *http.Request) {
	//Our target decode destination
	var input struct {
		Id       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}
	// Initialize new JSON.Decoder instance
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		app.errorRepsonse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	// Copy the values from the input struct to a new School struct
	user := &data.User{
		Username: input.Username,
		Email:    input.Email,
	}

	v := validator.New()

	//Check the map to determine if there were any validation errors
	if data.ValidateUser(v, user); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	//Create a user
	err = app.models.User.Insert(user)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
	// Create a Location header for the newly created resource/School
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/user/%d", user.Id))
	// Write the JSON response with 201 - Created status code with the body
	// being the School data and the header being the headers map
	err = app.writeJSON(w, http.StatusCreated, envelope{"user": user}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// showUserHandler for the "GET /v1/user/:id" endpoint
func (app *application) showUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	// Create a new instance of the school struct containing the ID we extracted
	// From our Url and some sample data
	user := data.User{
		Id:       id,
		Username: "Jalen",
		Email:    "2018118881@ub.edu.bz",
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"user": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
