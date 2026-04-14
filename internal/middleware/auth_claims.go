package middleware

import (
	"backend/internal/models"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = fmt.Appendf(nil, ":%s", os.Getenv("PORT"))

func GenerateToken(userID int, email string, full_name string, address string, phone string, pictures string, created_at time.Time, role string) (string, error) {

	expiration := time.Now().Add(time.Hour * 24)

	claims := &models.Claims{
		UserID: userID,
		// 	FullName string  `db:"full_name" json:"full_name"`
		// Address  *string `db:"address" json:"address"`
		// Phone    *string `db:"phone" json:"phone"`
		// Pictures *string `db:"pictures" json:"pictures"`
		Email:     email,
		FullName:  full_name,
		Address:   &address,
		Phone:     &phone,
		Pictures:  &pictures,
		CreatedAt: &created_at,
		Role:      role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}
