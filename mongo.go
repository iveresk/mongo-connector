package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func connectMongo(target string, ch chan string) {
	user := os.Getenv("user")
	pass := os.Getenv("pass")
	port := "27017"
	connections := "5"
	connect_url := "mongodb://" + user + ":" + pass + "@" + target + ":" + port + "/?maxPoolSize=" + connections

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connect_url))

	if err != nil {
		ch <- "No such mongoDB for the target " + target
		return
	}
	filter := bson.M{"name": primitive.Regex{Pattern: "^prefix_"}}

	dbs, err := client.ListDatabaseNames(context.TODO(), filter)

	if err != nil {
		ch <- "Can not connect and get list of tables for the target " + target
		return
	}
	fmt.Printf("%+v\n", dbs)
	ch <- "\nFor the target: " + target + " The list of DBs is higher"
	return
}

func dumpDB(target string) {
	// TODO Dump Database
}
