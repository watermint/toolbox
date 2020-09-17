package es_filesystem

type Path interface {
	// Base name of the path entry.
	Base() string

	// Cleaned absolute path in the file system.
	Path() string

	// Shard of the path.
	Shard() Shard

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
	EntryShard     ShardData              `json:"entry_shard"`
	Attributes     map[string]interface{} `json:"attributes"`
}

func (z PathData) Path() string {
	return z.EntryPath
}

func (z PathData) Shard() Shard {
	return z.EntryShard
}

func (z PathData) AsData() PathData {
	return z
}
