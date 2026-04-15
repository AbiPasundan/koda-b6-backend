package handler

import (
	"backend/internal/helper"
	"backend/internal/models"
	"backend/internal/service"
	"fmt"
	"net/http"
	"time"

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
func (h *ProfileHandler) UpdateProfile(ctx *gin.Context) {
	var updateUser models.UpdateProfile

	idRaw, exists := ctx.Get("user_id")
	if !exists {
		helper.CustomeError(ctx, http.StatusUnauthorized, "Unauthorized", nil, nil)
		return
	}
	id := idRaw.(int)

	if err := ctx.ShouldBind(&updateUser); err != nil {
		helper.BadRequest(ctx, "Invalid request", nil, err)
		return
	}

	file, err := ctx.FormFile("pictures")
	if err == nil {
		filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)

		path := "./uploads/" + filename
		if err := ctx.SaveUploadedFile(file, path); err != nil {
			helper.CustomeError(ctx, http.StatusInternalServerError, "Failed upload file", nil, err)
			return
		}

		updateUser.Pictures = &filename
	}

	result, err := h.ProfileService.UpdateUser(id, updateUser)
	if err != nil {
		helper.CustomeError(ctx, http.StatusInternalServerError, "Failed update profile", nil, err)
		return
	}

	helper.ResponseOk(ctx, "Success Update User", result)
}
