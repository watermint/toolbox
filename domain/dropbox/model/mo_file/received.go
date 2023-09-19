package mo_file

import "encoding/json"

type Received struct {
	Id          string          `path:"id" json:"id"`
	Name        string          `path:"name" json:"name"`
	PathLower   string          `path:"path_lower" json:"path_lower"`
	PathDisplay string          `path:"path_display" json:"path_display"`
	TimeInvited string          `path:"time_invited" json:"time_invited"`
	AccessType  string          `path:"access_type.\\.tag" json:"access_type"`
	Raw         json.RawMessage `json:"-"`
}
