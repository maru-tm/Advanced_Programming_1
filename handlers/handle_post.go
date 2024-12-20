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
	var rawData map[string]interface{}

	// Декодируем в map для проверки типа полей
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&rawData); err != nil {
		response := utils.Response{
			Status:  "fail",
			Message: "Invalid JSON format",
		}
		utils.RespondWithJSON(w, http.StatusBadRequest, response)
		return
	}

	// Проверяем, что поле Message существует и имеет тип string
	message, ok := rawData["message"].(string)
	if !ok || message == "" {
		response := utils.Response{
			Status:  "fail",
			Message: "`json:'message'` field is required and must be a string",
		}
		utils.RespondWithJSON(w, http.StatusBadRequest, response)
		return
	}

	// Обрабатываем сообщение, если всё корректно
	fmt.Printf("Received message: %s\n", message)
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
