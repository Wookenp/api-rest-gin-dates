package dates

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Wookenp/date-manager-api-golang/pkg/common/db"
	"github.com/Wookenp/date-manager-api-golang/pkg/common/models"
)

func RetrieveDate(c *gin.Context) {
	// Get the ID parameter from the URL.
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	// log the id
	fmt.Println(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	// Connect to the database.
	dbClient, dbErr := db.ConnectDB()
	if dbErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr.Error()})
		return
	}
	defer db.CloseDB(dbClient)

	// Get the "dates" collection from the database.
	datesCollection := dbClient.Database("date_manager").Collection("dates")

	// Find the date by ID.
	var date models.Date
	findErr := datesCollection.FindOne(c, map[string]interface{}{"id": id}).Decode(&date)
	if findErr != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "date not found"})
		return
	}

	// Return the date to the client.
	c.JSON(http.StatusOK, gin.H{"date": date})
}
