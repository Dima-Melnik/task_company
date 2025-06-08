package utils

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
)

func CreateUpdateProcessing(c *gin.Context, data interface{}) (interface{}, error) {
	if !BindJSON(c, data) {
		log.Print("Invalid request body")
		return nil, errors.New("invalid request body")
	}

	if !ValidateStruct(c, data) {
		log.Print("Invalid struct")
		return nil, errors.New("invalid struct")
	}

	return data, nil
}
