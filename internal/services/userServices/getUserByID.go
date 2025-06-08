package userservices

import (
	"tasks_company/app/db"
	"tasks_company/app/utils/logger"
	"tasks_company/internal/models"

	"github.com/gin-gonic/gin"
)

func (u *UsersServices) GetUserByID(c *gin.Context, id int) (models.User, error) {
	var userByID models.User

	result := db.DB.First(&userByID, id)
	if result.Error != nil {
		logger.LogService(c, "get by id", "GetUserByID", result.Error)
		return userByID, result.Error
	}

	return userByID, nil
}
