package data

import "example.com/todoApi/models"

var Todos = []models.Todo{
	{ID: 1, NAME: "Wash dirty clothes", COMPLETED: false},
	{ID: 2, NAME: "Watch season 5 of My Hero Acamedia", COMPLETED: false},
	{ID: 3, NAME: "Learn Backend Development with Go", COMPLETED: false},
	{ID: 4, NAME: "Buy Apple Vision Pro", COMPLETED: false},
	{ID: 5, NAME: "Watch season 2 of Jujutsu Kaisen", COMPLETED: true},
}
