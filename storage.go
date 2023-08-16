package main

import (
	"database/sql"
	"log"

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

func (s *PostgresStore) Init() error {
	return s.createTodoTable()
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

func (s *PostgresStore) CreateTodo(*Todo) error {
	return nil
}

func (s *PostgresStore) DeleteTodo(int) error {
	return nil
}

func (s *PostgresStore) UpdateTodo(*Todo) error {
	return nil
}

func (s *PostgresStore) GetTodoById(int) (*Todo, error) {
	return nil, nil
}

func (s *PostgresStore) createTodoTable() error {
	query := `CREATE TABLE IF NOT EXISTS todos (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			description TEXT,
			completed BOOLEAN,
			created_at TIMESTAMPTZ DEFAULT NOW(),
			updated_at TIMESTAMPTZ
	);
	`
	_, err := s.db.Exec(query)

	if err != nil {
		log.Fatal(err)
	}

	return err
}
