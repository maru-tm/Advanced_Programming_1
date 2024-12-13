package handlers

import (
	"encoding/json"
	"fmt"
	"json_responses_requests/utils"
	"net/http"
)

type Request struct {
	Message string `json:"message"`
}

func HandlePost(w http.ResponseWriter, r *http.Request) {
	var reqBody Request
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&reqBody); err != nil {
		response := utils.Response{
			Status:  "fail",
			Message: "Invalid JSON format",
		}
		utils.RespondWithJSON(w, http.StatusBadRequest, response)
		return
	}
	if reqBody.Message == "" {
		response := utils.Response{
			Status:  "fail",
			Message: "Invalid JSON message",
		}
		utils.RespondWithJSON(w, http.StatusBadRequest, response)
		return
	}
	fmt.Printf("Received message: %s\n", reqBody.Message)
	response := utils.Response{
		Status:  "success",
		Message: "Data successfully received",
	}
	utils.RespondWithJSON(w, http.StatusOK, response)
}

func HandleMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	response := utils.Response{
		Status:  "fail",
		Message: "Method not allowed",
	}
	utils.RespondWithJSON(w, http.StatusMethodNotAllowed, response)
}
