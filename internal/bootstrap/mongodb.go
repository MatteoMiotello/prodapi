package bootstrap

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDb() {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s",
		viper.GetString("MONGO_USERNAME"),
		viper.GetString("MONGO_PASSWORD"),
		viper.GetString("MONGO_HOST"),
		viper.GetString("MONGO_PORT"),
	)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	var result bson.M
	if err := client.
		Database(viper.GetString("MONGO_DB_NAME")).
		RunCommand(context.TODO(), bson.D{{"ping", 1}}).
		Decode(&result); err != nil {
		panic(err)
	}

	fmt.Println("Mongodb online")
}
