package response

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type ApiError struct {
	Param   string
	Message string
}

func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "statusValidation":
		fallthrough
	case "oneof":
		return "Invalid value"
	case "len":
		fallthrough
	case "lte":
		fallthrough
	case "gte":
		return "Invalid length"
	case "numeric":
		return "Must be numeric"
	case "datetime":
		return "Invalid date format"
	}
	return fe.Error() // default error
}

func ValidationErrorResponse(c *gin.Context, err error, body interface{}) {
	if err.Error() == "EOF" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid body type",
		})
		return
	}
	var je *json.SyntaxError
	if errors.As(err, &je) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid body type",
		})
		return
	}
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ApiError, len(ve))
		val := reflect.ValueOf(body)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}
		if val.Kind() != reflect.Struct {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid body type",
				"errors":  "Expected a struct",
			})
			return
		}
		for i, fe := range ve {
			fieldName := fe.Field()
			field, _ := val.Type().FieldByName(fieldName)
			fieldJSONName, _ := field.Tag.Lookup("json")
			out[i] = ApiError{fieldJSONName, msgForTag(fe)}
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Validation error",
			"errors":  out,
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"message": err.Error(),
	})
}
