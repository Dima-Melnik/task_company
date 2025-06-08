package auth

import (
	"net/http"
	"tasks_company/app/auth"
	"tasks_company/app/db"
	"tasks_company/app/utils"
	"tasks_company/internal/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid input")
		return
	}

	var user models.User
	if err := db.DB.Where("email = ?", input.Email).First(&user); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid email or password")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid email or password")
		return
	}

	tokenString, err := auth.GenerateToken(user.ID, user.Role)
	if err != nil {
		utils.SendError(c, 500, "Failed to create token")
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
