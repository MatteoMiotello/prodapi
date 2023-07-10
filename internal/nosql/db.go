package nosql

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

var dbName string

func init() {
	dbName = viper.GetString("MONGO_DB_NAME")
}

var Client *mongo.Client

func InitClient(client *mongo.Client) {
	Client = client
}

func Disconnect() {
	if err := Client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func TyreCollection() *mongo.Collection {
	return Client.Database(viper.GetString("MONGO_DB_NAME")).Collection("tyres")
}

func BrandCollection() *mongo.Collection {
	return Client.Database(viper.GetString("MONGO_DB_NAME")).Collection("brands")
}
