package es_size

import (
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/essentials/time/ut_compare"
	"github.com/watermint/toolbox/infra/security/sc_random"
	"go.uber.org/multierr"
	"time"
)

type FolderSize struct {
	Path            string     `json:"path"`
	Depth           int        `json:"depth"`
	Size            int64      `json:"size"`
	NumFile         int64      `json:"num_file"`
	NumFolder       int64      `json:"num_folder"`
	ModTimeEarliest *time.Time `json:"mod_time_earliest"`
	ModTimeLatest   *time.Time `json:"mod_time_latest"`
}

// Returns new instance of this instance plus given s.
// But keeps Path and Depth attributes.
func (z FolderSize) Add(s FolderSize) FolderSize {
	z.Size += s.Size
	z.NumFile += s.NumFile
	z.NumFolder += s.NumFolder
	z.ModTimeEarliest = ut_compare.EarliestPtr(z.ModTimeEarliest, s.ModTimeEarliest)
	z.ModTimeLatest = ut_compare.LatestPtr(z.ModTimeLatest, s.ModTimeLatest)
	return z
}

func Fold(path string, entries []es_filesystem.Entry) (size FolderSize) {
	size.Path = path
	modTimes := make([]time.Time, 0)
	for _, entry := range entries {
		switch {
		case entry.IsFile():
			size.NumFile++
			size.Size += entry.Size()
			modTimes = append(modTimes, entry.ModTime())

		case entry.IsFolder():
			size.NumFolder++
		}
	}
	if len(modTimes) > 0 {
		earliest := ut_compare.Earliest(modTimes...)
		latest := ut_compare.Latest(modTimes...)
		size.ModTimeEarliest = &earliest
		size.ModTimeLatest = &latest
	}
	return
}

type Traverse interface {
	Scan(path es_filesystem.Path, h func(s FolderSize)) error
}

func New(log esl.Logger, factory kv_storage.Factory, seq eq_sequence.Sequence, fs es_filesystem.FileSystem, depth int) (Traverse, error) {
	log.Debug("Create new traverse")
	folder, err := factory.New("folder_" + sc_random.MustGenerateRandomString(6))
	if err != nil {
		log.Debug("Unable to create new storage", esl.Error(err))
		return nil, err
	}
	sum, err := factory.New("sum_" + sc_random.MustGenerateRandomString(6))
	if err != nil {
		log.Debug("Unable to create new storage", esl.Error(err))
		return nil, err
	}

	return &traverseImpl{
		log:      log,
		folder:   folder,
		sum:      sum,
		sequence: seq,
		fs:       fs,
		depth:    depth,
	}, nil
}

const (
	queueIdScanFolder = "scan_folder"
)

type traverseImpl struct {
	log      esl.Logger
	sequence eq_sequence.Sequence
	fs       es_filesystem.FileSystem
	depth    int

	// folder structure (path -> descendants)
	folder kv_storage.Storage

	// folder sum (path -> *FolderSize)
	sum kv_storage.Storage
}

type TaskScanFolder struct {
	Path  es_filesystem.PathData `json:"path"`
	Depth int                    `json:"depth"`
}

// Descendant folder paths
type TaskScanFolderDescendants struct {
	Folders []string `json:"folders"`
}

func (z traverseImpl) scanFolder(task *TaskScanFolder, stg eq_sequence.Stage) error {
	l := z.log.With(esl.Any("task", task))
	l.Debug("Scan folder")

	path, fsErr := z.fs.Path(task.Path)
	if fsErr != nil {
		l.Debug("Unable to deserialize path", esl.Error(fsErr))
		return fsErr
	}

	entries, fsErr := z.fs.List(path)
	if fsErr != nil {
		l.Debug("Unable to list", esl.Error(fsErr))
		return fsErr
	}

	sum := Fold(task.Path.Path(), entries)
	sum.Depth = task.Depth
	kvErr := z.sum.Update(func(kvs kv_kvs.Kvs) error {
		return kvs.PutJsonModel(path.Path(), &sum)
	})
	if kvErr != nil {
		l.Debug("Unable to store result", esl.Error(kvErr))
		return kvErr
	}
	l.Debug("Folder summary", esl.Any("sum", sum))

	descendants := TaskScanFolderDescendants{
		Folders: make([]string, 0),
	}

	q := stg.Get(queueIdScanFolder).Batch(path.Shard().Id())
	for _, entry := range entries {
		if entry.IsFolder() {
			l.Debug("Enqueue descendant", esl.Any("entry", entry.AsData()))
			q.Enqueue(&TaskScanFolder{
				Path:  entry.Path().AsData(),
				Depth: task.Depth + 1,
			})
			descendants.Folders = append(descendants.Folders, entry.Path().Path())
		}
	}

	kvErr = z.folder.Update(func(kvs kv_kvs.Kvs) error {
		return kvs.PutJsonModel(path.Path(), &descendants)
	})
	if kvErr != nil {
		l.Debug("Unable to store result", esl.Error(kvErr))
		return kvErr
	}

	return nil
}

func (z traverseImpl) sumDescendants(sum *FolderSize) (total FolderSize, err error) {
	l := z.log.With(esl.Any("sum", sum))
	total.Path = sum.Path
	total.Depth = sum.Depth
	total = total.Add(*sum)

	l.Debug("Summarize descendants")
	descendants := &TaskScanFolderDescendants{}
	err = z.folder.View(func(kvd kv_kvs.Kvs) error {
		return kvd.GetJsonModel(sum.Path, descendants)
	})
	if err != nil {
		l.Debug("Unable to retrieve descendants", esl.Error(err))
		return total, err
	}

	for _, path := range descendants.Folders {
		descendant := &FolderSize{}
		err = z.sum.View(func(kvs kv_kvs.Kvs) error {
			return kvs.GetJsonModel(path, descendant)
		})
		if err != nil {
			l.Debug("Unable to retrieve descendant data", esl.String("path", path), esl.Error(err))
			return total, err
		}

		descendantTotal, err3 := z.sumDescendants(descendant)
		if err3 != nil {
			l.Debug("Unable to summarize descendants", esl.Error(err3))
			return total, err3
		}

		total = total.Add(descendantTotal)
		l.Debug("Sub total",
			esl.Any("total", total),
			esl.Any("descendantTotal", descendantTotal),
		)
	}

	l.Debug("summarize completed", esl.Any("total", total))
	return total, nil
}

func (z traverseImpl) Scan(path es_filesystem.Path, h func(s FolderSize)) error {
	l := z.log.With(esl.Any("path", path.AsData()))
	l.Debug("Start scanning")
	var lastErr error
	z.sequence.Do(func(s eq_sequence.Stage) {
		s.Define(queueIdScanFolder, z.scanFolder, s)
		q := s.Get(queueIdScanFolder)
		q.Enqueue(&TaskScanFolder{
			Path:  path.AsData(),
			Depth: 0,
		})
	}, eq_sequence.ErrorHandler(func(err error, mouldId, batchId string, p interface{}) {
		if err != nil {
			lastErr = err
		}
	}))

	l.Debug("Folder scan finished", esl.Error(lastErr))

	kvErr := z.sum.View(func(kvs kv_kvs.Kvs) error {
		return kvs.ForEachModel(&FolderSize{}, func(key string, m interface{}) error {
			sum := m.(*FolderSize)
			if sum.Depth <= z.depth {
				l.Debug("Reporting", esl.Any("sum", sum))
				total, sumErr := z.sumDescendants(sum)
				if sumErr != nil {
					l.Debug("Unable to summarize", esl.Error(sumErr))
					return sumErr
				}
				h(total)
			}
			return nil
		})
	})

	l.Debug("Scan finished", esl.Error(kvErr))
	if lastErr != nil || kvErr != nil {
		return multierr.Combine(lastErr, kvErr)
	}

	return nil
}
