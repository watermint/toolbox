package filesystem

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/security/sc_random"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"path/filepath"
	"strings"
	"sync"
)

type MsgFileSystemCached struct {
	ProgressPreScan app_msg.Message
}

var (
	MFileSystemCached = app_msg.Apply(&MsgFileSystemCached{}).(*MsgFileSystemCached)
)

func NewPreScanFileSystem(ctl app_control.Control, ctx dbx_context.Context, path mo_path.DropboxPath) (fs es_filesystem.FileSystem, err es_filesystem.FileSystemError) {
	l := ctl.Log().With(esl.String("Path", path.Path()))
	l.Debug("Prepare pre-scan file system")

	cacheName := sc_random.MustGenerateRandomString(6)
	cacheList, kvErr := ctl.NewKvs(cacheName + "list")
	if kvErr != nil {
		l.Debug("Unable to create the cache", esl.Error(kvErr))
		return nil, es_filesystem.NewLowLevelError(kvErr)
	}
	cacheEntry, kvErr := ctl.NewKvs(cacheName + "entry")
	if kvErr != nil {
		l.Debug("Unable to create the cache", esl.Error(kvErr))
		return nil, es_filesystem.NewLowLevelError(kvErr)
	}

	cfs := &cachedFs{
		l:            ctl.Log(),
		fs:           NewFileSystem(ctx),
		cacheInvalid: false,
		cacheMutex:   sync.Mutex{},
		cacheList:    cacheList,
		cacheEntries: cacheEntry,
		ctx:          ctx,
	}

	ctl.UI().Progress(MFileSystemCached.ProgressPreScan.With("Path", path.Path()))
	l.Debug("Start scanning")
	err = cfs.Scan(path)
	l.Debug("Scan finished", esl.Error(err))

	if err != nil {
		return nil, err
	} else {
		return cfs, nil
	}
}

type cachedFs struct {
	l            esl.Logger
	fs           es_filesystem.FileSystem
	cacheInvalid bool
	cacheMutex   sync.Mutex
	cacheList    kv_storage.Storage // path -> list of path (CacheEntryList)
	cacheEntries kv_storage.Storage // path -> entry (*mo_file.Metadata)
	ctx          dbx_context.Context
}

type CacheEntryList struct {
	Entries []string `json:"entries" path:"entries"`
}

func (z *cachedFs) updateCache(entry mo_file.Entry) {
	z.cacheMutex.Lock()
	defer z.cacheMutex.Unlock()

	l := z.l.With(esl.String("path", entry.PathDisplay()))

	// Add entry
	if f, ok := entry.File(); ok {
		l.Debug("Add a file entry")
		kvErr := z.cacheEntries.Update(func(kvs kv_kvs.Kvs) error {
			return kvs.PutJsonModel(f.EntryPathLower, f.Metadata())
		})
		if kvErr != nil {
			l.Debug("Unable to update cache")
			z.cacheInvalid = true
			return
		}
	} else if f, ok := entry.Folder(); ok {
		l.Debug("Add a folder entry")
		kvErr := z.cacheEntries.Update(func(kvs kv_kvs.Kvs) error {
			return kvs.PutJsonModel(f.EntryPathLower, f.Metadata())
		})
		if kvErr != nil {
			l.Debug("Unable to update cache")
			z.cacheInvalid = true
			return
		}
	} else {
		l.Debug("Ignore other entry types", esl.String("tag", entry.Tag()))
		return
	}

	// Add entry to the ancestor entry list
	{
		ancestor := filepath.Dir(entry.PathLower())
		entryList := &CacheEntryList{}
		_ = z.cacheList.View(func(kvs kv_kvs.Kvs) error {
			return kvs.GetJsonModel(ancestor, entryList)
		})
		entryList.Entries = append(entryList.Entries, entry.PathLower())
		kvErr := z.cacheList.Update(func(kvs kv_kvs.Kvs) error {
			return kvs.PutJsonModel(ancestor, entryList)
		})
		if kvErr != nil {
			l.Debug("Unable to update cache")
			z.cacheInvalid = true
			return
		}
	}
}

func (z *cachedFs) Scan(path mo_path.DropboxPath) (err es_filesystem.FileSystemError) {
	l := z.l.With(esl.String("path", path.Path()))
	l.Debug("Start scan")

	if err := sv_file.NewFiles(z.ctx).ListEach(path, z.updateCache, sv_file.Recursive(true)); err != nil {
		if dbx_error.NewErrors(err).Path().IsNotFound() {
			l.Debug("Path not found")
			return nil
		}
		l.Debug("Error on scanning files", esl.Error(err))
		return NewError(err)
	}
	return nil
}

func (z *cachedFs) List(path es_filesystem.Path) (entries []es_filesystem.Entry, err es_filesystem.FileSystemError) {
	l := z.l.With(esl.String("path", path.Path()))
	if z.cacheInvalid {
		l.Debug("Invalid cache, fallback to the orig fs")
		return z.fs.List(path)
	}

	entryList := &CacheEntryList{}
	kvErr := z.cacheList.View(func(kvs kv_kvs.Kvs) error {
		return kvs.GetJsonModel(strings.ToLower(path.Path()), entryList)
	})
	if kvErr != nil {
		l.Debug("Unable to retrieve info, fallback to not found", esl.Error(kvErr))
		return nil, NotFoundError()
	}

	entries = make([]es_filesystem.Entry, 0)
	for _, e := range entryList.Entries {
		entry, err := z.getEntry(e)
		if err != nil {
			l.Debug("Unable to find the entry, skip the entry", esl.Error(err))
			continue
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

func (z *cachedFs) getEntry(path string) (entry es_filesystem.Entry, err es_filesystem.FileSystemError) {
	l := z.l.With(esl.String("path", path))
	dbxEntry := &mo_file.Metadata{}
	kvErr := z.cacheEntries.View(func(kvs kv_kvs.Kvs) error {
		return kvs.GetJsonModel(strings.ToLower(path), dbxEntry)
	})
	if kvErr != nil {
		l.Debug("Unable to retrieve info, fallback to not found", esl.Error(kvErr))
		return nil, NotFoundError()
	}
	return NewEntry(dbxEntry), nil
}

func (z *cachedFs) Info(path es_filesystem.Path) (entry es_filesystem.Entry, err es_filesystem.FileSystemError) {
	l := z.l.With(esl.String("path", path.Path()))
	if z.cacheInvalid {
		l.Debug("Invalid cache, fallback to the orig fs")
		return z.fs.Info(path)
	}

	return z.getEntry(path.Path())
}

func (z *cachedFs) Delete(path es_filesystem.Path) (err es_filesystem.FileSystemError) {
	return z.fs.Delete(path)
}

func (z *cachedFs) CreateFolder(path es_filesystem.Path) (entry es_filesystem.Entry, err es_filesystem.FileSystemError) {
	return z.fs.CreateFolder(path)
}

func (z *cachedFs) Entry(data es_filesystem.EntryData) (entry es_filesystem.Entry, err es_filesystem.FileSystemError) {
	return z.fs.Entry(data)
}

func (z *cachedFs) Path(data es_filesystem.PathData) (path es_filesystem.Path, err es_filesystem.FileSystemError) {
	return z.fs.Path(data)
}

func (z *cachedFs) Shard(data es_filesystem.ShardData) (shard es_filesystem.Shard, err es_filesystem.FileSystemError) {
	return z.fs.Shard(data)
}

func (z *cachedFs) FileSystemType() string {
	return FileSystemTypeDropbox
}

func (z *cachedFs) OperationalComplexity(entries []es_filesystem.Entry) (complexity int64) {
	return z.fs.OperationalComplexity(entries)
}
