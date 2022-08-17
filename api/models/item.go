package models

import (
	"github.com/google/uuid"
	"github.com/rafaelbreno/todo-web/api/enum"
)

// Item is a item from a List.
type Item struct {
	ID     uuid.UUID       `json:"id"`
	ListID uuid.UUID       `json:"list_id"`
	Status enum.ItemStatus `json:"status"`
	Text   string          `json:"text"`
}
