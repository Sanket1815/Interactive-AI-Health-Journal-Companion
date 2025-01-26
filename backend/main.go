
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors" // Import the CORS package
	"go_health_sentiment/db"
	"go_health_sentiment/handlers"
	"go_health_sentiment/services"
	"go_health_sentiment/middleware" // Import your JWT middleware
)

func main() {
	// Connect to the database
	database, err := db.OpenConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
	chat := services.NewChatConversation()
	// Create a new ServeMux for better route handling
	mux := http.NewServeMux()

	// Define public routes (No authentication needed)
	mux.HandleFunc("/register", handlers.Register(database))
	mux.HandleFunc("/login", handlers.Login(database))

	// Define protected route (JWT authentication required)
	mux.Handle("/journal", middleware.JWTMiddleware(http.HandlerFunc(handlers.CreateJournalEntry(database, chat))))

	// Set up CORS to allow requests from your Vue frontend
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8081"}, // Allow only your frontend URL
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // HTTP methods allowed
		AllowedHeaders:   []string{"Authorization", "Content-Type"}, // Headers allowed
		AllowCredentials: true,  // Allow cookies (if needed)
	})

	// Wrap the mux with the CORS middleware
	handler := c.Handler(mux)

	// Start the server
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}