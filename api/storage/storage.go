package storage

// Storage wraps all actions related to storage.
type Storage interface {
	Create(any) error
	Delete(any) error
	HealthCheck() error
	Migrate(...any) error
	Open() error
	Read(any) error
	ReadAll(any) error
	Update(any) error
}
