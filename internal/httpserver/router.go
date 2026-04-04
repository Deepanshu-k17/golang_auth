package httpserver

import (
	"go-auth/internal/app"
	"go-auth/internal/middleware"
	"go-auth/internal/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(a *app.App) *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())

	r.Use(gin.Recovery())
	r.GET("/health", health)

	userRepo := user.NewRepo(a.DB)
	userSvc := user.NewService(userRepo, a.Config.JWTSecret)
	userHandler := user.NewHandler(userSvc)

	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)
	api := r.Group("/api")
	api.Use(middleware.AuthRequired(a.Config.JWTSecret))
	api.GET("/files", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Files endpoint - to be implemented"})
	})
	api.GET("/products", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Products endpoint - to be implemented"})
	})
	return r
}
