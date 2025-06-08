package userservices

import (
	"errors"
	"strconv"
	"tasks_company/app/db"
	"tasks_company/app/utils/logger"
	"tasks_company/internal/models"

	"github.com/gin-gonic/gin"
)

func (u *UsersServices) GetAllUsers(c *gin.Context) ([]models.User, error) {
	var users []models.User

	query := db.DB

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}

	if limit < 10 {
		page = 10
	}

	offset := (page - 1) * limit

	result := query.Limit(limit).Offset(offset).Find(&users)
	if result.Error != nil {
		logger.LogService(c, "get all users", "GetAllUsers", result.Error)
		return nil, result.Error
	}

	if len(users) == 0 {
		return nil, errors.New("users list is empty")
	}

	return users, nil
}
