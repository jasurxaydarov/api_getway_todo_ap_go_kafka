package dto

import "github.com/google/uuid"

type CreateTodoDto struct {
	Tsak string `json:"task"`
}

type UpdatedTodoDto struct {
	Id          uuid.UUID `json:"id"`
	Tsak        string    `json:"task"`
	IsCompleted bool      `json:"is_completed"`
}
