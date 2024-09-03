package models

// Model
type Document struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}
