package taskhandlers

import (
	"net/http"
	"tasks_company/app/utils"
	"tasks_company/app/utils/logger"

	"github.com/gin-gonic/gin"
)

func GetTaskByID(c *gin.Context) {
	id, ok := utils.CheckCorrectID(c)
	if !ok {
		return
	}

	task, err := taskHandlers.GetTaskByID(c, id)
	if err != nil {
		logger.LogHandler(c, "GetTaskByID", err)
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, task)
}
