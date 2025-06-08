package userhandlers

import (
	"net/http"
	"tasks_company/app/utils"
	"tasks_company/app/utils/logger"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	id, ok := utils.CheckCorrectID(c)
	if !ok {
		return
	}

	err := userHandlers.DeleteUser(c, id)
	if err != nil {
		logger.LogHandler(c, "DeleteUser", err)
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, "User successfully deleted")
}
