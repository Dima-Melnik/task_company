package taskhandlers

import (
	"net/http"
	"tasks_company/app/utils"
	"tasks_company/app/utils/logger"
	"tasks_company/internal/models"

	"github.com/gin-gonic/gin"
)

func UpdateTask(c *gin.Context) {
	id, ok := utils.CheckCorrectID(c)
	if !ok {
		return
	}

	var updatedTask *models.Tasks

	task, err := taskHandlers.UpdateTask(c, id, updatedTask)
	if err != nil {
		logger.LogHandler(c, "UpdateTask", err)
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, task)
}
