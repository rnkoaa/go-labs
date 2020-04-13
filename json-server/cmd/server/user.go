package main

import (
	"errors"
	"github.com/rnkoaa/go-labs/json-server/pkg/domain"
)

type UserService interface {
	FindAll() ([]*domain.User, error)
	Find(id int) (*domain.User, error)
	FindPosts(id int) ([]*domain.Post, error)
	FindTodos(id int) ([]*domain.Todo, error)
	FindAlbums(id int) ([]*domain.Album, error)
	//FindCommentsByUser(id int) ([]*domain.Comment, error)
}

type UserServiceOp struct {
	db *Database
}

func (u *UserServiceOp) FindAll() ([]*domain.User, error) {
	return u.db.Users, nil
}
func (u *UserServiceOp) Find(id int) (*domain.User, error) {
	for _, u := range db.Users {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, errors.New("user not found")
}
func (u *UserServiceOp) FindPosts(id int) ([]*domain.Post, error) {
	var posts = make([]*domain.Post, 0)
	for _, u := range db.Posts {
		if u.UserID == id {
			posts = append(posts, u)
		}
	}
	return posts, nil
}
func (u *UserServiceOp) FindTodos(id int) ([]*domain.Todo, error) {
	var todos = make([]*domain.Todo, 0)
	for _, u := range db.Todos {
		if u.UserID == id {
			todos = append(todos, u)
		}
	}
	return todos, nil
}
func (u *UserServiceOp) FindAlbums(id int) ([]*domain.Album, error) {
	var albums = make([]*domain.Album, 0)
	for _, u := range db.Albums {
		if u.UserID == id {
			albums = append(albums, u)
		}
	}
	return albums, nil
}
//func (u *UserServiceOp) FindCommentsByUser(id int) ([]*domain.Comment, error) {
//	var comments = make([]*domain.Comment, 0)
//	for _, u := range db.Comments {
//		if u. == id {
//			comments = append(comments, u)
//		}
//	}
//	return comments, nil
//}
