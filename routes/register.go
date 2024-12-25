package routes

import (
	"findmydoc-backend/database"
	"findmydoc-backend/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterParams struct {
	AccToken string `json:"acc-token"`
}

func RegisterHandler(c *gin.Context) {
	var body RegisterDoctorParams

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var id = helpers.Authenticate(body.AccToken)

	if id == nil {
		c.Status(http.StatusUnauthorized)

		return
	}

	var _, err = database.Db.Query("select register($1)", id)

	if err != nil {
		c.Status(500)
	}
}
