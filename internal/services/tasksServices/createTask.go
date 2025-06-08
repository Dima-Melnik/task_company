package tasksservices

import (
	"tasks_company/app/db"
	"tasks_company/app/utils"
	"tasks_company/app/utils/logger"
	"tasks_company/internal/models"

	"github.com/gin-gonic/gin"
)

func (t *TasksServices) CreateTasks(c *gin.Context, data *models.Tasks) (*models.Tasks, error) {
	utils.CreateUpdateProcessing(c, data)

	result := db.DB.Create(&data)
	if result.Error != nil {
		logger.LogService(c, "create user", "CreateUser", result.Error)
		return nil, result.Error
	}

	return data, nil
}
