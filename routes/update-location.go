package routes

import (
	"findmydoc-backend/database"
	"findmydoc-backend/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateLocationParams struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	AccToken  string  `json:"acc-token"`
}

func UpdateLocationHandler(c *gin.Context) {
	var body UpdateLocationParams

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var id = helpers.Authenticate(body.AccToken)

	if id == nil {
		c.Status(http.StatusUnauthorized)

		return
	}

	var _, err = database.Db.Query("select update_location($1, $2, $3)", id, body.Latitude, body.Longitude)

	if err != nil {
		c.Status(500)
	}
}
