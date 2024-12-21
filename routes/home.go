package routes

import "github.com/gin-gonic/gin"

func HomeHandler(c *gin.Context) {
	c.String(200, "hehe")
}
