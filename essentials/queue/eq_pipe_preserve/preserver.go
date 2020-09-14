package eq_pipe_preserve

type Preserver interface {
	// Start preserve session
	Start() error

	// Add data to the preserver
	// Persistent store will be cleaned up when the operation failure.
	Add(d []byte) error

	// Commit data to the persistent store
	// Persistent store will be cleaned up when the operation failure.
	Commit(info []byte) (sessionId string, err error)
}

type Factory interface {
	NewPreserver() Preserver

	NewRestorer(sessionId string) Restorer
}
