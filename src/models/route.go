package models

type Route struct {
	Name string
	Path string
}

type Navigation struct {
	Routes []Route
}
