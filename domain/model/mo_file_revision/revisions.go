package mo_file_revision

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/model/mo_file"
)

type Revisions struct {
	Raw           json.RawMessage
	IsDeleted     bool                     `json:"is_deleted"`
	ServerDeleted string                   `json:"server_deleted"`
	Entries       []*mo_file.ConcreteEntry `json:"entries"`
}
