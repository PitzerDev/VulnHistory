package models

// User schema of the user table
type Film struct {
	Film_id          int64  `json:"film_id"`
	Title            string `json:"title"`
	Description      string `json:"description"`
	Release_year     int64  `json:"release_year"`
	Language_id      int64  `json:"lanuage_id"`
	Rental_duration  int64  `json:"rental_duration"`
	Rental_rate      string `json:"rental_rate"`
	Length           int64  `json:"length"`
	Replacement_cost string `json:"replacement_cost"`
	Rating           string `json:"rating"`
	Last_update      string `json:"last_update"`
	Special_features string `json:"special_features"`
	Fulltext         string `json:"fulltext"`
}
