package blog

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"totmapi/internal/controllers"
	"totmapi/internal/db"
	"totmapi/internal/di"
	"totmapi/internal/dto"
	"totmapi/internal/models"
)

func SetRoutes(router *mux.Router) {
	router.HandleFunc("/blogposts/{id}", Handler).Methods(http.MethodGet)
	router.HandleFunc("/blogposts", Handler).Methods(http.MethodGet)
}

func init() {
	controllers.RegisterRouteSetter(SetRoutes)
}

func Handler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println(fmt.Sprintf("Blogpost ID: %s", id))

	switch httpmethod := r.Method; httpmethod {
	case http.MethodGet:
		listBlogPosts(w, r)
	case http.MethodPost:
		fmt.Println("posting")
	case http.MethodPut:
		fmt.Println("putting")
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func getSinglePost(id string) {

}

func listBlogPosts(w http.ResponseWriter, r *http.Request) {

	blogPostRepository := di.GetService[db.BlogPostRepository]()
	blogposts, err := dto.ConvertSlice[models.Blogposts, dto.BlogpostDTO](blogPostRepository.SelectAllPublished())
	if err != nil {
		log.Println("error converting model to dto")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	}

	json, err := json.Marshal(blogposts)
	if err != nil {
		log.Println("error marshalling Json")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	}

	fmt.Fprintf(w, fmt.Sprintf(string(json)))
}

//queryParams := r.URL.Query()
//published := queryParams.Get("published")

// blogposts, err := db.SelectBlogPosts(published)
//
//	if err != nil {
//		http.Error(w, "Bad request: "+err.Error(), http.StatusUnauthorized)
//		return
//	}
func create(w http.ResponseWriter, r *http.Request) {

}

func read(w http.ResponseWriter, r *http.Request) {

}

func update(w http.ResponseWriter, r *http.Request) {

}

func delete(w http.ResponseWriter, r *http.Request) {

}
