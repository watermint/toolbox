package mo_file_diff

const (
	DiffFileContent        = "file_content_diff"
	DiffFileMissingLeft    = "left_file_missing"
	DiffFileMissingRight   = "right_file_missing"
	DiffFolderMissingLeft  = "left_folder_missing"
	DiffFolderMissingRight = "right_folder_missing"
	DiffSkipped            = "skipped"
)

type Diff struct {
	DiffType  string `json:"diff_type"`
	LeftPath  string `json:"left_path"`
	LeftKind  string `json:"left_kind"`
	LeftSize  *int64 `json:"left_size"`
	LeftHash  string `json:"left_hash"`
	RightPath string `json:"right_path"`
	RightKind string `json:"right_kind"`
	RightSize *int64 `json:"right_size"`
	RightHash string `json:"right_hash"`
}
