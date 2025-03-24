package shop

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"totmapi/internal/controllers"
)

func SetRoutes(router *mux.Router) {
	router.HandleFunc("/shop", Handler)
}

func init() {
	controllers.RegisterRouteSetter(SetRoutes)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
