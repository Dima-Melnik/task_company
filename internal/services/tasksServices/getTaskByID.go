package tasksservices

import (
	"tasks_company/app/db"
	"tasks_company/app/utils/logger"
	"tasks_company/internal/models"

	"github.com/gin-gonic/gin"
)

func (t *TasksServices) GetTaskByID(c *gin.Context, id int) (models.Tasks, error) {
	var task models.Tasks

	result := db.DB.First(&task, id)
	if result.Error != nil {
		logger.LogService(c, "get by id", "GetTaskByID", result.Error)
		return task, result.Error
	}

	return task, nil
}
