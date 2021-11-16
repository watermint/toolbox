package batch

type UpMapping struct {
	MemberEmail string `json:"member_email"`
	LocalPath   string `json:"local_path"`
	DropboxPath string `json:"dropbox_path"`
}
