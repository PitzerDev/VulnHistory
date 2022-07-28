package router

import (
	"PostgresCRUD/middleware"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	//Routes for film
	router.HandleFunc("/api/film/{id}", middleware.GetFilm).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/film", middleware.GetAllFilm).Methods("GET", "OPTIONS")
	//Film filters
	router.HandleFunc("/api/film/title/{title}", middleware.GetFilmByTitle).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/film/rating/{rating}", middleware.GetFilmByRating).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/film/category/{category}", middleware.GetFilmByCategory).Methods("GET", "OPTIONS")

	//Routes for review
	router.HandleFunc("/api/review", middleware.GetAllReview).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/newreview", middleware.CreateReview).Methods("POST", "OPTIONS")
	//Review filters
	router.HandleFunc("/api/film/reviews/{id}", middleware.GetReviewsByFilm).Methods("GET", "OPTIONS")

	return router
}
