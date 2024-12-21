package main

import (
	"findmydoc-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.RegisterRoutes(r)
	r.Run(":8080")
}