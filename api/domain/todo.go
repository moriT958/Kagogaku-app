package domain

import "time"

type Todo struct {
	ID      string
	Title   string
	Status  bool
	DueDate time.Time
}

type TodoRepository interface {
	Save(Todo) (string, error)
	FindAll() ([]Todo, error)
	FindByID() (Todo, error)
}

type Usecase interface {
	CreateNewTodo(Todo) (string, error)
	ListAllTodos() ([]Todo, error)
	CompleteTodo(string) error
	DeleteTodo(string) error
}
