package mo_file_diff

const (
	DiffFileContent        = "file_content_diff"
	DiffFileMissingLeft    = "left_file_missing"
	DiffFileMissingRight   = "right_file_missing"
	DiffFolderMissingLeft  = "left_folder_missing"
	DiffFolderMissingRight = "right_folder_missing"
)

type Diff struct {
	DiffType  string `json:"diff_type"`
	LeftPath  string `json:"left_path,omitempty"`
	LeftKind  string `json:"left_kind,omitempty"`
	LeftSize  *int64 `json:"left_size,omitempty"`
	LeftHash  string `json:"left_hash,omitempty"`
	RightPath string `json:"right_path,omitempty"`
	RightKind string `json:"right_kind,omitempty"`
	RightSize *int64 `json:"right_size,omitempty"`
	RightHash string `json:"right_hash,omitempty"`
}
