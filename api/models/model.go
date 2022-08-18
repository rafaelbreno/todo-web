package models

// Model wraps up all methods related to models.
type Model interface {
	// SetDefault fill the fields with default values.
	SetDefault()
	// Validate the fields of a Model.
	Validate() error
	// Update will receive a model, and
	// apply those fields into the current
	// instance of a same type.
	Update(any)
}
