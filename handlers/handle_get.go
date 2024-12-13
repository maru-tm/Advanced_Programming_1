package handlers

import (
	"json_responses_requests/utils"
	"net/http"
)

func HandleGet(w http.ResponseWriter, r *http.Request) {
	response := utils.Response{
		Status:  "success",
		Message: "Send a POST request with JSON data to this endpoint",
	}
	utils.RespondWithJSON(w, http.StatusOK, response)
}
