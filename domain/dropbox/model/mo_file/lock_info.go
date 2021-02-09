package mo_file

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/log/esl"
)

type LockInfo struct {
	Raw                  json.RawMessage
	Id                   string `path:"id" json:"id"`
	Tag                  string `path:"\\.tag" json:"tag"`
	Name                 string `path:"name" json:"name"`
	PathLower            string `path:"path_lower" json:"path_lower"`
	PathDisplay          string `path:"path_display" json:"path_display"`
	ClientModified       string `path:"client_modified" json:"client_modified"`
	ServerModified       string `path:"server_modified" json:"server_modified"`
	Revision             string `path:"rev" json:"revision"`
	Size                 int64  `path:"size" json:"size"`
	ContentHash          string `path:"content_hash" json:"content_hash"`
	SharedFolderId       string `path:"sharing_info.shared_folder_id" json:"shared_folder_id"`
	ParentSharedFolderId string `path:"sharing_info.parent_shared_folder_id" json:"parent_shared_folder_id"`
	IsLockHolder         bool   `path:"file_lock_info.is_lockholder" json:"is_lock_holder"`
	LockHolderName       string `path:"file_lock_info.lockholder_name" json:"lock_holder_name"`
	LockHolderAccountId  string `path:"file_lock_info.lockholder_account_id" json:"lock_holder_account_id"`
	LockCreated          string `path:"file_lock_info.created" json:"lock_created"`
}

func newLockInfo(raw json.RawMessage) *LockInfo {
	l := esl.Default()
	e := &LockInfo{}
	j, err := es_json.Parse(raw)
	if err != nil {
		l.Debug("Unable to parse", esl.Error(err))
		return nil
	}
	_, found := j.Find("file_lock_info")
	if !found {
		l.Debug("File lock info does not found")
		return nil
	}
	if err = j.Model(e); err != nil {
		l.Debug("Model parse error", esl.Error(err))
		return nil
	} else {
		e.Raw = raw
		return e
	}
}
