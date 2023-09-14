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

// Defaults for the port and timeout
const (
	defaultPort    = "7090"
	defaultTimeout = time.Second * 30
)

func main() {
	// Create a new TODO list
	list := todo.NewList()

	// Example list of TODO items
	list.Add("Bake a cake")
	list.Add("Feed the cat")
	list.Add("Take out the trash")

	// Create a new Chi router
	router := newChiRouter(list)

	// Get the port to listen on
	port := getPort(os.Getenv)

	// Create a new server
	server := newServer(port, router)

	// Display the localhost address and port
	fmt.Printf("Listening on http://localhost:%s\n", port)

	// Start listening on the port
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func newChiRouter(repository todo.Repository) *chi.Mux {
	// Create a Chi router
	router := chi.NewRouter()

	// Wrap the repository in a service, add it to a handler mounted on the router
	app.Mount(router, app.NewHandler(app.NewService(repository)))

	// Mount the assets on the router
	assets.Mount(router)

	// Return the router
	return router
}

func getPort(getenv func(string) string) string {
	// Get the PORT from the environment, if present
	if port := getenv("PORT"); port != "" {
		return port
	}

	// Return the default port
	return defaultPort
}

func newServer(port string, handler http.Handler) *http.Server {
	// Create a HTTP server for the given port and handler.
	// The handler is wrapped in a timeout handler set to the defaultTimeout
	return &http.Server{
		Addr:    ":" + port,
		Handler: http.TimeoutHandler(handler, defaultTimeout, "Request timed out"),
	}
}
