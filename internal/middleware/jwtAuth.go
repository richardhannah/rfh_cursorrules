package middleware

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

// Key type for storing token claims in context
type contextKey string

const (
	// For retrieving claims or other data from request context
	claimsContextKey = contextKey("jwtClaims")

	// Example shared secret - in production, store this in config or environment
	jwtSecret = "mySecretKey"
)

// jwtAuthMiddleware checks for a Bearer token in the Authorization header,
// validates it, and either returns 401 or calls next handler.
func JwtAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// If path is /login, skip JWT validation
		if r.URL.Path == "/login" {
			next.ServeHTTP(w, r)
			return
		}

		// If path is /login, skip JWT validation
		if r.URL.Path == "/register" {
			next.ServeHTTP(w, r)
			return
		}

		// If OPTIONS Request, skip JWT validation
		if r.Method == http.MethodOptions {
			next.ServeHTTP(w, r)
			return
		}

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		// Expect authHeader to look like: "Bearer <token>"
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			// Ensure the signing method is HMAC
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Optionally, extract custom claims or standard claims
		// You might define a custom claims struct, but let's just parse the standard claims:
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Embed claims in context if desired
			ctx := context.WithValue(r.Context(), claimsContextKey, claims)
			// Continue with request
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}
	})
}
