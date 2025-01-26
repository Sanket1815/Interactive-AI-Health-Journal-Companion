// package handlers

// import (
// 	"database/sql"
// 	"encoding/json"
// 	"net/http"
// 	"go_health_sentiment/models"
// 	"go_health_sentiment/services"
// 	"go_health_sentiment/middleware"
// )

// // CreateJournalEntry handles journal entry creation
// func CreateJournalEntry(db *sql.DB) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// Extract user_id from context (set by JWT middleware)
// 		userID, ok := r.Context().Value(middleware.UserKey).(int)
// 		if !ok {
// 			http.Error(w, "User not authenticated", http.StatusUnauthorized)
// 			return
// 		}

// 		// Decode the journal entry from the request body
// 		var entry models.JournalEntry
// 		err := json.NewDecoder(r.Body).Decode(&entry)
// 		if err != nil {
// 			http.Error(w, "Invalid input", http.StatusBadRequest)
// 			return
// 		}

// 		// Assign the user ID from the context to the journal entry
// 		entry.UserID = userID

// 		// Create the journal entry in the database
// 		err = entry.CreateEntry(db)
// 		if err != nil {
// 			http.Error(w, "Error creating entry", http.StatusInternalServerError)
// 			return
// 		}

// 		// Send journal entry content to LLM for analysis
// 		analysis := services.AnalyzeJournalEntry(entry.Content)

// 		// Return a success response
// 		json.NewEncoder(w).Encode(map[string]string{
// 			"message":  "Journal entry created",
// 			"analysis": analysis,
// 		})
// 	}
// }


package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"go_health_sentiment/models"
	"go_health_sentiment/services"
	"go_health_sentiment/middleware"
	"log"
)

// CreateJournalEntry handles journal entry creation
func CreateJournalEntry(db *sql.DB, chat *services.ChatConversation) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract user_id from context (set by JWT middleware)
		userID, ok := r.Context().Value(middleware.UserKey).(int)
		if !ok {
			http.Error(w, "User not authenticated", http.StatusUnauthorized)
			return
		}

		// Decode the journal entry from the request body
		var entry models.JournalEntry
		err := json.NewDecoder(r.Body).Decode(&entry)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		// Validate the journal entry content
		if entry.Content == "" {
			http.Error(w, "Journal content cannot be empty", http.StatusBadRequest)
			return
		}

		// Assign the user ID from the context to the journal entry
		entry.UserID = userID

		// Create the journal entry in the database
		err = entry.CreateEntry(db)
		log.Println("Error creating journal entry:", entry)
		if err != nil {
			http.Error(w, "Error creating entry", http.StatusInternalServerError)
			return
		}

		// Send journal entry content to LLM for analysis
		analysis, err := chat.AnalyzeJournalEntry(entry.Content)
		if err != nil {
			http.Error(w, "Error analyzing journal entry: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Return a success response
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{
			"message":  "Journal entry created successfully",
			"analysis": analysis,
		})
	}
}
