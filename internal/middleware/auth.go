package middleware

import (
	"backend/internal/models"
	"fmt"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		auth := c.GetHeader("Authorization")
		// fmt.Printf("Debug: Received Authorization Header: '%s'\n", auth)

		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}

		parts := strings.Split(auth, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
			return
		}

		tokenStr := strings.Trim(parts[1], "\"' ")

		claims := &models.Claims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims,
			func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token", "details": err.Error()})
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token is not valid"})
			return
		}

		fmt.Println("DEBUG ROLE:", claims.Role)

		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)

		c.Set("userEmail", claims.Email)

		fmt.Println("NOW:", time.Now())
		fmt.Println("EXP:", claims.ExpiresAt.Time)

		c.Next()
	}
}

func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {

		role, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "role not found"})
			return
		}

		userRole := role.(string)

		if slices.Contains(allowedRoles, userRole) {
			c.Next()
			return
		}

		fmt.Println("CTX ROLE:", userRole)

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "access forbidden"})
	}
}
