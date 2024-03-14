package main

import (
	"errors"
	"fmt"
	"jsonHelpers"
	"net/http"
)

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := jsonHelpers.ReadJSON(w, r, &requestPayload)
	if err != nil {
		jsonHelpers.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		jsonHelpers.ErrorJSON(w, errors.New("invalid credentials"), http.StatusUnauthorized)
		return
	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		jsonHelpers.ErrorJSON(w, errors.New("invalid credentials"), http.StatusUnauthorized)
		return
	}

	responsePayload := jsonHelpers.JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("User %s logged in.", user.Email),
		Data:    user,
	}

	jsonHelpers.WriteJSON(w, http.StatusAccepted, responsePayload)
}
