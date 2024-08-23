package mo_release_asset

import "encoding/json"

type Asset struct {
	Raw           json.RawMessage
	Id            string `path:"id" json:"id"`
	Name          string `path:"name" json:"name"`
	Size          int64  `path:"size" json:"size"`
	State         string `path:"state" json:"state"`
	CreatedAt     string `path:"created_at" json:"-"`
	UpdatedAt     string `path:"updated_at" json:"-"`
	DownloadCount int64  `path:"download_count" json:"download_count"`
	DownloadUrl   string `path:"browser_download_url" json:"download_url"`
}
