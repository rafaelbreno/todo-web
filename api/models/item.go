package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rafaelbreno/todo-web/api/enum"
)

// Item - A data model which represents
// 	each item of a list.
type Item struct {
	ID        uuid.UUID       `json:"id"`
	ListID    uuid.UUID       `json:"list_id"`
	Status    enum.ItemStatus `json:"status"`
	Text      string          `json:"text"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt time.Time       `json:"deleted_at"`
}

// SetDefault fill the fields with default values.
func (i *Item) SetDefault() {
	i.Status = enum.ItemNotStarted
}

// Validate the fields of a Model.
func (i *Item) Validate() error {
	if i.ListID.ID() == 0 {
		return fmt.Errorf(errEmpty, "id")
	}
	if i.Text == "" {
		return fmt.Errorf(errEmpty, "text")
	}
	return nil
}

// Update the fields of a Model.
func (i *Item) Update(newItem Item) {
	if newItem.Text != "" {
		i.Text = newItem.Text
	}
	if newItem.Status != 0 {
		i.Status = newItem.Status
	}
}
