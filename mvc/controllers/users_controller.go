package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/albertoribeiro/golang-microservices/mvc/utils"

	"github.com/albertoribeiro/golang-microservices/mvc/services"
)

// GetUser - return  a user
func GetUser(resp http.ResponseWriter, req *http.Request) {
	userIDParam := req.URL.Query().Get("user_ID")
	userID, err := strconv.ParseInt(userIDParam, 10, 64)

	resp.Header().Set("Content-Type", "application/json")

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_ID must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	user, apiErr := services.GetUser(userID)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	jsonValue, _ := json.Marshal(user)

	resp.Write(jsonValue)

}
