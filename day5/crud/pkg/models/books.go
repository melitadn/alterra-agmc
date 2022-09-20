package models

import "time"

type Book struct {
	Name      string `json:"name"`
	CreatedAt time.Time
}
