package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/cors"
	"go_health_sentiment/auth"
	"go_health_sentiment/config"
	"go_health_sentiment/db"
	"go_health_sentiment/handlers"
	"go_health_sentiment/middleware"
	"go_health_sentiment/services"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize auth with JWT secret
	auth.InitializeAuth(cfg.JWTSecret)

	// Connect to database
	database, err := db.NewConnection(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer database.Close()

	// Initialize services
	chat := services.NewChatConversation()

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(database.DB)
	journalHandler := handlers.NewJournalHandler(database.DB, chat)

	// Initialize rate limiter (60 requests per minute, burst of 10)
	rateLimiter := middleware.NewRateLimiter(60, 10)

	// Create router
	mux := http.NewServeMux()

	// Health check endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if err := database.HealthCheck(); err != nil {
			http.Error(w, "Database connection failed", http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"healthy"}`))
	})

	// Public routes
	mux.HandleFunc("/register", authHandler.Register)
	mux.HandleFunc("/login", authHandler.Login)

	// Protected routes
	mux.Handle("/profile", middleware.JWTMiddleware(http.HandlerFunc(authHandler.GetProfile)))
	mux.Handle("/journal", middleware.JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			journalHandler.CreateJournalEntry(w, r)
		case http.MethodGet:
			journalHandler.GetJournalEntries(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})))

	// Individual journal entry route
	mux.Handle("/journal/", middleware.JWTMiddleware(http.HandlerFunc(journalHandler.GetJournalEntry)))

	// Setup CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   cfg.AllowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	// Apply middleware chain
	handler := middleware.LoggingMiddleware(
		rateLimiter.RateLimitMiddleware(
			c.Handler(mux),
		),
	)

	// Start server
	server := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: handler,
	}

	// Graceful shutdown
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		log.Println("Shutting down server...")
		if err := server.Close(); err != nil {
			log.Printf("Error during server shutdown: %v", err)
		}
	}()

	fmt.Printf("Server starting on port %s\n", cfg.ServerPort)
	fmt.Printf("Health check available at: http://localhost:%s/health\n", cfg.ServerPort)
	
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server failed to start:", err)
	}

	log.Println("Server stopped")
}