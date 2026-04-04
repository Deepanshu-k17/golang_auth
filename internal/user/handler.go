package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Register(ctx *gin.Context) {
	var input LoginInput

	// Use ShouldBindJSON for cleaner validation error handling
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	output, err := h.svc.Register(ctx.Request.Context(), input)
	if err != nil {
		// Consider logging the actual error internally and returning a friendlier message
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, output)
}

func (h *Handler) Login(ctx *gin.Context) {
	var input LoginInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials format"})
		return
	}

	output, err := h.svc.Login(ctx.Request.Context(), input)
	if err != nil {
		// Unauthorized is the correct choice here for failed logins
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	ctx.JSON(http.StatusOK, output)
}
