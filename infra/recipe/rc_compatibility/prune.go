package rc_compatibility

type PruneDefinition struct {
	Announcement        string   `json:"announcement,omitempty"`
	PruneAfterBuildDate string   `json:"prune_after_build_date,omitempty"`
	Current             PathPair `json:"current"`
}
