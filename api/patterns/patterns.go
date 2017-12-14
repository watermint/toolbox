package patterns

import (
	"github.com/cihub/seelog"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"github.com/watermint/toolbox/api"
	"io"
)

const (
	// Hard limit of API spec: 150MB.
	UPLOAD_CHUNK_THRESHOLD int64 = 128 * 1048576 // 128MB
	UPLOAD_CHUNK_SIZE      int64 = 128 * 1048576 // 128MB
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

func FilesUpload(ac *api.ApiContext, content io.Reader, size int64, ci *files.CommitInfo) (fm *files.FileMetadata, err error) {
	if size > UPLOAD_CHUNK_THRESHOLD {
		fm, err = filesUploadChunked(ac, content, size, ci)
	} else {
		fm, err = filesUploadSingle(ac, content, size, ci)
	}
	if fm != nil {
		seelog.Tracef("filesUpload: toPath[%s] id[%s] hash[%s]", fm.PathDisplay, fm.Id, fm.ContentHash)
	}
	return
}

func filesUploadSingle(ac *api.ApiContext, content io.Reader, size int64, ci *files.CommitInfo) (fm *files.FileMetadata, err error) {
	seelog.Tracef("filesUploadSingle: toPath[%s] size[%d]", ci.Path, size)

	return ac.FilesUpload(ci, content)
}

func filesUploadChunked(ac *api.ApiContext, content io.Reader, size int64, ci *files.CommitInfo) (fm *files.FileMetadata, err error) {
	seelog.Tracef("filesUploadChunked: toPath[%s] size[%d]", ci.Path, size)

	r := io.LimitReader(content, UPLOAD_CHUNK_SIZE)
	s, err := ac.FilesUploadSessionStart(files.NewUploadSessionStartArg(), r)
	if err != nil {
		seelog.Debugf("Unable to start upload session : error[%s]", err)
		return
	}

	var uploaded int64
	uploaded = UPLOAD_CHUNK_SIZE
	for (size - uploaded) > UPLOAD_CHUNK_SIZE {
		seelog.Tracef("filesUploadChunked: toPath[%s]: uploaded[%d] of size[%d]", ci.Path, uploaded, size)

		cursor := files.NewUploadSessionCursor(s.SessionId, uint64(uploaded))
		arg := files.NewUploadSessionAppendArg(cursor)
		r = io.LimitReader(r, int64(UPLOAD_CHUNK_SIZE))
		err = ac.FilesUploadSessionAppendV2(arg, r)
		if err != nil {
			seelog.Debugf("Unable to append upload session : error[%s]", err)
			return
		}
		uploaded += UPLOAD_CHUNK_SIZE
	}

	seelog.Tracef("filesUploadChunked: toPath[%s]: uploaded[%d] of size[%d]", ci.Path, uploaded, size)

	cursor := files.NewUploadSessionCursor(s.SessionId, uint64(uploaded))
	arg := files.NewUploadSessionFinishArg(cursor, ci)
	fm, err = ac.FilesUploadSessionFinish(arg, content)
	if err != nil {
		seelog.Debugf("Unable to finish upload session : error[%s]", err)
	}
	return
}
