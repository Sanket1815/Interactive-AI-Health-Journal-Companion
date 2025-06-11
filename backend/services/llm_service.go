package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatConversation struct {
	Messages []Message
	client   *http.Client
}

func NewChatConversation() *ChatConversation {
	return &ChatConversation{
		Messages: []Message{},
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *ChatConversation) AnalyzeJournalEntry(content string) (string, error) {
	apiURL := "https://api-inference.huggingface.co/models/meta-llama/Llama-3.2-1B-Instruct"
	apiKey := os.Getenv("OPENAI_API_KEY")

	if apiKey == "" {
		return "", fmt.Errorf("API key is not set")
	}

	// Create a more structured prompt for better analysis
	prompt := fmt.Sprintf(`You are an empathetic AI mental health companion. Analyze the following journal entry and provide supportive, insightful feedback. Focus on:
1. Emotional tone and sentiment
2. Potential patterns or themes
3. Supportive encouragement
4. Gentle suggestions for reflection or self-care

Journal Entry: "%s"

Response:`, content)

	payload := map[string]interface{}{
		"inputs": prompt,
		"parameters": map[string]interface{}{
			"max_new_tokens": 200,
			"temperature":    0.7,
			"do_sample":      true,
			"top_p":          0.9,
		},
		"options": map[string]interface{}{
			"wait_for_model": true,
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("error encoding request payload: %v", err)
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := c.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	var result []map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return "", fmt.Errorf("error decoding response: %v", err)
	}

	if len(result) > 0 {
		if generatedText, ok := result[0]["generated_text"].(string); ok {
			// Extract the response part after "Response:"
			if idx := strings.Index(generatedText, "Response:"); idx != -1 {
				response := strings.TrimSpace(generatedText[idx+9:])
				if response != "" {
					// Store the conversation for context
					c.Messages = append(c.Messages, Message{Role: "user", Content: content})
					c.Messages = append(c.Messages, Message{Role: "assistant", Content: response})
					
					// Keep only last 10 messages to prevent context from growing too large
					if len(c.Messages) > 10 {
						c.Messages = c.Messages[len(c.Messages)-10:]
					}
					
					return response, nil
				}
			}
			
			// Fallback: return the full generated text if we can't parse it
			return strings.TrimSpace(generatedText), nil
		}
	}

	return "", fmt.Errorf("failed to extract response from model")
}

func (c *ChatConversation) GetConversationHistory() []Message {
	return c.Messages
}

func (c *ChatConversation) ClearHistory() {
	c.Messages = []Message{}
}