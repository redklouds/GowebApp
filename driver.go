package main

/*
This file will be responsible for keeping track of the database connections
object controllers,
*/

import (
	"context"
	//"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/mongo/readpref"
	//"math/rand"
	//"time"
	//"encoding/json"
	"errors"
	"log"
)

//global database object/class
type DB struct {
	MongoClient  *mongo.Client
	MongoContext context.Context
}

//global database connection
var mongoClient = &DB{}

//global context variable

//connect to Mongo

func makeContext() context.Context {
	_ctx := context.Background()
	return _ctx
}
func connectToMongoDB(connectionString string) (*DB, error) {
	ctx := makeContext()
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, fmt.Errorf("Couldn't connect to Mongo: %v", err)
	}
	//reuse the err variable
	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("Mongo Client Couldn't Connect with background context: %v", err)
	}
	mongoClient.MongoClient = client
	mongoClient.MongoContext = ctx
	//return the client object pointer
	return mongoClient, err
}

func getLastInsertedTodo() Todo {
	/*
		NEEDS WORK I MADE A WORK AROUND TO SET A TODO AS  'COUNT' VARIALBE TO KEEP THE UNIQUE ID'S CLEAN
	*/
	coll := mongoClient.MongoClient.Database("ToDoDatabase").Collection("todos") //connect to the ToDoDatabase and teh todo Collection
	var savedData Todo
	filter := bson.M{"name": "Count"}

	update := bson.M{"$inc": bson.M{"id": 1}}
	updateRes, _err := coll.UpdateOne(mongoClient.MongoContext, filter, update)
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateRes.MatchedCount, updateRes.ModifiedCount)
	if _err != nil {
		fmt.Println("Errro updateig: ", _err)
	}
	res := coll.FindOne(mongoClient.MongoContext, filter)
	if res != nil {
		fmt.Println(res)
	}

	res.Decode(&savedData)
	return savedData

}
func InsertTodoIntoDB(ins_todo Todo) Todo {

	coll := mongoClient.MongoClient.Database("ToDoDatabase").Collection("todos") //connect to the ToDoDatabase and teh todo Collection

	//ins_todo.Id =
	insertResult, err := coll.InsertOne(mongoClient.MongoContext, ins_todo)
	if err != nil {
		log.Fatal("ERROR", err)
	}

	//return insertResult.InsertedID.(int) //type assertion https://stackoverflow.com/questions/18041334/convert-interface-to-int
	fmt.Println(insertResult)
	return ins_todo
}

func GetAllTodo() []*Todo {
	coll := mongoClient.MongoClient.Database("ToDoDatabase").Collection("todos") //connect to the ToDoDatabase and teh todo Collection
	cur, err := coll.Find(mongoClient.MongoContext, bson.M{})                    //query for all documents
	if err != nil {
		fmt.Println("ERROr getting all documents", err)
	}

	var results []*Todo

	for cur.Next(mongoClient.MongoContext) {
		var elem Todo
		err = cur.Decode(&elem)
		if err != nil {
			fmt.Println("Error decoding")
		}
		results = append(results, &elem)
	}
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	defer cur.Close(mongoClient.MongoContext)
	return results

}
func GetSpecificTodo(id int) (Todo, error) {
	//retunr a specific todo in the database
	coll := mongoClient.MongoClient.Database("ToDoDatabase").Collection("todos") //connect to the ToDoDatabase and teh todo Collection
	filter := bson.M{"id": id}
	var resultTodo Todo
	cur := coll.FindOne(mongoClient.MongoContext, filter) //returns a single result
	if cur.Err != nil {
		fmt.Println("ERROR getting single result, probably DNE: ", cur.Err)
		return resultTodo, errors.New("Record does not exist!")
	}

	cur.Decode(&resultTodo)
	return resultTodo, nil

}
func findAllRecords() {
	coll := mongoClient.MongoClient.Database("NumberOfHits").Collection("Hits")

	cur, err := coll.Find(mongoClient.MongoContext, bson.M{})

	/*
		findOptions := options.Find()
		findOptions.SetLimit(4)

		cur, err1 := coll.Find(mongoClient.MongoContext, nil, findOptions)
	*/
	if err != nil {
		fmt.Println("ERRO")
	}

	var results []*NumberOfHits

	for cur.Next(mongoClient.MongoContext) {
		var elem NumberOfHits
		err = cur.Decode(&elem)
		if err != nil {
			fmt.Println("Error decoding")
		}
		results = append(results, &elem)
	}

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	defer cur.Close(mongoClient.MongoContext)

}
