package routes

import "clean-go/src/models"

func GetMainNavigation() models.Navigation {
	return models.Navigation{
		Routes: []models.Route{
			{Name: "Variables", Path: "/variables"},
		},
	}
}
