package models

// User schema of the user table
type Customer struct {
	Customer_id int64  `json:"customer_id"`
	Store_id    string `json:"store_id"`
	First_name  string `json:"first_name"`
	Last_name   string `json:"last_name"`
	Email       string `json:"email"`
	Address_id  int64  `json:"address_id"`
	Activebool  string `json:"activebool"`
	Create_date string `json:"create_date"`
	Last_update string `json:"last_update"`
	Active      string `json:"active"`
}
