package patterns

import (
	"errors"
	"github.com/cihub/seelog"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"github.com/watermint/toolbox/api"
	"time"
)

type TreeWalk struct {
	ApiContext *api.ApiContext

	// Recursively walk down folder tree
	RecursiveWalk bool

	// API option: Recursive
	Recursive                       bool
	IncludeMediaInfo                bool
	IncludeDeleted                  bool
	IncludeHasExplicitSharedMembers bool
	IncludeMountedFolders           bool
	//	IncludePropertyGroups           bool
}

// New TreeWalk with Api default options
func NewTreeWalk(ac *api.ApiContext) *TreeWalk {
	return &TreeWalk{
		ApiContext:                      ac,
		RecursiveWalk:                   false,
		Recursive:                       false,
		IncludeMediaInfo:                false,
		IncludeDeleted:                  false,
		IncludeHasExplicitSharedMembers: false,
		IncludeMountedFolders:           true,
	}
}

func (t *TreeWalk) walkDispatch(isTop bool, entry files.IsMetadata, f func(files.IsMetadata) error) error {
	switch e := entry.(type) {
	case *files.FileMetadata:
		return f(e)
	case *files.DeletedMetadata:
		return f(e)
	case *files.FolderMetadata:
		if t.RecursiveWalk || isTop {
			t.walkFolder(e, f)
		}
		return f(e)
	default:
		seelog.Warn("Unknown metadata type detected")
		return errors.New("unknown metadata type detected")
	}
}

func (t *TreeWalk) walkFolder(folder *files.FolderMetadata, f func(files.IsMetadata) error) error {
	lfa := files.NewListFolderArg(folder.PathDisplay)
	lfa.Recursive = t.Recursive
	lfa.IncludeMediaInfo = t.IncludeMediaInfo
	lfa.IncludeDeleted = t.IncludeDeleted
	lfa.IncludeHasExplicitSharedMembers = t.IncludeHasExplicitSharedMembers
	lfa.IncludeMountedFolders = t.IncludeMountedFolders

	seelog.Tracef("ListFolder: Path[%s]", folder.PathDisplay)
	res, err := t.ApiContext.FilesListFolder(lfa)
	if err != nil {
		seelog.Debugf("Unable to list folder[%s] : error[%s]", folder.PathDisplay, err)
		return err
	}

	more := true

	allEntries := make([]files.IsMetadata, 0)
	for more {
		allEntries = append(allEntries, res.Entries...)
		if res.HasMore {
			contArg := files.NewListFolderContinueArg(res.Cursor)
			res, err = t.ApiContext.FilesListFolderContinue(contArg)
			if err != nil {
				seelog.Debugf("Unable to list folder(cont)[%s] : error[%s]", folder.PathDisplay, err)
				return err
			}
		}
		more = res.HasMore
	}

	for _, entry := range allEntries {
		err = t.walkDispatch(false, entry, f)
		if err != nil {
			seelog.Tracef("Error from operation : error[%s]", err)
			return err
		}
	}

	return nil
}

func (t *TreeWalk) Walk(path *api.DropboxPath, f func(files.IsMetadata) error) error {
	seelog.Tracef("Walk path[%s]", path.CleanPath())
	gma := files.NewGetMetadataArg(path.CleanPath())
	gma.IncludeDeleted = t.IncludeDeleted
	gma.IncludeMediaInfo = t.IncludeMediaInfo
	entry, err := t.ApiContext.FilesGetMetadata(gma)
	if err != nil {
		seelog.Debugf("Unable to get metadata for path[%s] : error[%s]", path.CleanPath(), err)
		return err
	}
	return t.walkDispatch(true, entry, f)
}

type BatchFileOper struct {
	BatchSize int
	TreeWalk  *TreeWalk
	Filter    func(m files.IsMetadata) bool
	BatchApi  func(m []files.IsMetadata) error
}

func (b *BatchFileOper) Oper(path *api.DropboxPath) error {
	batch := make([]files.IsMetadata, 0)
	seelog.Debugf("Walk tree for batch operation")

	bef := func(m []files.IsMetadata) error {
		retries := 0
		for {
			seelog.Tracef("Execute batch: size[%d]", len(m))
			be := b.BatchApi(m)
			re := api.IsRetriableError(be)
			if re != api.THINSDK_RETRY_REASON_NORETRY {
				return be
			}
			retries++
			if api.THINSDK_API_CALL_RETRY_MAX <= retries {
				seelog.Debugf("Reached to maximum retry[%d] error[%s]", be)
				return be
			}
			seelog.Debugf("Retry with reason[%d] retries[%d] error[%s]", re, retries, be)
			time.Sleep(api.THINSDK_API_CALL_RETRY_INTERVAL)
		}
	}

	err := b.TreeWalk.Walk(path, func(m files.IsMetadata) error {
		if !b.Filter(m) {
			return nil
		}
		batch = append(batch, m)
		if len(batch) < b.BatchSize {
			return nil
		}
		return bef(batch)
	})
	if err != nil {
		seelog.Tracef("Batch execution failed: error[%s]", err)
		return err
	}

	if 0 < len(batch) {
		seelog.Tracef("Exec remainder batch: size[%d]", len(batch))
		return bef(batch)
	}
	return nil
}
