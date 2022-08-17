package repository

import (
	"github.com/rafaelbreno/todo-web/api/models"
	"github.com/rafaelbreno/todo-web/api/storage"
)

// ListRepo - Makes the connection between the Handler and Services.
// 	- Services can be Storage, etc.
type ListRepo struct {
	st storage.Storage
}

// NewListRepo returns an instance of List Repository.
func NewListRepo(st storage.Storage) Repository[models.List] {
	return &ListRepo{
		st: st,
	}
}

// Create - inserts in the storage a new List.
func (r *ListRepo) Create(l *models.List) error {
	return r.st.Create(l)
}

// Delete - delete a list from the storage.
func (r *ListRepo) Delete(l *models.List) error {
	return r.st.Delete(l)
}

// Read - retrieve a list from the storage.
func (r *ListRepo) Read(l *models.List) error {
	return r.st.Read(l)
}

// ReadAll - retrieve all lists from the storage.
func (r *ListRepo) ReadAll(l *[]models.List) error {
	return r.st.ReadAll(l)
}

// Update - update a list from the storage.
func (r *ListRepo) Update(l *models.List) error {
	return r.st.Update(l)
}
