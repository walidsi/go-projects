package models

import "time"

type Blog struct {
	Id             int       `json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Date_published time.Time `json:"date_published"`
}
