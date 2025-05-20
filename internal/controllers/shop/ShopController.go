package shop

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"totmapi/internal/controllers"
)

func SetRoutes(router *mux.Router) {
	router.HandleFunc("/shop/{shopid}", Handler)
}

func init() {
	controllers.RegisterRouteSetter(SetRoutes)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	fmt.Println(queryParams)
	data := GetTestShop()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json, err := json.Marshal(data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Fprintf(w, fmt.Sprintf(string(json)))
}
