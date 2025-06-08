package taskhandlers

import (
	"net/http"
	"tasks_company/app/utils"
	"tasks_company/app/utils/logger"
	"tasks_company/internal/models"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var createdTask *models.Tasks

	task, err := taskHandlers.CreateTasks(c, createdTask)
	if err != nil {
		logger.LogHandler(c, "CreateTask", err)
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, task)
}
