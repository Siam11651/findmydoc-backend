package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterParams struct {
	Id       string `json:"id"`
	AccToken string `json:"acc-token"`
}

func RegisterHandler(c *gin.Context) {
	var body RegisterParams

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	println(body.AccToken)
}
