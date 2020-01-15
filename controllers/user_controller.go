package controllers

import (
	"encoding/json"
	"go_microservices/services"
	"go_microservices/utils"
	"net/http"
	"strconv"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
	userIdParam := req.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	res.Header().Set("Content-Type", "application/json")

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}

		jsonVal, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.StatusCode)
		res.Write(jsonVal)
		return
	}

	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		res.WriteHeader(apiErr.StatusCode)
		jsonVal, _ := json.Marshal(apiErr)
		res.Write(jsonVal)
		return
	}

	jsonVal, _ := json.Marshal(user)
	res.Write(jsonVal)
}