package config

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoDb *mongo.Database
	logger  *Logger
)

func Init() error {
	var err error

	// Inicializar MongoDB
	mongoDb, err = InitializeMongo()
	if err != nil {
		return fmt.Errorf("error initializing mongodb: %v", err)
	}

	return nil
}

func GetMongoDb() *mongo.Database {
	return mongoDb
}

func GetLogger(p string) *Logger {
	//Initialize Logger
	logger = NewLogger(p)
	return logger
}
