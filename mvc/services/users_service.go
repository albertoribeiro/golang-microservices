package services

import (
	"github.com/albertoribeiro/golang-microservices/mvc/domain"
	"github.com/albertoribeiro/golang-microservices/mvc/utils"
)

// GetUser - return a user from domain
func GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userID)
}
