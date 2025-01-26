// package services

// import (
//     "bytes"
//     "encoding/json"
//     "fmt"
//     "io/ioutil"
//     "net/http"
//     "os"
//     "log"

//     "github.com/joho/godotenv"
// )

// // Initialize the environment variables
// func init() {
//     err := godotenv.Load()
//     if err != nil {
//         log.Println("Error loading .env file")
//     }
// }

// func AnalyzeJournalEntry(content string) string {
//     // apiURL := "https://api.openai.com/v1/chat/completions"
//     // apiKey := os.Getenv("OPENAI_API_KEY")
//     apiURL := "meta-llama/Llama-2-7b"
//     apiKey := os.Getenv("OPENAI_API_KEY")

//     if apiKey == "" {
//         return "Error: API key is not set"
//     }

//     // Create the payload for the Chat Completion API
//     payload := map[string]interface{}{
//         "model": "gpt-3.5-turbo",  // Use the correct model name for chat completions
//         "messages": []map[string]string{
//             {
//                 "role":    "user",
//                 "content": "Analyze the following journal entry: " + content,
//             },
//         },
//         "max_tokens":   200,
//         "temperature":  0.7,
//     }

//     body, err := json.Marshal(payload)
//     if err != nil {
//         return "Error encoding request payload"
//     }

//     req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(body))
//     if err != nil {
//         return "Error creating request"
//     }

//     req.Header.Set("Content-Type", "application/json")
//     req.Header.Set("Authorization", "Bearer "+apiKey)

//     client := &http.Client{}
//     resp, err := client.Do(req)
//     if err != nil {
//         return "Error sending request"
//     }
//     defer resp.Body.Close()

//     // Read the response body
//     respBody, err := ioutil.ReadAll(resp.Body)
//     if err != nil {
//         return "Error reading response body"
//     }

//     // Log the response body for debugging
//     fmt.Println("OpenAI API response:", string(respBody))

//     // Check if the API returned an error
//     if resp.StatusCode != http.StatusOK {
//         // Parse the error message from the response
//         var errorResult map[string]interface{}
//         err = json.Unmarshal(respBody, &errorResult)
//         if err != nil {
//             return fmt.Sprintf("API request failed with status %d: %s", resp.StatusCode, string(respBody))
//         }
//         if errorData, ok := errorResult["error"].(map[string]interface{}); ok {
//             if errorMessage, ok := errorData["message"].(string); ok {
//                 return fmt.Sprintf("API error: %s", errorMessage)
//             }
//         }
//         return fmt.Sprintf("API request failed with status %d: %s", resp.StatusCode, string(respBody))
//     }

//     // Parse the successful response
//     var result map[string]interface{}
//     err = json.Unmarshal(respBody, &result)
//     if err != nil {
//         return "Error decoding response"
//     }

//     // Extract and return the analysis from the response
//     if choices, ok := result["choices"].([]interface{}); ok && len(choices) > 0 {
//         if choice, ok := choices[0].(map[string]interface{}); ok {
//             if message, ok := choice["message"].(map[string]interface{}); ok {
//                 if content, ok := message["content"].(string); ok {
//                     return content
//                 }
//             }
//         }
//     }

//     return "Analysis failed"
// }


package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"log"
    "strings"

	"github.com/joho/godotenv"
)

// Initialize the environment variables
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
}

// Message represents a single message in the conversation
type Message struct {
	Role    string `json:"role"`    // "user" or "assistant"
	Content string `json:"content"` // Message content
}

// ChatConversation manages the conversation between user and model
type ChatConversation struct {
	Messages []Message
}

// NewChatConversation initializes a new chat conversation
func NewChatConversation() *ChatConversation {
	return &ChatConversation{Messages: []Message{}}
}

// AnalyzeJournalEntry sends a journal entry to the model for analysis and maintains context
func (c *ChatConversation) AnalyzeJournalEntry(content string) (string, error) {
	// Hugging Face model endpoint
	apiURL := "https://api-inference.huggingface.co/models/meta-llama/Llama-3.2-1B-Instruct" // Replace with your Hugging Face model endpoint
	apiKey := os.Getenv("OPENAI_API_KEY")

	// Check if the API key is set
	if apiKey == "" {
		return "", fmt.Errorf("API key is not set")
	}

	// Append user message to the conversation
	c.Messages = append(c.Messages, Message{Role: "user", Content: content})

	// Create a single string representing the conversation so far
	chatHistory := ""
	for _, msg := range c.Messages {
		if msg.Role == "user" {
			chatHistory += fmt.Sprintf("User: %s\n", msg.Content)
		} else if msg.Role == "assistant" {
			chatHistory += fmt.Sprintf("Assistant: %s\n", msg.Content)
		}
	}

	// Add the current user input to the prompt
	prompt := fmt.Sprintf("The following is a conversation between a user and an empathetic assistant:\n%s\nAssistant:", chatHistory)

	// Prepare the payload
	payload := map[string]interface{}{
		"inputs": prompt,
		"parameters": map[string]interface{}{
			"max_length": 300, // Increase max_length for longer responses
			"temperature": 0.7,
		},
	}

	// Encode the payload to JSON
	body, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("Error encoding request payload: %v", err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("Error creating request: %v", err)
	}

	// Set the necessary headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Error reading response body: %v", err)
	}

	// Log the response body for debugging
	fmt.Println("Hugging Face API response:", string(respBody))

	// Parse the response
	var result []map[string]interface{}
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return "", fmt.Errorf("Error decoding response: %v", err)
	}

	// Extract the generated text from the response
	if len(result) > 0 {
		if generatedText, ok := result[0]["generated_text"].(string); ok {
			// Find all occurrences of "Assistant:"
			lines := strings.Split(generatedText, "\n")
			for _, line := range lines {
				if strings.HasPrefix(line, "Assistant:") {
					// Return the first "Assistant:" response
					return strings.TrimSpace(strings.TrimPrefix(line, "Assistant:")), nil
				}
			}
		}
	}
	return "", fmt.Errorf("Failed to extract response from model")
}

