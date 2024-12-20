package handlers

import (
	"encoding/json"
	"json_responses_requests/models"
	"json_responses_requests/storage"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateUserHandler создает нового пользователя
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Парсим JSON из тела запроса
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Создаем нового пользователя
	if err := CreateUser(storage.DB, &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем успешный ответ
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetUserHandler обрабатывает запрос на получение пользователя по ID
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, `{"error": "Invalid ID format"}`, http.StatusBadRequest)
		return
	}

	user, err := GetUserByID(storage.DB, uint(id))
	if err != nil {
		http.Error(w, `{"error": "User not found"}`, http.StatusNotFound)
		log.Printf("Error finding user with ID %d: %v", id, err)
		return
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, `{"error": "Failed to encode response"}`, http.StatusInternalServerError)
		log.Printf("Error encoding response: %v", err)
	}
}

// UpdateUserHandler обрабатывает запрос на обновление пользователя
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, `{"error": "Invalid ID format"}`, http.StatusBadRequest)
		return
	}

	// Получаем пользователя из базы
	user, err := GetUserByID(storage.DB, uint(id))
	if err != nil {
		http.Error(w, `{"error": "User not found"}`, http.StatusNotFound)
		log.Printf("Error finding user with ID %d: %v", id, err)
		return
	}

	// Декодируем новые данные
	var updatedUser models.User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, `{"error": "Invalid input data"}`, http.StatusBadRequest)
		return
	}

	// Обновляем данные пользователя
	user.Name = updatedUser.Name
	user.Email = updatedUser.Email

	if err := storage.DB.Save(user).Error; err != nil {
		http.Error(w, `{"error": "Failed to update user"}`, http.StatusInternalServerError)
		log.Printf("Error updating user with ID %d: %v", id, err)
		return
	}

	// Возвращаем обновленного пользователя
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, `{"error": "Failed to encode response"}`, http.StatusInternalServerError)
		log.Printf("Error encoding response: %v", err)
	}
}

// DeleteUserHandler обрабатывает запрос на удаление пользователя
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, `{"error": "Invalid ID format"}`, http.StatusBadRequest)
		return
	}

	// Используем функцию из service.go для удаления пользователя
	if err := DeleteUser(storage.DB, uint(id)); err != nil {
		http.Error(w, `{"error": "Error deleting user"}`, http.StatusNotFound)
		log.Printf("Error deleting user with ID %d: %v", id, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetAllUsersHandler обрабатывает запрос на получение списка всех пользователей
func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	if err := storage.DB.Find(&users).Error; err != nil {
		http.Error(w, `{"error": "Failed to retrieve users"}`, http.StatusInternalServerError)
		log.Printf("Error retrieving users: %v", err)
		return
	}

	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, `{"error": "Failed to encode response"}`, http.StatusInternalServerError)
		log.Printf("Error encoding response: %v", err)
	}
}
