package mo_filerequest

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/infra/api/api_parser"
	"github.com/watermint/toolbox/infra/control/app_root"
	"go.uber.org/zap"
)

type FileRequest struct {
	Raw                      json.RawMessage
	Id                       string `path:"id" json:"id"`
	Url                      string `path:"url" json:"url"`
	Title                    string `path:"title" json:"title"`
	Created                  string `path:"created" json:"created"`
	IsOpen                   bool   `path:"is_open" json:"is_open"`
	FileCount                int    `path:"file_count" json:"file_count"`
	Destination              string `path:"destination" json:"destination"`
	Deadline                 string `path:"deadline.deadline" json:"deadline"`
	DeadlineAllowLateUploads string `path:"deadline.allow_late_uploads.\\.tag" json:"deadline_allow_late_uploads"`
}

func (z *FileRequest) IsSame(other *FileRequest) bool {
	return z.Title == other.Title &&
		z.Destination == other.Destination &&
		z.IsOpen == other.IsOpen &&
		z.Deadline == other.Deadline &&
		z.DeadlineAllowLateUploads == other.DeadlineAllowLateUploads
}

type MemberFileRequest struct {
	Raw                      json.RawMessage
	AccountId                string `path:"member.profile.account_id" json:"account_id"`
	TeamMemberId             string `path:"member.profile.team_member_id" json:"team_member_id"`
	Email                    string `path:"member.profile.email" json:"email"`
	Status                   string `path:"member.profile.status.\\.tag" json:"status"`
	Surname                  string `path:"member.profile.name.surname" json:"surname"`
	GivenName                string `path:"member.profile.name.given_name" json:"given_name"`
	FileRequestId            string `path:"file_request.id" json:"file_request_id"`
	Url                      string `path:"file_request.url" json:"url"`
	Title                    string `path:"file_request.title" json:"title"`
	Created                  string `path:"file_request.created" json:"created"`
	IsOpen                   bool   `path:"file_request.is_open" json:"is_open"`
	FileCount                int    `path:"file_request.file_count" json:"file_count"`
	Destination              string `path:"file_request.destination" json:"destination"`
	Deadline                 string `path:"file_request.deadline.deadline" json:"deadline"`
	DeadlineAllowLateUploads string `path:"file_request.deadline.allow_late_uploads.\\.tag" json:"deadline_allow_late_uploads"`
}

func NewMemberFileRequest(fr *FileRequest, member *mo_member.Member) *MemberFileRequest {
	raws := make(map[string]json.RawMessage)
	raws["file_request"] = fr.Raw
	raws["member"] = member.Raw
	raw := api_parser.CombineRaw(raws)

	mfr := &MemberFileRequest{}
	if err := api_parser.ParseModelRaw(mfr, raw); err != nil {
		app_root.Log().Warn("unexpected data format", zap.Error(err))
		// return empty
		return mfr
	}
	return mfr
}
