package storage

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/rafaelbreno/todo-web/api/models"
)

// LocalMap is the local representation of storage.
type LocalMap struct {
	lists map[string]models.List
	items map[string]models.Item
	Storage
}

// NewLocalMap returns a new instace of LocalMap.
func NewLocalMap() *LocalMap {
	return &LocalMap{
		lists: make(map[string]models.List),
		items: make(map[string]models.Item),
	}
}

// Create inserts a model instance to the local DB.
func (l *LocalMap) Create(v any) error {
	switch m := v.(type) {
	case *models.Item:
		m.ID = uuid.New()
		l.items[m.ID.String()] = *m
		return nil
	case *models.List:
		m.ID = uuid.New()
		l.lists[m.ID.String()] = *m
		return nil
	default:
		return fmt.Errorf("type '%T' not supported", v)
	}
}

// Delete deletes an existent model instance in the local DB.
func (l *LocalMap) Delete(v any) error {
	switch m := v.(type) {
	case *models.Item:
		if _, ok := l.items[m.ID.String()]; !ok {
			return fmt.Errorf("item with id '%s' does not exists", m.ID.String())
		}
		delete(l.items, m.ID.String())
		return nil
	case *models.List:
		if _, ok := l.lists[m.ID.String()]; !ok {
			return fmt.Errorf("list with id '%s' does not exists", m.ID.String())
		}
		delete(l.lists, m.ID.String())
		return nil
	default:
		return fmt.Errorf("type '%T' not supported", v)
	}
}

// HealthCheck just mocks a health check.
func (l *LocalMap) HealthCheck() error { return nil }

// Update updates an existent model instance in the local DB.
func (l *LocalMap) Read(v any) error {
	switch m := v.(type) {
	case *models.Item:
		if _, ok := l.items[m.ID.String()]; !ok {
			return fmt.Errorf("item with id '%s' does not exists", m.ID.String())
		}
		*m = l.items[m.ID.String()]
		return nil
	case *models.List:
		if _, ok := l.lists[m.ID.String()]; !ok {
			return fmt.Errorf("list with id '%s' does not exists", m.ID.String())
		}
		*m = l.lists[m.ID.String()]
		return nil
	default:
		return fmt.Errorf("type '%T' not supported", v)
	}
}

// ReadAll retrieves all values from the DB of a given Model.
func (l *LocalMap) ReadAll(v any) error {
	switch m := v.(type) {
	case *[]models.Item:
		if len(*m) == 0 {
			for _, v := range l.items {
				*m = append(*m, v)
			}
			return nil
		}
		listID := (*m)[0].ListID
		if _, ok := l.lists[listID.String()]; !ok {
			return fmt.Errorf("list with id '%s' does not exists", listID.String())
		}
		*m = []models.Item{}
		for _, v := range l.items {
			if v.ListID == listID {
				*m = append(*m, v)
			}
		}

		return nil
	case *[]models.List:
		for _, v := range l.lists {
			*m = append(*m, v)
		}

		return nil
	default:
		return fmt.Errorf("type '%T' not supported", v)
	}
}

// Update updates an existent model instance in the local DB.
func (l *LocalMap) Update(v any) error {
	switch m := v.(type) {
	case *models.Item:
		if _, ok := l.items[m.ID.String()]; !ok {
			return fmt.Errorf("item with id '%s' does not exists", m.ID.String())
		}
		l.items[m.ID.String()] = *m
		return nil
	case *models.List:
		if _, ok := l.lists[m.ID.String()]; !ok {
			return fmt.Errorf("list with id '%s' does not exists", m.ID.String())
		}
		l.lists[m.ID.String()] = *m
		return nil
	default:
		return fmt.Errorf("type '%T' not supported", v)
	}
}
