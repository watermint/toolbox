package mo_legalhold

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
)

type Policy struct {
	Raw                     json.RawMessage
	Id                      string `path:"id" json:"id"`
	Name                    string `path:"name" json:"name"`
	Description             string `path:"description" json:"description"`
	Status                  string `path:"status.\\.tag" json:"status"`
	StartDate               string `path:"start_date" json:"start_date"`
	EndDate                 string `path:"end_date" json:"end_date"`
	ActivationTime          string `path:"activation_time" json:"activation_time"`
	PermanentlyDeletedUsers int    `path:"members.permanently_deleted_users" json:"permanently_deleted_users"`
}

func (z Policy) TeamMemberIds() (teamMemberIds []string) {
	teamMemberIds = make([]string, 0)
	j := es_json.MustParse(z.Raw)
	members, found := j.FindArray("members.team_member_ids")
	if !found {
		return
	}
	for _, member := range members {
		id, found := member.String()
		if found {
			teamMemberIds = append(teamMemberIds, id)
		}
	}
	return
}

type Revision struct {
	Raw                json.RawMessage
	NewFilename        string `path:"new_filename" json:"new_filename,omitempty"`
	OriginalRevisionId string `path:"original_revision_id" json:"original_revision_id,omitempty"`
	OriginalFilePath   string `path:"original_file_path" json:"original_file_path,omitempty"`
	ServerModified     string `path:"server_modified" json:"server_modified,omitempty"`
	AuthorMemberId     string `path:"author_member_id" json:"author_member_id,omitempty"`
	AuthorMemberStatus string `path:"author_member_status" json:"author_member_status,omitempty"`
	AuthorEmail        string `path:"author_email" json:"author_email,omitempty"`
	FileType           string `path:"file_type" json:"file_type,omitempty"`
	Size               uint64 `path:"size" json:"size,omitempty"`
	ContentHash        string `path:"content_hash" json:"content_hash,omitempty"`
}
