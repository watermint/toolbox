package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/async"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/file_properties"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

var (
	API_RPC_ENDPOINT                                  = "api.dropboxapi.com"
	API_REQ_HEADER_SELECT_USER                        = "Dropbox-API-Select-User"
	API_RES_HEADER_RETRY_AFTER                        = "Retry-After"
	API_DEFAULT_UPLOAD_CHUNKED_UPLOAD_THRESHOLD int64 = 150 * 1048576
	API_DEFAULT_UPLOAD_CHUNK_SIZE               int64 = 150 * 1048576
	API_DEFAULT_CLIENT_TIMEOUT                        = 60
	API_DEFAULT_CLIENT_RETRY                          = 1
)

type DropboxPath struct {
	Path string
}

func NewDropboxPath(path string) *DropboxPath {
	return &DropboxPath{
		Path: path,
	}
}

func (d *DropboxPath) CleanPath() string {
	p := filepath.ToSlash(filepath.Clean(d.Path))
	if p == "/" {
		return ""
	} else {
		return p
	}
}

func RebaseTimeForAPI(t time.Time) time.Time {
	return t.UTC().Round(time.Second)
}

type ApiConfig struct {
	Timeout                      time.Duration
	Retry                        int
	UploadChunkedUploadThreshold int64
	UploadChunkedUploadChunkSize int64
	ErrorCallback                ApiErrorCallback
}

func NewDefaultApiConfig() *ApiConfig {
	return &ApiConfig{
		Timeout: time.Duration(API_DEFAULT_CLIENT_TIMEOUT) * time.Second,
		Retry:   API_DEFAULT_CLIENT_RETRY,
		UploadChunkedUploadThreshold: API_DEFAULT_UPLOAD_CHUNKED_UPLOAD_THRESHOLD,
		UploadChunkedUploadChunkSize: API_DEFAULT_UPLOAD_CHUNK_SIZE,
	}
}

type ApiContext struct {
	Token      string
	AsMemberId string
	Client     *http.Client
	Config     *ApiConfig
}

func (a *ApiContext) Files() files.Client {
	return &ApiFiles{
		Context: a,
	}
}

func (a *ApiContext) PatternsFile() *ApiPatternFiles {
	return &ApiPatternFiles{
		Context: a,
	}
}

type ApiFiles struct {
	Context *ApiContext
}

func (a *ApiFiles) Compat() files.Client {
	return files.New(
		dropbox.Config{
			Token:      a.Context.Token,
			AsMemberID: a.Context.AsMemberId,
		},
	)
}

func NewDefaultApiContext(token string) *ApiContext {
	config := NewDefaultApiConfig()
	return &ApiContext{
		Token:  token,
		Client: &http.Client{Timeout: config.Timeout},
		Config: config,
	}
}

func (c *ApiContext) PrepareHeader(req *http.Request) *http.Request {
	if c.AsMemberId != "" {
		req.Header.Add(API_REQ_HEADER_SELECT_USER, c.AsMemberId)
	}
	return req
}

type ApiRpcResponse struct {
	StatusCode int
	Body       []byte
}

type ApiRpcRequest struct {
	Param               interface{}
	AuthHeader          bool
	Route               string
	Context             *ApiContext
	EndpointErrorParser ApiEndpointSpecificErrorParser
}

type ApiEndpointSpecificErrorParser func([]byte) error
type ApiErrorCallback func(*http.Response, []byte)

func (a *ApiRpcRequest) requestUrl() string {
	return fmt.Sprintf("https://%s/2/%s", API_RPC_ENDPOINT, a.Route)
}

func (a *ApiRpcRequest) rpcRequest() (req *http.Request, err error) {
	url := a.requestUrl()

	// param
	requestParam, err := json.Marshal(a.Param)
	if err != nil {
		seelog.Debugf("Route[%s] Unable to marshal params. error[%s]", a.Route, err)
		return nil, err
	}

	req, err = http.NewRequest("POST", url, bytes.NewReader(requestParam))
	if err != nil {
		seelog.Debugf("Route[%s] Unable create request. error[%s]", a.Route, err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	if a.AuthHeader {
		req.Header.Add("Authorization", "Bearer "+a.Context.Token)
	}
	a.Context.PrepareHeader(req)
	return
}

func (a *ApiRpcRequest) Call() (apiRes *ApiRpcResponse, err error) {
	req, err := a.rpcRequest()
	if err != nil {
		seelog.Tracef("Route[%s] Unable to prepare request : error[%s]", a.Route, err)
		return
	}

	var lastErr error
	// call and retry
	for retry := 0; retry < a.Context.Config.Retry; retry++ {
		seelog.Tracef("Route[%s] Do try[%d of %d]", a.Route, retry+1, a.Context.Config.Retry)
		res, err := a.Context.Client.Do(req)

		if err != nil {
			seelog.Debugf("Route[%s] Transport error[%s]", a.Route, err)
			lastErr = err
			continue
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			seelog.Debugf("Route[%s] Unable to read body. error[%s]", a.Route, err)
			return nil, err
		}
		res.Body.Close()

		if res.StatusCode == http.StatusOK {
			return &ApiRpcResponse{
				StatusCode: res.StatusCode,
				Body:       body,
			}, nil
		}

		if a.Context.Config.ErrorCallback != nil {
			seelog.Tracef("Route[%s] raise error callback", a.Route)
			a.Context.Config.ErrorCallback(res, body)
		}

		switch res.StatusCode {
		case 400: // Bad input param
			err := dropbox.APIError{
				ErrorSummary: string(body),
			}
			seelog.Debugf("Route[%s] Bad input param. error[%s]", a.Route, err)
			return nil, err

		case 401: // Bad or expired token
			seelog.Debugf("Route[%s] Bad or expired token.", a.Route)
			return nil, errors.New("token err")

		case 409: // Endpoint specific error
			seelog.Debugf("Route[%s] Endpoint specific error. error[%s]", a.Route)
			return nil, a.EndpointErrorParser(body)

		case 429: // Rate limit
			retryAfter := res.Header.Get(API_RES_HEADER_RETRY_AFTER)
			retryAfterSec, err := strconv.Atoi(retryAfter)
			if err != nil {
				seelog.Debugf("Route[%s] Unable to parse '%s' header. HeaderContent[%s] error[%s]", a.Route, retryAfter, err)
				return nil, errors.New("unknown retry param")
			}
			seelog.Debugf("Route[%s] Wait for retry [%d] seconds.", retryAfterSec)
			time.Sleep(time.Duration(retryAfterSec) * time.Second)
			lastErr = dropbox.APIError{
				ErrorSummary: string(body),
			}

			continue

		default:
			// Try parse as APIError
			var apiErr dropbox.APIError
			err = json.Unmarshal(body, &apiErr)
			if err != nil {
				seelog.Debugf("Route[%s] unknown or server error. response body[%s], unmarshal err[%s]", a.Route, string(body), err)
				return nil, err
			}
			seelog.Debugf("Route[%s] unknown or server error[%s]", a.Route, err)
			return nil, apiErr
		}
	}

	return nil, lastErr
}

func (a *ApiContext) NewApiRpcRequest(route string, errParser ApiEndpointSpecificErrorParser, arg interface{}) *ApiRpcRequest {
	return &ApiRpcRequest{
		Param:               arg,
		Route:               route,
		AuthHeader:          true,
		Context:             a,
		EndpointErrorParser: errParser,
	}
}

func parseResponseFilesSearch(res *ApiRpcResponse) (r *files.SearchResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}

func parseErrorFilesSearch(body []byte) error {
	var apiErr files.SearchAPIError
	if err := json.Unmarshal(body, &apiErr); err != nil {
		return err
	}
	return apiErr
}

func parseResponseListFolder(res *ApiRpcResponse) (r *files.ListFolderResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}

func parseErrorListFolder(body []byte) error {
	var apiErr files.ListFolderAPIError
	if err := json.Unmarshal(body, &apiErr); err != nil {
		return err
	}
	return apiErr
}

func (a *ApiFiles) Search(arg *files.SearchArg) (sr *files.SearchResult, err error) {
	if res, err := a.Context.NewApiRpcRequest("files/search", parseErrorFilesSearch, arg).Call(); err != nil {
		return nil, err
	} else {
		return parseResponseFilesSearch(res)
	}
}

func (a *ApiFiles) ListFolder(arg *files.ListFolderArg) (lr *files.ListFolderResult, err error) {
	if res, err := a.Context.NewApiRpcRequest("files/list_folder", parseErrorListFolder, arg).Call(); err != nil {
		return nil, err
	} else {
		return parseResponseListFolder(res)
	}
}

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
func (a *ApiFiles) CopyV2(arg *files.RelocationArg) (res *files.RelocationResult, err error) {
	return a.Compat().CopyV2(arg)
}
func (a *ApiFiles) CreateFolder(arg *files.CreateFolderArg) (res *files.FolderMetadata, err error) {
	return a.Compat().CreateFolder(arg)
}
func (a *ApiFiles) CreateFolderV2(arg *files.CreateFolderArg) (res *files.CreateFolderResult, err error) {
	return a.Compat().CreateFolderV2(arg)
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
func (a *ApiFiles) GetMetadata(arg *files.GetMetadataArg) (res files.IsMetadata, err error) {
	return a.Compat().GetMetadata(arg)
}
func (a *ApiFiles) GetPreview(arg *files.PreviewArg) (res *files.FileMetadata, content io.ReadCloser, err error) {
	return a.Compat().GetPreview(arg)
}
func (a *ApiFiles) GetTemporaryLink(arg *files.GetTemporaryLinkArg) (res *files.GetTemporaryLinkResult, err error) {
	return a.Compat().GetTemporaryLink(arg)
}
func (a *ApiFiles) GetThumbnail(arg *files.ThumbnailArg) (res *files.FileMetadata, content io.ReadCloser, err error) {
	return a.Compat().GetThumbnail(arg)
}
func (a *ApiFiles) GetThumbnailBatch(arg *files.GetThumbnailBatchArg) (res *files.GetThumbnailBatchResult, err error) {
	return a.Compat().GetThumbnailBatch(arg)
}

//func (a *ApiFiles) ListFolder(arg *files.ListFolderArg) (res *files.ListFolderResult, err error) {
//	return a.Compat().ListFolder(arg)
//}
func (a *ApiFiles) ListFolderContinue(arg *files.ListFolderContinueArg) (res *files.ListFolderResult, err error) {
	return a.Compat().ListFolderContinue(arg)
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
func (a *ApiFiles) MoveV2(arg *files.RelocationArg) (res *files.RelocationResult, err error) {
	return a.Compat().MoveV2(arg)
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

//func (a *ApiFiles) Search(arg *files.SearchArg) (res *files.SearchResult, err error) {
//	return a.Compat().Search(arg)
//}
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

type ApiPatternFiles struct {
	Context     *ApiContext
	FilesClient files.Client
}

func (a *ApiPatternFiles) ListFolder(lfa *files.ListFolderArg) (entries []files.IsMetadata, err error) {
	seelog.Tracef("ListFolder: Path[%s]", lfa.Path)
	res, err := a.FilesClient.ListFolder(lfa)
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
		res, err = a.FilesClient.ListFolderContinue(contArg)
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

func (a *ApiPatternFiles) Upload(content io.Reader, size int64, ci *files.CommitInfo) (fm *files.FileMetadata, err error) {
	if size > a.Context.Config.UploadChunkedUploadThreshold {
		fm, err = a.filesUploadChunked(content, size, ci)
	} else {
		fm, err = a.filesUploadSingle(content, size, ci)
	}
	if fm != nil {
		seelog.Tracef("filesUpload: toPath[%s] id[%s] hash[%s]", fm.PathDisplay, fm.Id, fm.ContentHash)
	}
	return
}

func (a *ApiPatternFiles) filesUploadSingle(content io.Reader, size int64, ci *files.CommitInfo) (fm *files.FileMetadata, err error) {
	seelog.Tracef("filesUploadSingle: toPath[%s] size[%d]", ci.Path, size)

	return a.FilesClient.Upload(ci, content)
}

func (a *ApiPatternFiles) filesUploadChunked(content io.Reader, size int64, ci *files.CommitInfo) (fm *files.FileMetadata, err error) {
	seelog.Tracef("filesUploadChunked: toPath[%s] size[%d]", ci.Path, size)

	r := io.LimitReader(content, a.Context.Config.UploadChunkedUploadChunkSize)
	s, err := a.FilesClient.UploadSessionStart(files.NewUploadSessionStartArg(), r)
	if err != nil {
		seelog.Debugf("Unable to start upload session : error[%s]", err)
		return
	}

	var uploaded int64
	uploaded = a.Context.Config.UploadChunkedUploadChunkSize
	for (size - uploaded) > a.Context.Config.UploadChunkedUploadChunkSize {
		seelog.Tracef("filesUploadChunked: toPath[%s]: uploaded[%d] of size[%d]", ci.Path, uploaded, size)

		cursor := files.NewUploadSessionCursor(s.SessionId, uint64(uploaded))
		arg := files.NewUploadSessionAppendArg(cursor)
		r = io.LimitReader(content, int64(a.Context.Config.UploadChunkedUploadChunkSize))
		err = a.FilesClient.UploadSessionAppendV2(arg, r)
		if err != nil {
			seelog.Debugf("Unable to append upload session : error[%s]", err)
			return
		}
		uploaded += a.Context.Config.UploadChunkedUploadChunkSize
	}

	seelog.Tracef("filesUploadChunked: toPath[%s]: uploaded[%d] of size[%d]", ci.Path, uploaded, size)

	cursor := files.NewUploadSessionCursor(s.SessionId, uint64(uploaded))
	arg := files.NewUploadSessionFinishArg(cursor, ci)
	fm, err = a.FilesClient.UploadSessionFinish(arg, content)
	if err != nil {
		seelog.Debugf("Unable to finish upload session : error[%s]", err)
	}
	return
}
