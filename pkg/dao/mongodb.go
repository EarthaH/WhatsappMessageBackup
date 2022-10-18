package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"hogsback.com/whatsapp/pkg/logger"
)

func Init() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://cluster0.yhlijvd.mongodb.net/?authSource=%24external&authMechanism=MONGODB-X509&retryWrites=true&w=majority"))

	if err != nil {
		logger.Error(err)
	}

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		logger.Error(err)
	}
}
