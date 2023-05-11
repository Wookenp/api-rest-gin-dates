// https://www.makeuseof.com/golang-crud-api-mongodb-gin/
package dates

import (
	"net/http"

	"github.com/Wookenp/date-manager-api-golang/pkg/common/db"
	"github.com/Wookenp/date-manager-api-golang/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func CreateDate(c *gin.Context) {
	// Connect to the database.
	dbClient, dbErr := db.ConnectDB()
	if dbErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr.Error()})
		return
	}
	defer db.CloseDB(dbClient)

	// Get the "dates" collection from the database.
	datesCollection := dbClient.Database("date_manager").Collection("dates")

	// Bind the date data from the request body.
	var date models.Date
	if bindErr := c.BindJSON(&date); bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": bindErr.Error()})
		return
	}

	// Validate the date data.
	if date.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title is required"})
		return
	}
	if date.User.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user email is required"})
		return
	}
	if date.StartDate.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "start date is required"})
		return
	}
	if date.EndDate.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "end date is required"})
		return
	}

	// Create a new user object.
	user := models.NewUser(date.User.FullName, date.User.Email)

	// Create a new date object.
	newDate := models.NewDate(date.Title, date.Details, date.StartDate, date.EndDate, *user)

	// Insert the new date into the database.
	insertResult, insertErr := datesCollection.InsertOne(c, newDate)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": insertErr.Error()})
		return
	}

	// Retrieve the inserted date from the database.
	var insertedDate models.Date
	findErr := datesCollection.FindOne(c, map[string]interface{}{"_id": insertResult.InsertedID}).Decode(&insertedDate)
	if findErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": findErr.Error()})
		return
	}

	// Return the new date to the client.
	c.JSON(http.StatusOK, gin.H{"message": "date created successfully", "date": insertedDate})
}
