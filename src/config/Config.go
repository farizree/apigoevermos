package config

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Environment() (string, error) {
	env := "development"

	return env, nil
}
func Hostname() (string, error) {
	host := "localhost"

	return host, nil
}

func DetermineListenAddressEvermos() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		//return "", fmt.Errorf("$PORT not set")
		port = "2021"
	}
	return ":" + port, nil
}

//docker connection
// func Connectmongo() (*mongo.Database, error) {
// 	clientOptions := options.Client()
// 	clientOptions.ApplyURI("mongodb://docker.for.mac.localhost:27017")
// 	client, err := mongo.NewClient(clientOptions)
// 	if err != nil {
// 		return nil, err
// 	}
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	err = client.Connect(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return client.Database("evermos"), nil
// }

//end docker connection

func Connectmongo() (*mongo.Database, error) {
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	//fmt.Println(ctx)
	//fmt.Println(err)

	return client.Database("evermos"), nil
}
