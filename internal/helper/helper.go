package helper

import (
	"strconv"
)

func GetId(i string) int {
	id, err := strconv.Atoi(i)
	if err != nil {
		// ctx.JSON(http.StatusBadRequest, models.Response{
		// 	Success: false,
		// 	Message: "Invalid Id" + err.Error(),
		// 	Results: nil,
		// })
		return 0
	}
	return id
}
