package models

// User schema of the user table
type Review struct {
	Review_id   int64  `json:"review_id,string"`
	Review_text string `json:"Review_text"`
	Rating      string `json:"rating"`
	Customer_id int64  `json:"customer_id,string"`
	Film_id     int64  `json:"film_id,string"`
}
