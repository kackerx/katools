package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LogRecord struct {
	JobName  string `bson:"job_name" json:"job_name"`
	Command  string `bson:"command"`
	Err      string `bson:"err"`
	Conotent string `bson:"conotent"`
}

type FindByName struct {
	JobName string `bson:"job_name"`
}

func (l *LogRecord) Test(s string) (int, error) {
	return len(l.JobName), nil
}

func main() {
	client, err := mongo.Connect(context.TODO(), &options.ClientOptions{
		Hosts: []string{"localhost:27017"},
		Auth:  &options.Credential{Username: "root", Password: "root"},
	})
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	db := client.Database("my_db")
	col := db.Collection("my_col")

	// 插入
	result, err := col.InsertOne(context.TODO(), LogRecord{JobName: "jobname", Command: "command", Err: "err", Conotent: "content"})
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	docId := result.InsertedID.(primitive.ObjectID)
	fmt.Println(docId)

	// 查找
	cursor, err := col.Find(context.TODO(), &FindByName{JobName: "jobname"}, &options.FindOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}

	defer cursor.Close(context.TODO())

	var record *LogRecord
	for cursor.Next(context.TODO()) {
		record = &LogRecord{}
		if err := cursor.Decode(record); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(record)
	}

	col.DeleteOne(context.TODO(), &FindByName{})

}
