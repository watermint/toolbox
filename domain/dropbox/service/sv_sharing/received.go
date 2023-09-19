package sv_sharing

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_list"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
)

type Received interface {
	List() (entries []*mo_file.Received, err error)
}

func NewReceived(ctx dbx_client.Client) Received {
	return &receivedImpl{
		client: ctx,
	}
}

type receivedImpl struct {
	client dbx_client.Client
}

func (z receivedImpl) List() (entries []*mo_file.Received, err error) {
	entries = make([]*mo_file.Received, 0)
	res := z.client.List("sharing/list_received_files").Call(
		dbx_list.Continue("sharing/list_received_files/continue"),
		dbx_list.ResultTag("entries"),
		dbx_list.OnEntry(func(entry es_json.Json) error {
			r := &mo_file.Received{}
			if err := entry.Model(r); err != nil {
				return err
			}
			entries = append(entries, r)
			return nil
		}),
	)
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	return
}
