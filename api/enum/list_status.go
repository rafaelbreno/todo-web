package enum

// ListStatus represents the status of an list.
type ListStatus uint8

const (
	// ListNotStarted is when the list still haven't been started yet.
	ListNotStarted ListStatus = iota + 1
	// ListInProgress is when the list still in InProgress
	ListInProgress
	// ListCompleted is when the list was completed.
	ListCompleted
	// ListCanceled is when the list was canceled.
	ListCanceled
)

var (
	// ListStatusMap is a map to get the string value
	// of an ListStatus.
	ListStatusMap = map[ListStatus]string{
		ListNotStarted: "not_started",
		ListInProgress: "in_progress",
		ListCompleted:  "completed",
		ListCanceled:   "canceled",
	}
)
