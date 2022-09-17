package sv_file_tag

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
)

type Tag interface {
	Add(path mo_path.DropboxPath, tag string) error
	Delete(path mo_path.DropboxPath, tag string) error
	Resolve(path mo_path.DropboxPath) (tags []string, err error)
}

func New(client dbx_client.Client) Tag {
	return &tagImpl{
		client: client,
	}
}

type tagImpl struct {
	client dbx_client.Client
}

func (z tagImpl) Add(path mo_path.DropboxPath, tag string) error {
	//TODO implement me
	panic("implement me")
}

func (z tagImpl) Delete(path mo_path.DropboxPath, tag string) error {
	//TODO implement me
	panic("implement me")
}

func (z tagImpl) Resolve(path mo_path.DropboxPath) (tags []string, err error) {
	//TODO implement me
	panic("implement me")
}
