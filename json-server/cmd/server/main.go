package main

import (
	"context"
	"fmt"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	ver = "v1.0.0"
	db  *Database
	repo *RepositoryOp
)

func main() {
	db = NewDatabase("")
	// https://jsonplaceholder.typicode.com/posts/1/comments
	//https://jsonplaceholder.typicode.com/albums/1/photos
	//https://jsonplaceholder.typicode.com/users/1/albums
	//https://jsonplaceholder.typicode.com/users/1/todos
	//https://jsonplaceholder.typicode.com/users/1/posts

	repo = NewRepository(db)


	router := mux.NewRouter()
	//router.HandleFunc("/webhook", handlers.HandleWebhook).Methods("POST")
	router.HandleFunc("/users", users)
	router.HandleFunc("/users/{userId}", user)
	router.HandleFunc("/users/{userId}/albums", userAlbums)
	router.HandleFunc("/users/{userId}/todos", userTodos)
	router.HandleFunc("/users/{userId}/posts", userPosts)
	//router.HandleFunc("/users/{userId}/comments", userComments)
	router.HandleFunc("/posts", posts)
	router.HandleFunc("/posts/{postId}", post)
	router.HandleFunc("/posts/{postId}/comments", postComments)
	router.HandleFunc("/photos", photos)
	router.HandleFunc("/photos/{photoId}", photo)
	router.HandleFunc("/comments", comments)
	router.HandleFunc("/comments/{commentId}", comment)
	router.HandleFunc("/todos", todos)
	router.HandleFunc("/todos/{todoId}", todo)
	router.HandleFunc("/albums", albums)
	router.HandleFunc("/albums/{albumId}", album)
	router.HandleFunc("/albums/{albumId}/photos", albumPhotos)
	router.HandleFunc("/healthz", healthz)
	router.HandleFunc("/version", version)
	//
	//// use PORT environment variable, or default to 8080
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	//
	// CORS
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	srv := &http.Server{
		Handler: ch(router), // set the default handler
		Addr:    fmt.Sprintf("0.0.0.0:%s", port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	serveAndGracefulShutdown(srv)
}

// endpoint to test the health of the app
func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "OK")
}

// version returns the service version
func version(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, ver)
}

// Starts server with gracefully shutdown semantics
func serveAndGracefulShutdown(svr *http.Server) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// wait for requests and serve
	serveAndWait := make(chan error)
	go func() {
		log.Printf("Server listening on port %s", svr.Addr)
		serveAndWait <- svr.ListenAndServe()
	}()

	// block until either an error or OS-level signals
	// to shutdown gracefully
	select {
	case err := <-serveAndWait:
		log.Fatal(err)
	case <-sigChan:
		log.Printf("Shutdown signal received... closing server")
		// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		_ = svr.Shutdown(ctx)
	}
}
