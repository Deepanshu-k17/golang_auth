package httpserver

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"service": "go_auth",
		"time":    time.Now().UTC(),
	})
}
