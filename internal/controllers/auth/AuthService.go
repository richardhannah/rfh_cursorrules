package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"totmapi/internal/controllers"
	"totmapi/internal/db"
	"totmapi/internal/di"
	"totmapi/internal/dto"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var jwtSecret = []byte("mySecretKey")

func SetRoutes(router *mux.Router) {
	router.HandleFunc("/login", LoginJwt)
	router.HandleFunc("/register", Register)
	router.HandleFunc("/changepass", Changepass)
}

func init() {
	controllers.RegisterRouteSetter(SetRoutes)
}

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

	// Get UserRepository from DI container
	userRepo := di.GetService[db.UserRepository]()

	// Find user by username
	users := userRepo.SelectByUsername(reqData.Username)
	if len(users) == 0 {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	user := users[0]

	// Check if user is enabled
	if !user.Enabled.Bool {
		http.Error(w, "User account is disabled", http.StatusUnauthorized)
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
				Role:     user.Role.String,
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

	// Get UserRepository from DI container
	userRepo := di.GetService[db.UserRepository]()

	// Check if user already exists
	existingUsers := userRepo.SelectByUsername(reqData.Username)
	if len(existingUsers) > 0 {
		http.Error(w, "Username already exists", http.StatusBadRequest)
		return
	}

	user := reqData.Username
	pass := reqData.Password
	salt := uuid.New().String()
	saltedPass := hashit(pass + salt)
	token, err := generateToken(reqData.Username)

	// Create new user DTO
	newUser := dto.UserDTO{
		ID:        uuid.New().String(),
		Username:  user,
		Password:  saltedPass,
		Salt:      salt,
		Ipaddress: r.RemoteAddr,
		Enabled:   true,
		Role:      "standard",
	}

	userRepo.Insert(newUser)

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

	// Get UserRepository from DI container
	userRepo := di.GetService[db.UserRepository]()

	// Find user by username
	users := userRepo.SelectByUsername(reqData.Username)
	if len(users) == 0 {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	user := users[0]

	inputPass := hashit(reqData.OldPassword + user.Salt)

	if user.Password == inputPass {
		// Update user with new password
		updatedUser := dto.UserDTO{
			ID:        user.ID,
			Username:  user.Username,
			Password:  hashit(reqData.NewPassword + user.Salt),
			Salt:      user.Salt,
			Ipaddress: user.Ipaddress.String,
			Enabled:   user.Enabled.Bool,
			Role:      user.Role.String,
		}

		userRepo.Update(updatedUser)

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
