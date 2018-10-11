package api

import (
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/async"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/file_properties"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"io"
)

func (a *ApiFiles) AlphaGetMetadata(arg *files.AlphaGetMetadataArg) (res files.IsMetadata, err error) {
	return a.Compat().AlphaGetMetadata(arg)
}
func (a *ApiFiles) AlphaUpload(arg *files.CommitInfoWithProperties, content io.Reader) (res *files.FileMetadata, err error) {
	return a.Compat().AlphaUpload(arg, content)
}
func (a *ApiFiles) Copy(arg *files.RelocationArg) (res files.IsMetadata, err error) {
	return a.Compat().Copy(arg)
}
func (a *ApiFiles) CopyBatch(arg *files.RelocationBatchArg) (res *files.RelocationBatchLaunch, err error) {
	return a.Compat().CopyBatch(arg)
}
func (a *ApiFiles) CopyBatchCheck(arg *async.PollArg) (res *files.RelocationBatchJobStatus, err error) {
	return a.Compat().CopyBatchCheck(arg)
}
func (a *ApiFiles) CopyReferenceGet(arg *files.GetCopyReferenceArg) (res *files.GetCopyReferenceResult, err error) {
	return a.Compat().CopyReferenceGet(arg)
}
func (a *ApiFiles) CopyReferenceSave(arg *files.SaveCopyReferenceArg) (res *files.SaveCopyReferenceResult, err error) {
	return a.Compat().CopyReferenceSave(arg)
}
func (a *ApiFiles) CopyV2(arg *files.RelocationArg) (*files.RelocationResult, error) {
	if r, err := a.Context.NewApiRpcRequest("files/copy_v2", nil, arg).Call(); err != nil {
		return nil, err
	} else {
		return parseRelocationResult(r)
	}
}
func (a *ApiFiles) CreateFolder(arg *files.CreateFolderArg) (res *files.FolderMetadata, err error) {
	return a.Compat().CreateFolder(arg)
}
func (a *ApiFiles) CreateFolderV2(arg *files.CreateFolderArg) (res *files.CreateFolderResult, err error) {
	return a.Compat().CreateFolderV2(arg)
}
func (a *ApiFiles) CreateFolderBatch(arg *files.CreateFolderBatchArg) (res *files.CreateFolderBatchLaunch, err error) {
	return a.Compat().CreateFolderBatch(arg)
}
func (a *ApiFiles) 	CreateFolderBatchCheck(arg *async.PollArg) (res *files.CreateFolderBatchJobStatus, err error) {
	return a.Compat().CreateFolderBatchCheck(arg)
}
func (a *ApiFiles) Delete(arg *files.DeleteArg) (res files.IsMetadata, err error) {
	return a.Compat().Delete(arg)
}
func (a *ApiFiles) DeleteBatch(arg *files.DeleteBatchArg) (res *files.DeleteBatchLaunch, err error) {
	return a.Compat().DeleteBatch(arg)
}
func (a *ApiFiles) DeleteBatchCheck(arg *async.PollArg) (res *files.DeleteBatchJobStatus, err error) {
	return a.Compat().DeleteBatchCheck(arg)
}
func (a *ApiFiles) DeleteV2(arg *files.DeleteArg) (res *files.DeleteResult, err error) {
	return a.Compat().DeleteV2(arg)
}
func (a *ApiFiles) Download(arg *files.DownloadArg) (res *files.FileMetadata, content io.ReadCloser, err error) {
	return a.Compat().Download(arg)
}
func (a *ApiFiles) DownloadZip(arg *files.DownloadZipArg) (res *files.DownloadZipResult, content io.ReadCloser, err error) {
	return a.Compat().DownloadZip(arg)
}
func (a *ApiFiles) GetMetadata(arg *files.GetMetadataArg) (res files.IsMetadata, err error) {
	return a.Compat().GetMetadata(arg)
}
func (a *ApiFiles) GetPreview(arg *files.PreviewArg) (res *files.FileMetadata, content io.ReadCloser, err error) {
	return a.Compat().GetPreview(arg)
}
func (a *ApiFiles) GetTemporaryLink(arg *files.GetTemporaryLinkArg) (res *files.GetTemporaryLinkResult, err error) {
	return a.Compat().GetTemporaryLink(arg)
}
func (a *ApiFiles) 	GetTemporaryUploadLink(arg *files.GetTemporaryUploadLinkArg) (res *files.GetTemporaryUploadLinkResult, err error) {
	return a.Compat().GetTemporaryUploadLink(arg)
}
func (a *ApiFiles) GetThumbnail(arg *files.ThumbnailArg) (res *files.FileMetadata, content io.ReadCloser, err error) {
	return a.Compat().GetThumbnail(arg)
}
func (a *ApiFiles) GetThumbnailBatch(arg *files.GetThumbnailBatchArg) (res *files.GetThumbnailBatchResult, err error) {
	return a.Compat().GetThumbnailBatch(arg)
}
func (a *ApiFiles) ListFolder(arg *files.ListFolderArg) (*files.ListFolderResult, error) {
	if r, err := a.Context.NewApiRpcRequest("files/list_folder", parseErrorListFolder, arg).Call(); err != nil {
		return nil, err
	} else {
		return parseListFolderResult(r)
	}
}
func (a *ApiFiles) ListFolderContinue(arg *files.ListFolderContinueArg) (res *files.ListFolderResult, err error) {
	if r, err := a.Context.NewApiRpcRequest("files/list_folder/continue", parseErrorListFolder, arg).Call(); err != nil {
		return nil, err
	} else {
		return parseListFolderResult(r)
	}
}
func (a *ApiFiles) ListFolderGetLatestCursor(arg *files.ListFolderArg) (res *files.ListFolderGetLatestCursorResult, err error) {
	return a.Compat().ListFolderGetLatestCursor(arg)
}
func (a *ApiFiles) ListFolderLongpoll(arg *files.ListFolderLongpollArg) (res *files.ListFolderLongpollResult, err error) {
	return a.Compat().ListFolderLongpoll(arg)
}
func (a *ApiFiles) ListRevisions(arg *files.ListRevisionsArg) (res *files.ListRevisionsResult, err error) {
	return a.Compat().ListRevisions(arg)
}
func (a *ApiFiles) Move(arg *files.RelocationArg) (res files.IsMetadata, err error) {
	return a.Compat().Move(arg)
}
func (a *ApiFiles) MoveBatch(arg *files.RelocationBatchArg) (res *files.RelocationBatchLaunch, err error) {
	return a.Compat().MoveBatch(arg)
}
func (a *ApiFiles) MoveBatchCheck(arg *async.PollArg) (res *files.RelocationBatchJobStatus, err error) {
	return a.Compat().MoveBatchCheck(arg)
}
func (a *ApiFiles) MoveV2(arg *files.RelocationArg) (*files.RelocationResult, error) {
	if r, err := a.Context.NewApiRpcRequest("files/move_v2", nil, arg).Call(); err != nil {
		return nil, err
	} else {
		return parseRelocationResult(r)
	}
}
func (a *ApiFiles) PermanentlyDelete(arg *files.DeleteArg) (err error) {
	return a.Compat().PermanentlyDelete(arg)
}
func (a *ApiFiles) PropertiesAdd(arg *file_properties.AddPropertiesArg) (err error) {
	return a.Compat().PropertiesAdd(arg)
}
func (a *ApiFiles) PropertiesOverwrite(arg *file_properties.OverwritePropertyGroupArg) (err error) {
	return a.Compat().PropertiesOverwrite(arg)
}
func (a *ApiFiles) PropertiesRemove(arg *file_properties.RemovePropertiesArg) (err error) {
	return a.Compat().PropertiesRemove(arg)
}
func (a *ApiFiles) PropertiesTemplateGet(arg *file_properties.GetTemplateArg) (res *file_properties.GetTemplateResult, err error) {
	return a.Compat().PropertiesTemplateGet(arg)
}
func (a *ApiFiles) PropertiesTemplateList() (res *file_properties.ListTemplateResult, err error) {
	return a.Compat().PropertiesTemplateList()
}
func (a *ApiFiles) PropertiesUpdate(arg *file_properties.UpdatePropertiesArg) (err error) {
	return a.Compat().PropertiesUpdate(arg)
}
func (a *ApiFiles) Restore(arg *files.RestoreArg) (res *files.FileMetadata, err error) {
	return a.Compat().Restore(arg)
}
func (a *ApiFiles) SaveUrl(arg *files.SaveUrlArg) (res *files.SaveUrlResult, err error) {
	return a.Compat().SaveUrl(arg)
}
func (a *ApiFiles) SaveUrlCheckJobStatus(arg *async.PollArg) (res *files.SaveUrlJobStatus, err error) {
	return a.Compat().SaveUrlCheckJobStatus(arg)
}
func (a *ApiFiles) Search(arg *files.SearchArg) (sr *files.SearchResult, err error) {
	if r, err := a.Context.NewApiRpcRequest("files/search", parseErrorFilesSearch, arg).Call(); err != nil {
		return nil, err
	} else {
		return parseSearchResult(r)
	}
}
func (a *ApiFiles) Upload(arg *files.CommitInfo, content io.Reader) (res *files.FileMetadata, err error) {
	return a.Compat().Upload(arg, content)
}
func (a *ApiFiles) UploadSessionAppend(arg *files.UploadSessionCursor, content io.Reader) (err error) {
	return a.Compat().UploadSessionAppend(arg, content)
}
func (a *ApiFiles) UploadSessionAppendV2(arg *files.UploadSessionAppendArg, content io.Reader) (err error) {
	return a.Compat().UploadSessionAppendV2(arg, content)
}
func (a *ApiFiles) UploadSessionFinish(arg *files.UploadSessionFinishArg, content io.Reader) (res *files.FileMetadata, err error) {
	return a.Compat().UploadSessionFinish(arg, content)
}
func (a *ApiFiles) UploadSessionFinishBatch(arg *files.UploadSessionFinishBatchArg) (res *files.UploadSessionFinishBatchLaunch, err error) {
	return a.Compat().UploadSessionFinishBatch(arg)
}
func (a *ApiFiles) UploadSessionFinishBatchCheck(arg *async.PollArg) (res *files.UploadSessionFinishBatchJobStatus, err error) {
	return a.Compat().UploadSessionFinishBatchCheck(arg)
}
func (a *ApiFiles) UploadSessionStart(arg *files.UploadSessionStartArg, content io.Reader) (res *files.UploadSessionStartResult, err error) {
	return a.Compat().UploadSessionStart(arg, content)
}
