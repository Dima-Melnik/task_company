package utils

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(c *gin.Context, data interface{}) bool {
	if err := validate.Struct(data); err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, err.Field()+"is validate"+err.Tag())
			log.Printf("Error validation struct")
			SendError(c, http.StatusBadRequest, "Error validate struct")
			return false
		}
	}
	return true
}
