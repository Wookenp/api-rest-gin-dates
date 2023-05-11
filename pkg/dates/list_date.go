package dates

import (
	"net/http"

	"github.com/Wookenp/date-manager-api-golang/pkg/common/db"
	"github.com/Wookenp/date-manager-api-golang/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func ListDates(c *gin.Context) {
	// Connect to the database.
	dbClient, dbErr := db.ConnectDB()
	if dbErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr.Error()})
		return
	}
	defer db.CloseDB(dbClient)

	// Get the "dates" collection from the database.
	datesCollection := dbClient.Database("date_manager").Collection("dates")

	// Retrieve all dates from the database.
	cursor, findErr := datesCollection.Find(c, map[string]interface{}{})
	if findErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": findErr.Error()})
		return
	}
	defer cursor.Close(c)

	// Initialize a slice to hold the dates.
	var dates []models.Date

	// Iterate over the cursor and append each date to the slice.
	for cursor.Next(c) {
		var date models.Date
		if err := cursor.Decode(&date); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		dates = append(dates, date)
	}

	// If an error occurred while iterating over the cursor, return it.
	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the list of dates to the client.
	c.JSON(http.StatusOK, gin.H{"dates": dates})
}
