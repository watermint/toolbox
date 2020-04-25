package sv_file_revision

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file_revision"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"go.uber.org/zap"
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
	l := z.ctx.Log().With(zap.String("path", path.Path()), zap.String("mode", mode))
	p := struct {
		Path  string `json:"path"`
		Mode  string `json:"mode"`
		Limit int    `json:"limit"`
	}{
		Path:  path.Path(),
		Mode:  mode,
		Limit: z.opts.Limit,
	}
	revs = &mo_file_revision.Revisions{}
	res, err := z.ctx.Post("files/list_revisions").Param(p).Call()
	if err != nil {
		return nil, err
	}
	j := res.Success().Json()
	entries, found := j.FindArray("entries")
	if !found {
		l.Debug("Response `entries` was not an array")
		return nil, ErrorUnexpectedResponseFormat
	}
	if x, found := j.FindBool("is_deleted"); found {
		revs.IsDeleted = x
	}
	if x, found := j.FindString("server_deleted"); found {
		revs.ServerDeleted = x
	}
	revs.Entries = make([]*mo_file.ConcreteEntry, 0)
	for _, e := range entries {
		ce := &mo_file.ConcreteEntry{}
		if _, err := e.Model(ce); err != nil {
			l.Debug("Unable to parse entry", zap.Error(err))
			return nil, err
		}
		revs.Entries = append(revs.Entries, ce)
	}
	return revs, nil
}

func (z *revisionImpl) List(path mo_path.DropboxPath) (revs *mo_file_revision.Revisions, err error) {
	return z.doList(path, "path")
}

func (z *revisionImpl) ListById(path mo_path.DropboxPath) (revs *mo_file_revision.Revisions, err error) {
	return z.doList(path, "id")
}
