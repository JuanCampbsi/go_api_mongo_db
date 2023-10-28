package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @BasePath /api/v1

// @Summary Delete opening
// @Description Delete a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [delete]
func DeleteOpeningHandler(ctx *gin.Context) {
	idStr := ctx.Query("id")

	if idStr == "" {
		sendError(ctx, http.StatusBadRequest, errParamRequired("id", "queryParameter").Error())
		return
	}

	// Convert string ID to MongoDB ObjectID
	objectID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, "Invalid ID format")
		return
	}

	collection := mongoDb.Collection("opportunities")

	deleteCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Delete the document by ID
	result, err := collection.DeleteOne(deleteCtx, bson.M{"_id": objectID})
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error deleting from database")
		return
	}

	// Check if no document was found and deleted
	if result.DeletedCount == 0 {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("Opening with ID: %s not found", idStr))
		return
	}

	// Note: You might want to return some sort of confirmation instead of the deleted object's ID
	sendSuccess(ctx, "delete-opening", bson.M{"id": idStr})
}
