package handler

import (
	"github.com/JuanCampbsi/go-opportunities/config"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	logger  *config.Logger
	mongoDb *mongo.Database
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	mongoDb = config.GetMongoDb()
}
