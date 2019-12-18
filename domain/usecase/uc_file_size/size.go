package uc_file_size

import (
	"github.com/watermint/toolbox/domain/model/mo_file_size"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_file"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/recpie/rc_kitchen"
	"go.uber.org/zap"
	"sync"
)

const (
	apiComplexityThreshold = 10_000
)

type Scale interface {
	Size(path mo_path.Path, depth int) (sizes map[mo_path.Path]mo_file_size.Size, errors map[mo_path.Path]error)
}

func New(ctx api_context.Context, k rc_kitchen.Kitchen) Scale {
	return &scaleImpl{
		ctx: ctx,
		k:   k,
	}
}

func newErrorDict() *errorDict {
	return &errorDict{
		lastError: make(map[mo_path.Path]error),
	}
}

type errorDict struct {
	lastError map[mo_path.Path]error
	mutex     sync.Mutex
}

func (z *errorDict) add(path mo_path.Path, err error) {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	z.lastError[path] = err
}

func newSizeDict() *sizeDict {
	return &sizeDict{
		sizes: make(map[mo_path.Path]mo_file_size.Size),
	}
}

type sizeDict struct {
	sizes map[mo_path.Path]mo_file_size.Size
	mutex sync.Mutex
}

func (z *sizeDict) add(path mo_path.Path, size mo_file_size.Size) {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	if s, ok := z.sizes[path]; ok {
		z.sizes[path] = s.Plus(path.Path(), size)
	} else {
		z.sizes[path] = size
	}
}

type scaleWorker struct {
	k        rc_kitchen.Kitchen
	ctx      api_context.Context
	svc      sv_file.Files
	keyPaths []mo_path.Path
	path     mo_path.Path
	curDepth int
	maxDepth int
	sd       *sizeDict
	ed       *errorDict
}

func (z *scaleWorker) Exec() error {
	ns, _ := z.path.Namespace()
	l := z.ctx.Log().With(zap.String("ns", ns),
		zap.String("path", z.path.Path()),
		zap.Int("curDepth", z.curDepth),
	)
	current := mo_file_size.Size{
		Path: z.path.Path(),
	}
	entries, err := z.svc.List(z.path)
	if err != nil {
		l.Debug("Unable to fetch list", zap.Error(err))
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

	q := z.k.NewQueue()
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
			kps := make([]mo_path.Path, 0)
			kps = append(kps, z.keyPaths...)
			if nd < z.maxDepth {
				kps = append(kps, np)
			}
			kpsDebug := make([]string, 0)
			for _, k := range kps {
				kpsDebug = append(kpsDebug, k.Path())
			}
			l.Debug("Process into child",
				zap.String("childPath", np.Path()),
				zap.Strings("keyPaths", kpsDebug),
				zap.Int("childDepth", nd),
			)
			q.Enqueue(&scaleWorker{
				k:        z.k,
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
	k   rc_kitchen.Kitchen
	ctx api_context.Context
}

func (z *scaleImpl) Size(path mo_path.Path, depth int) (sizes map[mo_path.Path]mo_file_size.Size, errors map[mo_path.Path]error) {
	sd := newSizeDict()
	ed := newErrorDict()
	svc := sv_file.NewFiles(z.ctx)

	q := z.k.NewQueue()
	q.Enqueue(&scaleWorker{
		k:        z.k,
		ctx:      z.ctx,
		svc:      svc,
		keyPaths: []mo_path.Path{path},
		path:     path,
		curDepth: 0,
		maxDepth: depth,
		sd:       sd,
		ed:       ed,
	})
	q.Wait()

	return sd.sizes, ed.lastError
}
