package postgres

import (
	"log"
	"tourist-blog/internal/domain/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dbconn := "host=localhost user=admintblog password=system dbname=tblog port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dbconn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных!")
	}

	db.AutoMigrate(&models.User{})

	DB = db
	log.Println("Успех! Подключение к бд + миграция")
}
