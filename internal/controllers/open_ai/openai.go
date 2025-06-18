package open_ai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"totmapi/internal/controllers"
	"totmapi/internal/logger"

	"github.com/gorilla/mux"
	"github.com/sashabaranov/go-openai"
)

type PromptRequest struct {
	Prompt string `json:"Prompt"`
}

type OpenAIResponse struct {
	Response string `json:"response"`
}

func SetRoutes(router *mux.Router) {
	router.HandleFunc("/openai/prompt", Handler)
}

func init() {
	controllers.RegisterRouteSetter(SetRoutes)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var reqBody PromptRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		logger.Error("Error decoding prompt request", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	logger.Info("Processing prompt request", logger.String("prompt", reqBody.Prompt))

	// Use the prompt from the request body
	prompt := reqBody.Prompt

	apiKey := "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"
	client := openai.NewClient(apiKey)

	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
		// Adjust temperature, max_tokens, etc., as needed
		MaxTokens:   200,
		Temperature: 0.7,
	}
	ctx := context.Background()

	// Send the request to the Chat Completion API
	resp, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		logger.Error("ChatCompletion error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// The response contains an array of choices; typically you just need the first one.
	if len(resp.Choices) == 0 {
		logger.Error("No response received from the API", nil)
		return
	}

	respnse := OpenAIResponse{Response: resp.Choices[0].Message.Content}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(respnse)
	if err != nil {
		logger.Error("Failed to encode response as JSON", err)
		http.Error(w, "Failed to encode response as JSON", http.StatusInternalServerError)
		return
	}
}

func callOpenAI(message string) (string, error) {
	// OpenAI API endpoint
	url := "https://api.openai.com/v1/chat/completions"

	// Request payload
	payload := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]string{
			{"role": "user", "content": message},
		},
		"max_tokens": 150,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("error marshaling payload: %w", err)
	}

	// Create request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+getOpenAIKey())

	// Make request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %w", err)
	}

	// Parse response
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("error unmarshaling response: %w", err)
	}

	// Extract the response text
	if choices, ok := result["choices"].([]interface{}); ok && len(choices) > 0 {
		if choice, ok := choices[0].(map[string]interface{}); ok {
			if message, ok := choice["message"].(map[string]interface{}); ok {
				if content, ok := message["content"].(string); ok {
					return content, nil
				}
			}
		}
	}

	logger.Error("No response received from the API", nil)
	return "Sorry, I couldn't generate a response.", nil
}

func getOpenAIKey() string {
	// In a real application, you would get this from environment variables or configuration
	return "your-openai-api-key-here"
}
