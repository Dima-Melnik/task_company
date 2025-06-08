package tasksservices

import (
	"tasks_company/app/db"
	"tasks_company/app/utils/logger"
	"tasks_company/internal/models"

	"github.com/gin-gonic/gin"
)

func (y *TasksServices) DeleteTask(c *gin.Context, id int) error {
	err := db.DB.Delete(&models.Tasks{}, id)
	if err.Error != nil {
		logger.LogService(c, "delete task", "DeleteTask", err.Error)
		return err.Error
	}

	return nil
}
