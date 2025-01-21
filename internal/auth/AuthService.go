package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"net/http"
	"time"
	"totmapi/internal/db"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegistrationRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
	Role     string `json:"role"`
}

type ChangePassRequest struct {
	Username    string `json:"username"`
	OldPassword string `json:"oldpassword"`
	NewPassword string `json:"newpassword"`
}

var jwtSecret = []byte("mySecretKey")

func LoginJwt(w http.ResponseWriter, r *http.Request) {
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

	user, err := db.SelectUser(reqData.Username)
	if err != nil {
		http.Error(w, "Bad request: "+err.Error(), http.StatusUnauthorized)
		return
	}

	inputPass := hashit(reqData.Password + user.Salt)

	if user.Password == inputPass {

		token, err := generateToken(reqData.Username)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			// Write a response
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			authresponse := AuthResponse{
				Username: reqData.Username,
				Token:    token,
				Role:     user.Role,
			}

			json, err := json.Marshal(authresponse)
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
			}

			fmt.Fprintf(w, fmt.Sprintf(string(json)))
		}

	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
	}

}

func Register(w http.ResponseWriter, r *http.Request) {
	// Ensure this is a POST request
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the JSON body
	var reqData RegistrationRequest
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		http.Error(w, "Bad request: "+err.Error(), http.StatusUnauthorized)
		return
	}

	user := reqData.Username
	pass := reqData.Password
	salt := uuid.New().String()
	saltedPass := hashit(pass + salt)
	token, err := generateToken(reqData.Username)

	err = db.InsertUser(user, saltedPass, salt)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	} else {
		// Write a response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		authresponse := AuthResponse{
			Username: reqData.Username,
			Token:    token,
			Role:     "standard",
		}

		json, err := json.Marshal(authresponse)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
		}

		fmt.Fprintf(w, fmt.Sprintf(string(json)))
	}

}

func Changepass(w http.ResponseWriter, r *http.Request) {
	// Ensure this is a POST request
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the JSON body
	var reqData ChangePassRequest
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		http.Error(w, "Bad request: "+err.Error(), http.StatusUnauthorized)
		return
	}

	user, err := db.SelectUser(reqData.Username)
	if err != nil {
		http.Error(w, "Bad request: "+err.Error(), http.StatusUnauthorized)
		return
	}

	inputPass := hashit(reqData.OldPassword + user.Salt)

	if user.Password == inputPass {
		db.UpdatePassword(reqData.Username, hashit(reqData.NewPassword+user.Salt))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
	}

}

// generateToken creates a JWT token with standard claims plus a subject
func generateToken(username string) (string, error) {
	// Add claims
	claims := jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(1 * time.Hour).Unix(), // token expires in 1 hour
		"iat": time.Now().Unix(),                    // issued at
	}

	// Create a token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign with your secret
	return token.SignedString(jwtSecret)
}

func hashit(saltedPass string) string {
	// Compute the SHA-256 hash
	hash := sha256.New()
	hash.Write([]byte(saltedPass))
	hashedBytes := hash.Sum(nil)

	// Encode the hashed bytes to a hex string
	hashHex := hex.EncodeToString(hashedBytes)
	return hashHex
}
