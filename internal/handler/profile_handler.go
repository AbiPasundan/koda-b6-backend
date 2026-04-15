package handler

import (
	"backend/internal/helper"
	"backend/internal/models"
	"backend/internal/service"
	"net/http"

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

	// ambil user_id dari middleware JWT
	idRaw, exists := ctx.Get("user_id")
	if !exists {
		helper.CustomeError(ctx, http.StatusUnauthorized, "Unauthorized", nil, nil)
		return
	}

	id := idRaw.(int)

	// bind body
	if err := ctx.ShouldBindJSON(&updateUser); err != nil {
		helper.BadRequest(ctx, "Invalid request body", nil, err)
		return
	}

	// CALL SERVICE (langsung update, bukan create dulu!)
	result, err := h.ProfileService.UpdateUser(id, updateUser)
	if err != nil {
		helper.CustomeError(ctx, http.StatusInternalServerError, "Failed update profile", nil, err)
		return
	}

	helper.ResponseOk(ctx, "Success Update User", result)
}

// func (h *ProfileHandler) UpdateProfile(ctx *gin.Context) {
// 	var updateUser models.UpdateProfile
// 	id, exists := ctx.Get("user_id")

// 	if !exists {
// 		helper.CustomeError(ctx, http.StatusUnauthorized, "Unauthorized", nil, nil)
// 		return
// 	}
// 	updateUser.Id = id.(int)

// 	if err := ctx.ShouldBindJSON(&updateUser); err != nil {
// 		helper.BadRequest(ctx, "Invalid request body", nil, err)
// 		return
// 	}

// 	createUser, err := h.ProfileService.AddUser(updateUser)
// 	if helper.NotFoundError(ctx, err) {
// 		return
// 	}

// 	h.ProfileService.UpdateUserById(updateUser.Id, createUser)

// 	helper.ResponseOk(ctx, "Succes Updated User", createUser)
// }
