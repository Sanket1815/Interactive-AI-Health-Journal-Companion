package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"go_health_sentiment/middleware"
	"go_health_sentiment/models"
	"go_health_sentiment/services"
	"go_health_sentiment/utils"
)

type JournalHandler struct {
	db   *sql.DB
	chat *services.ChatConversation
}

func NewJournalHandler(db *sql.DB, chat *services.ChatConversation) *JournalHandler {
	return &JournalHandler{
		db:   db,
		chat: chat,
	}
}

type CreateJournalRequest struct {
	Content string `json:"content"`
}

type JournalResponse struct {
	Entry    models.JournalEntryResponse `json:"entry"`
	Analysis string                      `json:"analysis"`
}

func (h *JournalHandler) CreateJournalEntry(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserKey).(int)
	if !ok {
		utils.WriteError(w, http.StatusUnauthorized, "User not authenticated")
		return
	}

	var req CreateJournalRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Validate journal content
	validationErrors := utils.ValidateJournalContent(req.Content)
	if len(validationErrors) > 0 {
		utils.WriteValidationError(w, validationErrors)
		return
	}

	req.Content = utils.SanitizeInput(req.Content)

	// Analyze journal entry
	analysis, err := h.chat.AnalyzeJournalEntry(req.Content)
	if err != nil {
		// Don't fail the request if analysis fails, just log it
		analysis = "Analysis temporarily unavailable. Your entry has been saved successfully."
	}

	// Determine sentiment (simple implementation)
	sentiment := h.determineSentiment(analysis)

	// Create journal entry
	entry := models.JournalEntry{
		Content:   req.Content,
		UserID:    userID,
		Analysis:  analysis,
		Sentiment: sentiment,
	}

	if err := entry.CreateEntry(h.db); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error creating journal entry")
		return
	}

	response := JournalResponse{
		Entry:    entry.ToResponse(),
		Analysis: analysis,
	}

	utils.WriteCreated(w, "Journal entry created successfully", response)
}

func (h *JournalHandler) GetJournalEntries(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserKey).(int)
	if !ok {
		utils.WriteError(w, http.StatusUnauthorized, "User not authenticated")
		return
	}

	// Parse pagination parameters
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit := 10 // default
	offset := 0 // default

	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 && l <= 100 {
			limit = l
		}
	}

	if offsetStr != "" {
		if o, err := strconv.Atoi(offsetStr); err == nil && o >= 0 {
			offset = o
		}
	}

	entries, err := models.GetEntriesByUser(h.db, userID, limit, offset)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error retrieving journal entries")
		return
	}

	// Convert to response format
	var responses []models.JournalEntryResponse
	for _, entry := range entries {
		responses = append(responses, entry.ToResponse())
	}

	utils.WriteSuccess(w, "Journal entries retrieved successfully", responses)
}

func (h *JournalHandler) GetJournalEntry(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserKey).(int)
	if !ok {
		utils.WriteError(w, http.StatusUnauthorized, "User not authenticated")
		return
	}

	// Extract entry ID from URL path
	entryIDStr := r.URL.Path[len("/journal/"):]
	entryID, err := strconv.Atoi(entryIDStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid entry ID")
		return
	}

	entry, err := models.GetEntryByID(h.db, entryID, userID)
	if err != nil {
		if err.Error() == "journal entry not found" {
			utils.WriteError(w, http.StatusNotFound, "Journal entry not found")
			return
		}
		utils.WriteError(w, http.StatusInternalServerError, "Error retrieving journal entry")
		return
	}

	utils.WriteSuccess(w, "Journal entry retrieved successfully", entry.ToResponse())
}

func (h *JournalHandler) determineSentiment(analysis string) string {
	// Simple sentiment analysis based on keywords
	// In a production app, you might want to use a more sophisticated approach
	analysis = strings.ToLower(analysis)
	
	positiveWords := []string{"positive", "happy", "good", "great", "excellent", "wonderful", "joy", "grateful"}
	negativeWords := []string{"negative", "sad", "bad", "terrible", "awful", "depressed", "anxious", "worried"}
	
	positiveCount := 0
	negativeCount := 0
	
	for _, word := range positiveWords {
		if strings.Contains(analysis, word) {
			positiveCount++
		}
	}
	
	for _, word := range negativeWords {
		if strings.Contains(analysis, word) {
			negativeCount++
		}
	}
	
	if positiveCount > negativeCount {
		return "positive"
	} else if negativeCount > positiveCount {
		return "negative"
	}
	
	return "neutral"
}