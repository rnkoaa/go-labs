package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rnkoaa/go-labs/json-server/pkg/domain"
	"net/http"
	"strconv"
)

func sendResponse(res interface{}, w http.ResponseWriter) {
	js, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}

func posts(w http.ResponseWriter, r *http.Request) {
	posts, _ := repo.Post.FindAll()

	sendResponse(posts, w)
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	postId := vars["postId"]
	if postId == "" || postId == "0" {
		w.WriteHeader(404)
		_, _ = fmt.Fprintf(w, `{"message": "post not found"}`)
		return
	}
	uID, err := strconv.Atoi(postId)
	if err != nil {
		w.WriteHeader(404)
		_, _ = fmt.Fprintf(w, `{"message": "post not found"}`)
		return
	}
	post, _ := repo.Post.Find(uID)
	if post != nil {
		sendResponse(&post, w)
		return
	}
	w.WriteHeader(404)
	_, _ = fmt.Fprintf(w, `{"message": "post not found"}`)
}

func postComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	postId := vars["postId"]
	if postId == "" || postId == "0" {
		w.WriteHeader(404)
		_, _ = fmt.Fprintf(w, `{"message": "post not found"}`)
		return
	}
	uID, err := strconv.Atoi(postId)
	if err != nil {
		w.WriteHeader(404)
		_, _ = fmt.Fprintf(w, `{"message": "post not found"}`)
		return
	}
	post, _ := repo.Post.FindComments(uID)
	if post != nil {
		sendResponse(&post, w)
		return
	}
	w.WriteHeader(404)
	_, _ = fmt.Fprintf(w, `{"message": "post not found"}`)
}

func users(w http.ResponseWriter, r *http.Request) {
	users, _ := repo.User.FindAll()
	sendResponse(users, w)
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId := vars["userId"]
	if userId == "" || userId == "0" {
		w.WriteHeader(404)
		_, _ = fmt.Fprintf(w, `{"message": "user not found"}`)
		return
	}
	uID, err := strconv.Atoi(userId)
	if err != nil {
		w.WriteHeader(404)
		_, _ = fmt.Fprintf(w, `{"message": "user not found"}`)
		return
	}
	user, _ := repo.User.Find(uID)
	if user != nil {
		sendResponse(&user, w)
		return
	}
	w.WriteHeader(404)
	_, _ = fmt.Fprintf(w, `{"message": "user not found"}`)

}

func userAlbums(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId := vars["userId"]
	if userId == "" || userId == "0" {
		w.WriteHeader(404)
		_, _ = fmt.Fprintf(w, `{"message": "user not found"}`)
		return
	}
	uID, err := strconv.Atoi(userId)
	if err != nil {
		w.WriteHeader(404)
		_, _ = fmt.Fprintf(w, `{"message": "user not found"}`)
		return
	}

	albums, _ := repo.User.FindAlbums(uID)

	if albums != nil {
		sendResponse(&albums, w)
		return
	}
	w.WriteHeader(404)
	_, _ = fmt.Fprintf(w, `{"message": "user not found"}`)

}

func userTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId := vars["userId"]
	if userId == "" || userId == "0" {
		w.WriteHeader(404)
		_, _ = fmt.Fprintf(w, `{"message": "user not found"}`)
		return
	}
	uID, err := strconv.Atoi(userId)
	if err != nil {
		w.WriteHeader(404)
		_, _ = fmt.Fprintf(w, `{"message": "user not found"}`)
		return
	}
	todos, _ := repo.User.FindTodos(uID)

	if todos != nil {
		sendResponse(&todos, w)
		return
	}
	w.WriteHeader(404)
	_, _ = fmt.Fprintf(w, `{"message": "user not found"}`)
}

//func userComments(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	vars := mux.Vars(r)
//	userId := vars["userId"]
//	if userId == "" || userId == "0" {
//		w.WriteHeader(404)
//		_, _ = fmt.Fprintf(w, `{"message": "user not found"}`)
//		return
//	}
//	uID, err := strconv.Atoi(userId)
//	if err != nil {
//		w.WriteHeader(404)
//		_, _ = fmt.Fprintf(w, `{"message": "user not found"}`)
//		return
//	}
//	comments, _ := repo.User.FindCommentsByUser(uID)
//	if comments != nil {
//		sendResponse(&comments, w)
//		return
//	}
//	w.WriteHeader(404)
//	_, _ = fmt.Fprintf(w, `{"message": "user not found"}`)
//}

func userPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId := vars["userId"]
	if userId == "" || userId == "0" {
		w.WriteHeader(404)
		_, _ = fmt.Fprintf(w, `{"message": "user not found"}`)
		return
	}
	uID, err := strconv.Atoi(userId)
	if err != nil {
		w.WriteHeader(404)
		_, _ = fmt.Fprintf(w, `{"message": "user not found"}`)
		return
	}
	posts, _ := repo.User.FindPosts(uID)
	sendResponse(&posts, w)
}

func albums(w http.ResponseWriter, r *http.Request) {
	sendResponse(db.Albums, w)
}

func album(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	albumId := vars["albumId"]
	if albumId == "" || albumId == "0" {
		w.WriteHeader(404)
		_, _ = fmt.Fprintf(w, `{"message": "album not found"}`)
		return
	}
	aID, err := strconv.Atoi(albumId)
	if err != nil {
		w.WriteHeader(404)
		_, _ = fmt.Fprintf(w, `{"message": "album not found"}`)
		return
	}
	//db.Users
	var album *domain.Album
	for _, a := range db.Albums {
		if a.ID == aID {
			album = a
			break
		}
	}

	if album != nil {
		sendResponse(&album, w)
		return
	}
	w.WriteHeader(404)
	_, _ = fmt.Fprintf(w, `{"message": "album not found"}`)
}

func albumPhotos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	albumId := vars["albumId"]
	if albumId == "" || albumId == "0" {
		w.WriteHeader(404)
		_, _ = fmt.Fprintf(w, `{"message": "album not found"}`)
		return
	}
	aID, err := strconv.Atoi(albumId)
	if err != nil {
		w.WriteHeader(404)
		_, _ = fmt.Fprintf(w, `{"message": "album not found"}`)
		return
	}
	//db.Users
	albumPhotos, _ := repo.Album.FindPhotos(aID)
	if albumPhotos != nil {
		sendResponse(&albumPhotos, w)
		return
	}
	w.WriteHeader(404)
	_, _ = fmt.Fprintf(w, `{"message": "album not found"}`)
}

func comments(w http.ResponseWriter, r *http.Request) {
	sendResponse(db.Comments, w)
}

func comment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	commentId := vars["commentId"]
	if commentId == "" || commentId == "0" {
		w.WriteHeader(404)
		_, _ = fmt.Fprintf(w, `{"message": "comment not found"}`)
		return
	}
	uID, err := strconv.Atoi(commentId)
	if err != nil {
		w.WriteHeader(404)
		_, _ = fmt.Fprintf(w, `{"message": "comment not found"}`)
		return
	}
	//db.Users
	var comment *domain.Comment
	for _, u := range db.Comments {
		if u.ID == uID {
			comment = u
			break
		}
	}

	if comment != nil {
		sendResponse(&comment, w)
		return
	}
	w.WriteHeader(404)
	_, _ = fmt.Fprintf(w, `{"message": "comment not found"}`)
}

func photos(w http.ResponseWriter, r *http.Request) {
	sendResponse(db.Photos, w)
}

func photo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	photoId := vars["photoId"]
	if photoId == "" || photoId == "0" {
		w.WriteHeader(404)
		_, _ = fmt.Fprintf(w, `{"message": "photo not found"}`)
	}
	aID, err := strconv.Atoi(photoId)
	if err != nil {
		w.WriteHeader(404)
		_, _ = fmt.Fprintf(w, `{"message": "photo not found, invalid key"}`)
	}
	var photo *domain.Photo
	for _, a := range db.Photos {
		if a.ID == aID {
			photo = a
			break
		}
	}

	if photo != nil {
		sendResponse(&photo, w)
		return
	}
	w.WriteHeader(404)
	_, _ = fmt.Fprintf(w, `{"message": "photo not found"}`)
}

func todos(w http.ResponseWriter, r *http.Request) {
	sendResponse(db.Todos, w)
}

func todo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	if todoId == "" || todoId == "0" {
		w.WriteHeader(404)
		_, _ = fmt.Fprintf(w, `{"message": "todo not found"}`)
	}
	aID, err := strconv.Atoi(todoId)
	if err != nil {
		w.WriteHeader(404)
		_, _ = fmt.Fprintf(w, `{"message": "todo not found, invalid key"}`)
	}
	var todo *domain.Todo
	for _, a := range db.Todos {
		if a.ID == aID {
			todo = a
			break
		}
	}

	if todo != nil {
		sendResponse(&todo, w)
		return
	}
	w.WriteHeader(404)
	_, _ = fmt.Fprintf(w, `{"message": "todo not found"}`)
}
