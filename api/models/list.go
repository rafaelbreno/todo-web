package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/rafaelbreno/todo-web/api/enum"
)

// List - A data model which is used for
// 	declaring lists and connecting it's items.
type List struct {
	ID        uuid.UUID       `json:"id"`
	Name      string          `json:"name"`
	Status    enum.ListStatus `json:"status"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt time.Time       `json:"deleted_at"`
}
