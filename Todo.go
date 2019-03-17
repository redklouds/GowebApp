package main

import "time"

/*
This file stores the various models we will be using in this service
*/
/*
***************** Models **********************
 */
type Todo struct {
	Id        int       `json:"id"`             //ID of the record
	Name      string    `json:"name,omitempty"` //doing this will serialize the json to the lowercase Name
	Completed bool      `json:"completed"`      //don't moit becuase default is false we want that!
	Due       time.Time `json:"due"`            //don't omit because we want the date time
}

//def <Name> <Type>
type Todos []Todo

/*
***************** End Models **********************
 */
