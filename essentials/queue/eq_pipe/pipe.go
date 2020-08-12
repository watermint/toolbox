package eq_pipe

type Pipe interface {
	// Enqueue the data
	Enqueue(d []byte)

	// Dequeue the data from the pipe. Returns nil if the pipe is empty.
	Dequeue() (d []byte)

	// Delete the data.
	Delete(d []byte)

	// Size of the pipe.
	Size() int

	// Close & clean up this pipe.
	Close()

	// Close & preserve the state
	Preserve() (id SessionId, err error)
}

// Pipe factory
type Factory interface {
	// Create a new pipe
	New(batchId string) Pipe

	// Restore a pipe from the session
	Restore(id SessionId) (pipe Pipe, err error)
}

type SessionId string
