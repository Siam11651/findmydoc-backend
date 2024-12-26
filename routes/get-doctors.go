package routes

import (
	"findmydoc-backend/database"
	"findmydoc-backend/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetDoctorsParams struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	AccToken  string  `json:"acc-token"`
}

type GetDoctorsColums struct {
	Id        string  `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func GetDoctorsHandler(c *gin.Context) {
	var body GetDoctorsParams

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	var id = helpers.Authenticate(body.AccToken)

	if id == nil {
		c.Status(http.StatusUnauthorized)

		return
	}

	var rows, err = database.Db.Query(
		"select * from get_doctors($1, point($2, $3))",
		id,
		body.Latitude,
		body.Longitude,
	)

	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	defer rows.Close()

	const COLUMN_COUNT = 3
	rowValue := GetDoctorsColums{}
	result := []GetDoctorsColums{}

	for rows.Next() {
		err := rows.Scan(&rowValue.Id, &rowValue.Latitude, &rowValue.Longitude)

		if err != nil {
			c.Status(http.StatusInternalServerError)

			return
		}

		result = append(result, rowValue)
	}

	c.JSON(http.StatusOK, result)
}
