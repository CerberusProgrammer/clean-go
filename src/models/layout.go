package models

type Layout struct {
	Title      string   `json:"title"`
	Body       string   `json:"body"`
	Examples   []string `json:"examples"`
	Navigation Navigation
}
