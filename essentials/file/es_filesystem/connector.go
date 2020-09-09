package es_filesystem

type Connector interface {
	// Copy source to target system. This operation expects an file.
	// Returns error if an entry is a file.
	// Target path must include file name.
	Copy(source Entry, target Path) (err FileSystemError)
}
