package userhandlers

import (
	"net/http"
	"tasks_company/app/utils"
	"tasks_company/app/utils/logger"

	"github.com/gin-gonic/gin"
)

func GetUserByID(c *gin.Context) {
	id, ok := utils.CheckCorrectID(c)
	if !ok {
		return
	}

	user, err := userHandlers.GetUserByID(c, id)
	if err != nil {
		logger.LogHandler(c, "GetUserByID", err)
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}
