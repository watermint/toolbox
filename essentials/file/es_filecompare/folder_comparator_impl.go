package es_filecompare

import (
	"errors"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"strings"
	"sync"
)

const (
	queueIdCompareFolder = "compare_folder"
)

var (
	ErrorUnknownEntryState = errors.New("unknown entry state")
)

func NewFolderComparator(source, target es_filesystem.FileSystem, seq eq_sequence.Sequence, opt ...Opt) FolderComparator {
	comparator := New(opt...)
	return &folderComparator{
		sourceFs:   source,
		targetFs:   target,
		opts:       Opts{}.Apply(opt),
		comparator: comparator,
		seq:        seq,
	}
}

func CompareLevel(
	logger esl.Logger,
	sourceFs es_filesystem.FileSystem,
	targetFs es_filesystem.FileSystem,
	sourcePath es_filesystem.Path,
	targetPath es_filesystem.Path,
	comparator FileComparator,
	handlerMissingSource MissingSource,
	handlerMissingTarget MissingTarget,
	handlerFileDiff FileDiff,
	handlerTypeDiff TypeDiff,
	handlerSameFile SameFile,
	handlerCompareFolder CompareFolder,
) error {
	l := logger.With(esl.Any("source", sourcePath.AsData()), esl.Any("target", targetPath.AsData()))
	sourceEntries := make(map[string]es_filesystem.Entry)
	sourceFiles := make(map[string]es_filesystem.Entry)
	targetFiles := make(map[string]es_filesystem.Entry)
	targetEntries := make(map[string]es_filesystem.Entry)
	sourceFolders := make(map[string]es_filesystem.Entry)
	targetFolders := make(map[string]es_filesystem.Entry)
	base := PathPair{
		Source: sourcePath,
		Target: targetPath,
	}

	// scan source folder
	l.Debug("Scan source folder")
	{
		entries, err := sourceFs.List(sourcePath)
		if err != nil {
			l.Debug("Unable to list source folder", esl.Error(err))
			return err
		}

		for _, e := range entries {
			entryName := strings.ToLower(e.Name())
			if e.IsFile() {
				sourceFiles[entryName] = e
			}
			if e.IsFolder() {
				sourceFolders[entryName] = e
			}
			sourceEntries[entryName] = e
		}
	}

	// scan target folder
	l.Debug("Scan target folder")
	{
		entries, err := targetFs.List(targetPath)
		if err != nil {
			l.Debug("Unable to list target folder", esl.Error(err))
			return err
		}

		for _, e := range entries {
			entryName := strings.ToLower(e.Name())
			if e.IsFile() {
				targetFiles[entryName] = e
			}
			if e.IsFolder() {
				targetFolders[entryName] = e
			}
			targetEntries[entryName] = e
		}
	}

	var lastErr error
	// compare files source -> target
	l.Debug("Compare file: source -> target")
	for sn, se := range sourceFiles {
		if te, ok := targetFiles[sn]; ok {
			ll := l.With(esl.Any("sourceEntry", se.AsData()), esl.Any("targetEntry", te.AsData()))
			same, err := comparator.Compare(se, te)
			if err != nil {
				ll.Debug("Unable to compare files", esl.Error(err))
				lastErr = err
			} else if same {
				ll.Debug("Same content")
				handlerSameFile(base, se, te)
			} else {
				ll.Debug("Content diff found")
				handlerFileDiff(base, se, te)
			}
		} else if _, ok := targetEntries[sn]; !ok {
			ll := l.With(esl.Any("sourceEntry", se.AsData()))
			ll.Debug("Target missing")
			handlerMissingTarget(base, se)
		}
	}

	// compare files target -> source
	l.Debug("Compare file: target -> source")
	for tn, te := range targetFiles {
		if _, ok := sourceFiles[tn]; !ok {
			if _, ok := sourceEntries[tn]; !ok {
				ll := l.With(esl.Any("targetEntry", te.AsData()))
				ll.Debug("Source missing")

				handlerMissingSource(base, te)
			}
		}
	}

	// compare folders: source -> target
	for sn, se := range sourceFolders {
		if te, ok := targetFolders[sn]; ok {
			ll := l.With(esl.Any("sourceEntry", se.AsData()), esl.Any("targetEntry", te.AsData()))
			ll.Debug("Recurse into descendant")

			handlerCompareFolder(base, se, te)
		} else if _, ok := targetEntries[sn]; !ok {
			l.Debug("Target missing")
			handlerMissingTarget(base, se)
		}
	}

	// compare folders: target -> source
	for tn, te := range targetFolders {
		if _, ok := sourceFolders[tn]; !ok {
			if _, ok := sourceEntries[tn]; !ok {
				ll := l.With(esl.Any("targetEntry", te.AsData()))
				ll.Debug("Source missing")
				handlerMissingSource(base, te)
			}
		}
	}

	// detect type diff
	for sn, se := range sourceEntries {
		if te, ok := targetEntries[sn]; ok {
			ll := l.With(esl.Any("sourceEntry", se.AsData()), esl.Any("targetEntry", te.AsData()))
			switch {
			case se.IsFolder() && te.IsFolder():
				ll.Debug("both folder")
			case se.IsFolder() && te.IsFile():
				ll.Debug("type diff : source (folder), target (file)")
				handlerTypeDiff(base, se, te)
			case se.IsFile() && te.IsFolder():
				ll.Debug("type diff : source (file), target (folder)")
				handlerTypeDiff(base, se, te)
			case se.IsFile() && te.IsFile():
				ll.Debug("both file")
			default:
				ll.Debug("unknown state")
				lastErr = ErrorUnknownEntryState
			}
		}
	}

	l.Debug("Complete")
	return lastErr
}

type TaskCompareEntry struct {
	Source es_filesystem.PathData `json:"source"`
	Target es_filesystem.PathData `json:"target"`
}

type folderComparator struct {
	sourceFs   es_filesystem.FileSystem
	targetFs   es_filesystem.FileSystem
	opts       Opts
	comparator FileComparator
	seq        eq_sequence.Sequence
}

func (z folderComparator) batchId(source, target es_filesystem.Path) string {
	return source.Namespace().Id() + ":" + target.Namespace().Id()
}

func (z folderComparator) compareLevel(entry *TaskCompareEntry, stg eq_sequence.Stage) error {
	l := z.opts.Log().With(esl.Any("source", entry.Source), esl.Any("target", entry.Target))
	sourcePath, err := z.sourceFs.Path(entry.Source)
	if err != nil {
		l.Debug("Unable to deserialize source path", esl.Error(err))
		return err
	}
	targetPath, err := z.targetFs.Path(entry.Target)
	if err != nil {
		l.Debug("Unable to deserialize target path", esl.Error(err))
		return err
	}

	enqueueDescendant := func(base PathPair, source, target es_filesystem.Entry) {
		q := stg.Get(queueIdCompareFolder).Batch(z.batchId(source.Path(), target.Path()))
		q.Enqueue(&TaskCompareEntry{
			Source: source.Path().AsData(),
			Target: target.Path().AsData(),
		})
	}

	return CompareLevel(z.opts.Log(),
		z.sourceFs,
		z.targetFs,
		sourcePath,
		targetPath,
		z.comparator,
		z.opts.ReportMissingSource,
		z.opts.ReportMissingTarget,
		z.opts.ReportFileDiff,
		z.opts.ReportTypeDiff,
		z.opts.ReportSameFile,
		enqueueDescendant,
	)
}

func (z folderComparator) Compare(source, target es_filesystem.Path) (err error) {
	var lastError error
	lastErrorMutex := sync.Mutex{}
	onError := func(err error, mouldId, batchId string, p interface{}) {
		lastErrorMutex.Lock()
		if err != nil {
			lastError = err
		}
		lastErrorMutex.Unlock()
	}
	z.seq.Do(func(s eq_sequence.Stage) {
		s.Define(queueIdCompareFolder, z.compareLevel, s)
		q := s.Get(queueIdCompareFolder).Batch(z.batchId(source, target))
		q.Enqueue(&TaskCompareEntry{
			Source: source.AsData(),
			Target: target.AsData(),
		})
	}, eq_sequence.ErrorHandler(onError))

	return lastError
}

func (z folderComparator) CompareAndSummarize(source, target es_filesystem.Path) (missingSource, missingTarget []es_filesystem.Entry, fileDiffs, typeDiffs []EntryDataPair, err error) {
	missingTarget = make([]es_filesystem.Entry, 0)
	missingTargetMutex := sync.Mutex{}
	missingSource = make([]es_filesystem.Entry, 0)
	missingSourceMutex := sync.Mutex{}
	fileDiffs = make([]EntryDataPair, 0)
	fileDiffMutex := sync.Mutex{}
	typeDiffs = make([]EntryDataPair, 0)
	typeDiffMutex := sync.Mutex{}

	onMissingTarget := func(base PathPair, target es_filesystem.Entry) {
		missingTargetMutex.Lock()
		missingTarget = append(missingTarget, target)
		missingTargetMutex.Unlock()
	}
	onMissingSource := func(base PathPair, source es_filesystem.Entry) {
		missingSourceMutex.Lock()
		missingSource = append(missingSource, source)
		missingSourceMutex.Unlock()
	}
	onFileDiff := func(base PathPair, source, target es_filesystem.Entry) {
		fileDiffMutex.Lock()
		fileDiffs = append(fileDiffs, EntryDataPair{
			SourceData: source.AsData(),
			TargetData: target.AsData(),
		})
		fileDiffMutex.Unlock()
	}
	onTypeDiff := func(base PathPair, source, target es_filesystem.Entry) {
		typeDiffMutex.Lock()
		typeDiffs = append(typeDiffs, EntryDataPair{
			SourceData: source.AsData(),
			TargetData: target.AsData(),
		})
		typeDiffMutex.Unlock()
	}

	newOpts := make([]Opt, 0)
	newOpts = append(newOpts,
		HandlerMissingSource(onMissingSource),
		HandlerMissingTarget(onMissingTarget),
		HandlerFileDiff(onFileDiff),
		HandlerTypeDiff(onTypeDiff),
	)
	z.opts = z.opts.Apply(newOpts)
	err = z.Compare(source, target)
	return
}
