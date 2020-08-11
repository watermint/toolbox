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

	// Close this pipe.
	Close()
}

// Pipe factory
type Factory interface {
	New(batchId string) Pipe
}
