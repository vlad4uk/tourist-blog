package main

import (
	"log"
	"time"

	handlers "tourist-blog/internal/handlers/auth"
	database "tourist-blog/internal/storage/postgres"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Подключение к БД
	// !!! Установите переменные окружения для подключения к PostgreSQL
	// os.Setenv("DB_HOST", "localhost")
	// os.Setenv("DB_USER", "postgres")
	// os.Setenv("DB_PASSWORD", "mysecretpassword")
	// os.Setenv("DB_NAME", "tourist_blog_db")
	// os.Setenv("DB_PORT", "5432")
	database.ConnectDB()

	// 2. Инициализация Gin
	router := gin.Default()

	// 3. Настройка CORS
	// ЭТО КЛЮЧЕВОЙ ШАГ для связи React (порт 3000) и Go (порт 8080)
	router.Use(cors.New(cors.Config{
		// Разрешаем ваш фронтенд
		AllowOrigins: []string{"http://localhost:3000"},
		// Разрешаем HTTP-методы
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		// Разрешаем заголовки
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
		// ЭТО КЛЮЧЕВОЙ МОМЕНТ для работы с HTTP-Only Cookies!
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 4. Роутинг
	api := router.Group("/api/auth")
	{
		api.POST("/register", handlers.Register)
		api.POST("/login", handlers.Login)
		api.POST("/logout", handlers.Logout)

		// Пример защищенного маршрута (требует middleware, который мы не писали, но он должен быть здесь)
		// api.GET("/user", middleware.AuthRequired(), handlers.GetCurrentUser)
	}

	// 5. Запуск сервера
	port := ":8080"
	log.Printf("Сервер запущен на порту %s...", port)
	if err := router.Run(port); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
