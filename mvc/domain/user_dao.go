package domain

import (
	"fmt"
	"net/http"

	"github.com/albertoribeiro/golang-microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {ID: 123, FirstName: "Alberto", LastName: "Flavio", Email: "alberto.r.flavio@gmail.com"},
	}
)

// GetUser - return the user from the database
func GetUser(userID int64) (*User, *utils.ApplicationError) {
	if user := users[userID]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
