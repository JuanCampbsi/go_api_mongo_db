package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/JuanCampbsi/go-opportunities/schemas"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// @BasePath /api/v1

// @Summary List openings
// @Description List all job openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningHandler(ctx *gin.Context) {
	var openings []schemas.Opening

	collection := mongoDb.Collection("opportunities")
	findCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(findCtx, bson.M{}, options.Find())
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error listing openings")
		return
	}
	defer cursor.Close(findCtx)

	if err := cursor.All(findCtx, &openings); err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error decoding openings")
		return
	}

	length := len(openings)
	if length == 0 {
		sendSuccess(ctx, "listings openings is empty", openings)
	} else {
		sendSuccess(ctx, "listings openings", openings)
	}
}
