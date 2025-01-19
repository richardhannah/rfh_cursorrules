package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/microcosm-cc/bluemonday"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// jsonSanitizeMiddleware reads JSON from the body (for POST/PUT requests),
// sanitises any string fields using Bluemonday, then rewrites the request body.
func JsonSanitizeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Only handle JSON-based POST or PUT
		if (r.Method == http.MethodPost || r.Method == http.MethodPut) &&
			strings.Contains(r.Header.Get("Content-Type"), "application/json") {

			// Read the entire body
			body, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Failed to read request body", http.StatusBadRequest)
				return
			}
			// Important to close the original body
			r.Body.Close()

			// Parse the JSON into an interface{} (or any)
			var data any
			if err := json.Unmarshal(body, &data); err != nil {
				http.Error(w, "Invalid JSON", http.StatusBadRequest)
				return
			}

			// Sanitize recursively
			policy := bluemonday.StrictPolicy()
			sanitized := sanitizeData(data, policy)

			// Re-encode the sanitized data
			newBody, err := json.Marshal(sanitized)
			if err != nil {
				http.Error(w, "Failed to encode sanitized JSON", http.StatusInternalServerError)
				return
			}

			// Replace the request body with the sanitized version
			r.Body = io.NopCloser(bytes.NewReader(newBody))
			r.ContentLength = int64(len(newBody))
			r.Header.Set("Content-Length", strconv.Itoa(len(newBody)))
		}

		// Pass to the next handler
		next.ServeHTTP(w, r)
	})
}

// sanitizeData recursively traverses the parsed JSON (maps, slices, strings),
// and sanitizes string fields using the provided bluemonday policy.
func sanitizeData(v any, policy *bluemonday.Policy) any {
	switch val := v.(type) {
	case map[string]any:
		for k, sub := range val {
			val[k] = sanitizeData(sub, policy)
		}
		return val
	case []any:
		for i, sub := range val {
			val[i] = sanitizeData(sub, policy)
		}
		return val
	case string:
		return policy.Sanitize(val)
	default:
		// Int, float, bool, nil, etc. remain unchanged
		return v
	}
}
