package handler

import (
	"context"
	"net/http"
	"time"

	"fmt"

	"github.com/JuanCampbsi/go-opportunities/schemas"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @BasePath /api/v1

// @Summary Show opening
// @Description Show a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [get]
func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamRequired("id", "queryParameter").Error())
		return
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, "Invalid ID format")
		return
	}

	collection := mongoDb.Collection("opportunities")
	findCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var opening schemas.Opening
	err = collection.FindOne(findCtx, bson.M{"_id": objectID}).Decode(&opening)
	if err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	sendSuccess(ctx, "show-opening", opening)
}
