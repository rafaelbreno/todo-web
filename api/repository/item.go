package repository

import (
	"github.com/rafaelbreno/todo-web/api/models"
	"github.com/rafaelbreno/todo-web/api/storage"
)

// ItemRepo - Makes the connection between the Handler and Services.
// 	- Services can be Storage, etc.
type ItemRepo struct {
	st storage.Storage
}

// NewItemRepo returns an instance of Item Repository.
func NewItemRepo(st storage.Storage) Repository[models.Item] {
	return &ItemRepo{
		st: st,
	}
}

// Create - inserts in the storage a new Item.
func (r *ItemRepo) Create(l *models.Item) error {
	if err := r.st.Read(&models.List{
		ID: l.ListID,
	}); err != nil {
		return err
	}
	return r.st.Create(l)
}

// Delete - delete a list from the storage.
func (r *ItemRepo) Delete(l *models.Item) error {
	return r.st.Delete(l)
}

// Read - retrieve a list from the storage.
func (r *ItemRepo) Read(l *models.Item) error {
	return r.st.Read(l)
}

// ReadAll - retrieve all lists from the storage.
func (r *ItemRepo) ReadAll(l *[]models.Item) error {
	return r.st.ReadAll(l)
}

// Update - update a list from the storage.
func (r *ItemRepo) Update(l *models.Item) error {
	return r.st.Update(l)
}
