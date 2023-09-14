package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/c7/todo.c7.se/app"
	"github.com/c7/todo.c7.se/assets"
	"github.com/c7/todo.c7.se/todo"
)

const (
	defaultPort    = "7090"
	defaultTimeout = time.Second * 30
)

func main() {
	// Create a new TODO list
	list := todo.NewList()

	// Example list
	list.Add("Bake a cake")
	list.Add("Feed the cat")
	list.Add("Take out the trash")

	// Create a Chi router
	router := chi.NewRouter()

	// Wrap the list in a service and add it to a handler mounted on the router
	app.Mount(router, app.NewHandler(app.NewService(list)))

	// Mount the assets on the router
	assets.Mount(router)

	// Get the port to listen on
	port := getPort(os.Getenv)

	// Create a new server
	server := newServer(port, router)

	// Display the localhost address and port
	fmt.Printf("Listening on http://localhost:%s\n", port)

	// Start listening
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func newServer(port string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    ":" + port,
		Handler: http.TimeoutHandler(handler, defaultTimeout, "Request timed out"),
	}
}

func getPort(getenv func(string) string) string {
	if port := getenv("PORT"); port != "" {
		return port
	}

	return defaultPort
}
