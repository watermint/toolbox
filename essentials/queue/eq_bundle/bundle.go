package eq_bundle

import (
	"github.com/watermint/toolbox/essentials/queue/eq_pipe"
)

// Bundle interface will not return error.
// If there any error happens, the impl. should recover from the error themself or
// raise panic() to tell critical issue to the caller.
type Bundle interface {
	// Enqueue data with batchId
	Enqueue(b Barrel)

	// Fetch data from the queue.
	Fetch() (b Barrel, found bool)

	// Mark data as completed
	Complete(b Barrel)

	// Queue storage size per batchId
	Size() (sizes map[string]int, total int)

	// Size of the InProgress queue
	SizeInProgress() int

	// Close this bundle.
	Close()

	// Preserve this bundle.
	Preserve() (session Session, err error)
}

type OnCompleteHandler func(batchId string, completed, total int)

type Session struct {
	Pipes      map[string]eq_pipe.SessionId `json:"pipes"`
	InProgress eq_pipe.SessionId            `json:"in_progress"`
}
