package handler

import (
	"backend/internal/models"
	"backend/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthService *service.AuthService
}

func NewAuthService(repo *service.AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: repo,
	}
}

func (h *AuthHandler) FindEmail(ctx *gin.Context) {
	// return h.AuthService.FindEmail(email)
	var password models.Users

	ctx.ShouldBindJSON(&password)

	// argon := argon2.DefaultConfig()

	// encoded, err := argon.HashEncoded([]byte(password))
	// if err != nil {
	// 	return
	// }
	// fmt.Println(string(encoded))

}

// password := "secret"

// argon := argon2.DefaultConfig()

// encoded, err := argon.HashEncoded([]byte(password))
// if err != nil {
// 	return nil, err
// }
// fmt.Println(string(encoded))
