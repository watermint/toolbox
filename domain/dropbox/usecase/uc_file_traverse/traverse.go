package uc_file_traverse

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_namespace"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
)

type TraverseHandler func(te TraverseEntry, entries []mo_file.Entry)
type TraverseErrorHandler func(te TraverseEntry, err error)

type TraverseEntry struct {
	Namespace *mo_namespace.Namespace `json:"namespace"`
	Path      string                  `json:"path"`
}

func NewTraverse(ctx dbx_client.Client, ctl app_control.Control, queueId string, handler TraverseHandler, errHandler TraverseErrorHandler, opts ...sv_file.ListOpt) *Traverse {
	return &Traverse{
		ctx:          ctx,
		ctl:          ctl,
		listOpts:     opts,
		queueId:      queueId,
		handler:      handler,
		handlerError: errHandler,
	}
}

type Traverse struct {
	ctx          dbx_client.Client
	ctl          app_control.Control
	listOpts     []sv_file.ListOpt
	queueId      string
	handler      TraverseHandler
	handlerError TraverseErrorHandler
}

func (z Traverse) Traverse(te TraverseEntry, stage eq_sequence.Stage) error {
	l := z.ctx.Log().With(esl.String("namespace", te.Namespace.NamespaceId), esl.String("path", te.Path))
	var ctn dbx_client.Client
	if te.Namespace.NamespaceId != "" {
		ctn = z.ctx.WithPath(dbx_client.Namespace(te.Namespace.NamespaceId))
	} else {
		ctn = z.ctx
	}

	l.Debug("Retrieve path")
	opts := append(z.listOpts, sv_file.Recursive(false))
	entries, err := sv_file.NewFiles(ctn).List(mo_path.NewDropboxPath(te.Path), opts...)
	if err != nil {
		l.Debug("Unable to scan path", esl.Error(err))
		z.handlerError(te, err)
		return err
	}
	z.handler(te, entries)

	q := stage.Get(z.queueId).Batch(te.Namespace.NamespaceId)
	for _, entry := range entries {
		ll := l.With(esl.String("descendantPath", entry.PathDisplay()))
		ll.Debug("Process descendant")
		if _, ok := entry.Folder(); ok {
			l.Debug("Recurse")
			q.Enqueue(TraverseEntry{
				Namespace: te.Namespace,
				Path:      entry.PathDisplay(),
			})
		}
	}
	return nil
}
