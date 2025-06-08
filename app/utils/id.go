package utils

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func CheckCorrectID(c *gin.Context) (int, bool) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"id":    "invalid",
			"error": err,
		}).Error("Invalid ID value")
		SendError(c, http.StatusBadRequest, "Invalid ID value")
		return id, false
	}
	return id, true
}
