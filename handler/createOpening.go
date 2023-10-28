package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/JuanCampbsi/go-opportunities/schemas"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @BasePath /api/v1

// @Summary Create opening
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreatedOpeningRequest true "Request body"
// @Success 200 {object} CreatedOpeningRequest
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
func CreateOpeningHandler(ctx *gin.Context) {

	request := CreatedOpeningRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	collection := mongoDb.Collection("opportunities")

	insertCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(insertCtx, bson.M{
		"role":     opening.Role,
		"company":  opening.Company,
		"location": opening.Location,
		"remote":   opening.Remote,
		"link":     opening.Link,
		"salary":   opening.Salary,
	})

	if err != nil {
		logger.ErrorF("error creating opening: %+v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	opening.ID = result.InsertedID.(primitive.ObjectID)

	sendSuccess(ctx, "create-opening", opening)
}
