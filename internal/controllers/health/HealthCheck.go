package health

import (
	"fmt"
	"net/http"
	"totmapi/internal/controllers"
	"totmapi/internal/db"
	"totmapi/internal/di"
	"totmapi/internal/logger"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func SetRoutes(router *mux.Router) {
	router.HandleFunc("/health", HealthCheck)
}

func init() {
	controllers.RegisterRouteSetter(SetRoutes)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	// Get the database connection from the container
	dbContext := di.GetService[db.DbContext]()
	if dbContext == nil {
		logger.Error("Failed to get database context", &healthError{message: "database context not available"})
		http.Error(w, "Database context not available", http.StatusInternalServerError)
		return
	}

	// Test the database connection
	db := dbContext.DB.(*sqlx.DB)
	if err := db.Ping(); err != nil {
		logger.Error("Error pinging database", err)
		http.Error(w, "Database connection failed", http.StatusServiceUnavailable)
		return
	}

	// Test a simple query
	var result int
	err := db.Get(&result, "SELECT 1")
	if err != nil {
		logger.Error("Error executing simple database query", err)
		http.Error(w, "Database query failed", http.StatusServiceUnavailable)
		return
	}

	// Log the result for debugging
	logger.Info("Health check completed successfully",
		logger.Int("query_result", result),
	)

	// Return a simple JSON response
	response := fmt.Sprintf(`{"status": "healthy", "database": "connected", "query_result": %d}`, result)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(response))
}

// healthError implements the error interface
type healthError struct {
	message string
}

func (e *healthError) Error() string {
	return e.message
}
