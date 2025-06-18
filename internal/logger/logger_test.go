package logger

import (
	"testing"
)

func TestLogger(t *testing.T) {
	// Initialize the logger
	Init()

	// Test different log levels with structured fields
	Info("Application started",
		String("service", "TOTM API"),
		Int("port", 5150),
		Bool("debug", false),
	)

	Warn("Database connection slow",
		String("database", "postgres"),
		Int("response_time_ms", 1500),
	)

	Error("Failed to process request",
		&testError{message: "connection timeout"},
		String("endpoint", "/api/users"),
		String("method", "POST"),
		Int("status_code", 500),
	)

	Debug("Processing user request",
		String("user_id", "12345"),
		String("action", "login"),
	)
}

// testError implements the error interface for testing
type testError struct {
	message string
}

func (e *testError) Error() string {
	return e.message
}
