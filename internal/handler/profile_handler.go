package handler

import (
	"backend/internal/helper"
	"backend/internal/models"
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

// GetMyProfile godoc
//
//	@Summary		Get My Profile
//	@Description	Get the profile of the currently logged-in user
//	@Tags			profile
//	@Produce		json
//	@Success		200	{object}	models.Response
//	@Failure		401	{object}	models.Response
//	@Failure		404	{object}	models.Response
//	@Router			/profile [get]
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

// UpdateProfile godoc
//
//	@Summary		Update Profile
//	@Description	Update the profile of the currently logged-in user
//	@Tags			profile
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.User	true	"User Data"
//	@Success		200		{object}	models.Response
//	@Failure		400		{object}	models.Response
//	@Router			/profile [put]
func (h *UserHandler) UpdateProfile(ctx *gin.Context) {
	id := 1

	var updateUser models.User
	if err := ctx.ShouldBindJSON(&updateUser); err != nil {
		helper.BadRequest(ctx, "Invalid request body", nil, err)
		return
	}

	createUser, err := h.UserService.AddUser(updateUser)
	if helper.NotFoundError(ctx, err) {
		return
	}

	h.UserService.UpdateUserById(id, createUser)

	helper.ResponseOk(ctx, "Succes Updated User", createUser)
}
