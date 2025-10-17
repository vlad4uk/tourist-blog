package auth

import (
	"net/http"
	"time"
	"tourist-blog/internal/domain/models"
	database "tourist-blog/internal/storage/postgres"
	utils "tourist-blog/utils/auth"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: All fields are required and password must be at least 6 characters."})
		return
	}

	// 1. Хеширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// 2. Создание пользователя с инициализацией полей
	user := models.User{
		Username:     input.Username,
		Email:        input.Email,
		PasswordHash: string(hashedPassword), // Используем PasswordHash
		RoleID:       1,                      // Значение по умолчанию для нового пользователя
		Is_blocked:   false,                  // По умолчанию не заблокирован
		Created_at:   time.Now(),
	}

	if result := database.DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User with this email or username already exists."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Registration successful. Please log in."})
}

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input."})
		return
	}

	// 1. Поиск пользователя
	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials: Email not found."})
		return
	}

	// 1.1 Проверка на блокировку
	if user.Is_blocked {
		c.JSON(http.StatusForbidden, gin.H{"error": "Your account has been blocked."})
		return
	}

	// 2. Проверка пароля (сравнение с PasswordHash)
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials: Password incorrect."})
		return
	}

	// 3. Генерация JWT-токена
	// Приводим ID к uint для JWT claims
	token, err := utils.GenerateJWT(uint(user.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token."})
		return
	}

	// 4. Установка токена в HTTP-Only Cookie
	maxAge := int(time.Hour * 24)
	c.SetCookie("jwt", token, maxAge, "/", "localhost", false, true)

	// 5. Ответ пользователю
	userRes := models.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		RoleID:    user.RoleID,
		IsBlocked: user.Is_blocked,
	}
	c.JSON(http.StatusOK, gin.H{"user": userRes, "message": "Login successful."})
}

func Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out."})
}
