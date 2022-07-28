package middleware

import (
	"PostgresCRUD/models" // models package where Review schema is defined
	_ "database/sql"
	"encoding/json" // package to encode and decode the json into struct and vice versa
	"fmt"
	"log"
	"net/http" // used to access the request and reviewResponse object of the api

	// used to read the environment variable
	_ "strconv" // package used to covert string into int type

	"github.com/gorilla/mux" // used to get the params from the route

	_ "github.com/lib/pq" // postgres golang driver
)

// reviewResponse format
type reviewResponse struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// create connection with postgres db

// CreateReview create a review in the postgres db
func CreateReview(w http.ResponseWriter, r *http.Request) {
	// set the header to content type x-www-form-urlencoded
	// Allow all origin to handle cors issue
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// create an empty review of type models.Review
	var review models.Review

	// decode the json request to Review
	err := json.NewDecoder(r.Body).Decode(&review)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	// call insert review function and pass the review
	insertID := insertReview(review)

	// format a reviewResponse object
	res := reviewResponse{
		ID:      insertID,
		Message: "Review created successfully",
	}

	// send the reviewResponse
	json.NewEncoder(w).Encode(res)
}

// GetAllReview will return all the reviews
func GetAllReview(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// get all the reviews in the db
	reviews, err := getAllReviews()

	if err != nil {
		log.Fatalf("Unable to get all review. %v", err)
	}

	// send all the reviews as reviewResponse
	json.NewEncoder(w).Encode(reviews)
}

// GetFilm will return any film by its category
func GetReviewsByFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// get the id from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id type from string to int
	film := params["id"]

	// call the getFilm function with category to retrieve variable number of films
	films, err := getReviewsByFilm(string(film))

	if err != nil {
		log.Fatalf("Unable to get film. %v", err)
	}

	// send the filmResponse
	json.NewEncoder(w).Encode(films)

}

//------------------------- handler functions ----------------
// insert one review in the DB
func insertReview(review models.Review) int64 {

	// create the postgres db connection
	db := SetupDB()

	// close the db connection
	defer db.Close()

	// create the insert sql query
	// returning Review_id will return the id of the inserted review
	sqlStatement := `INSERT INTO review (Review_id, Review_text, Rating, Customer_id, Film_id) VALUES ($1, $2, $3, $4, $5) RETURNING Review_id`

	// the inserted id will store in this id
	var id int64

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, review.Review_id, review.Review_text, review.Rating, review.Customer_id, review.Film_id).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	// return the inserted id
	return id
}

// get one review from the DB by its Review_id
func getAllReviews() ([]models.Review, error) {
	// create the postgres db connection
	db := SetupDB()

	// close the db connection
	defer db.Close()

	var reviews []models.Review

	// create the select sql query
	sqlStatement := `SELECT * FROM Review`

	// execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var review models.Review

		// unmarshal the row object to review
		err = rows.Scan(&review.Review_id, &review.Review_text, &review.Rating, &review.Customer_id, &review.Film_id)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the review in the reviews slice
		reviews = append(reviews, review)

	}

	// return empty review on error
	return reviews, err
}

// get all reviews for film by id
func getReviewsByFilm(id string) ([]models.Review, error) {
	// create the postgres db connection
	db := SetupDB()

	// close the db connection
	defer db.Close()

	var reviews []models.Review

	// create the select sql query
	sqlStatement := `SELECT r.* FROM review r INNER JOIN film f ON f.film_id = r.film_id WHERE r.film_id=$1`

	// execute the sql statement
	rows, err := db.Query(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var review models.Review

		// unmarshal the row object to film
		err = rows.Scan(&review.Review_id, &review.Review_text, &review.Rating, &review.Customer_id, &review.Film_id)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the film in the films slice
		reviews = append(reviews, review)

	}

	// return empty film on error
	return reviews, err
}
