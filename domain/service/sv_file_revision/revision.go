package sv_file_revision

import (
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_file_revision"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_parser"
	"go.uber.org/zap"
)

type Revision interface {
	// Returns revisions with the same file path as identified by the latest file entry at the given file path or id.
	List(path mo_path.Path) (revs *mo_file_revision.Revisions, err error)

	// Returns revisions with the same file id as identified by the latest file entry at the given file path or id.
	ListById(path mo_path.Path) (revs *mo_file_revision.Revisions, err error)
}

type RevisionOpt func(o *RevisionOpts) *RevisionOpts
type RevisionOpts struct {
	Limit int
}

func New(ctx api_context.Context, opt ...RevisionOpt) Revision {
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
	ctx  api_context.Context
	opts *RevisionOpts
}

func (z *revisionImpl) doList(path mo_path.Path, mode string) (revs *mo_file_revision.Revisions, err error) {
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
	res, err := z.ctx.Request("files/list_revisions").Param(p).Call()
	if err != nil {
		return nil, err
	}
	j, err := res.Json()
	if err != nil {
		l.Debug("Unable to get JSON response", zap.Error(err))
		return nil, err
	}
	entries := j.Get("entries")
	if !entries.IsArray() {
		l.Debug("Response `entries` was not an array")
		return nil, err
	}
	revs.IsDeleted = j.Get("is_deleted").Bool()
	revs.ServerDeleted = j.Get("server_deleted").String()
	revs.Entries = make([]*mo_file.ConcreteEntry, 0)
	for _, e := range entries.Array() {
		ce := &mo_file.ConcreteEntry{}
		if err := api_parser.ParseModel(ce, e); err != nil {
			l.Debug("Unable to parse entry", zap.Error(err))
			return nil, err
		}
		revs.Entries = append(revs.Entries, ce)
	}
	return revs, nil
}

func (z *revisionImpl) List(path mo_path.Path) (revs *mo_file_revision.Revisions, err error) {
	return z.doList(path, "path")
}

func (z *revisionImpl) ListById(path mo_path.Path) (revs *mo_file_revision.Revisions, err error) {
	return z.doList(path, "id")
}
