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
	"github.com/robfig/cron/v3"
)

// Config - config object for server
type Config struct {
	Address string      `mapstructure:"address"`
	Port    string      `mapstructure:"port"`
	Vault   VaultConfig `mapstructure:"vault"`
	// cronSchedule string      `mapstructure:"cron.schedule"`
	Cron struct {
		Schedule string `mapstructure:"schedule"`
	} `mapstructure:"cron"`
}

type VaultConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Address  string `mapstructure:"address"`
}

// Server - struct to handle server functions
type Server struct {
	config     *Config
	Address    string
	Port       string
	httpServer *http.Server
	router     *mux.Router
	cron       *cron.Cron
}

// NewServer - create server object
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
		cron:       cron.New(),
		router:     router,
		httpServer: srv,
	}

	server.SetupRoutes()
	return server
}

// Start - starts the server application
func (s Server) Start() (chan os.Signal, chan error) {
	fmt.Printf("Starting server at Address: %s:%s\n", s.Address, s.Port)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		s.cron.AddFunc(s.config.Cron.Schedule, func() {
			currentTime := time.Now().Format("2006-01-02T15:04:05-0700")
			fmt.Printf("[%s] Starting cron job\n", currentTime)
		})
		s.cron.Run()
	}()

	// wait for requests and serve
	serveAndWait := make(chan error)
	go func() {
		fmt.Println(fmt.Sprintf("Server listening on port %s:%s", s.Address, s.Port))
		serveAndWait <- s.httpServer.ListenAndServe()
	}()

	return sigChan, serveAndWait
}

// SetupRoutes - adds routes functionalities to server
func (s *Server) SetupRoutes() {
	s.router.HandleFunc("/health", Health)
	// router.Use(muxMetrics.Middleware())
}

// Health - indicates that the server is running
func Health(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("content-type", "application/json")
	status := map[string]string{
		"status": "Up",
	}
	json.NewEncoder(res).Encode(status)
}

// Shutdown - allows graceful shutdown of server
func (s Server) Shutdown(sigChan chan os.Signal, serveAndWait chan error) {
	select {
	case err := <-serveAndWait:
		fmt.Println(err)
	case <-sigChan:
		fmt.Println("Shutdown signal received... closing server")
		s.cron.Stop()
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		s.httpServer.Shutdown(ctx)
	}
}
