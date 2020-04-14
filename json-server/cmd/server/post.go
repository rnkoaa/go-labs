package main

import (
	"errors"
	"github.com/rnkoaa/go-labs/json-server/pkg/domain"
)

type PostService interface {
	FindAll() ([]*domain.Post, error)
	Find(id int) (*domain.Post, error)
	FindByUser(id int) ([]*domain.Post, error)
	FindComments(id int) ([]*domain.Comment, error)
}

type PostServiceOp struct {
	db *Database
}

func (p *PostServiceOp) FindAll() ([]*domain.Post, error) {
	return p.db.Posts, nil
}

func (p *PostServiceOp) Find(id int) (*domain.Post, error) {
	for _, u := range p.db.Posts {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, errors.New("post not found")
}

func (p *PostServiceOp) FindByUser(id int) ([]*domain.Post, error) {
	var posts = make([]*domain.Post, 0)
	for _, u := range p.db.Posts {
		if u.UserID == id {
			posts = append(posts, u)
		}
	}
	return posts, nil
}

func (p *PostServiceOp) FindComments(id int) ([]*domain.Comment, error) {
	var comments = make([]*domain.Comment, 0)
	for _, u := range p.db.Comments {
		if u.PostID == id {
			comments = append(comments, u)
		}
	}
	return comments, nil
}

var _ PostService = &PostServiceOp{}
