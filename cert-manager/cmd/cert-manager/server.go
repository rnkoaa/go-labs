package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Config struct {
	Address      string
	Port         string
	cronSchedule string
}

type Server struct {
	config     *Config
	Address    string
	Port       string
	httpServer *http.Server
	router     *mux.Router
}

func NewServer(config *Config) *Server {
	router := mux.NewRouter()
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))
	srv := &http.Server{
		Handler: ch(router), // set the default handler
		Addr:    fmt.Sprintf("%s:%s", config.Address, config.Port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	server := &Server{
		config:     config,
		Address:    config.Address,
		Port:       config.Port,
		router:     router,
		httpServer: srv,
	}

	server.SetupRoutes()
	return server
}

func (s Server) Start() (chan os.Signal, chan error) {
	fmt.Printf("Starting server at Address: %s:%s\n", s.Address, s.Port)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// wait for requests and serve
	serveAndWait := make(chan error)
	go func() {
		fmt.Println(fmt.Sprintf("Server listening on port %s:%s", s.Address, s.Port))
		serveAndWait <- s.httpServer.ListenAndServe()
	}()

	return sigChan, serveAndWait
}

func (s *Server) SetupRoutes() {
	s.router.HandleFunc("/health", Health)
	// router.Use(muxMetrics.Middleware())
}

func Health(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("content-type", "application/json")
	status := map[string]string{
		"status": "Up",
	}
	json.NewEncoder(res).Encode(status)
}

func (s Server) Shutdown(sigChan chan os.Signal, serveAndWait chan error) {
	// sigChan := make(chan os.Signal, 1)
	// signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// // wait for requests and serve
	// serveAndWait := make(chan error)
	// go func() {
	// 	fmt.Println(fmt.Sprintf("Server listening on port %s:%s", s.Address, s.Port))
	// 	serveAndWait <- s.httpServer.ListenAndServe()
	// }()

	// block until either an error or OS-level signals
	// to shutdown gracefully
	select {
	case err := <-serveAndWait:
		fmt.Println(err)
	case <-sigChan:
		fmt.Println("Shutdown signal received... closing server")
		// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		s.httpServer.Shutdown(ctx)
	}
}
