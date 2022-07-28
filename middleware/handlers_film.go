package middleware

import (
	"PostgresCRUD/models" // models package where Film schema is defined
	"database/sql"
	"encoding/json" // package to encode and decode the json into struct and vice versa
	"fmt"
	"log"
	"net/http" // used to access the request and object of the api

	// used to read the environment variable
	"strconv" // package used to covert string into int type

	"github.com/gorilla/mux" // used to get the params from the route

	_ "github.com/lib/pq" // postgres golang driver
)

// GetFilm will return a single film by its id
func GetFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// get the Film_id from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	// call the getFilm function with film id to retrieve a single film
	film, err := getFilm(int64(id))

	if err != nil {
		log.Fatalf("Unable to get film. %v", err)
	}

	// send the filmResponse
	json.NewEncoder(w).Encode(film)
}

// GetAllFilm will return all the films
func GetAllFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// get all the films in the db
	films, err := getAllFilms()

	if err != nil {
		log.Fatalf("Unable to get all film. %v", err)
	}

	// send all the films as filmResponse
	json.NewEncoder(w).Encode(films)
}

// GetFilm will return a single film by its title
func GetFilmByTitle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// get the Title from the request params, key is "title"
	params := mux.Vars(r)

	title := params["title"]

	// call the getFilm function with title to retrieve a single film
	film, err := getFilmByTitle(string(title))

	if err != nil {
		log.Fatalf("Unable to get film. %v", err)
	}

	// send the filmResponse
	json.NewEncoder(w).Encode(film)
}

// GetFilm will return any film by its rating
func GetFilmByRating(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// get the Rating from the request params, key is "rating"
	params := mux.Vars(r)

	rating := params["rating"]

	// call the getFilm function with rating to retrieve a single film
	films, err := getFilmByRating(string(rating))

	if err != nil {
		log.Fatalf("Unable to get film. %v", err)
	}

	// send the filmResponse
	json.NewEncoder(w).Encode(films)

}

// GetFilm will return any film by its category
func GetFilmByCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// get the Category from the request params, key is "category"
	params := mux.Vars(r)

	category := params["category"]

	// call the getFilm function with category to retrieve variable number of films
	films, err := getFilmByCategory(string(category))

	if err != nil {
		log.Fatalf("Unable to get film. %v", err)
	}

	// send the filmResponse
	json.NewEncoder(w).Encode(films)

}

//------------------------- handler functions ----------------

// get one film from the DB by its Film_id
func getFilm(id int64) (models.Film, error) {
	// create the postgres db connection
	db := SetupDB()

	// close the db connection
	defer db.Close()

	// create a film of models.Film type
	var film models.Film

	// create the select sql query
	sqlStatement := `SELECT * FROM film WHERE Film_id=$1`

	// execute the sql statement
	row := db.QueryRow(sqlStatement, id)

	// unmarshal the row object to film
	err := row.Scan(&film.Film_id, &film.Title, &film.Description, &film.Release_year, &film.Language_id, &film.Rental_duration, &film.Rental_rate, &film.Length, &film.Replacement_cost, &film.Rating, &film.Last_update, &film.Special_features, &film.Fulltext)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return film, nil
	case nil:
		return film, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty film on error
	return film, err
}

// get one film from the DB by its Film_id
func getAllFilms() ([]models.Film, error) {
	// create the postgres db connection
	db := SetupDB()

	// close the db connection
	defer db.Close()

	var films []models.Film

	// create the select sql query
	sqlStatement := `SELECT * FROM Film`

	// execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var film models.Film

		// unmarshal the row object to film
		err = rows.Scan(&film.Film_id, &film.Title, &film.Description, &film.Release_year, &film.Language_id, &film.Rental_duration, &film.Rental_rate, &film.Length, &film.Replacement_cost, &film.Rating, &film.Last_update, &film.Special_features, &film.Fulltext)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the film in the films slice
		films = append(films, film)

	}

	// return empty film on error
	return films, err
}

// get one film from the DB by its Title
func getFilmByTitle(title string) (models.Film, error) {
	// create the postgres db connection
	db := SetupDB()

	// close the db connection
	defer db.Close()

	// create a film of models.Film type
	var film models.Film

	// create the select sql query
	sqlStatement := `SELECT * FROM Film WHERE Title LIKE $1`

	// execute the sql statement
	row := db.QueryRow(sqlStatement, title)

	// unmarshal the row object to film
	err := row.Scan(&film.Film_id, &film.Title, &film.Description, &film.Release_year, &film.Language_id, &film.Rental_duration, &film.Rental_rate, &film.Length, &film.Replacement_cost, &film.Rating, &film.Last_update, &film.Special_features, &film.Fulltext)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return film, nil
	case nil:
		return film, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty film on error
	return film, err
}

// get films based off rating
func getFilmByRating(rating string) ([]models.Film, error) {
	// create the postgres db connection
	db := SetupDB()

	// close the db connection
	defer db.Close()

	var films []models.Film

	// create the select sql query
	sqlStatement := `SELECT * FROM Film WHERE Rating::text LIKE $1`

	// execute the sql statement
	rows, err := db.Query(sqlStatement, rating)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var film models.Film

		// unmarshal the row object to film
		err = rows.Scan(&film.Film_id, &film.Title, &film.Description, &film.Release_year, &film.Language_id, &film.Rental_duration, &film.Rental_rate, &film.Length, &film.Replacement_cost, &film.Rating, &film.Last_update, &film.Special_features, &film.Fulltext)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the film in the films slice
		films = append(films, film)

	}

	// return empty film on error
	return films, err
}

// get films by category
func getFilmByCategory(category string) ([]models.Film, error) {
	// create the postgres db connection
	db := SetupDB()

	// close the db connection
	defer db.Close()

	var films []models.Film

	// create the select sql query
	sqlStatement := `SELECT f.* FROM category c INNER JOIN film_category fc ON fc.category_id = c.category_id INNER JOIN film f ON fc.film_id = f.film_id WHERE c.name LIKE $1 AND fc.category_id=c.category_id AND f.film_id=fc.film_id`

	// execute the sql statement
	rows, err := db.Query(sqlStatement, category)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var film models.Film

		// unmarshal the row object to film
		err = rows.Scan(&film.Film_id, &film.Title, &film.Description, &film.Release_year, &film.Language_id, &film.Rental_duration, &film.Rental_rate, &film.Length, &film.Replacement_cost, &film.Rating, &film.Last_update, &film.Special_features, &film.Fulltext)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the film in the films slice
		films = append(films, film)

	}

	// return empty film on error
	return films, err
}
