package middleware

import (
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		origin := ctx.GetHeader("Origin")

		urlEnv := os.Getenv("FRONTEND_URL")
		if urlEnv == "" {
			urlEnv = "http://localhost:5173"
		}

		allowOrigins := []string{
			urlEnv,
			"http://68.183.226.223:20609",
			"http://68.183.226.223:20610",
			"http://localhost:5173",
			"http://localhost:5174",
		}

		allowHeaders := []string{
			"Origin",
			"Content-Type",
			"Authorization",
		}

		allowMethods := []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		}

		if slices.Contains(allowOrigins, origin) {
			ctx.Header("Access-Control-Allow-Origin", origin)
		}
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-Headers", strings.Join(allowHeaders, ", "))
		ctx.Header("Access-Control-Allow-Methods", strings.Join(allowMethods, ", "))

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}

		ctx.Next()
	}
}
