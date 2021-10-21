package es_size

import (
	"encoding/base32"
	"errors"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/lang"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/essentials/time/ut_compare"
	"go.uber.org/multierr"
	"sync"
	"time"
)

type FolderSize struct {
	Path                  string     `json:"path"`
	Depth                 int        `json:"depth"`
	Size                  int64      `json:"size"`
	NumFile               int64      `json:"num_file"`
	NumFolder             int64      `json:"num_folder"`
	ModTimeEarliest       *time.Time `json:"mod_time_earliest"`
	ModTimeLatest         *time.Time `json:"mod_time_latest"`
	OperationalComplexity int64      `json:"operational_complexity"`
}

var (
	ErrorSessionNotFound = errors.New("session not found")
)

// Returns new instance of this instance plus given s.
// But keeps Path and Depth attributes.
func (z FolderSize) Add(s FolderSize) FolderSize {
	z.Size += s.Size
	z.NumFile += s.NumFile
	z.NumFolder += s.NumFolder
	z.ModTimeEarliest = ut_compare.EarliestPtr(z.ModTimeEarliest, s.ModTimeEarliest)
	z.ModTimeLatest = ut_compare.LatestPtr(z.ModTimeLatest, s.ModTimeLatest)
	z.OperationalComplexity += s.OperationalComplexity
	return z
}

func Fold(path string, fs es_filesystem.FileSystem, entries []es_filesystem.Entry) (size FolderSize) {
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
	size.OperationalComplexity = fs.OperationalComplexity(entries)
	return
}

func New(log esl.Logger, queueIdScanFolder string, factory kv_storage.Factory) Context {
	return &ctxImpl{
		log:               log,
		queueIdScanFolder: queueIdScanFolder,
		sessions:          make(map[string]Session),
		sessionsMutex:     sync.Mutex{},
		sessionPath:       make(map[string]es_filesystem.Path),
		factory:           factory,
	}
}

type Context interface {
	QueueIdScanFolder() string

	New(sessionId string, path es_filesystem.Path, stg eq_sequence.Stage, fs es_filesystem.FileSystem, meta interface{}) Session
	StartSession(sessionId string, stg eq_sequence.Stage) error

	Get(sessionId string) (Session, error)

	Log() esl.Logger

	ListEach(depth int, h func(sessionId string, meta interface{}, size FolderSize)) error
}

type ctxImpl struct {
	log               esl.Logger
	queueIdScanFolder string
	sessions          map[string]Session
	sessionsMutex     sync.Mutex
	sessionPath       map[string]es_filesystem.Path
	factory           kv_storage.Factory
}

func (z *ctxImpl) ListEach(depth int, h func(sessionId string, meta interface{}, size FolderSize)) error {
	z.sessionsMutex.Lock()
	defer z.sessionsMutex.Unlock()

	var lastErr error
	for _, session := range z.sessions {
		err := session.ListEach(depth, func(size FolderSize) {
			h(session.SessionId(), session.Metadata(), size)
		})
		if err != nil {
			lastErr = err
		}
	}
	return lastErr
}

func (z *ctxImpl) QueueIdScanFolder() string {
	return z.queueIdScanFolder
}

func (z *ctxImpl) StartSession(sessionId string, stg eq_sequence.Stage) error {
	l := z.Log().With(esl.String("sessionId", sessionId))

	path, ok := z.sessionPath[sessionId]
	if !ok {
		l.Debug("Session path not found")
		return ErrorSessionNotFound
	}

	if session, ok := z.sessions[sessionId]; !ok {
		l.Debug("Session not found")
		return ErrorSessionNotFound
	} else {
		if err := session.Open(); err != nil {
			l.Debug("Unable to open the session", esl.Error(err))
			return err
		}

		session.Enqueue(path, 0)
		return nil
	}
}

func (z *ctxImpl) New(sessionId string, path es_filesystem.Path, stg eq_sequence.Stage, fs es_filesystem.FileSystem, meta interface{}) Session {
	z.sessionsMutex.Lock()
	defer z.sessionsMutex.Unlock()
	l := z.Log()

	if session, ok := z.sessions[sessionId]; ok {
		l.Debug("Session already exists", esl.String("sessionId", sessionId))
		return session
	}
	session := newSession(z, sessionId, stg, fs, z.factory, meta)

	z.sessionPath[sessionId] = path
	z.sessions[sessionId] = session

	return session
}

func (z *ctxImpl) Get(sessionId string) (Session, error) {
	z.sessionsMutex.Lock()
	defer z.sessionsMutex.Unlock()

	if session, ok := z.sessions[sessionId]; ok {
		return session, nil
	}

	return nil, ErrorSessionNotFound
}

func (z *ctxImpl) Log() esl.Logger {
	return z.log
}

type Session interface {
	// Session Id of this scan
	SessionId() string

	// Logger
	Log() esl.Logger

	// Queue stage
	Stage() eq_sequence.Stage

	// Target file system
	FileSystem() es_filesystem.FileSystem

	// Storage for folder tree (path -> descendant paths)
	Folder() kv_storage.Storage

	// Storage for sum (path -> *FolderSize)
	Sum() kv_storage.Storage

	// List each results. This function must call after stage finish.
	// depth == 0 for the root folder.
	ListEach(depth int, h func(size FolderSize)) error

	// Session metadata
	Metadata() interface{}

	// Open session
	Open() error

	// Enqueue
	Enqueue(path es_filesystem.Path, depth int)
}

func newSession(ctx Context, sessionId string, stg eq_sequence.Stage, fs es_filesystem.FileSystem, factory kv_storage.Factory, meta interface{}) Session {
	return &sessionImpl{
		ctx:       ctx,
		sessionId: sessionId,
		stg:       stg,
		fs:        fs,
		factory:   factory,
		metadata:  meta,
	}
}

type sessionImpl struct {
	sessionId string
	stg       eq_sequence.Stage
	fs        es_filesystem.FileSystem
	metadata  interface{}

	factory kv_storage.Factory

	// folder structure (path -> descendants)
	folder kv_storage.Storage

	// folder sum (path -> *FolderSize)
	sum kv_storage.Storage

	ctx Context
}

func (z *sessionImpl) Metadata() interface{} {
	return z.metadata
}

func (z *sessionImpl) Open() (err error) {
	l := z.Log()

	key := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString([]byte(z.SessionId()))
	l.Debug("Prepare traverse resources", esl.String("sessionId", z.SessionId()), esl.String("sessionKey", key))
	z.folder, err = z.factory.New("folder_" + key)
	if err != nil {
		l.Debug("Unable to create new storage", esl.Error(err))
		return err
	}
	z.sum, err = z.factory.New("sum_" + key)
	if err != nil {
		l.Debug("Unable to create new storage", esl.Error(err))
		return err
	}

	return nil
}

func (z *sessionImpl) Enqueue(path es_filesystem.Path, depth int) {
	q := z.Stage().Get(z.ctx.QueueIdScanFolder()).Batch(z.SessionId() + path.Shard().Id())
	q.Enqueue(&TaskScanFolder{
		SessionId: z.SessionId(),
		Path:      path.AsData(),
		Depth:     depth,
	})
}

func (z *sessionImpl) SessionId() string {
	return z.sessionId
}

func (z *sessionImpl) ListEach(depth int, h func(size FolderSize)) error {
	l := z.Log()
	var lastErr error
	viewErr := z.sum.View(func(kvs kv_kvs.Kvs) error {
		return kvs.ForEachModel(&FolderSize{}, func(key string, m interface{}) error {
			sum := m.(*FolderSize)
			if sum.Depth <= depth {
				l.Debug("Reporting", esl.Any("sum", sum))
				total, sumErr := sumDescendants(sum, z)
				if sumErr != nil {
					l.Debug("Unable to summarize", esl.Error(sumErr))
					lastErr = sumErr
					return nil
				}
				h(total)
			}
			return nil
		})
	})
	if lastErr != nil || viewErr != nil {
		return multierr.Combine(lastErr, viewErr)
	}
	return nil
}

func (z *sessionImpl) Stage() eq_sequence.Stage {
	return z.stg
}

func (z *sessionImpl) Log() esl.Logger {
	return z.ctx.Log().With(esl.String("sessionId", z.SessionId()))
}

func (z *sessionImpl) FileSystem() es_filesystem.FileSystem {
	return z.fs
}

func (z *sessionImpl) Folder() kv_storage.Storage {
	return z.folder
}

func (z *sessionImpl) Sum() kv_storage.Storage {
	return z.sum
}

type TaskScanFolder struct {
	SessionId string                 `json:"session_id"`
	Path      es_filesystem.PathData `json:"path"`
	Depth     int                    `json:"depth"`
}

// Descendant folder paths
type TaskScanFolderDescendants struct {
	Folders []string `json:"folders"`
}

func ScanFolder(task *TaskScanFolder, ctx Context) error {
	l := ctx.Log().With(esl.Any("task", task))
	l.Debug("Scan folder")

	session, err := ctx.Get(task.SessionId)
	if err != nil {
		l.Debug("Unable to find the session", esl.Error(err))
		return err
	}

	path, fsErr := session.FileSystem().Path(task.Path)
	if fsErr != nil {
		l.Debug("Unable to deserialize path", esl.Error(fsErr))
		return fsErr
	}

	entries, fsErr := session.FileSystem().List(path)
	if fsErr != nil {
		l.Debug("Unable to list", esl.Error(fsErr))
		return fsErr
	}

	sum := Fold(task.Path.Path(), session.FileSystem(), entries)
	sum.Depth = task.Depth
	kvErr := session.Sum().Update(func(kvs kv_kvs.Kvs) error {
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

	for _, entry := range entries {
		if entry.IsFolder() {
			l.Debug("Enqueue descendant", esl.Any("entry", entry.AsData()))
			session.Enqueue(entry.Path(), task.Depth+1)
			descendants.Folders = append(descendants.Folders, entry.Path().Path())
		}
	}

	kvErr = session.Folder().Update(func(kvs kv_kvs.Kvs) error {
		return kvs.PutJsonModel(path.Path(), &descendants)
	})
	if kvErr != nil {
		l.Debug("Unable to store result", esl.Error(kvErr))
		return kvErr
	}

	return nil
}

func sumDescendants(sum *FolderSize, session Session) (total FolderSize, err error) {
	l := session.Log().With(esl.Any("sum", sum))
	total.Path = sum.Path
	total.Depth = sum.Depth
	total = total.Add(*sum)

	l.Debug("Summarize descendants")
	descendants := &TaskScanFolderDescendants{}
	err = session.Folder().View(func(kvd kv_kvs.Kvs) error {
		return kvd.GetJsonModel(sum.Path, descendants)
	})
	if err != nil {
		l.Debug("Unable to retrieve descendants", esl.Error(err))
		return total, err
	}

	for _, path := range descendants.Folders {
		descendant := &FolderSize{}
		err = session.Sum().View(func(kvs kv_kvs.Kvs) error {
			return kvs.GetJsonModel(path, descendant)
		})
		if err != nil {
			l.Debug("Unable to retrieve descendant data", esl.String("path", path), esl.Error(err))
			continue
		}

		descendantTotal, err3 := sumDescendants(descendant, session)
		if err3 != nil {
			l.Debug("Unable to summarize descendants", esl.Error(err3))
			continue
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

func ScanSingleFileSystem(
	log esl.Logger,
	seq eq_sequence.Sequence,
	factory kv_storage.Factory,
	fs es_filesystem.FileSystem,
	path es_filesystem.Path,
	depth int,
	h func(s FolderSize),
) error {
	ctx := New(log, "scan_folder", factory)
	sessionId := "single"
	var lastErr error
	seq.Do(func(s eq_sequence.Stage) {
		ctx.New(sessionId, path, s, fs, nil)
		s.Define("scan_folder", ScanFolder, ctx)
		s.Define("scan_session", ctx.StartSession, s)
		s.Get("scan_session").Enqueue(sessionId)
	}, eq_sequence.ErrorHandler(func(err error, mouldId, batchId string, p interface{}) {
		if err != nil {
			lastErr = err
		}
	}))

	listErr := ctx.ListEach(depth, func(sessionId string, meta interface{}, size FolderSize) {
		h(size)
	})

	if lastErr != nil || listErr != nil {
		return lang.NewMultiErrorOrNull(lastErr, listErr)
	}
	return nil
}
