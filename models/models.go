package models

import "gorm.io/gorm"

type Todo struct {
	ID        int    `json:"id" binding:"required"`
	NAME      string `json:"name" binding:"required"`
	COMPLETED bool   `json:"completed"`
}

type TodoCompletedStatus struct {
	COMPLETED bool `json:"completed"`
}

type User struct {
	gorm.Model
	ID        int    `json:"id" gorm:"primary_key"`
	USER_NAME string `json:"user_name"`
	PASSWORD  string `json:"password"`
}
