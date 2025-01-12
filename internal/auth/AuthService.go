package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"totmapi/internal/db"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	// Ensure this is a POST request
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the JSON body
	var reqData LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		http.Error(w, "Bad request: "+err.Error(), http.StatusUnauthorized)
		return
	}

	// Process the data (e.g., validation, saving to DB, etc.)
	// For demonstration, just print it out
	log.Printf("Received: Username=%s, Password=%s\n", reqData.Username, reqData.Password)

	user, err := db.SelectUser(reqData.Username)
	if err != nil {
		http.Error(w, "Bad request: "+err.Error(), http.StatusUnauthorized)
		return
	}
	log.Printf(reqData.Password)
	log.Printf(user.Password)
	if user.Password == reqData.Password {
		// Write a response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status":"success","message":"Data received"}`)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
	}

}
