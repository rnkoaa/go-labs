package main

import "github.com/rnkoaa/go-labs/json-server/pkg/domain"

type TodoService interface {
	FindAll() ([]*domain.Todo, error)
	Find(id int) (*domain.Todo, error)
	FindByUser(id int) ([]*domain.Todo, error)
}

type TodoServiceOp struct {
	db *Database
}

func (t *TodoServiceOp) FindAll() ([]*domain.Todo, error) {
	panic("implement me")
}

func (t *TodoServiceOp) Find(id int) (*domain.Todo, error) {
	panic("implement me")
}

func (t *TodoServiceOp) FindByUser(id int) ([]*domain.Todo, error) {
	panic("implement me")
}

var _ TodoService = &TodoServiceOp{}
