package tasksservices

import (
	"tasks_company/app/db"
	"tasks_company/app/utils"
	"tasks_company/app/utils/logger"
	"tasks_company/internal/models"

	"github.com/gin-gonic/gin"
)

func (t *TasksServices) UpdateTask(c *gin.Context, id int, data *models.Tasks) (*models.Tasks, error) {
	utils.CreateUpdateProcessing(c, data)

	result := db.DB.Model(&models.Tasks{}).Where("id = ?", id).Updates(data)
	if result.Error != nil {
		logger.LogService(c, "update task", "UpdateTask", result.Error)
		return nil, result.Error
	}

	return data, nil
}
