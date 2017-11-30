package patterns

import (
	"github.com/cihub/seelog"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"github.com/watermint/toolbox/api"
	"time"
)

type TreeWalk struct {
	ApiContext                      *api.ApiContext
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
		Recursive:                       false,
		IncludeMediaInfo:                false,
		IncludeDeleted:                  false,
		IncludeHasExplicitSharedMembers: false,
		IncludeMountedFolders:           true,
	}
}

func (t *TreeWalk) Walk(path api.DropboxPath, f func(files.IsMetadata) error) error {
	arg := files.NewListFolderArg(path.CleanPath())
	arg.Recursive = t.Recursive
	arg.IncludeMediaInfo = t.IncludeMediaInfo
	arg.IncludeDeleted = t.IncludeDeleted
	arg.IncludeHasExplicitSharedMembers = t.IncludeHasExplicitSharedMembers
	arg.IncludeMountedFolders = t.IncludeMountedFolders

	seelog.Tracef("ListFolder: Path[%s]", path)
	res, err := t.ApiContext.FilesListFolder(arg)
	if err != nil {
		seelog.Debugf("Unable to list folder[%s] : error[%s]", path.CleanPath(), err)
		return err
	}

	more := true
	for more {
		for _, r := range res.Entries {
			err = f(r)
			if err != nil {
				seelog.Tracef("Error from operation : error[%s]", err)
				return err
			}
		}
		if res.HasMore {
			contArg := files.NewListFolderContinueArg(res.Cursor)
			res, err = t.ApiContext.FilesListFolderContinue(contArg)
			if err != nil {
				seelog.Debugf("Unable to list folder(cont)[%s] : error[%s]", path.CleanPath(), err)
				return err
			}
		}
		more = res.HasMore
	}
	return nil
}

type BatchFileOper struct {
	BatchSize int
	TreeWalk  *TreeWalk
	Filter    func(m files.IsMetadata) bool
	BatchApi  func(m []files.IsMetadata) error
}

func (b *BatchFileOper) Oper(path api.DropboxPath) error {
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
