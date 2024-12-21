package routes

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up the application's routes
func RegisterRoutes(r *gin.Engine) {
	r.GET("/", HomeHandler)
}
