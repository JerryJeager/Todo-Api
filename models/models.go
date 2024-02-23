package models

import "gorm.io/gorm"

type Todo struct {
	ID        int    `json:"id" binding:"required"`
	NAME      string `json:"name" binding:"required"`
	COMPLETED bool   `json:"completed"`
	UserID    int
}

type TodoCompletedStatus struct {
	COMPLETED bool `json:"completed"`
}

type User struct {
	gorm.Model
	ID        int    `json:"id" gorm:"primary_key"`
	USER_NAME string `json:"user_name" gorm:"typevarchar(100);unique_index"`
	PASSWORD  string `json:"password"`
}
