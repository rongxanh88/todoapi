package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenerAddress string
}

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

var todos []Todo

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func NewAPIServer(listenerAddress string) *APIServer {
	return &APIServer{listenerAddress: listenerAddress}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/todos", s.handleTodos).Methods("GET")
	router.HandleFunc("/todos", s.handleCreateTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", s.handleUpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", s.handleDeleteTodo).Methods("DELETE")

	fmt.Println("Server started at %d", s.listenerAddress)
	http.ListenAndServe(s.listenerAddress, router)
}

func (s *APIServer) handleTodos(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusOK, todos)
}

func (s *APIServer) handleCreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todos = append(todos, todo)
	WriteJSON(w, http.StatusCreated, todos)
}

func (s *APIServer) handleUpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var updatedTodo Todo
	_ = json.NewDecoder(r.Body).Decode(&updatedTodo)
	for index, todo := range todos {
		if todo.ID == id {
			todos[index] = updatedTodo
			WriteJSON(w, http.StatusOK, updatedTodo)
			return
		}
	}
	WriteJSON(w, http.StatusNoContent, nil)
}

func (s *APIServer) handleDeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for index, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:index], todos[index+1:]...)
			break
		}
	}
	WriteJSON(w, http.StatusNoContent, nil)
}
