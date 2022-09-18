package sv_file_content

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/essentials/io/es_rewinder"
	"time"
)

type UploadStream interface {
	Upload(path mo_path.DropboxPath, f es_rewinder.ReadRewinder, clientModified time.Time) (entry mo_file.Entry, err error)
}

func NewUploadStream(client dbx_client.Client, autoRename, mute bool) UploadStream {
	return &streamImpl{
		client:     client,
		autoRename: autoRename,
		mute:       mute,
	}
}

type streamImpl struct {
	client     dbx_client.Client
	autoRename bool
	mute       bool
}

func (z streamImpl) Upload(path mo_path.DropboxPath, f es_rewinder.ReadRewinder, clientModified time.Time) (entry mo_file.Entry, err error) {
	mode := "overwrite"
	if z.autoRename {
		mode = "add"
	}
	p := struct {
		AutoRename     bool   `json:"auto_rename,omitempty"`
		ClientModified string `json:"client_modified,omitempty"`
		Mode           string `json:"mode,omitempty"`
		Mute           bool   `json:"mute,omitempty"`
		Path           string `json:"path,omitempty"`
	}{
		AutoRename:     z.autoRename,
		ClientModified: dbx_util.ToApiTimeString(clientModified),
		Mode:           mode,
		Mute:           z.mute,
		Path:           path.Path(),
	}

	res := z.client.Upload("files/upload",
		api_request.Content(f),
		api_request.Param(&p))
	if err, f := res.Failure(); f {
		return nil, err
	}

	entry = &mo_file.Metadata{}
	err = res.Success().Json().Model(entry)
	return
}
