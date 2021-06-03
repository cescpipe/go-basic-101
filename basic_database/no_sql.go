package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type AirBNB struct {
	Id   string `bson:"_id"`
	Name string `bson:"name"`
}

func connectMongo() error {
	password := "HDrLv7gRVMJ8BJaX"
	conn := fmt.Sprintf("mongodb+srv://teerapat:%s@cluster0.5m0mj.mongodb.net/%s?retryWrites=true&w=majority", password, "sample_airbnb")
	log.Println("conn:", conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conn))
	if err != nil {
		return err
	}

	coll := client.Database("sample_airbnb").Collection("listingsAndReviews")
	cur, err := coll.Find(ctx, bson.M{}, options.Find().SetLimit(1))
	if err != nil {
		return errors.Wrap(err, "find error")
	}

	var abs []AirBNB
	if err := cur.All(ctx, &abs); err != nil {
		return err
	}

	log.Println(len(abs))
	for _, a := range abs {
		log.Println("airbnb: ", a.Id, a.Name)
	}
	return nil
}
