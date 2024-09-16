package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateTodoDto struct {
	Task string `json:"task"`
}

type DeleteByID struct {
	ID uuid.UUID `json:"id"`
}

type GetByID struct {
	ID uuid.UUID `json:"id"`
}

type Gets struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type UpdateReqTodo struct {
	Id          uuid.UUID `json:"id" gorm:"primaryKey"`
	Task        string    `json:"task"`
	IsCompleted bool      `json:"is_completed"`
	CompletedAt time.Time `json:"completed_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	// UserID      uuid.UUID `json:"user_id" gorm:"index"` // Foreign key reference
	// User        User      `json:"user" gorm:"foreignKey:UserID"`
}
