package main

import (
	"authentication/data"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

type UserDBMock struct {
}

func (u UserDBMock) GetByEmail(email string) (*data.User, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("test123"), 12)
	user := data.User{
		Email:    "test123",
		Password: string(hashedPassword),
	}
	return &user, nil
}

func TestAuthenticateSuccessful(t *testing.T) {
	app := Config{
		Models: data.Models{
			User: UserDBMock{},
		},
		DB: nil,
	}
	response := httptest.NewRecorder()
	requestPayload := `
{
	"email": "test123",
	"password": "test123"
}`
	body := strings.NewReader(requestPayload)

	request := httptest.NewRequest("POST", "/authenticate", body)
	app.Authenticate(response, request)

	if response.Result().StatusCode != http.StatusOK {
		t.Fatalf(`Authenticate() with status code %d, want %d\nerror: %s`, response.Result().StatusCode, http.StatusOK, response.Body)
	}
}
