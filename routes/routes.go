package routes

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up the application's routes
func RegisterRoutes(r *gin.Engine) {
	r.GET("/", HomeHandler)
	r.POST("/register", RegisterHandler)
	r.POST("/register-doctor", RegisterDoctorHandler)
	r.POST("/is-doctor", IsDoctorHandler)
	r.POST("/update-location", UpdateLocationHandler)
	r.POST("/llm", LlmHandler)
}
