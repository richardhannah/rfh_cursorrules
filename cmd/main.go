package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"totmapi/internal/auth"
	"totmapi/internal/blog"
	"totmapi/internal/config"
	"totmapi/internal/controllers"
	_ "totmapi/internal/controllers/shop"
	"totmapi/internal/health"
	"totmapi/internal/middleware"
	"totmapi/internal/open_ai"
)

func main() {

	connectionString := os.Getenv("TOTM_CONN_STRING")
	config.SetDBConfig(&connectionString)

	authmap := make(map[string]string)
	mux := mux.NewRouter()
	controllers.SetAllRoutes(mux)
	//shop.SetRoutes(mux)
	//mux.HandleFunc("test", shop.Handler)
	mux.HandleFunc("/openai/prompt", open_ai.Handler)
	mux.HandleFunc("/hello", health.Hello)
	mux.HandleFunc("/database", health.DatabaseHealth)
	mux.HandleFunc("/login", auth.LoginJwt)
	mux.HandleFunc("/register", auth.Register)
	mux.HandleFunc("/blogposts/{id}", blog.BlogHandler).Methods(http.MethodGet)
	mux.HandleFunc("/blogposts", blog.BlogHandler).Methods(http.MethodGet)
	mux.HandleFunc("/changepass", auth.Changepass)
	mux.HandleFunc("shops/{shopid}", auth.LoginJwt)

	handlerCORS := middleware.CorsMiddleware(mux)
	handlerAuth := middleware.JwtAuthMiddleware(handlerCORS, authmap)
	handlerSanitize := middleware.JsonSanitizeMiddleware(handlerAuth)

	http.ListenAndServe(":5150", handlerSanitize)

}
