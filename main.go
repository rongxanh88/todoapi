package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	// router := todo.Routes()

	// // Define API endpoints
	// router.HandleFunc("/todos", GetTodos).Methods("GET")
	// router.HandleFunc("/todos/{id}", GetTodo).Methods("GET")
	// router.HandleFunc("/todos", CreateTodo).Methods("POST")
	// router.HandleFunc("/todos/{id}", UpdateTodo).Methods("PUT")
	// router.HandleFunc("/todos/{id}", DeleteTodo).Methods("DELETE")

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", router)
}
