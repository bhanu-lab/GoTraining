package dao

import (
	"context"
	"fmt"
	"log"

	coll "github.com/mylibrary/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBNAME declaring db name
var DBNAME = "mylibrary"
var db *mongo.Database
var client *mongo.Client

func init() {
	fmt.Println("inside init dao")
	clientOptions := options.Client().ApplyURI("mongodb://192.168.0.107:27018")
	//	client, err := mongo.Connect(context.TODO(), clientOptions)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("created client")
	/* 	err = client.Ping(context.TODO(), nil)
	   	fmt.Println("ping complete")
	   	if err != nil {
	   		log.Fatal(err)
	   	} */
	fmt.Println("Connected to MongoDB")

	db = client.Database(DBNAME)
}

func disconnect() {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
}

// GetAllBooks get all the documents present in Collection BOOKS
func GetAllBooks() ([]*coll.Book, error) {
	var books []*coll.Book
	findOptions := options.Find()
	findOptions.SetLimit(2)
	fmt.Println("performing db query")
	cur, err := db.Collection("BOOKS").Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("waiting on results")
	for cur.Next(context.TODO()) {
		var b coll.Book
		err := cur.Decode(&b)
		if err != nil {
			log.Fatal(err)
		}

		books = append(books, &b)
	}
	return books, err
}

// InsertBook inserts Book document into collection BOOKS
func InsertBook(book coll.Book) {
	collection := client.Database("mylibrary").Collection("Books")
	result, err := collection.InsertOne(context.TODO(), book)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a book : ", result.InsertedID)
}
