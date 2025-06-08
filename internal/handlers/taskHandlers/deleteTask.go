package taskhandlers

import (
	"net/http"
	"tasks_company/app/utils"
	"tasks_company/app/utils/logger"

	"github.com/gin-gonic/gin"
)

func DeleteTask(c *gin.Context) {
	id, ok := utils.CheckCorrectID(c)
	if !ok {
		return
	}

	err := taskHandlers.DeleteTask(c, id)
	if err != nil {
		logger.LogHandler(c, "DeleteTask", err)
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "task is deleted"})
}
