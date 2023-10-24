/*
 * Namf_Communication
 *
 * AMF Communication Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package communication

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/free5gc/amf/internal/logger"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NonUeN2InfoUnSubscribe - Namf_Communication Non UE N2 Info UnSubscribe service Operation
func HTTPNonUeN2InfoUnSubscribe(c *gin.Context) {
	logger.CommLog.Warnf("Handle Non Ue N2 Info UnSubscribe is being implemented.")
	id := c.Param("n2NotifySubscriptionId")
	deleteObjectByID(id)
	c.JSON(http.StatusOK, gin.H{})
}

func deleteObjectByID(id string) error {
	// Convert the string ID to MongoDB's ObjectId type
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	collection := client.Database("local").Collection("AMFNonUeN2MessageSubscriptions")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("failed to convert ID to ObjectId: %v", err)
	}

	// Create a filter to match the document with the given ID
	filter := bson.M{"_id": objectID}

	// Delete the document matching the filter
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		fmt.Println("failed to delete document: %v", err)
	}

	// Check the result to see if any documents were deleted
	if result.DeletedCount == 0 {
		fmt.Println("No documents were deleted")
	}

	fmt.Printf("Deleted %d document(s) with ID %s", result.DeletedCount, id)
	return nil
}
