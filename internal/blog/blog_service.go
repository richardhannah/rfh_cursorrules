package blog

import (
	"encoding/json"
	"fmt"
	"net/http"
	"totmapi/internal/db"
)

func BlogHandler(w http.ResponseWriter, r *http.Request) {

	blogposts, err := db.SelectBlogPosts()
	if err != nil {
		http.Error(w, "Bad request: "+err.Error(), http.StatusUnauthorized)
		return
	}

	json, err := json.Marshal(blogposts)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	}

	fmt.Fprintf(w, fmt.Sprintf(string(json)))

}
