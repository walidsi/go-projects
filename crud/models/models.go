package models

import "time"

type Blog struct {
	Id             int
	Title          string
	Description    string
	Date_published time.Time
}
