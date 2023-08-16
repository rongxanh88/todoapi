package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateTodo(*Todo) error
	DeleteTodo(int) error
	UpdateTodo(*Todo) error
	GetTodoById(int) (*Todo, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=vorto_test sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{db: db}, nil
}
