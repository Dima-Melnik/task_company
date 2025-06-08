package userservices

import (
	"tasks_company/app/db"
	"tasks_company/app/utils/logger"
	"tasks_company/internal/models"

	"github.com/gin-gonic/gin"
)

func (u *UsersServices) DeleteUser(c *gin.Context, id int) error {
	err := db.DB.Delete(&models.User{}, id)
	if err != nil {
		logger.LogService(c, "delete user", "DeleteUser", err.Error)
		return err.Error
	}

	return nil
}
