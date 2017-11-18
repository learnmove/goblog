package ViewModel

import (
	"time"
)

type BaseModel struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  ` json:"created_at"`
	UpdatedAt time.Time  ` json:"updated_at"`
	DeletedAt *time.Time `json:"delete_at"`
}
