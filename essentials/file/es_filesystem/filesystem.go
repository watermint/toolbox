package es_filesystem

// Abstract file system interface.
type FileSystem interface {
	// List entries of the path.
	List(path Path) (entries []Entry, err FileSystemError)

	// Retrieve entry of the path.
	Info(path Path) (entry Entry, err FileSystemError)

	// Delete path.
	Delete(path Path) (err FileSystemError)

	// Create folder. Returns created entry on success, otherwise returns nil for entry.
	CreateFolder(path Path) (entry Entry, err FileSystemError)

	// Deserialize entry from entry data.
	// Returns err if the format is not valid for this file system.
	Entry(data EntryData) (entry Entry, err FileSystemError)

	// Deserialize path from path data.
	Path(data PathData) (path Path, err FileSystemError)

	// Deserialize shard from shard data.
	Shard(data ShardData) (shard Shard, err FileSystemError)

	// Type of file system
	FileSystemType() string

	// Operation complexity parameter in this file system
	OperationalComplexity(entries []Entry) (complexity int64)
}
