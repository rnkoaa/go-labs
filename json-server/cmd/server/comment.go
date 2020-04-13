package main

import "github.com/rnkoaa/go-labs/json-server/pkg/domain"

type CommentService interface {
	FindAll() ([]*domain.Comment, error)
	Find(id int) (*domain.Comment, error)
	FindByUserId(userID int) ([]*domain.Comment, error)
	FindByPostId(postID int) ([]*domain.Comment, error)
}

type CommentServiceOp struct {
	db *Database
}

func (c *CommentServiceOp) FindAll() ([]*domain.Comment, error) {
	return nil, nil
}
func (c *CommentServiceOp) Find(id int) (*domain.Comment, error) {
	return nil, nil
}
func (c *CommentServiceOp) FindByUserId(userID int) ([]*domain.Comment, error) {
	return nil, nil
}
func (c *CommentServiceOp) FindByPostId(postID int) ([]*domain.Comment, error) {
	return nil, nil
}
