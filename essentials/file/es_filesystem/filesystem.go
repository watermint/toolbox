package es_filesystem

// Abstract file system interface.
type FileSystem interface {
	// List entries of the path.
	List(path Path) (entries []Entry, err FileSystemError)

	// Retrieve entry of the path.
	Info(path Path) (entry Entry, err FileSystemError)

	// Delete path.
	Delete(path Path) (err FileSystemError)

	// Create folder.
	CreateFolder(path Path) (err FileSystemError)

	// Deserialize entry from entry data.
	// Returns err if the format is not valid for this file system.
	Entry(data EntryData) (entry Entry, err FileSystemError)

	// Deserialize path from path data.
	Path(data PathData) (path Path, err FileSystemError)

	// Deserialize namespace from namespace data.
	Namespace(data NamespaceData) (namespace Namespace, err FileSystemError)

	// Type of file system
	FileSystemType() string
}
