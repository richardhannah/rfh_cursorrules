package shop

import (
	"encoding/json"
	"fmt"
	"net/http"
	"totmapi/internal/controllers"
	"totmapi/internal/logger"

	"github.com/gorilla/mux"
)

func SetRoutes(router *mux.Router) {
	router.HandleFunc("/shop", Handler)
}

func init() {
	controllers.RegisterRouteSetter(SetRoutes)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	logger.Debug("Shop handler called", logger.String("query_params", fmt.Sprintf("%v", queryParams)))

	// Create a simple response
	response := map[string]interface{}{
		"message": "Shop endpoint",
		"status":  "active",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
