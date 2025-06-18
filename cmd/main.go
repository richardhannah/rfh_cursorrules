package main

import (
	"net/http"
	"os"
	"totmapi/internal/config"
	"totmapi/internal/controllers"
	_ "totmapi/internal/controllers/auth"
	_ "totmapi/internal/controllers/blog"
	_ "totmapi/internal/controllers/health"
	_ "totmapi/internal/controllers/open_ai"
	_ "totmapi/internal/controllers/shop"
	"totmapi/internal/di"
	"totmapi/internal/logger"
	"totmapi/internal/middleware"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	// Initialize structured logger
	logger.Init()

	logger.Info("Starting TOTM API",
		logger.String("version", "1.0.0"),
		logger.String("environment", "development"),
	)

	connectionString := os.Getenv("TOTM_CONN_STRING")
	if connectionString == "" {
		logger.Error("Database connection string not found",
			&configError{message: "TOTM_CONN_STRING environment variable is required"},
			logger.String("env_var", "TOTM_CONN_STRING"),
		)
		os.Exit(1)
	}

	config.SetDBConfig(&connectionString)
	logger.Info("Database configuration loaded")

	di.InitializeServices()
	logger.Info("Dependency injection container initialized")

	authmap := make(map[string]string)
	mux := mux.NewRouter()
	controllers.SetAllRoutes(mux)

	handlerCORS := middleware.CorsMiddleware(mux)
	handlerAuth := middleware.JwtAuthMiddleware(handlerCORS, authmap)
	handlerSanitize := middleware.JsonSanitizeMiddleware(handlerAuth)

	logger.Info("Server starting",
		logger.String("port", "5150"),
		logger.String("address", "0.0.0.0:5150"),
	)

	if err := http.ListenAndServe(":5150", handlerSanitize); err != nil {
		logger.Error("Failed to start server", err,
			logger.String("port", "5150"),
		)
		os.Exit(1)
	}
}

// configError implements the error interface
type configError struct {
	message string
}

func (e *configError) Error() string {
	return e.message
}
