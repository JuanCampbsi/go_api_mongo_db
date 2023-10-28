package config

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitializeMongo() (*mongo.Database, error) {
	logger := GetLogger("mongo")
	// Configurar o cliente para se conectar ao MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://root:root@localhost:27017/")

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	logger.Info("connect sucessful...")
	// Verificar a conexão
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	// Retornar a referência para o banco de dados
	db := client.Database("local")
	return db, nil
}
