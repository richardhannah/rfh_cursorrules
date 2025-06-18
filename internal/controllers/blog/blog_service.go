package blog

import (
	"encoding/json"
	"net/http"
	"totmapi/internal/controllers"
	"totmapi/internal/db"
	"totmapi/internal/di"
	"totmapi/internal/dto"
	"totmapi/internal/logger"
	"totmapi/internal/models"

	"github.com/gorilla/mux"
)

func SetRoutes(router *mux.Router) {
	router.HandleFunc("/blogposts", GetBlogPosts)
	router.HandleFunc("/blogposts/{id}", GetBlogPost)
	router.HandleFunc("/blogposts", PostBlogPost).Methods("POST")
	router.HandleFunc("/blogposts", PutBlogPost).Methods("PUT")
}

func init() {
	controllers.RegisterRouteSetter(SetRoutes)
}

func GetBlogPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	logger.Debug("Getting blog post by ID", logger.String("blogpost_id", id))

	blogPostRepository := di.GetService[db.BlogPostRepository]()
	blogposts := blogPostRepository.SelectById(id)

	if len(blogposts) > 0 {
		blogpost := blogposts[0]
		blogpostDTO, err := dto.ConvertToDTO[models.Blogposts, dto.BlogpostDTO](blogpost)
		if err != nil {
			logger.Error("Error converting model to dto", err,
				logger.String("blogpost_id", id),
			)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		json, err := json.Marshal(blogpostDTO)
		if err != nil {
			logger.Error("Error marshalling JSON", err,
				logger.String("blogpost_id", id),
			)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	} else {
		http.Error(w, "Blog post not found", http.StatusNotFound)
	}
}

func GetBlogPosts(w http.ResponseWriter, r *http.Request) {
	logger.Debug("Getting all published blog posts")

	blogPostRepository := di.GetService[db.BlogPostRepository]()
	blogposts, err := dto.ConvertSlice[models.Blogposts, dto.BlogpostDTO](blogPostRepository.SelectAllPublished())
	if err != nil {
		logger.Error("Error converting blogposts to DTOs", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json, err := json.Marshal(blogposts)
	if err != nil {
		logger.Error("Error marshalling blogposts JSON", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func PostBlogPost(w http.ResponseWriter, r *http.Request) {
	logger.Info("Creating new blog post")

	var blogPost dto.BlogpostDTO
	if err := json.NewDecoder(r.Body).Decode(&blogPost); err != nil {
		logger.Error("Error decoding blog post request", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	blogPostRepository := di.GetService[db.BlogPostRepository]()
	blogPostRepository.Insert(blogPost)

	logger.Info("Blog post created successfully",
		logger.String("title", blogPost.Title),
	)

	w.WriteHeader(http.StatusCreated)
}

func PutBlogPost(w http.ResponseWriter, r *http.Request) {
	logger.Info("Updating blog post")

	var blogPost dto.BlogpostDTO
	if err := json.NewDecoder(r.Body).Decode(&blogPost); err != nil {
		logger.Error("Error decoding blog post update request", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	blogPostRepository := di.GetService[db.BlogPostRepository]()
	blogPostRepository.Update(blogPost)

	logger.Info("Blog post updated successfully",
		logger.String("blogpost_id", blogPost.BlogpostID),
		logger.String("title", blogPost.Title),
	)

	w.WriteHeader(http.StatusOK)
}

func create(w http.ResponseWriter, r *http.Request) {

}

func read(w http.ResponseWriter, r *http.Request) {

}

func update(w http.ResponseWriter, r *http.Request) {

}

func delete(w http.ResponseWriter, r *http.Request) {

}
