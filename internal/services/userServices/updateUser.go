package userservices

import (
	"tasks_company/app/db"
	"tasks_company/app/utils"
	"tasks_company/app/utils/logger"
	"tasks_company/internal/models"

	"github.com/gin-gonic/gin"
)

func (u *UsersServices) UpdateUser(c *gin.Context, id int, data *models.User) (*models.User, error) {
	utils.CreateUpdateProcessing(c, data)

	result := db.DB.Model(&models.User{}).Where("id = ?", id).Updates(data)
	if result.Error != nil {
		logger.LogService(c, "update user", "UpdateUser", result.Error)
		return nil, result.Error
	}

	return data, nil
}
