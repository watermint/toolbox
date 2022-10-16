package dbx_fs_copier_batch

const (
	queueIdBlockCommit = "upload_commit"
	queueIdBlockUpload = "upload_block"
	queueIdBlockBatch  = "upload_batch"
	queueIdBlockCheck  = "upload_check"
)

type CommitInfo struct {
	Path           string `json:"path"`
	Mode           string `json:"mode"`
	Autorename     bool   `json:"autorename"`
	ClientModified string `json:"client_modified"`
	Mute           bool   `json:"mute"`
	StrictConflict bool   `json:"strict_conflict"`
}

type UploadCursor struct {
	SessionId string `json:"session_id" path:"session_id"`
	Offset    int64  `json:"offset" path:"offset"`
}
type SessionId struct {
	SessionId string `json:"session_id" path:"session_id"`
}
type UploadAppend struct {
	Cursor UploadCursor `json:"cursor"`
	Close  bool         `json:"close"`
}
type UploadFinish struct {
	Cursor UploadCursor `json:"cursor"`
	Commit CommitInfo   `json:"commit"`
}
type UploadFinishBatch struct {
	Entries []UploadFinish `json:"entries"`
}
type FinishBatch struct {
	Batch []string `json:"batch"`
}

type SessionCheck struct {
	SessionId string `json:"session_id"`
	Path      string `json:"path"`
}
