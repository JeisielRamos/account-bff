package entities

import "time"

type Transfer struct {
	ID                     int       `json:"id"`
	Account_origin_id      int       `json:"account_origin_id"`
	Account_destination_id int       `json:"account_destination_id"`
	Amount                 float64   `json:"amount"`
	Created_at             time.Time `json:"created_at"`
}
