package main

import (
	"encoding/base64"
	"net/http"
	"os"
	"strings"
	"totmapi/internal/auth"
	"totmapi/internal/config"
	"totmapi/internal/health"
	"totmapi/internal/middleware"
	"totmapi/internal/open_ai"
)

type MyData struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func main() {

	connectionString := os.Getenv("TOTM_CONN_STRING")
	config.SetDBConfig(&connectionString)

	mux := http.NewServeMux()
	mux.HandleFunc("/openai/prompt", open_ai.Handler)
	mux.HandleFunc("/hello", health.Hello)
	mux.HandleFunc("/database", health.DatabaseHealth)
	mux.HandleFunc("/login", auth.LoginJwt)
	mux.HandleFunc("/register", auth.Register)

	handlerCORS := middleware.CorsMiddleware(mux)
	handlerAuth := middleware.JwtAuthMiddleware(handlerCORS)
	handlerSanitize := middleware.JsonSanitizeMiddleware(handlerAuth)

	http.ListenAndServe(":5150", handlerSanitize)

}

// basicAuthMiddleware enforces Basic Authentication. Only requests with valid credentials
// pass through to the next handler.
func basicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check for an "Authorization" header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			// No auth header, deny
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Authorization header typically looks like: "Basic base64encodedUserPass"
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Basic" {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Invalid authorization header", http.StatusUnauthorized)
			return
		}

		// Decode the base64
		decoded, err := base64.StdEncoding.DecodeString(parts[1])
		if err != nil {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Invalid base64 encoding", http.StatusUnauthorized)
			return
		}

		// The decoded string is typically "username:password"
		userPass := strings.SplitN(string(decoded), ":", 2)
		if len(userPass) != 2 {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Invalid credentials format", http.StatusUnauthorized)
			return
		}

		username := userPass[0]
		password := userPass[1]

		// Check if username/password is correct
		// For example, you can hard-code, or check DB, or integrate with some vault.
		if !validateUser(username, password) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// If valid, pass control to the next handler
		next.ServeHTTP(w, r)
	})
}

// validateUser is a placeholder for checking the user's credentials.
// This could be a database lookup, an LDAP check, etc.
// For demonstration, it just checks for "admin" / "secret".
func validateUser(username, password string) bool {
	// Example only. Replace with real credential check.
	return username == "guest" && password == "secret"
}
