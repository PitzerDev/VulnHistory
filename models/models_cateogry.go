package models

// User schema of the user table
type Category struct {
	Film_id     int64  `json:"film_id"`
	Category_id int64  `json:"customer_id"`
	Last_update string `json:"last_update"`
}
