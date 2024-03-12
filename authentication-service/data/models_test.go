package data

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

type passwordTest struct {
	description  string
	userPassword string
	testPassword string
	expected     bool
}

var testCases = []passwordTest{
	{
		description:  "user.PasswordMatches() expected to match",
		userPassword: "test123",
		testPassword: "test123",
		expected:     true,
	},
	{
		description:  "user.PasswordMatches() expected to not match",
		userPassword: "test123",
		testPassword: "test321",
		expected:     false,
	},
}

func TestPassword(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(testCase.userPassword), 12)
			user := User{
				Password: string(hashedPassword),
			}
			result, _ := user.PasswordMatches(testCase.testPassword)
			if result != testCase.expected {
				t.Errorf("user.PasswordMatches(%s) = %v, expected %v", testCase.testPassword, testCase.expected, result)
			}
		})
	}
}
