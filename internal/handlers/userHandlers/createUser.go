package userhandlers

import (
	"net/http"
	"tasks_company/app/utils"
	"tasks_company/app/utils/logger"
	"tasks_company/internal/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var createdUser *models.User

	user, err := userHandlers.CreateUser(c, createdUser)
	if err != nil {
		logger.LogHandler(c, "CreateUser", err)
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, user)
}
