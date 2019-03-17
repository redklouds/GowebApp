package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	DBStruct, err := connectToMongoDB("http://localhost:27017") //in the driver file

	findAllRecords()

	if err != nil {
		fmt.Println("SOMETHING wel wrong")
	}
	fmt.Println(DBStruct)
	routers := NewRouter()
	fmt.Println("Starting the servivce ....")
	log.Fatal(http.ListenAndServe(":8080", routers))
}
