package enum

// ItemStatus represents the status of an item.
type ItemStatus int8

const (
	// ItemNotStarted is when the item still haven't been started yet.
	ItemNotStarted ItemStatus = iota
	// ItemInProgress is when the item still in InProgress
	ItemInProgress
	// ItemCompleted is when the item was completed.
	ItemCompleted
	// ItemCanceled is when the item was canceled.
	ItemCanceled
)

var (
	// ItemStatusMap is a map to get the string value
	// of an ItemStatus.
	ItemStatusMap = map[ItemStatus]string{
		ItemNotStarted: "not_started",
		ItemInProgress: "in_progress",
		ItemCompleted:  "completed",
		ItemCanceled:   "canceled",
	}
)
