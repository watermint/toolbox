package es_filesystem

type Path interface {
	// Base name of the path entry.
	Base() string

	// Cleaned absolute path in the file system.
	Path() string

	// Namespace of the path.
	Namespace() Namespace

	// Ancestor path. Returns root path if this instance is root.
	Ancestor() Path

	// Returns descendant path with pathFragment.
	Descendant(pathFragment ...string) Path

	// Relative path.
	//Rel(path Path) (string, FileSystemError)

	// True if the path indicates root path of the namespace.
	IsRoot() bool

	// Serialize
	AsData() PathData
}

type PathData struct {
	FileSystemType string                 `json:"file_system_type"`
	EntryPath      string                 `json:"entry_path"`
	EntryNamespace NamespaceData          `json:"entry_namespace"`
	Attributes     map[string]interface{} `json:"attributes"`
}
