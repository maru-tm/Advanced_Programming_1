package handlers

import (
	"errors"
	"json_responses_requests/models"
	"log"

	"gorm.io/gorm"
)

// CreateUser создает нового пользователя и контролирует проблему с автоинкрементом ID
func CreateUser(db *gorm.DB, user *models.User) error {
	// Начинаем транзакцию
	tx := db.Begin()

	// Убедимся, что транзакция не завершена раньше времени
	if tx.Error != nil {
		return tx.Error
	}

	// Проверяем, существует ли уже пользователь с таким email
	existingUser, err := GetUserByEmail(tx, user.Email)
	if err != nil {
		tx.Rollback() // Откатываем транзакцию при ошибке
		return err
	}
	if existingUser != nil {
		tx.Rollback() // Откатываем транзакцию при конфликте по email
		return errors.New("user with this email already exists")
	}

	// Теперь создаем нового пользователя и пытаемся избежать блокировки ID
	if err := tx.Create(user).Error; err != nil {
		// Если произошла ошибка при создании пользователя, откатываем транзакцию
		tx.Rollback()
		// Возможно, стоит сбросить автоинкремент ID, чтобы избежать блокировки
		if resetErr := resetAutoIncrement(tx); resetErr != nil {
			log.Printf("Failed to reset auto-increment: %v", resetErr)
		}
		return err
	}

	// Фиксируем транзакцию
	return tx.Commit().Error
}

// resetAutoIncrement сбрасывает автоинкремент в базе данных
func resetAutoIncrement(db *gorm.DB) error {
	// Сбрасываем автоинкремент, если в записи не было успеха
	// Для PostgreSQL
	return db.Exec("SELECT setval(pg_get_serial_sequence('users', 'id'), (SELECT MAX(id) FROM users))").Error
}

// GetUserByEmail ищет пользователя по email
func GetUserByEmail(db *gorm.DB, email string) (*models.User, error) {
	var user models.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // Не найден пользователь
		}
		return nil, err // Ошибка при поиске пользователя
	}
	return &user, nil
}

// GetUserByID извлекает пользователя по ID
func GetUserByID(db *gorm.DB, id uint) (*models.User, error) {
	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// DeleteUser удаляет пользователя по ID
func DeleteUser(db *gorm.DB, id uint) error {
	if err := db.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
