package es_sync

import (
	"github.com/vbauerster/mpb/v5"
	"github.com/vbauerster/mpb/v5/decor"
	"github.com/watermint/toolbox/essentials/ambient/ea_indicator"
	"github.com/watermint/toolbox/essentials/file/es_filecompare"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_filter"
	"github.com/watermint/toolbox/essentials/queue/eq_queue"
	"sync"
)

func New(log esl.Logger, qd eq_queue.Definition, source, target es_filesystem.FileSystem, conn es_filesystem.Connector, opt ...Opt) Syncer {
	opts := Opts{
		entryNameFilter: mo_filter.New(""),
		progress:        ea_indicator.Global(),
	}.Apply(opt)
	cmp := es_filecompare.New(
		es_filecompare.DontCompareContent(opts.syncDontCompareContent),
		es_filecompare.DontCompareTime(opts.syncDontCompareTime),
	)

	indicator := opts.Progress().NewIndicator(0,
		mpb.PrependDecorators(
			decor.Name("Completed ", decor.WC{W: 20}),
			decor.AverageSpeed(decor.UnitKiB, "% 1.f"),
		),
		mpb.AppendDecorators(
			decor.CountersKibiByte(" % .1f / % .1f"),
			decor.OnComplete(
				decor.Spinner(mpb.DefaultSpinnerStyle, decor.WC{W: 5}), "DONE",
			),
		),
	)

	return &syncImpl{
		log:       log,
		qd:        qd,
		source:    source,
		target:    target,
		fileCmp:   cmp,
		conn:      conn,
		opts:      opts,
		indicator: indicator,
	}
}

const (
	queueIdSyncFolder          = "sync_folder"
	queueIdCopyFile            = "sync_file"
	queueIdDelete              = "delete"
	queueIdReplaceFolderByFile = "replace_folder_by_file"
	queueIdReplaceFileByFolder = "replace_file_by_folder"
)

type syncImpl struct {
	log       esl.Logger
	qd        eq_queue.Definition
	source    es_filesystem.FileSystem
	target    es_filesystem.FileSystem
	conn      es_filesystem.Connector
	fileCmp   es_filecompare.FileComparator
	indicator ea_indicator.Indicator
	opts      Opts
	wg        sync.WaitGroup
}

func (z *syncImpl) enqueueTask(queueId string, source, target es_filesystem.Path, data interface{}) {
	z.wg.Add(1)
	q := z.qd.Current().MustGet(queueId).Batch(z.computeBatchId(source, target))
	q.Enqueue(data)
}

func (z *syncImpl) computeBatchId(source, target es_filesystem.Path) string {
	return source.Shard().Id() + "/" + target.Shard().Id()
}

func (z *syncImpl) copy(source es_filesystem.Entry, target es_filesystem.Path) error {
	l := z.log.With(esl.Any("source", source.AsData()), esl.Any("target", target.AsData()))
	l.Debug("Copy")

	if !z.opts.entryNameFilter.Accept(source.Name()) {
		l.Debug("Filter applied, skip")
		z.opts.OnSkip(SkipFilter, source, target)
		return nil
	}

	z.indicator.AddTotal(source.Size())

	z.conn.Copy(
		source,
		target,
		func(pair es_filesystem.CopyPair, copied es_filesystem.Entry) {
			z.opts.OnCopySuccess(source, copied)
			z.indicator.AddProgress(source.Size())
		}, func(pair es_filesystem.CopyPair, err es_filesystem.FileSystemError) {
			l.Debug("Unable to copy data from source to target", esl.Error(err))
			z.opts.OnCopyFailure(source.Path(), err)
		},
	)

	return nil
}

func (z *syncImpl) delete(target es_filesystem.Path) error {
	l := z.log.With(esl.Any("target", target.AsData()))
	l.Debug("Delete")
	err := z.target.Delete(target)
	if err != nil {
		l.Debug("unable to delete", esl.Error(err))
		z.opts.OnDeleteFailure(target, err)
		return err
	}
	z.opts.OnDeleteSuccess(target)
	return nil
}

func (z *syncImpl) createFolder(target es_filesystem.Path) error {
	l := z.log.With(esl.Any("target", target.AsData()))
	l.Debug("Create folder")
	_, err := z.target.CreateFolder(target)
	if err != nil {
		l.Debug("unable to create folder", esl.Error(err))
		z.opts.OnCreateFolderFailure(target, err)
		return err
	}
	z.opts.OnCreateFolderSuccess(target)
	return nil
}

type TaskCopyFile struct {
	Source es_filesystem.EntryData `json:"source"`
	Target es_filesystem.PathData  `json:"target"`
}

// copy or overwrite the file. This task overwrites the file if exists in target fs.
func (z *syncImpl) taskCopyFile(task *TaskCopyFile, qd eq_queue.Definition) error {
	l := z.log.With(esl.Any("task", task))
	defer z.wg.Done()

	sourceEntry, err := z.source.Entry(task.Source)
	if err != nil {
		l.Debug("Unable to create an entry data", esl.Error(err))
		return err
	}
	targetPath, err := z.target.Path(task.Target)
	if err != nil {
		l.Debug("Unable to create a path data", esl.Error(err))
		z.opts.OnCopyFailure(sourceEntry.Path(), err)
		return err
	}
	targetEntry, err := z.target.Info(targetPath)
	switch {
	case err == nil:
		l.Debug("Successfully retrieved target file info, compare those")
		same, cmpErr := z.fileCmp.Compare(sourceEntry, targetEntry)
		if cmpErr != nil {
			l.Debug("Unable to compare files", esl.Error(err))
			z.opts.OnCopyFailure(sourceEntry.Path(), err)
			return cmpErr
		}

		if same {
			l.Debug("Same skip")
			z.opts.OnSkip(SkipSame, sourceEntry, targetPath)
			return nil
		}
		if sourceEntry.ModTime().Before(targetEntry.ModTime()) && z.opts.SyncOverwrite() {
			l.Debug("Same old file")
			z.opts.OnSkip(SkipOld, sourceEntry, targetPath)
			return nil
		}

	case err.IsPathNotFound():
		l.Debug("File not found in the target path: Skip comparison", esl.Error(err))

	default:
		l.Debug("Unable to retrieve target file info", esl.Error(err))
		z.opts.OnCopyFailure(sourceEntry.Path(), err)
		return err
	}

	l.Debug("Copy file")
	return z.copy(sourceEntry, targetPath)
}

type TaskReplaceFolderByFile struct {
	Source es_filesystem.EntryData `json:"source"`
	Target es_filesystem.PathData  `json:"target"`
}

func (z *syncImpl) taskReplaceFolderByFile(task *TaskReplaceFolderByFile, qd eq_queue.Definition) error {
	l := z.log.With(esl.Any("task", task))
	defer z.wg.Done()

	sourceEntry, err := z.source.Entry(task.Source)
	if err != nil {
		l.Debug("Unable to create an entry data", esl.Error(err))
		return err
	}
	targetPath, err := z.target.Path(task.Target)
	if err != nil {
		l.Debug("Unable to create a path data", esl.Error(err))
		return err
	}

	l.Debug("Remove the target folder")
	if err := z.delete(targetPath); err != nil {
		return err
	}

	z.enqueueTask(queueIdCopyFile, sourceEntry.Path(), targetPath, &TaskCopyFile{
		Source: sourceEntry.AsData(),
		Target: targetPath.AsData(),
	})
	return nil
}

type TaskReplaceFileByFolder struct {
	Source es_filesystem.EntryData `json:"source"`
	Target es_filesystem.PathData  `json:"target"`
}

func (z *syncImpl) taskReplaceFileByFolder(task *TaskReplaceFileByFolder, qd eq_queue.Definition) error {
	l := z.log.With(esl.Any("task", task))
	defer z.wg.Done()

	sourceEntry, err := z.source.Entry(task.Source)
	if err != nil {
		l.Debug("Unable to create an entry data", esl.Error(err))
		return err
	}
	targetPath, err := z.target.Path(task.Target)
	if err != nil {
		l.Debug("Unable to create a path data", esl.Error(err))
		return err
	}

	l.Debug("Remove the target file")
	if err := z.delete(targetPath); err != nil {
		return err
	}

	l.Debug("Create the folder")
	if err := z.createFolder(targetPath); err != nil {
		return err
	}

	l.Debug("enqueue copy")
	z.enqueueTask(queueIdSyncFolder, sourceEntry.Path(), targetPath, &TaskSyncFolder{
		Source: sourceEntry.Path().AsData(),
		Target: targetPath.AsData(),
	})
	return nil
}

type TaskSyncFolder struct {
	Source es_filesystem.PathData `json:"source"`
	Target es_filesystem.PathData `json:"target"`
}

func (z *syncImpl) taskSyncFolder(task *TaskSyncFolder, qd eq_queue.Definition) error {
	l := z.log.With(esl.Any("source", task.Source), esl.Any("target", task.Target))
	defer z.wg.Done()

	sourcePath, err := z.source.Path(task.Source)
	if err != nil {
		l.Debug("Unable to create an path data", esl.Error(err))
		return err
	}
	targetPath, err := z.target.Path(task.Target)
	if err != nil {
		l.Debug("Unable to create a path data", esl.Error(err))
		return err
	}

	handlerMissingSource := func(base es_filecompare.PathPair, target es_filesystem.Entry) {
		ll := l.With(esl.Any("target", target.AsData()))
		if !z.opts.SyncDelete() {
			ll.Debug("skip delete")
			return
		}

		ll.Debug("sync delete")
		z.enqueueTask(queueIdDelete, base.Source, base.Target, &TaskDelete{
			Target: target.Path().AsData(),
		})
	}
	handlerMissingTarget := func(base es_filecompare.PathPair, source es_filesystem.Entry) {
		ll := l.With(esl.Any("source", source.AsData()))
		newTargetPath := base.Target.Descendant(source.Path().Base())

		if source.IsFolder() {
			if z.opts.OptimizeReduceCreateFolder() {
				ll.Debug("Skip create folder", esl.Any("targetFolderPath", newTargetPath.AsData()))
			} else {
				ll.Debug("create folder", esl.Any("targetFolderPath", newTargetPath.AsData()))
				if err := z.createFolder(newTargetPath); err != nil {
					ll.Debug("unable to create folder", esl.Error(err))
					return
				}
			}

			ll.Debug("Enqueue sync folder")
			z.enqueueTask(queueIdSyncFolder, source.Path(), newTargetPath, &TaskSyncFolder{
				Source: source.Path().AsData(),
				Target: newTargetPath.AsData(),
			})
		} else {
			ll.Debug("Copy file")
			z.enqueueTask(queueIdCopyFile, source.Path(), newTargetPath, &TaskCopyFile{
				Source: source.AsData(),
				Target: newTargetPath.AsData(),
			})
		}
	}
	handlerFileDiff := func(base es_filecompare.PathPair, source, target es_filesystem.Entry) {
		ll := l.With(esl.Any("source", source.AsData()), esl.Any("target", target.AsData()))
		if !z.opts.SyncOverwrite() {
			ll.Debug("Don't overwrite")
			z.opts.OnSkip(SkipExists, source, target.Path())
			return
		}

		ll.Debug("Overwrite")
		z.enqueueTask(queueIdCopyFile, source.Path(), target.Path(), &TaskCopyFile{
			Source: source.AsData(),
			Target: target.Path().AsData(),
		})
	}
	handlerSameFile := func(base es_filecompare.PathPair, source, target es_filesystem.Entry) {
		ll := l.With(esl.Any("source", source.AsData()), esl.Any("target", target.AsData()))
		ll.Debug("Same file, skip sync")
		z.opts.OnSkip(SkipSame, source, target.Path())
	}
	handlerTypeDiff := func(base es_filecompare.PathPair, source, target es_filesystem.Entry) {
		ll := l.With(esl.Any("source", source.AsData()), esl.Any("target", target.AsData()))
		if !z.opts.SyncOverwrite() {
			ll.Debug("Don't overwrite")
			z.opts.OnSkip(SkipExists, source, target.Path())
			return
		}

		switch {
		case source.IsFile() && target.IsFolder():
			ll.Debug("Replace folder with file")
			z.enqueueTask(queueIdReplaceFolderByFile, source.Path(), target.Path(), &TaskReplaceFolderByFile{
				Source: source.AsData(),
				Target: target.Path().AsData(),
			})

		case source.IsFolder() && target.IsFile():
			ll.Debug("Replace file with folder")
			z.enqueueTask(queueIdReplaceFileByFolder, source.Path(), target.Path(), &TaskReplaceFileByFolder{
				Source: source.AsData(),
				Target: target.Path().AsData(),
			})

		case source.IsFile() && target.IsFile():
			ll.Debug("Same type, compare both")
			same, cmpErr := z.fileCmp.Compare(source, target)
			if cmpErr != nil {
				ll.Debug("Unable to compare files", esl.Error(err))
				z.opts.OnCopyFailure(source.Path(), err)
				return
			}

			if same {
				ll.Debug("Same skip")
				z.opts.OnSkip(SkipSame, source, target.Path())
				return
			}
			if source.ModTime().Before(target.ModTime()) && z.opts.SyncOverwrite() {
				l.Debug("Same old file")
				z.opts.OnSkip(SkipOld, source, targetPath)
				return
			}

			ll.Debug("Copy")
			z.enqueueTask(queueIdCopyFile, source.Path(), target.Path(), &TaskCopyFile{
				Source: source.AsData(),
				Target: target.Path().AsData(),
			})

		case source.IsFolder() && target.IsFolder():
			ll.Debug("Same type, but do as regular sync folder")
			z.enqueueTask(queueIdSyncFolder, source.Path(), target.Path(), &TaskSyncFolder{
				Source: source.Path().AsData(),
				Target: target.Path().AsData(),
			})
		}
	}
	handlerDescendant := func(base es_filecompare.PathPair, source, target es_filesystem.Entry) {
		l.Debug("Sync descendant", esl.Any("source", source.AsData()), esl.Any("target", target.AsData()))
		z.enqueueTask(queueIdSyncFolder, source.Path(), target.Path(), &TaskSyncFolder{
			Source: source.Path().AsData(),
			Target: target.Path().AsData(),
		})
	}

	errCmp := es_filecompare.CompareLevel(
		l,
		z.source,
		z.target,
		sourcePath,
		targetPath,
		z.fileCmp,
		handlerMissingSource,
		handlerMissingTarget,
		handlerFileDiff,
		handlerTypeDiff,
		handlerSameFile,
		handlerDescendant,
	)
	if errCmp != nil {
		l.Debug("comparison failed", esl.Error(errCmp))
	}
	return errCmp
}

type TaskDelete struct {
	Target es_filesystem.PathData `json:"target"`
}

func (z *syncImpl) taskDelete(task *TaskDelete, qd eq_queue.Definition) error {
	l := z.log.With(esl.Any("target", task.Target))
	l.Debug("Delete")
	defer z.wg.Done()

	targetPath, err := z.target.Path(task.Target)
	if err != nil {
		l.Debug("Unable to deserialize", esl.Error(err))
		return err
	}

	return z.delete(targetPath)
}

func (z *syncImpl) Sync(source es_filesystem.Path, target es_filesystem.Path) error {
	l := z.log.With(esl.String("sourcePath", source.Path()), esl.String("targetPath", target.Path()))
	l.Debug("Start sync")

	var lastError error
	lastErrorMutex := sync.Mutex{}
	onError := func(err error, mouldId, batchId string, p interface{}) {
		lastErrorMutex.Lock()
		defer lastErrorMutex.Unlock()
		switch e := err.(type) {
		case es_filesystem.FileSystemError:
			if e.IsMockError() {
				return
			}
		}
		lastError = err
	}
	z.qd.AddErrorListener(onError)

	srcEntry, err := z.source.Info(source)
	if err != nil {
		l.Debug("Unable to retrieve source entry", esl.Error(err))
		return err
	}
	l.Debug("Source entry found", esl.Any("sourceEntry", srcEntry.AsData()))

	tgtEntry, err := z.target.Info(target)
	switch {
	case err == nil:
		l.Debug("Target folder found", esl.Any("targetEntry", tgtEntry.AsData()))

	case err.IsPathNotFound():
		var errCreateFolder es_filesystem.FileSystemError
		tgtEntry, errCreateFolder = z.target.CreateFolder(target)
		if errCreateFolder != nil {
			l.Debug("Unable to create folder", esl.Error(errCreateFolder))
			z.opts.OnCreateFolderFailure(target, err)
			return errCreateFolder
		}
		z.opts.OnCreateFolderSuccess(target)

	case err.IsMockError():
		l.Debug("Ignore mock error")
		return nil

	default:
		l.Debug("Unable to retrieve target entry", esl.Error(err))
		return err
	}

	z.qd.Define(queueIdCopyFile, z.taskCopyFile, z.qd)
	z.qd.Define(queueIdSyncFolder, z.taskSyncFolder, z.qd)
	z.qd.Define(queueIdDelete, z.taskDelete, z.qd)
	z.qd.Define(queueIdReplaceFolderByFile, z.taskReplaceFolderByFile, z.qd)
	z.qd.Define(queueIdReplaceFileByFolder, z.taskReplaceFileByFolder, z.qd)

	if cnErr := z.conn.Startup(z.qd); cnErr != nil {
		l.Debug("Unable to startup the connector", esl.Error(cnErr))
		return cnErr
	}

	if srcEntry.IsFile() {
		z.enqueueTask(queueIdCopyFile, srcEntry.Path(), tgtEntry.Path(), &TaskCopyFile{
			Source: srcEntry.AsData(),
			Target: tgtEntry.Path().Descendant(srcEntry.Name()).AsData(),
		})
	} else {
		z.enqueueTask(queueIdSyncFolder, srcEntry.Path(), tgtEntry.Path(), &TaskSyncFolder{
			Source: srcEntry.Path().AsData(),
			Target: tgtEntry.Path().AsData(),
		})
	}

	l.Debug("Waiting for task completion")
	z.wg.Wait()

	l.Debug("Shutdown connector")
	if cnErr := z.conn.Shutdown(); cnErr != nil {
		l.Debug("Unable to shutdown the connector", esl.Error(cnErr))
		return cnErr
	}

	z.qd.Current().Wait()

	l.Debug("Sync finished", esl.Error(lastError))
	return lastError
}
