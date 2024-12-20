package storage

import (
	"json_responses_requests/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase initializes the database connection
func ConnectDatabase() {
	// Обновите строку подключения с вашей реальной конфигурацией
	dsn := "user=postgres password=admin dbname=mydb port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	// Удалим таблицу при запуске, если она существует
	ResetUserTable()

	// Миграция базы данных
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}

// ResetUserTable удаляет таблицу и пересоздает её
func ResetUserTable() {
	// Проверим, существует ли таблица
	hasTable := DB.Migrator().HasTable(&models.User{})

	if hasTable {
		// Удаляем таблицу, если она существует
		if err := DB.Migrator().DropTable(&models.User{}); err != nil {
			log.Printf("Error dropping User table: %v", err)
		} else {
			log.Println("User table dropped successfully")
		}
	}

	// Пересоздадим таблицу
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Error creating User table: %v", err)
	}
	log.Println("User table has been created")
}
