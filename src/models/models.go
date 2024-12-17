package models

type SimpleLayoutStructure struct {
	Title    string   `json:"title"`
	Body     string   `json:"body"`
	Examples []string `json:"examples"`
}
