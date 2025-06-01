package main

import (
	"github.com/gorilla/mux"
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
	"totmapi/internal/middleware"
)

func main() {

	connectionString := os.Getenv("TOTM_CONN_STRING")
	config.SetDBConfig(&connectionString)

	di.InitializeServices()

	authmap := make(map[string]string)
	mux := mux.NewRouter()
	controllers.SetAllRoutes(mux)

	handlerCORS := middleware.CorsMiddleware(mux)
	handlerAuth := middleware.JwtAuthMiddleware(handlerCORS, authmap)
	handlerSanitize := middleware.JsonSanitizeMiddleware(handlerAuth)

	http.ListenAndServe(":5150", handlerSanitize)

}
