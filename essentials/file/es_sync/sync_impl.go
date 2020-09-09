package es_sync

import (
	"github.com/watermint/toolbox/essentials/file/es_filecompare"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"sync"
)

func New(log esl.Logger, seq eq_sequence.Sequence, source, target es_filesystem.FileSystem, conn es_filesystem.Connector, opt ...Opt) Syncer {
	return &syncImpl{
		log:    log,
		seq:    seq,
		source: source,
		target: target,
		conn:   conn,
		opts:   Opts{}.Apply(opt),
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
	log    esl.Logger
	seq    eq_sequence.Sequence
	source es_filesystem.FileSystem
	target es_filesystem.FileSystem
	conn   es_filesystem.Connector
	fcmp   es_filecompare.FileComparator
	opts   Opts
}

func (z syncImpl) computeBatchId(source, target es_filesystem.Path) string {
	return source.Namespace().Id() + "/" + target.Namespace().Id()
}

func (z syncImpl) copy(source es_filesystem.Entry, target es_filesystem.Path) error {
	l := z.log.With(esl.Any("source", source.AsData()), esl.Any("target", target.AsData()))
	l.Debug("Copy")
	err := z.conn.Copy(source, target)
	if err != nil {
		l.Debug("Unable to copy data from source to target", esl.Error(err))
		z.opts.OnCopyFailure(source.Path(), err)
		return err
	}
	z.opts.OnCopySuccess(source, target)
	return nil
}

func (z syncImpl) delete(target es_filesystem.Path) error {
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

func (z syncImpl) createFolder(target es_filesystem.Path) error {
	l := z.log.With(esl.Any("target", target.AsData()))
	l.Debug("Create folder")
	err := z.target.CreateFolder(target)
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
func (z syncImpl) taskCopyFile(task *TaskCopyFile, stg eq_sequence.Stage) error {
	l := z.log.With(esl.Any("task", task))

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

	l.Debug("Copy file")
	return z.copy(sourceEntry, targetPath)
}

type TaskReplaceFolderByFile struct {
	Source es_filesystem.EntryData `json:"source"`
	Target es_filesystem.PathData  `json:"target"`
}

func (z syncImpl) taskReplaceFolderByFile(task *TaskReplaceFolderByFile, stg eq_sequence.Stage) error {
	l := z.log.With(esl.Any("task", task))

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

	q := stg.Get(queueIdCopyFile).Batch(z.computeBatchId(sourceEntry.Path(), targetPath))
	q.Enqueue(&TaskCopyFile{
		Source: sourceEntry.AsData(),
		Target: targetPath.AsData(),
	})
	return nil
}

type TaskReplaceFileByFolder struct {
	Source es_filesystem.EntryData `json:"source"`
	Target es_filesystem.PathData  `json:"target"`
}

func (z syncImpl) taskReplaceFileByFolder(task *TaskReplaceFileByFolder, stg eq_sequence.Stage) error {
	l := z.log.With(esl.Any("task", task))

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
	q := stg.Get(queueIdSyncFolder).Batch(z.computeBatchId(sourceEntry.Path(), targetPath))
	q.Enqueue(&TaskSyncFolder{
		Source: sourceEntry.Path().AsData(),
		Target: targetPath.AsData(),
	})
	return nil
}

type TaskSyncFolder struct {
	Source es_filesystem.PathData `json:"source"`
	Target es_filesystem.PathData `json:"target"`
}

func (z syncImpl) taskSyncFolder(task *TaskSyncFolder, stg eq_sequence.Stage) error {
	l := z.log.With(esl.Any("source", task.Source), esl.Any("target", task.Target))
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
		q := stg.Get(queueIdDelete).Batch(z.computeBatchId(base.Source, base.Target))
		q.Enqueue(&TaskDelete{
			Target: target.Path().AsData(),
		})
	}
	handlerMissingTarget := func(base es_filecompare.PathPair, source es_filesystem.Entry) {
		ll := l.With(esl.Any("source", source.AsData()))
		newTargetPath := base.Target.Descendant(source.Path().Base())

		if source.IsFolder() {
			ll.Debug("create folder", esl.Any("targetFolderPath", newTargetPath.AsData()))
			if err := z.createFolder(newTargetPath); err != nil {
				ll.Debug("unable to create folder")
				return
			}

			ll.Debug("Enqueue sync folder")
			q := stg.Get(queueIdSyncFolder).Batch(z.computeBatchId(source.Path(), newTargetPath))
			q.Enqueue(&TaskSyncFolder{
				Source: source.Path().AsData(),
				Target: newTargetPath.AsData(),
			})
		} else {
			ll.Debug("Copy file")
			q := stg.Get(queueIdCopyFile).Batch(z.computeBatchId(source.Path(), newTargetPath))
			q.Enqueue(&TaskCopyFile{
				Source: source.AsData(),
				Target: newTargetPath.AsData(),
			})
		}
	}
	handlerFileDiff := func(base es_filecompare.PathPair, source, target es_filesystem.Entry) {
		ll := l.With(esl.Any("source", source.AsData()), esl.Any("target", target.AsData()))
		if !z.opts.SyncOverwrite() {
			ll.Debug("Don't overwrite")
			z.opts.OnSkip(SkipExists, source, target)
			return
		}

		ll.Debug("Overwrite")
		q := stg.Get(queueIdCopyFile).Batch(z.computeBatchId(source.Path(), target.Path()))
		q.Enqueue(&TaskCopyFile{
			Source: source.AsData(),
			Target: target.Path().AsData(),
		})
	}
	handlerSameFile := func(base es_filecompare.PathPair, source, target es_filesystem.Entry) {
		ll := l.With(esl.Any("source", source.AsData()), esl.Any("target", target.AsData()))
		ll.Debug("Same file, skip sync")
		z.opts.OnSkip(SkipSame, source, target)
	}
	handlerTypeDiff := func(base es_filecompare.PathPair, source, target es_filesystem.Entry) {
		ll := l.With(esl.Any("source", source.AsData()), esl.Any("target", target.AsData()))
		if !z.opts.SyncOverwrite() {
			ll.Debug("Don't overwrite")
			z.opts.OnSkip(SkipExists, source, target)
			return
		}

		switch {
		case source.IsFile() && target.IsFolder():
			ll.Debug("Replace folder with file")
			q := stg.Get(queueIdReplaceFolderByFile).Batch(z.computeBatchId(source.Path(), target.Path()))
			q.Enqueue(&TaskReplaceFolderByFile{
				Source: source.AsData(),
				Target: target.Path().AsData(),
			})
		case source.IsFolder() && target.IsFile():
			ll.Debug("Replace file with folder")
			q := stg.Get(queueIdReplaceFileByFolder).Batch(z.computeBatchId(source.Path(), target.Path()))
			q.Enqueue(&TaskReplaceFileByFolder{
				Source: source.AsData(),
				Target: target.Path().AsData(),
			})
		case source.IsFile() && target.IsFile():
			ll.Debug("Same type, but do as copy")
			q := stg.Get(queueIdCopyFile).Batch(z.computeBatchId(source.Path(), target.Path()))
			q.Enqueue(&TaskCopyFile{
				Source: source.AsData(),
				Target: target.Path().AsData(),
			})
		case source.IsFolder() && target.IsFolder():
			ll.Debug("Same type, but do as regular sync folder")
			q := stg.Get(queueIdSyncFolder).Batch(z.computeBatchId(source.Path(), target.Path()))
			q.Enqueue(&TaskSyncFolder{
				Source: source.Path().AsData(),
				Target: target.Path().AsData(),
			})
		}
	}
	handlerDescendant := func(base es_filecompare.PathPair, source, target es_filesystem.Entry) {
		l.Debug("Sync descendant", esl.Any("source", source.AsData()), esl.Any("target", target.AsData()))
		q := stg.Get(queueIdSyncFolder).Batch(z.computeBatchId(source.Path(), target.Path()))
		q.Enqueue(&TaskSyncFolder{
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
		z.fcmp,
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

func (z syncImpl) taskDelete(task *TaskDelete, stg eq_sequence.Stage) error {
	l := z.log.With(esl.Any("target", task.Target))
	l.Debug("Delete")

	targetPath, err := z.target.Path(task.Target)
	if err != nil {
		l.Debug("Unable to deserialize", esl.Error(err))
		return err
	}

	return z.delete(targetPath)
}

func (z syncImpl) Sync(source es_filesystem.Path, target es_filesystem.Path) error {
	l := z.log.With(esl.String("sourcePath", source.Path()), esl.String("targetPath", target.Path()))
	l.Debug("Start sync")

	var lastError error
	lastErrorMutex := sync.Mutex{}
	onError := func(err error, mouldId, batchId string, p interface{}) {
		lastErrorMutex.Lock()
		if err != nil {
			lastError = err
		}
		lastErrorMutex.Unlock()
	}

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
		errCreateFolder := z.target.CreateFolder(target)
		if errCreateFolder != nil {
			l.Debug("Unable to create folder", esl.Error(errCreateFolder))
			z.opts.OnCreateFolderFailure(target, err)
			return errCreateFolder
		}
		z.opts.OnCreateFolderSuccess(target)

	default:
		l.Debug("Unable to retrieve target entry", esl.Error(err))
		return err
	}

	z.seq.Do(func(s eq_sequence.Stage) {
		s.Define(queueIdCopyFile, z.taskCopyFile, s)
		s.Define(queueIdSyncFolder, z.taskSyncFolder, s)
		s.Define(queueIdDelete, z.taskDelete, s)
		s.Define(queueIdReplaceFolderByFile, z.taskReplaceFolderByFile, s)
		s.Define(queueIdReplaceFileByFolder, z.taskReplaceFileByFolder, s)

		q := s.Get(queueIdSyncFolder).Batch(z.computeBatchId(srcEntry.Path(), tgtEntry.Path()))
		q.Enqueue(&TaskSyncFolder{
			Source: srcEntry.Path().AsData(),
			Target: tgtEntry.Path().AsData(),
		})
	}, eq_sequence.ErrorHandler(onError))

	l.Debug("Sync finished", esl.Error(lastError))
	return lastError
}
