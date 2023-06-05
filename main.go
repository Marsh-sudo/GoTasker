package main

import (
	"context"
	"errors"
	"time"
	"log"
	"os"
    
	"github.com/urfave/cli/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var ctx = context.TODO()

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client,err := mongo.Connect(ctx,clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx,nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("tasker").Collection("tasks")
}

func main() {
	app := &cli.App{
		Name: "tasker",
		Usage: "A simple todo list manager",
		Commands: []*cli.Command{},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func createTask(task *Task) error{
	_,err := collection.InsertOne(ctx,task)
	return err
}

type Task struct{
	ID primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
	Text string `bson:"text"`
	Completed bool `bson:"completed"`

}