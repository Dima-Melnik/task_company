package taskhandlers

import (
	"net/http"
	"tasks_company/app/utils"
	"tasks_company/app/utils/logger"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(c *gin.Context) {
	tasks, err := taskHandlers.GetAllTasks(c)
	if err != nil {
		logger.LogHandler(c, "GetAllTasks", err)
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, tasks)
}
