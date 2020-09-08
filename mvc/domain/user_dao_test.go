package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Using testify libraly -  github.com/stretchr/testify: more readable
// A Test have 3 parts : Initialization, Execution and Validation

func TestGetUserNoUserFound(t *testing.T) {
	// Initialization

	// Execution
	user, err := GetUser(0)

	// Validation
	assert.Nil(t, user, "we were not expecting a user with id 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user 0 was not found", err.Message)

}

func TestGetUserNoError(t *testing.T) {
	// Initialization

	// Execution
	user, err := GetUser(123)

	// Validation
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.ID)
	assert.EqualValues(t, "Alberto", user.FirstName)
	assert.EqualValues(t, "Flavio", user.LastName)
	assert.EqualValues(t, "alberto.r.flavio@gmail.com", user.Email)

}
