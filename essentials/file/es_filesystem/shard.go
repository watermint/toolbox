package es_filesystem

// Shard is for load balancing factor.
type Shard interface {
	// string representation of Shard ID.
	// Drive letter, or server name for windows file system (e.g. C:, \\SERVER\).
	// Or namespace ID (shared_folder_id/team_folder_id) of Dropbox file system.
	Id() string

	// Serialize
	AsData() ShardData
}

// Serializable form of Shard
type ShardData struct {
	FileSystemType string                 `json:"file_system_type"`
	ShardId        string                 `json:"shard_id"`
	Attributes     map[string]interface{} `json:"attributes"`
}

func (z ShardData) AsData() ShardData {
	return z
}

func (z ShardData) Id() string {
	return z.ShardId
}
