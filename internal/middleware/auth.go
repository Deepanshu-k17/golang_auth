package middleware

import (
	"go-auth/internal/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	ctxUserIDKey   = "auth.userID"
	ctxUserRoleKey = "auth.role"
)

func AuthRequired(jwtSecret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := strings.TrimSpace(ctx.GetHeader("Authorization"))
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or invalid"})
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			return
		}
		scheme := strings.TrimSpace(parts[0])
		tokenString := strings.TrimSpace(parts[1])
		if scheme != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization scheme must be Bearer"})
			return
		}
		if tokenString == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token is missing"})
			return
		}
		claims, err := auth.ParseToken(jwtSecret, tokenString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		ctx.Set(ctxUserIDKey, claims.Subject)
		ctx.Set(ctxUserRoleKey, claims.Role)
		ctx.Next()

	}

}

func GetUserID(ctx *gin.Context) (string, bool) {
	res, ok := ctx.Get(ctxUserIDKey)
	if !ok {
		return "", false
	}
	userID, ok := res.(string)
	return userID, ok

}

func GetUserRole(ctx *gin.Context) (string, bool) {
	res, ok := ctx.Get(ctxUserRoleKey)
	if !ok {
		return "", false
	}
	role, ok := res.(string)
	return role, ok

}
