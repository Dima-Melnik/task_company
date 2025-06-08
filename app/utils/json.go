package utils

import (
	"net/http"
	"tasks_company/app/utils/logger"

	"github.com/gin-gonic/gin"
)

func BindJSON(c *gin.Context, data interface{}) bool {
	if err := c.ShouldBindJSON(&data); err != nil {
		logger.LogWithFieldsErrorf("json", "binding", err)
		SendError(c, http.StatusBadRequest, "Error binding jsob")
		return false
	}

	return true
}
