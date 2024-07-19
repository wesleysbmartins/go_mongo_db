package services

import (
	"context"
	"fmt"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDb struct {
	host     string
	port     int
	user     string
	password string
	database string
}

type IMongoDB interface {
	Connect()
}

var Database *mongo.Database

func (m *MongoDb) Connect() {
	if Database == nil {
		credentials := &MongoDb{
			user:     "admin",
			password: "1234",
			database: "go_mongo_db",
			port:     27017,
			host:     "localhost",
		}

		uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/", credentials.user, credentials.password, credentials.host, strconv.Itoa(credentials.port))

		serverAPI := options.ServerAPI(options.ServerAPIVersion1)
		opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

		ctx := context.TODO()
		client, err := mongo.Connect(ctx, opts)

		if err != nil {
			fmt.Println("URI:", uri)
			panic(err)
		}

		pingErr := client.Ping(ctx, nil)

		if pingErr != nil {
			fmt.Println("PING ERROR:", pingErr.Error())
			panic(err)
		}

		Database = client.Database(credentials.database)

		fmt.Println("DATABASE CONNECTION SUCCESS!")
	}
}
