package models

type Todo struct {
	ID        int    `json:"id" binding:"required"`
	NAME      string `json:"name" binding:"required"`
	COMPLETED bool   `json:"completed"`
}

type TodoCompletedStatus struct {
	COMPLETED bool `json:"completed"`
}
