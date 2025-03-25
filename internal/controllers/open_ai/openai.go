package open_ai

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sashabaranov/go-openai"
	"log"
	"net/http"
	"totmapi/internal/controllers"
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

	// Ensure we only accept POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the JSON request body
	var reqBody PromptRequest
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

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
		log.Fatalf("ChatCompletion error: %v", err)
	}

	// The response contains an array of choices; typically you just need the first one.
	if len(resp.Choices) == 0 {
		log.Println("No response received from the API.")
		return
	}

	respnse := OpenAIResponse{Response: resp.Choices[0].Message.Content}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(respnse)
	if err != nil {
		http.Error(w, "Failed to encode response as JSON", http.StatusInternalServerError)
		return
	}

}
