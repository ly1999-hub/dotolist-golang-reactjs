package module

import (
	"context"
	"fmt"
	"github.com/logrusorgru/aurora"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func ConnectDB(host, user, password, dbName, mechanism, source string) error {
	connectOptions := options.ClientOptions{}

	if user != "" && password != "" {
		connectOptions.Auth = &options.Credential{
			AuthMechanism: mechanism,
			AuthSource:    source,
			Username:      user,
			Password:      password,
		}
	}

	// Connect
	client, err := mongo.Connect(context.Background(), connectOptions.ApplyURI(host))
	if err != nil {
		fmt.Println("Error when connect to MongoDB database", host, err)
		return err
	}

	fmt.Println(aurora.Green("*** CONNECTED TO MONGODB: " + host + " - DB NAME: " + dbName))

	// Set data
	db = client.Database(dbName)

	return nil
}

// GetInstance ...
func GetInstance() *mongo.Database {
	return db
}
