package main

import "github.com/rnkoaa/go-labs/json-server/pkg/domain"

type PhotoService interface {
	FindAll() ([]*domain.Photo, error)
	Find(id int) (*domain.Photo, error)
	FindByAlbum(albumID int) ([]*domain.Photo, error)
}

type PhotoServiceOp struct {
	db *Database
}

func (p *PhotoServiceOp) FindAll() ([]*domain.Photo, error) {
	panic("implement me")
}

func (p *PhotoServiceOp) Find(id int) (*domain.Photo, error) {
	panic("implement me")
}

func (p *PhotoServiceOp) FindByAlbum(albumID int) ([]*domain.Photo, error) {
	panic("implement me")
}

var _ PhotoService = &PhotoServiceOp{}
