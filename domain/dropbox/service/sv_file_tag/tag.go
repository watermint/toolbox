package sv_file_tag

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
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

type AddDeleteTags struct {
	Path    string `json:"path"`
	TagText string `json:"tag_text"`
}

type tagImpl struct {
	client dbx_client.Client
}

func (z tagImpl) Add(path mo_path.DropboxPath, tag string) error {
	res := z.client.Post("files/tags/add", api_request.Param(&AddDeleteTags{
		Path:    path.Path(),
		TagText: tag,
	}))
	err, _ := res.Failure()
	return err
}

func (z tagImpl) Delete(path mo_path.DropboxPath, tag string) error {
	res := z.client.Post("files/tags/remove", api_request.Param(&AddDeleteTags{
		Path:    path.Path(),
		TagText: tag,
	}))
	err, _ := res.Failure()
	return err

}

func (z tagImpl) Resolve(path mo_path.DropboxPath) (tags []string, err error) {
	q := struct {
		Paths []string `json:"paths"`
	}{
		Paths: []string{path.Path()},
	}
	res := z.client.Post("files/tags/get", api_request.Param(&q))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	return parseTags(res.Success().Json())
}

func parseTags(json es_json.Json) (tags []string, err error) {
	tags = make([]string, 0)
	err = json.FindArrayEach("paths_to_tags.0.tags", func(e es_json.Json) error {
		t, found := e.FindString("tag_text")
		if !found {
			return nil
		}
		tags = append(tags, t)
		return nil
	})
	return tags, err
}
