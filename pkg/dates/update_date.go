package dates

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Wookenp/date-manager-api-golang/pkg/common/db"
)

func UpdateDate(c *gin.Context) {
	// Get the ID parameter from the URL.
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	// Parse the JSON payload from the request body.
	var date map[string]interface{}
	if err := c.ShouldBindJSON(&date); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Connect to the database.
	dbClient, dbErr := db.ConnectDB()
	if dbErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr.Error()})
		return
	}
	defer db.CloseDB(dbClient)

	// Update the date in the database.
	collection := dbClient.Database("date_manager").Collection("dates")
	filter := bson.M{"id": id}
	update := bson.M{"$set": date}
	result, updateErr := collection.UpdateOne(c.Request.Context(), filter, update)
	if updateErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": updateErr.Error()})
		return
	}

	if result.ModifiedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "date not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "date updated successfully"})
}
