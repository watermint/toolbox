package es_filesystem

// FileSystem is the abstract file system interface.
type FileSystem interface {
	// List entries of the path.
	List(path Path) (entries []Entry, err FileSystemError)

	// Info retrieves entry of the path.
	Info(path Path) (entry Entry, err FileSystemError)

	// Delete path.
	Delete(path Path) (err FileSystemError)

	// CreateFolder returns created entry on success, otherwise returns nil for entry.
	CreateFolder(path Path) (entry Entry, err FileSystemError)

	// Entry deserializes entry from entry data.
	// Returns err if the format is not valid for this file system.
	Entry(data EntryData) (entry Entry, err FileSystemError)

	// Path deserializes path from path data.
	Path(data PathData) (path Path, err FileSystemError)

	// Shard deserializes shard from shard data.
	Shard(data ShardData) (shard Shard, err FileSystemError)

	// FileSystemType returns type of file system
	FileSystemType() string

	// OperationalComplexity returns operation complexity parameter in this file system
	OperationalComplexity(entries []Entry) (complexity int64)
}
