package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "time"

    "github.com/gorilla/mux"
    "gobookstore/data"
    "gobookstore/handlers"
)


func main() {
	// Initialize MongoDB connection
	if err := data.ConnectDB(); err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer data.DisconnectDB()

	// Set up router
	r := mux.NewRouter()

	// Register routes
	r.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", handlers.GetBook).Methods("GET")
	r.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")

	// Graceful shutdown setup
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		log.Println("Server running on port 8080...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on :8080: %v\n", err)
		}
	}()

	// Wait for interrupt signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	log.Println("Shutting down gracefully...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server exited.")
}
