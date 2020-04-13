package main

import "github.com/rnkoaa/go-labs/json-server/pkg/domain"

type AlbumService interface {
	FindAll() ([]*domain.Album, error)
	Find(id int) (*domain.Album, error)
	FindByUserId(userID int) ([]*domain.Album, error)
	FindPhotos(ID int) ([]*domain.Photo, error)
}

type AlbumServiceOp struct {
	db *Database
}

func (a *AlbumServiceOp) FindAll() ([]*domain.Album, error) {
	return a.db.Albums, nil
}
func (a *AlbumServiceOp) Find(id int) (*domain.Album, error) {
	return nil, nil
}
func (a *AlbumServiceOp) FindByUserId(userID int) ([]*domain.Album, error) {
	return nil, nil
}
func (a *AlbumServiceOp) FindPhotos(id int) ([]*domain.Photo, error) {
	// could probably do a func ()
	var photos = make([]*domain.Photo, 0)
	for _, u := range db.Photos {
		if u.AlbumID == id {
			photos = append(photos, u)
		}
	}
	return photos, nil
}
