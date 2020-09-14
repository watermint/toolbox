package es_filesystem

type Namespace interface {
	// string representation of Namespace ID.
	// Drive letter, or server name for windows file system (e.g. C:, \\SERVER\).
	// Or namespace ID (shared_folder_id/team_folder_id) of Dropbox file system.
	Id() string

	// Serialize
	AsData() NamespaceData
}

// Serializable form of Namespace
type NamespaceData struct {
	FileSystemType string                 `json:"file_system_type"`
	NamespaceId    string                 `json:"namespace_id"`
	Attributes     map[string]interface{} `json:"attributes"`
}

func (z NamespaceData) AsData() NamespaceData {
	return z
}

func (z NamespaceData) Id() string {
	return z.NamespaceId
}
