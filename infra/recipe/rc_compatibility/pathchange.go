package rc_compatibility

type PathChangeDefinition struct {
	Announcement        string     `json:"announcement,omitempty"`
	PruneAfterBuildDate string     `json:"prune_after_build_date,omitempty"`
	Current             PathPair   `json:"current"`
	FormerPaths         []PathPair `json:"former_paths"`
}
