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
