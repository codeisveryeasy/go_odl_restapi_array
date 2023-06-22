package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"restapi_array/models"
	"strconv"

	"github.com/gorilla/mux"
)

var todos = []models.Todo{
	{
		ID:        1,
		Task:      "Do something 1",
		Status:    "progress",
		Completed: "yes",
	},
	{
		ID:        2,
		Task:      "Do something 2",
		Status:    "done",
		Completed: "no",
	},
	{
		ID:        3,
		Task:      "Do something 3",
		Status:    "yettostart",
		Completed: "yes",
	},
}

func Hello(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Root endpoint called -> /")

	n, error := fmt.Fprint(response, "Hello from REST API in Go Lang")
	if error != nil {
		log.Fatal("Not able to write the response")
	} else {
		fmt.Printf("Written %v bytes in response", n)
	}

}

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome endpoint called -> /welcome")

	n, error := fmt.Fprint(w, "Welcome to ToDo REST API!")
	if error != nil {
		log.Fatal("Not able to write the response")
	} else {
		fmt.Printf("Written %v bytes in response", n)
	}
}

//create all CRUD handlers

// get all todos from array
func GetAllTodos(response http.ResponseWriter, r *http.Request) {
	log.Print("endpoint /todos/all called....")

	//Newencoder returns the encoder which writes to response
	//Encode writes the JSON encoded data to response
	json.NewEncoder(response).Encode(todos)
}

/*
	{
		id:       3,
		task:      "Do something 1",
		status:    "progress",
		completed: false,
	}

*/
//add new todo
func AddNewTodo(response http.ResponseWriter, request *http.Request) {
	log.Print("endpoint called -> /todos/add")

	//extract requestbody from the incoming request
	bytesRequestBody, error := ioutil.ReadAll(request.Body)
	if error != nil {
		log.Fatal("Error while reading request body")
	}

	//print requestBody in the console
	log.Print(string(bytesRequestBody))

	//create an instance of type Todo which can be inserted in todos array
	var newTodo = models.Todo{}
	log.Print(newTodo)
	//unmarshal -> parses the JSON encoded data i.e. bytesRequestBody  and stores the result in &newTodo
	json.Unmarshal(bytesRequestBody, &newTodo)
	log.Print(&newTodo)
	todos = append(todos, newTodo)
	//return the lenght of todos array
	log.Print("Length: ", len(todos))
	json.NewEncoder(response).Encode(len(todos))

}

// get todo by id
func GetTodoById(response http.ResponseWriter, request *http.Request) {
	log.Print("endpoint called -> /get/todo/myid")
	//log.Print(request)
	//extract path/route variable myid
	log.Print(mux.Vars(request)["myid"])
	//convert path variable value to int64
	lookforid, error := strconv.ParseInt(mux.Vars(request)["myid"], 10, 64)
	if error != nil {
		log.Fatal("Error converting path/route variable from string to int64 in  findbyid handler")
	}
	//iterate on todos to find lookforid
	for _, todo := range todos {
		if todo.ID == lookforid {
			json.NewEncoder(response).Encode(todo)
		}
	}

	/*
		Activity 01:
		1. Return the valid response as below if the todo with the lookforid is not found
	*/
	// json.NewEncoder(response).Encode(`{
	// 	message:"Not found"
	// }`)

}

// delete todo by id
func DeleteTodoById(response http.ResponseWriter, request *http.Request) {
	log.Print("endpoint called -> /todos/delete/myid")
	//extract myid from request and convert it to int64
	lookforid, error := strconv.ParseInt(mux.Vars(request)["myid"], 10, 64)
	if error != nil {
		log.Fatal("Error converting path/route variable from string to int64 in delete handler")
	}
	//iterate on todos to find lookforid
	for index, todo := range todos {
		if todo.ID == lookforid {
			log.Print(todos[:index])
			log.Print(todos[index+1:])
			//log.Print(todos[index+1:]...)
			todos = append(todos[:index], todos[index+1:]...)
			/*
				Activity 03:
				1. Modify the below response to JSON format before sending
			*/
			json.NewEncoder(response).Encode("Deleted by id: " + string(lookforid))
		}
	}

	/*
		Activity 02:
		1. Return the valid response as below if the todo with the lookforid is not found
	*/
	// json.NewEncoder(response).Encode(`{
	// 	message:"Not found"
	// }`)
}

// update todo by id
func UpdateTodoById(response http.ResponseWriter, request *http.Request) {
	log.Print("endpoint called -> todo/update/myid")
	//get id for which todo needs to be updated
	lookforid, error := strconv.ParseInt(mux.Vars(request)["myid"], 10, 64)
	if error != nil {
		log.Fatal("Error converting path/route variable from string to int64 in update handler")
	}
	log.Print("Edit id:", lookforid)
	//get JSON object from request body with new todo data
	log.Print("Request Body")
	log.Print(request.Body)
	bytesRequestBody, error := ioutil.ReadAll(request.Body)
	if error != nil {
		log.Fatal("Error while reading request body in update handler")
	}
	//str1 := fmt.Sprintf("%s", bytesRequestBody)
	//log.Print(str1)
	fmt.Println(bytesRequestBody)
	//create a new instance of type ToDo
	var updateTodo = models.Todo{}
	// err := json.NewDecoder(request.Body).Decode(&updateTodo)
	// if err != nil {
	// 	log.Fatal("Error while reading request body in update handler")
	// }
	//unmarshal values from bytesRequestBody to updateTodo
	json.Unmarshal(bytesRequestBody, &updateTodo)
	log.Print(updateTodo)
	log.Print(lookforid)
	//iterate on todos array to find lookforid
	//if matching id is found, then repla	ce the values

	/*
				Activity 05: What will happen if you send partial request body?
				e.g.
		 		{
				    "task": "Do something new with 1",
				  }

				How will you resolve this issue?
	*/
	for index, todo := range todos {
		if todo.ID == lookforid {
			todo.Completed = updateTodo.Completed
			todo.Status = updateTodo.Status
			todo.Task = updateTodo.Task
			todo.ID = lookforid
			//replace the updated todo in todos array
			todos[index] = todo
			json.NewEncoder(response).Encode(todo)
		}
	}

	/*
		Activity 04:
		1. Return the valid response as below if the todo with the lookforid is not found
	*/
	// json.NewEncoder(response).Encode(`{
	// 	message:"Not found"
	// }`)

}
