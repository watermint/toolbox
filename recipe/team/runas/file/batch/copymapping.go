package batch

type CopyMapping struct {
	MemberEmail string `json:"member_email"`
	SrcPath     string `json:"src_path"`
	DstPath     string `json:"dst_path"`
}
