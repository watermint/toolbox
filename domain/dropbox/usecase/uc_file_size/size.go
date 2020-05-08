package uc_file_size

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file_size"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_control"
	"sync"
)

const (
	apiComplexityThreshold = 10_000
)

type Scale interface {
	Size(path mo_path.DropboxPath, depth int) (sizes map[mo_path.DropboxPath]mo_file_size.Size, errors map[mo_path.DropboxPath]error)
}

func New(ctx dbx_context.Context, ctl app_control.Control) Scale {
	return &scaleImpl{
		ctx: ctx,
		ctl: ctl,
	}
}

func newErrorDict() *errorDict {
	return &errorDict{
		lastError: make(map[mo_path.DropboxPath]error),
	}
}

type errorDict struct {
	lastError map[mo_path.DropboxPath]error
	mutex     sync.Mutex
}

func (z *errorDict) add(path mo_path.DropboxPath, err error) {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	z.lastError[path] = err
}

func newSizeDict() *sizeDict {
	return &sizeDict{
		sizes: make(map[mo_path.DropboxPath]mo_file_size.Size),
	}
}

type sizeDict struct {
	sizes map[mo_path.DropboxPath]mo_file_size.Size
	mutex sync.Mutex
}

func (z *sizeDict) add(path mo_path.DropboxPath, size mo_file_size.Size) {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	if s, ok := z.sizes[path]; ok {
		z.sizes[path] = s.Plus(path.Path(), size)
	} else {
		z.sizes[path] = size
	}
}

type scaleWorker struct {
	ctl      app_control.Control
	ctx      api_context.Context
	svc      sv_file.Files
	keyPaths []mo_path.DropboxPath
	path     mo_path.DropboxPath
	curDepth int
	maxDepth int
	sd       *sizeDict
	ed       *errorDict
}

func (z *scaleWorker) Exec() error {
	ns, _ := z.path.Namespace()
	l := z.ctx.Log().With(esl.String("ns", ns),
		esl.String("path", z.path.Path()),
		esl.Int("curDepth", z.curDepth),
	)
	current := mo_file_size.Size{
		Path: z.path.Path(),
	}
	entries, err := z.svc.List(z.path)
	if err != nil {
		l.Debug("Unable to fetch list", esl.Error(err))
		for _, kp := range z.keyPaths {
			z.ed.add(kp, err)
		}
		return err
	}
	numEntries := len(entries)

	if numEntries >= apiComplexityThreshold {
		current.ApiComplexity = int64(numEntries)
	} else {
		current.ApiComplexity = 1
	}

	q := z.ctl.NewQueue()
	for _, entry := range entries {
		current.CountDescendant++
		if f, e := entry.File(); e {
			current.CountFile++
			current.Size += f.Size
		}
		if f, e := entry.Folder(); e {
			current.CountFolder++
			nd := z.curDepth + 1
			np := z.path.ChildPath(f.Name())
			kps := make([]mo_path.DropboxPath, 0)
			kps = append(kps, z.keyPaths...)
			if nd < z.maxDepth {
				kps = append(kps, np)
			}
			kpsDebug := make([]string, 0)
			for _, k := range kps {
				kpsDebug = append(kpsDebug, k.Path())
			}
			l.Debug("Process into child",
				esl.String("childPath", np.Path()),
				esl.Strings("keyPaths", kpsDebug),
				esl.Int("childDepth", nd),
			)
			q.Enqueue(&scaleWorker{
				ctl:      z.ctl,
				ctx:      z.ctx,
				svc:      z.svc,
				keyPaths: kps,
				path:     np,
				curDepth: nd,
				maxDepth: z.maxDepth,
				sd:       z.sd,
				ed:       z.ed,
			})
		}
	}
	q.Wait()
	for _, kp := range z.keyPaths {
		z.sd.add(kp, current)
	}

	return nil
}

type scaleImpl struct {
	ctl app_control.Control
	ctx dbx_context.Context
}

func (z *scaleImpl) Size(path mo_path.DropboxPath, depth int) (sizes map[mo_path.DropboxPath]mo_file_size.Size, errors map[mo_path.DropboxPath]error) {
	sd := newSizeDict()
	ed := newErrorDict()
	svc := sv_file.NewFiles(z.ctx)

	q := z.ctl.NewQueue()
	q.Enqueue(&scaleWorker{
		ctl:      z.ctl,
		ctx:      z.ctx,
		svc:      svc,
		keyPaths: []mo_path.DropboxPath{path},
		path:     path,
		curDepth: 0,
		maxDepth: depth,
		sd:       sd,
		ed:       ed,
	})
	q.Wait()

	return sd.sizes, ed.lastError
}
