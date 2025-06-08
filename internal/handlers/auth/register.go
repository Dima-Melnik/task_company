package auth

import (
	"net/http"
	"tasks_company/app/db"
	"tasks_company/app/utils"
	"tasks_company/internal/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid input")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	user := models.User{
		Email:    input.Email,
		Password: string(hashedPassword),
		Role:     "user",
	}

	if err := db.DB.Create(&user); err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to create user")
		return
	}

	c.JSON(http.StatusCreated, "User registered successfully")
}
