package models

type Route struct {
	Name string
	Path string
}

type Navigation struct {
	Routes []Route
}

func GetMainNavigation() Navigation {
	return Navigation{
		Routes: []Route{
			{Name: "Variables", Path: "/variables"},
		},
	}
}
