package storage

// Storage wraps all actions related to storage.
type Storage interface {
	Create(any) (any, error)
	Delete(string, any) error
	HealthCheck() error
	Migrate(...any) error
	Open() error
	Read(string, any) (any, error)
	ReadAll(any) (any, error)
	Update(string, any) (any, error)
}
