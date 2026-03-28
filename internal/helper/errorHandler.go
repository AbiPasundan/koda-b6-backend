package helper

import (
	"net/http"

	"backend/internal/models"

	"github.com/gin-gonic/gin"
)

func NotFoundError(ctx *gin.Context, err error) bool {
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: "404 Not Found: " + err.Error(),
			Results: nil,
		})
		return true
	}

	return false
}

func InternalServerError(ctx *gin.Context, message string, result any, err error) bool {
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: message + err.Error(),
			Results: result,
		})
		return true
	}

	return false
}

func BadRequest(ctx *gin.Context, message string, result any, err error) bool {
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: message + err.Error(),
			Results: result,
		})
		return true
	}

	return false
}

// custom error
func CustomeError(ctx *gin.Context, statusResponse int, message string, result any, err error) bool {
	if err != nil {
		ctx.JSON(statusResponse, models.Response{
			Success: false,
			Message: message + err.Error(),
			Results: result,
		})
		return true
	}

	return false
}
