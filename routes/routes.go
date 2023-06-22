package routes

import (
	"fmt"
	"log"
	"net/http"
	"restapi_array/handlers"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func SetupRoutes() {
	fmt.Println("Setup routes here....")

	//configure CORS
	allowedOrigin := gorillaHandlers.AllowedOrigins([]string{"http://localhost:3000"})
	allowedMethods := gorillaHandlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"})
	allwedHeaders := gorillaHandlers.AllowedHeaders([]string{"Access-Control-Allow-Origin", "X-Requested-With", "Accept", "Content-Type", "Origin"})

	//intialize router
	muxRouter := mux.NewRouter()

	//setup api endpoints
	muxRouter.HandleFunc("/", handlers.Hello).Methods("GET")
	muxRouter.HandleFunc("/welcome", handlers.Welcome).Methods("GET")
	muxRouter.HandleFunc("/todos/all", handlers.GetAllTodos).Methods("GET")
	muxRouter.HandleFunc("/todos/add", handlers.AddNewTodo).Methods("POST")
	muxRouter.HandleFunc("/todos/get/{myid}", handlers.GetTodoById).Methods("GET")
	muxRouter.HandleFunc("/todos/delete/{myid}", handlers.DeleteTodoById).Methods("DELETE")
	muxRouter.HandleFunc("/todos/update/{myid}", handlers.UpdateTodoById).Methods("PUT")

	//listen to the incoming request on port 4000
	log.Print("Starting API server on port :4000")
	log.Fatal(http.ListenAndServe(":4000", gorillaHandlers.CORS(allowedOrigin, allwedHeaders, allowedMethods)(muxRouter)))
}
