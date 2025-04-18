package es_open

type Desktop interface {
	// Open Launches the associated application to open a file or a URL
	// Returns nil on success, or an error if the operation failed
	Open(p string) error
}
