package sv_file_revision

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file_revision"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/api/api_request"
)

var (
	ErrorUnexpectedResponseFormat = errors.New("unexpected response format")
)

type Revision interface {
	// Returns revisions with the same file path as identified by the latest file entry at the given file path or id.
	List(path mo_path.DropboxPath) (revs *mo_file_revision.Revisions, err error)

	// Returns revisions with the same file id as identified by the latest file entry at the given file path or id.
	ListById(path mo_path.DropboxPath) (revs *mo_file_revision.Revisions, err error)
}

type RevisionOpt func(o *RevisionOpts) *RevisionOpts
type RevisionOpts struct {
	Limit int
}

func New(ctx dbx_context.Context, opt ...RevisionOpt) Revision {
	opts := &RevisionOpts{Limit: 10}
	for _, o := range opt {
		o(opts)
	}
	return &revisionImpl{
		ctx:  ctx,
		opts: opts,
	}
}

type revisionImpl struct {
	ctx  dbx_context.Context
	opts *RevisionOpts
}

func (z *revisionImpl) doList(path mo_path.DropboxPath, mode string) (revs *mo_file_revision.Revisions, err error) {
	l := z.ctx.Log().With(es_log.String("path", path.Path()), es_log.String("mode", mode))
	p := struct {
		Path  string `json:"path"`
		Mode  string `json:"mode"`
		Limit int    `json:"limit"`
	}{
		Path:  path.Path(),
		Mode:  mode,
		Limit: z.opts.Limit,
	}
	res := z.ctx.Post("files/list_revisions", api_request.Param(p))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	j := res.Success().Json()
	revs = &mo_file_revision.Revisions{}
	if x, found := j.FindBool("is_deleted"); found {
		revs.IsDeleted = x
	}
	if x, found := j.FindString("server_deleted"); found {
		revs.ServerDeleted = x
	}
	revs.Entries = make([]*mo_file.ConcreteEntry, 0)
	err = j.FindArrayEach("entries", func(e es_json.Json) error {
		ce := &mo_file.ConcreteEntry{}
		if err := e.Model(ce); err != nil {
			l.Debug("Unable to parse entry", es_log.Error(err))
			return err
		}
		revs.Entries = append(revs.Entries, ce)
		return nil
	})
	return
}

func (z *revisionImpl) List(path mo_path.DropboxPath) (revs *mo_file_revision.Revisions, err error) {
	return z.doList(path, "path")
}

func (z *revisionImpl) ListById(path mo_path.DropboxPath) (revs *mo_file_revision.Revisions, err error) {
	return z.doList(path, "id")
}
