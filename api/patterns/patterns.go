package patterns

import (
	"github.com/cihub/seelog"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"github.com/watermint/toolbox/api"
)

func FilesListFolder(ac *api.ApiContext, lfa *files.ListFolderArg) (entries []files.IsMetadata, err error) {
	seelog.Tracef("ListFolder: Path[%s]", lfa.Path)
	res, err := ac.FilesListFolder(lfa)
	if err != nil {
		seelog.Debugf("Unable to list folder[%s] : error[%s]", lfa.Path, err)
		return
	}

	entries = make([]files.IsMetadata, 0)
	entries = append(entries, res.Entries...)

	if !res.HasMore {
		return
	}
	for {
		contArg := files.NewListFolderContinueArg(res.Cursor)
		res, err = ac.FilesListFolderContinue(contArg)
		if err != nil {
			seelog.Debugf("Unable to list folder(cont)[%s] : error[%s]", lfa.Path, err)
			return
		}
		entries = append(entries, res.Entries...)
		if !res.HasMore {
			return
		}
	}
}
