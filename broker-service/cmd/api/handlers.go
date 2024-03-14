package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"jsonHelpers"
	"log"
	"net/http"
)

const AuthenticationServiceURL = "http://authentication-service/authenticate"

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonHelpers.JsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	err := jsonHelpers.WriteJSON(w, http.StatusOK, payload)
	if err != nil {
		log.Print(err.Error())
	}
}

func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload

	err := jsonHelpers.ReadJSON(w, r, &requestPayload)
	if err != nil {
		log.Print(err.Error())
		jsonHelpers.ErrorJSON(w, err)
		return
	}

	switch requestPayload.Action {
	case "auth":
		app.authenticate(w, requestPayload.Auth)
	default:
		jsonHelpers.ErrorJSON(w, errors.New("unknown action"))
	}
}

func (app *Config) authenticate(w http.ResponseWriter, authPayload AuthPayload) {
	jsonData, err := json.Marshal(authPayload)
	if err != nil {
		log.Printf("Error on json.Marshal(%v): %s", authPayload, err.Error())
		jsonHelpers.ErrorJSON(w, fmt.Errorf("could not authenticate with %s", string(jsonData)))
		return
	}

	request, err := http.NewRequest("POST", AuthenticationServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error on http.NewRequest: %s", err)
		jsonHelpers.ErrorJSON(w, errors.New("could not authenticate"))
		return
	}

	httpClient := http.Client{}
	response, err := httpClient.Do(request)
	if err != nil {
		log.Printf("Error on httpClient.Do: %s", err)
		jsonHelpers.ErrorJSON(w, errors.New("could not authenticate"))
		return
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		jsonHelpers.ErrorJSON(w, errors.New("invalid credentials"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		log.Printf("Error on authentication service")
		jsonHelpers.ErrorJSON(w, errors.New("could not authenticate"), response.StatusCode)
		return
	}

	var authResponse jsonHelpers.JsonResponse
	err = json.NewDecoder(response.Body).Decode(&authResponse)
	if err != nil {
		log.Printf("Error on json Decode %s", err)
		jsonHelpers.ErrorJSON(w, errors.New("could not authenticate"))
		return
	}
	if authResponse.Error {
		log.Printf("Error on auth response: %s", authResponse.Message)
		jsonHelpers.ErrorJSON(w, errors.New("could not authenticate"))
		return
	}

	payloadResponse := jsonHelpers.JsonResponse{
		Error:   false,
		Message: "authenticated",
		Data:    authResponse.Data,
	}

	jsonHelpers.WriteJSON(w, http.StatusAccepted, payloadResponse)
}
