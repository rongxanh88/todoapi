package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateTodo(*Todo) error
	DeleteTodo(int) error
	UpdateTodo(*Todo) error
	GetTodoById(int) (*Todo, error)
	GetTodos() ([]*Todo, error)
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

func (s *PostgresStore) GetTodos() ([]*Todo, error) {
	query := "SELECT * FROM todos;"

	rows, err := s.db.Query(query)

	if err != nil {
		return nil, err
	}

	todos := []*Todo{}

	for rows.Next() {
		todo := new(Todo)

		err := rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Description,
			&todo.Completed,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

func (s *PostgresStore) CreateTodo(todo *Todo) error {
	query := `INSERT INTO todos (title, description, completed, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5);
	`
	_, err := s.db.Query(query, todo.Title, todo.Description, todo.Completed, time.Now().UTC(), time.Now().UTC())
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) DeleteTodo(id int) error {
	query := "DELETE FROM Todos WHERE ID = $1;"

	if _, err := s.db.Query(query, id); err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) UpdateTodo(*Todo) error {
	return nil
}

func (s *PostgresStore) GetTodoById(id int) (*Todo, error) {
	query := "SELECT * FROM Todos WHERE ID = $1;"

	rows, err := s.db.Query(query, id)

	if err != nil {
		return nil, err
	}

	todo := new(Todo)
	for rows.Next() {
		err := rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Description,
			&todo.Completed,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
	}

	if todo.ID != id {
		return nil, nil
	} else {
		return todo, nil
	}
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
