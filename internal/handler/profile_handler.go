package handler

import (
	"backend/internal/service"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	ProfileService *service.ProfileService
}

func NewProfileHandler(service *service.ProfileService) *ProfileHandler {
	return &ProfileHandler{
		ProfileService: service,
	}
}

func (h *ProfileHandler) GetMyProfile(ctx *gin.Context) {
	email, exists := ctx.Get("userEmail")
	if !exists {
		ctx.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	user, err := h.ProfileService.GetProfile(email.(string))
	if err != nil {
		ctx.JSON(404, gin.H{"error": "user not found"})
		return
	}

	ctx.JSON(200, user)
}
