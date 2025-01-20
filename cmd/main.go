package main

import (
	"net/http"
	"os"
	"totmapi/internal/auth"
	"totmapi/internal/blog"
	"totmapi/internal/config"
	"totmapi/internal/health"
	"totmapi/internal/middleware"
	"totmapi/internal/open_ai"
)

type MyData struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func main() {

	connectionString := os.Getenv("TOTM_CONN_STRING")
	config.SetDBConfig(&connectionString)

	mux := http.NewServeMux()
	mux.HandleFunc("/openai/prompt", open_ai.Handler)
	mux.HandleFunc("/hello", health.Hello)
	mux.HandleFunc("/database", health.DatabaseHealth)
	mux.HandleFunc("/login", auth.LoginJwt)
	mux.HandleFunc("/register", auth.Register)
	mux.HandleFunc("/blogposts", blog.BlogHandler)

	handlerCORS := middleware.CorsMiddleware(mux)
	handlerAuth := middleware.JwtAuthMiddleware(handlerCORS)
	handlerSanitize := middleware.JsonSanitizeMiddleware(handlerAuth)

	http.ListenAndServe(":5150", handlerSanitize)

}
