package models

import (
	"fmt"
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

// SetDefault fill the fields with default values.
func (l *List) SetDefault() {
	l.Status = enum.ListNotStarted
}

// Validate the fields of a Model.
func (l *List) Validate() error {
	if l.Name == "" {
		return fmt.Errorf(errEmpty, "name")
	}
	return nil
}

// Update the fields of a Model.
func (l *List) Update(newList List) {
	if newList.Name != "" {
		l.Name = newList.Name
	}
	if newList.Status != 0 {
		l.Status = newList.Status
	}
}
