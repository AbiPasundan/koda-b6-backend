package handler

import (
	"backend/internal/helper"
	"backend/internal/middleware"
	"backend/internal/models"
	"backend/internal/service"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthService *service.AuthService
}

func NewAuthHandler(repo *service.AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: repo,
	}
}

// Login godoc
//
//	@Summary		Login user
//	@Description	Authenticate user and generate JWT token
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			req	body		models.AuthLogin	true	"Login credentials"
//	@Success		200	{object}	models.Response
//	@Failure		400	{object}	models.Response
//	@Failure		500	{object}	models.Response
//	@Router			/auth/login [post]
func (h *AuthHandler) Login(ctx *gin.Context) {
	var req models.AuthLogin

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.BadRequest(ctx, "Invalid request", nil, err)
		return
	}

	user, err := h.AuthService.Login(req.Email, req.Password)
	if err != nil {
		helper.BadRequest(ctx, "Wrong Email or Password", nil, err)
		return
	}

	var address, phone, pictures string

	if user.Address != nil {
		address = *user.Address
	}
	if user.Phone != nil {
		phone = *user.Phone
	}
	if user.Pictures != nil {
		pictures = *user.Pictures
	}

	var createdAt time.Time
	if user.CreatedAt != nil {
		createdAt = *user.CreatedAt
	}

	token, err := middleware.GenerateToken(user.Id, user.Email, user.Full_Name, address, phone, pictures, createdAt, user.Role)

	if err != nil {
		helper.CustomeError(ctx, http.StatusInternalServerError, "Failed generate token", nil, err)
		return
	}

	helper.ResponseOk(ctx, "Success Login", token)
}

// Register godoc
//
//	@Summary		Register new user
//	@Description	Register a new user to the system
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			req	body		models.AuthRegister	true	"Register credentials"
//	@Success		200	{object}	models.Response
//	@Failure		400	{object}	models.Response
//	@Failure		500	{object}	models.Response
//	@Router			/auth/register [post]
func (h *AuthHandler) Register(ctx *gin.Context) {

	var req models.AuthRegister

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.BadRequest(ctx, "Invalid request body", nil, err)
		return
	}

	err := h.AuthService.Register(&req)
	if err != nil {

		// handle duplicate email
		if strings.Contains(err.Error(), "duplicate key") {
			helper.BadRequest(ctx, "Email already exists", nil, err)
			return
		}

		helper.CustomeError(ctx, http.StatusInternalServerError, "Failed register", nil, err)
		return
	}

	helper.ResponseOk(ctx, "Success Create User", nil)
}

// RequestForgotPassword godoc
//
//	@Summary		Request forgot password
//	@Description	Request OTP for password reset
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			req	body		models.AuthForgotPassword	true	"Forgot password request"
//	@Success		200	{object}	models.Response
//	@Failure		400	{object}	models.Response
//	@Router			/auth/forgot-password [post]
func (h *AuthHandler) RequestForgotPassword(ctx *gin.Context) {
	var req models.AuthForgotPassword
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.BadRequest(ctx, "Invalid request body", nil, err)
		return
	}
	err := h.AuthService.ForgotPasswordRequest(&req)
	if err != nil {
		helper.BadRequest(ctx, err.Error(), nil, err)
		return
	}
	helper.ResponseOk(ctx, "OTP sent", nil)
}

func (h *AuthHandler) ResetPassword(ctx *gin.Context) {
	var req models.ResetPasswordRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.BadRequest(ctx, "Invalid request", nil, err)
		return
	}

	err := h.AuthService.ResetPassword(req)
	if err != nil {
		helper.BadRequest(ctx, err.Error(), nil, err)
		return
	}

	helper.ResponseOk(ctx, "Password reset success", nil)
}
