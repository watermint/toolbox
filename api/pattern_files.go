package api

//
//type ApiPatternFiles struct {
//	Context     *ApiContext
//	FilesClient files.Client
//}
//
//func (a *ApiPatternFiles) ListFolder(lfa *files.ListFolderArg) (entries []files.IsMetadata, err error) {
//	seelog.Tracef("ListFolder: Path[%s]", lfa.Path)
//	res, err := a.FilesClient.ListFolder(lfa)
//	if err != nil {
//		seelog.Debugf("Unable to list folder[%s] : error[%s]", lfa.Path, err)
//		return
//	}
//
//	entries = make([]files.IsMetadata, 0)
//	entries = append(entries, res.Entries...)
//
//	if !res.HasMore {
//		return
//	}
//	for {
//		contArg := files.NewListFolderContinueArg(res.Cursor)
//		res, err = a.FilesClient.ListFolderContinue(contArg)
//		if err != nil {
//			seelog.Debugf("Unable to list folder(cont)[%s] : error[%s]", lfa.Path, err)
//			return
//		}
//		entries = append(entries, res.Entries...)
//		if !res.HasMore {
//			return
//		}
//	}
//}
//
//func (a *ApiPatternFiles) Upload(content io.Reader, size int64, ci *files.CommitInfo) (fm *files.FileMetadata, err error) {
//	if size > a.Context.Config.UploadChunkedUploadThreshold {
//		fm, err = a.filesUploadChunked(content, size, ci)
//	} else {
//		fm, err = a.filesUploadSingle(content, size, ci)
//	}
//	if fm != nil {
//		seelog.Tracef("filesUpload: toPath[%s] id[%s] hash[%s]", fm.PathDisplay, fm.Id, fm.ContentHash)
//	}
//	return
//}
//
//func (a *ApiPatternFiles) filesUploadSingle(content io.Reader, size int64, ci *files.CommitInfo) (fm *files.FileMetadata, err error) {
//	seelog.Tracef("filesUploadSingle: toPath[%s] size[%d]", ci.Path, size)
//
//	return a.FilesClient.Upload(ci, content)
//}
//
//func (a *ApiPatternFiles) filesUploadChunked(content io.Reader, size int64, ci *files.CommitInfo) (fm *files.FileMetadata, err error) {
//	seelog.Tracef("filesUploadChunked: toPath[%s] size[%d]", ci.Path, size)
//
//	r := io.LimitReader(content, a.Context.Config.UploadChunkedUploadChunkSize)
//	s, err := a.FilesClient.UploadSessionStart(files.NewUploadSessionStartArg(), r)
//	if err != nil {
//		seelog.Debugf("Unable to start upload session : error[%s]", err)
//		return
//	}
//
//	var uploaded int64
//	uploaded = a.Context.Config.UploadChunkedUploadChunkSize
//	for (size - uploaded) > a.Context.Config.UploadChunkedUploadChunkSize {
//		seelog.Tracef("filesUploadChunked: toPath[%s]: uploaded[%d] of size[%d]", ci.Path, uploaded, size)
//
//		cursor := files.NewUploadSessionCursor(s.SessionId, uint64(uploaded))
//		arg := files.NewUploadSessionAppendArg(cursor)
//		r = io.LimitReader(content, int64(a.Context.Config.UploadChunkedUploadChunkSize))
//		err = a.FilesClient.UploadSessionAppendV2(arg, r)
//		if err != nil {
//			seelog.Debugf("Unable to append upload session : error[%s]", err)
//			return
//		}
//		uploaded += a.Context.Config.UploadChunkedUploadChunkSize
//	}
//
//	seelog.Tracef("filesUploadChunked: toPath[%s]: uploaded[%d] of size[%d]", ci.Path, uploaded, size)
//
//	cursor := files.NewUploadSessionCursor(s.SessionId, uint64(uploaded))
//	arg := files.NewUploadSessionFinishArg(cursor, ci)
//	fm, err = a.FilesClient.UploadSessionFinish(arg, content)
//	if err != nil {
//		seelog.Debugf("Unable to finish upload session : error[%s]", err)
//	}
//	return
//}
