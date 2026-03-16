package helper

import (
	"backend/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetID(ctx *gin.Context) (int, bool) {

	i := ctx.Param("id")

	id, err := strconv.Atoi(i)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid ID: " + err.Error(),
			Results: nil,
		})
		return 0, false
	}

	return id, true
}
