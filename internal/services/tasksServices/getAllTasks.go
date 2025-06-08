package tasksservices

import (
	"errors"
	"strconv"
	"tasks_company/app/db"
	"tasks_company/app/utils/logger"
	"tasks_company/internal/models"

	"github.com/gin-gonic/gin"
)

func (t *TasksServices) GetAllTasks(c *gin.Context) ([]models.Tasks, error) {
	var tasks []models.Tasks

	query := db.DB

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}

	if limit < 10 {
		limit = 10
	}

	offset := (page - 1) * limit

	result := query.Offset(offset).Limit(limit).Find(&tasks)
	if result.Error != nil {
		logger.LogService(c, "get tasks", "GetAllTasks", result.Error)
		return nil, result.Error
	}

	if len(tasks) == 0 {
		return nil, errors.New("tasks list is empty")
	}

	return tasks, nil
}
