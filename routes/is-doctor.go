package routes

import (
	"findmydoc-backend/database"
	"findmydoc-backend/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IsDoctorParams struct {
	AccToken string `json:"acc-token"`
}

func IsDoctorHandler(c *gin.Context) {
	var body IsDoctorParams

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	var id = helpers.Authenticate(body.AccToken)

	if id == nil {
		c.Status(http.StatusUnauthorized)

		return
	}

	var rows, err = database.Db.Query("select is_doctor($1)", id)

	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	defer rows.Close()

	var isDoctor bool

	if rows.Next() {
		if err := rows.Scan(&isDoctor); err != nil {
			c.Status(http.StatusInternalServerError)

			return
		}
	} else {
		c.Status(http.StatusNotFound)

		return
	}

	c.JSON(http.StatusOK, isDoctor)
}
