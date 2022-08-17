package repository

// Repository wraps up all methods related to
// repository actions.
type Repository[T any] interface {
	Create(*T) error
	Update(*T) error
	Read(*T) error
	ReadAll(*[]T) error
	Delete(*T) error
}
