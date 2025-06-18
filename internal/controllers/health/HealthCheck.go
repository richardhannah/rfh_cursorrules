package health

import (
	"fmt"
	"net/http"
	"totmapi/internal/controllers"
	"totmapi/internal/db"
	"totmapi/internal/di"
	"totmapi/internal/logger"
	"totmapi/internal/models"

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
		logger.Fatal("Failed to get database context", &healthError{message: "database context not available"})
	}

	// Test the database connection
	db := dbContext.DB.(*sqlx.DB)
	if err := db.Ping(); err != nil {
		logger.Fatal("Error pinging database", err)
	}

	// Test a simple query
	var person models.Person
	err := db.Get(&person, "SELECT * FROM person LIMIT 1")
	if err != nil {
		logger.Fatal("Error querying 'person' table", err)
	}

	// Log the result for debugging
	logger.Info("Health check completed successfully",
		logger.Int("person_id", person.ID),
		logger.String("person_name", person.Name),
	)

	// Return a simple JSON response
	response := fmt.Sprintf(`{"status": "healthy", "database": "connected", "person_id": %d, "person_name": "%s"}`, person.ID, person.Name)
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
