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
		WriteJSON(w, http.StatusInternalServerError, ErrMsg{Error: err.Error()})
	}

	WriteJSON(w, http.StatusOK, todos)
}

func (s *APIServer) handleTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		WriteJSON(w, http.StatusNotFound, ErrMsg{Error: err.Error()})
		return
	}

	todo, err := s.store.GetTodoById(id)

	if err != nil {
		WriteJSON(w, http.StatusNotFound, ErrMsg{Error: err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, todo)
}

func (s *APIServer) handleCreateTodo(w http.ResponseWriter, r *http.Request) {
	req := new(CreateTodoRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrMsg{Error: err.Error()})
		return
	}
	todo := NewTodo(req.Title, req.Description, req.Completed)

	resp, err := s.store.CreateTodo(todo)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrMsg{Error: err.Error()})
		return
	}
	WriteJSON(w, http.StatusCreated, resp)
}

func (s *APIServer) handleUpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		WriteJSON(w, http.StatusNotFound, ErrMsg{Error: err.Error()})
		return
	}
	todoTemp := Todo{ID: id}
	if err := json.NewDecoder(r.Body).Decode(&todoTemp); err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrMsg{Error: err.Error()})
		return
	}
	updatedTodo, err := s.store.UpdateTodo(&todoTemp)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrMsg{Error: err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, updatedTodo)
}

func (s *APIServer) handleDeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		WriteJSON(w, http.StatusNotFound, ErrMsg{Error: err.Error()})
		return
	}

	if err := s.store.DeleteTodo(id); err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrMsg{Error: err.Error()})
		return
	}
	WriteJSON(w, http.StatusNoContent, nil)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
