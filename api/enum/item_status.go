package enum

// ItemStatus represents the status of an item.
type ItemStatus int8

const (
	// NotStarted is when the item still haven't been started yet.
	NotStarted ItemStatus = iota
	// InProgress is when the item still in InProgress
	InProgress
	// Completed is when the item was completed.
	Completed
	// Canceled is when the item was canceled.
	Canceled
)

var (
	// ItemStatusMap is a map to get the string value
	// of an ItemStatus.
	ItemStatusMap = map[ItemStatus]string{
		NotStarted: "not_started",
		InProgress: "in_progress",
		Completed:  "completed",
		Canceled:   "canceled",
	}
)
