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
	store           Storage
}

func NewAPIServer(listenerAddress string, store Storage) *APIServer {
	return &APIServer{
		listenerAddress: listenerAddress,
		store:           store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/todos", s.handleTodos).Methods("GET")
	router.HandleFunc("/todos", s.handleCreateTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", s.handleTodo).Methods("GET")
	router.HandleFunc("/todos/{id}", s.handleUpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", s.handleDeleteTodo).Methods("DELETE")

	fmt.Println("Server started at port", s.listenerAddress)
	http.ListenAndServe(s.listenerAddress, router)
}

func (s *APIServer) handleTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := s.store.GetTodos()

	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, err)
	}

	WriteJSON(w, http.StatusOK, todos)
}

func (s *APIServer) handleTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	todo, err := s.store.GetTodoById(id)

	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, err)
		return
	}

	if todo == nil {
		WriteJSON(w, http.StatusNotFound, nil)
	} else {
		WriteJSON(w, http.StatusOK, todo)
	}
}

func (s *APIServer) handleCreateTodo(w http.ResponseWriter, r *http.Request) {
	req := new(CreateTodoRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Printf("%+v\n", req)
		WriteJSON(w, http.StatusBadRequest, err)
		return
	}
	todo := NewTodo(req.Title, req.Description, req.Completed)

	if err := s.store.CreateTodo(todo); err != nil {
		WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	WriteJSON(w, http.StatusCreated, todo)
}

func (s *APIServer) handleUpdateTodo(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	// id, _ := strconv.Atoi(params["id"])
	// var updatedTodo Todo
	// _ = json.NewDecoder(r.Body).Decode(&updatedTodo)
	// for index, todo := range todos {
	// 	if todo.ID == id {
	// 		todos[index] = updatedTodo
	// 		WriteJSON(w, http.StatusOK, updatedTodo)
	// 		return
	// 	}
	// }
	// WriteJSON(w, http.StatusNotFound, nil)
}

func (s *APIServer) handleDeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	if err := s.store.DeleteTodo(id); err != nil {
		WriteJSON(w, http.StatusInternalServerError, nil)
	}
	WriteJSON(w, http.StatusNoContent, nil)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
