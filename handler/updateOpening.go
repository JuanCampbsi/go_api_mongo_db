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

// @Summary Update opening
// @Description Update a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening Identification"
// @Param opening body UpdateOpeningRequest true "Opening data to Update"
// @Success 200 {object} UpdateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [put]
func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

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

	updateFields := bson.M{}

	if request.Role != "" {
		updateFields["role"] = request.Role
	}
	if request.Company != "" {
		updateFields["company"] = request.Company
	}
	if request.Location != "" {
		updateFields["location"] = request.Location
	}
	if request.Remote != nil {
		updateFields["remote"] = *request.Remote
	}
	if request.Link != "" {
		updateFields["link"] = request.Link
	}
	if request.Salary > 0 {
		updateFields["salary"] = request.Salary
	}

	collection := mongoDb.Collection("opportunities")

	updateCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Update the document
	result, err := collection.UpdateOne(updateCtx, bson.M{"_id": objectID}, bson.M{"$set": updateFields})

	if err != nil {
		logger.ErrorF("error updating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "Error updating opening in the database")
		return
	}

	if result.ModifiedCount == 0 {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("Opening with ID: %s not found or not modified", idStr))
		return
	}

	sendSuccess(ctx, "update-opening", bson.M{"id": idStr})
}
