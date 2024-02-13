package utils

import "example.com/todoApi/data"

func IsIdUnique(id int) bool {
	for _, t := range data.Todos {
		if t.ID == id {
			return false
		}
	}
	return true
}

func TodoExits(id int) bool {
	for _, t := range data.Todos {
		if t.ID == id {
			return true
		}
	}
	return false
}
