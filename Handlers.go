package main

/*
Handlers.go
Good practice to organize the API directory,
this file handles all the handlers of course service
*/

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"

	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Index(resp http.ResponseWriter, req *http.Request) {
	//just a welcome screen for who ever enters the index
	fmt.Fprintf(resp, "Welcome to My API Service in Go-Lang!")
}

func TodoIndex(resp http.ResponseWriter, req *http.Request) {
	//this will handle returning the current Todo in the system

	//** ALRIGHT LETS DO SOME RESPONSIBILITY CHECKING
	//stating the header for poeple requesting this data

	/*todos := Todos{
		Todo{
			Name:      "Pick up Hunny",
			Completed: false,
		},
		Todo{
			Name: "Go watch a movie",
		},
	}
	*/
	/*
		Lets set the header
		//if we dont set it, go's net/http will always try to guess
		what's the content type we are sending.. and it is NOT
		always right, so if we know the content type just set it here
	*/

	//this is basically telling the client, that `hey expect some json fucker`
	resp.Header().Set("Content-Type", "application/json; charset=UTF-8")
	resp.WriteHeader(http.StatusOK)
	//encode/serialize the above structure
	//List(Todos)

	//OLD(Static): err := json.NewEncoder(resp).Encode(GetAllTodo())
	err := json.NewEncoder(resp).Encode(GetAllTodo())
	if err != nil {
		//error encoding to json, halt here
		panic(err)
	}
}

func TodoCreate(resp http.ResponseWriter, req *http.Request) {
	//insert into the Todo's datastore
	var new_todo Todo

	//notice the limit reader, its a good way to protect again malicous
	//attacks imagine if somone wanted to send 500GB of data...
	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))
	if err != nil {
		panic(err)
	}

	err = req.Body.Close()
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &new_todo) //deserialized the payload(json)

	if err != nil {
		resp.Header().Set("Content-Type", "application/json; charset=UTF-8")
		resp.WriteHeader(422) //unprocessable entitiy
		//send the error back to the client in json format
		err = json.NewEncoder(resp).Encode(err)
		if err != nil {
			panic(err) //now we panic everything went to shit
		}
		return

	}
	//IF WE ARE GUGGIC proceed to fund
	//OLD(Static) temp := RepoCreateTodo(new_todo) //returns a Todo with an updated ID number
	//new_todo.Id = bson.Ob
	lastInserted := getLastInsertedTodo()
	//fmt.Println(lastInserted.Id)
	//set the new ID
	new_todo.Id = lastInserted.Id + 1

	//i need somway to give these Todos a unique ID that will 'auto increment each insert'
	//the current way of thinking is to get the latest inserted record and increment it's ID
	InsertedTodo := InsertTodoIntoDB(new_todo) //NEED A BETTER WORK AROUND

	resp.Header().Set("Content-Type", "application/json; charset=UTF-8")
	resp.WriteHeader(http.StatusOK)
	err = json.NewEncoder(resp).Encode(InsertedTodo)

	if err != nil {
		panic(err)
	}

}

func TodoShow(resp http.ResponseWriter, req *http.Request) {
	//show the todo id
	vars := mux.Vars(req) //get the url parameters from the url

	todoId := vars["todoId"]
	//fmt.Fprintf(resp, "Todo Show: ", todoId)

	resp.Header().Set("Content-Type", "application/json; charset=UTF-8")
	convertedInt, _ := strconv.Atoi(todoId)
	resultTodo, err := GetSpecificTodo(convertedInt)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(resp).Encode(err)
		if err != nil {
			panic(err)
		}
	} else {
		//if we can get and find the specific record lets return it without error
		resp.WriteHeader(http.StatusOK)
		err = json.NewEncoder(resp).Encode(resultTodo)
		if err != nil {
			panic(err)
		}
	}

}
