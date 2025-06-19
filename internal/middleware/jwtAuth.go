package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
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
func JwtAuthMiddleware(next http.Handler, x map[string]string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Public endpoints that don't require JWT validation
		publicPaths := []string{
			"/login",
			"/register",
			"/health",
			"/shop",
		}

		// Check if the current path is public
		for _, path := range publicPaths {
			if r.URL.Path == path {
				next.ServeHTTP(w, r)
				return
			}
		}

		// GET /blogposts is public (read-only)
		if r.URL.Path == "/blogposts" && r.Method == http.MethodGet {
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
