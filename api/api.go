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
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/file_requests"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/paper"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/sharing"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/team"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/team_log"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/users"
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

func (a *ApiContext) compatConfig() dropbox.Config {
	return dropbox.Config{
		Token:      a.Token,
		AsMemberID: a.AsMemberId,
	}
}

func (a *ApiContext) Files() files.Client {
	return &ApiFiles{
		Context: a,
	}
}

func (a *ApiContext) Team() team.Client {
	return &ApiTeam{
		Context: a,
	}
}

func (a *ApiContext) TeamLog() team_log.Client {
	return &ApiTeamLog{
		Context: a,
	}
}

func (a *ApiContext) TeamLogImpl() *ApiTeamLog {
	return &ApiTeamLog{
		Context: a,
	}
}

func (a *ApiContext) FileProperties() file_properties.Client {
	return &ApiFileProperties{
		Context: a,
	}
}
func (a *ApiContext) FileRequests() file_requests.Client {
	return &ApiFileRequests{
		Context: a,
	}
}
func (a *ApiContext) Paper() paper.Client {
	return &ApiPaper{
		Context: a,
	}
}
func (a *ApiContext) Sharing() sharing.Client {
	return &ApiSharing{
		Context: a,
	}
}
func (a *ApiContext) Users() users.Client {
	return &ApiUsers{
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
	return files.New(a.Context.compatConfig())
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

type ApiTeam struct {
	Context *ApiContext
}

func (a *ApiTeam) Compat() team.Client {
	return team.New(a.Context.compatConfig())
}

func (a *ApiTeam) DevicesListMemberDevices(arg *team.ListMemberDevicesArg) (res *team.ListMemberDevicesResult, err error) {
	return a.Compat().DevicesListMemberDevices(arg)
}
func (a *ApiTeam) DevicesListMembersDevices(arg *team.ListMembersDevicesArg) (res *team.ListMembersDevicesResult, err error) {
	return a.Compat().DevicesListMembersDevices(arg)
}
func (a *ApiTeam) DevicesListTeamDevices(arg *team.ListTeamDevicesArg) (res *team.ListTeamDevicesResult, err error) {
	return a.Compat().DevicesListTeamDevices(arg)
}
func (a *ApiTeam) DevicesRevokeDeviceSession(arg *team.RevokeDeviceSessionArg) (err error) {
	return a.Compat().DevicesRevokeDeviceSession(arg)
}
func (a *ApiTeam) DevicesRevokeDeviceSessionBatch(arg *team.RevokeDeviceSessionBatchArg) (res *team.RevokeDeviceSessionBatchResult, err error) {
	return a.Compat().DevicesRevokeDeviceSessionBatch(arg)
}
func (a *ApiTeam) FeaturesGetValues(arg *team.FeaturesGetValuesBatchArg) (res *team.FeaturesGetValuesBatchResult, err error) {
	return a.Compat().FeaturesGetValues(arg)
}
func (a *ApiTeam) GetInfo() (res *team.TeamGetInfoResult, err error) {
	return a.Compat().GetInfo()
}
func (a *ApiTeam) GroupsCreate(arg *team.GroupCreateArg) (res *team.GroupFullInfo, err error) {
	return a.Compat().GroupsCreate(arg)
}
func (a *ApiTeam) GroupsDelete(arg *team.GroupSelector) (res *async.LaunchEmptyResult, err error) {
	return a.Compat().GroupsDelete(arg)
}
func (a *ApiTeam) GroupsGetInfo(arg *team.GroupsSelector) (res []*team.GroupsGetInfoItem, err error) {
	return a.Compat().GroupsGetInfo(arg)
}
func (a *ApiTeam) GroupsJobStatusGet(arg *async.PollArg) (res *async.PollEmptyResult, err error) {
	return a.Compat().GroupsJobStatusGet(arg)
}
func (a *ApiTeam) GroupsList(arg *team.GroupsListArg) (res *team.GroupsListResult, err error) {
	return a.Compat().GroupsList(arg)
}
func (a *ApiTeam) GroupsListContinue(arg *team.GroupsListContinueArg) (res *team.GroupsListResult, err error) {
	return a.Compat().GroupsListContinue(arg)
}
func (a *ApiTeam) GroupsMembersAdd(arg *team.GroupMembersAddArg) (res *team.GroupMembersChangeResult, err error) {
	return a.Compat().GroupsMembersAdd(arg)
}
func (a *ApiTeam) GroupsMembersList(arg *team.GroupsMembersListArg) (res *team.GroupsMembersListResult, err error) {
	return a.Compat().GroupsMembersList(arg)
}
func (a *ApiTeam) GroupsMembersListContinue(arg *team.GroupsMembersListContinueArg) (res *team.GroupsMembersListResult, err error) {
	return a.Compat().GroupsMembersListContinue(arg)
}
func (a *ApiTeam) GroupsMembersRemove(arg *team.GroupMembersRemoveArg) (res *team.GroupMembersChangeResult, err error) {
	return a.Compat().GroupsMembersRemove(arg)
}
func (a *ApiTeam) GroupsMembersSetAccessType(arg *team.GroupMembersSetAccessTypeArg) (res []*team.GroupsGetInfoItem, err error) {
	return a.Compat().GroupsMembersSetAccessType(arg)
}
func (a *ApiTeam) GroupsUpdate(arg *team.GroupUpdateArgs) (res *team.GroupFullInfo, err error) {
	return a.Compat().GroupsUpdate(arg)
}
func (a *ApiTeam) LinkedAppsListMemberLinkedApps(arg *team.ListMemberAppsArg) (res *team.ListMemberAppsResult, err error) {
	return a.Compat().LinkedAppsListMemberLinkedApps(arg)
}
func (a *ApiTeam) LinkedAppsListMembersLinkedApps(arg *team.ListMembersAppsArg) (res *team.ListMembersAppsResult, err error) {
	return a.Compat().LinkedAppsListMembersLinkedApps(arg)
}
func (a *ApiTeam) LinkedAppsListTeamLinkedApps(arg *team.ListTeamAppsArg) (res *team.ListTeamAppsResult, err error) {
	return a.Compat().LinkedAppsListTeamLinkedApps(arg)
}
func (a *ApiTeam) LinkedAppsRevokeLinkedApp(arg *team.RevokeLinkedApiAppArg) (err error) {
	return a.Compat().LinkedAppsRevokeLinkedApp(arg)
}
func (a *ApiTeam) LinkedAppsRevokeLinkedAppBatch(arg *team.RevokeLinkedApiAppBatchArg) (res *team.RevokeLinkedAppBatchResult, err error) {
	return a.Compat().LinkedAppsRevokeLinkedAppBatch(arg)
}
func (a *ApiTeam) MemberSpaceLimitsExcludedUsersAdd(arg *team.ExcludedUsersUpdateArg) (res *team.ExcludedUsersUpdateResult, err error) {
	return a.Compat().MemberSpaceLimitsExcludedUsersAdd(arg)
}
func (a *ApiTeam) MemberSpaceLimitsExcludedUsersList(arg *team.ExcludedUsersListArg) (res *team.ExcludedUsersListResult, err error) {
	return a.Compat().MemberSpaceLimitsExcludedUsersList(arg)
}
func (a *ApiTeam) MemberSpaceLimitsExcludedUsersListContinue(arg *team.ExcludedUsersListContinueArg) (res *team.ExcludedUsersListResult, err error) {
	return a.Compat().MemberSpaceLimitsExcludedUsersListContinue(arg)
}
func (a *ApiTeam) MemberSpaceLimitsExcludedUsersRemove(arg *team.ExcludedUsersUpdateArg) (res *team.ExcludedUsersUpdateResult, err error) {
	return a.Compat().MemberSpaceLimitsExcludedUsersRemove(arg)
}
func (a *ApiTeam) MemberSpaceLimitsGetCustomQuota(arg *team.CustomQuotaUsersArg) (res []*team.CustomQuotaResult, err error) {
	return a.Compat().MemberSpaceLimitsGetCustomQuota(arg)
}
func (a *ApiTeam) MemberSpaceLimitsRemoveCustomQuota(arg *team.CustomQuotaUsersArg) (res []*team.RemoveCustomQuotaResult, err error) {
	return a.Compat().MemberSpaceLimitsRemoveCustomQuota(arg)
}
func (a *ApiTeam) MemberSpaceLimitsSetCustomQuota(arg *team.SetCustomQuotaArg) (res []*team.CustomQuotaResult, err error) {
	return a.Compat().MemberSpaceLimitsSetCustomQuota(arg)
}
func (a *ApiTeam) MembersAdd(arg *team.MembersAddArg) (res *team.MembersAddLaunch, err error) {
	return a.Compat().MembersAdd(arg)
}
func (a *ApiTeam) MembersAddJobStatusGet(arg *async.PollArg) (res *team.MembersAddJobStatus, err error) {
	return a.Compat().MembersAddJobStatusGet(arg)
}
func (a *ApiTeam) MembersGetInfo(arg *team.MembersGetInfoArgs) (res []*team.MembersGetInfoItem, err error) {
	return a.Compat().MembersGetInfo(arg)
}
func (a *ApiTeam) MembersList(arg *team.MembersListArg) (res *team.MembersListResult, err error) {
	return a.Compat().MembersList(arg)
}
func (a *ApiTeam) MembersListContinue(arg *team.MembersListContinueArg) (res *team.MembersListResult, err error) {
	return a.Compat().MembersListContinue(arg)
}
func (a *ApiTeam) MembersRecover(arg *team.MembersRecoverArg) (err error) {
	return a.Compat().MembersRecover(arg)
}
func (a *ApiTeam) MembersRemove(arg *team.MembersRemoveArg) (res *async.LaunchEmptyResult, err error) {
	return a.Compat().MembersRemove(arg)
}
func (a *ApiTeam) MembersRemoveJobStatusGet(arg *async.PollArg) (res *async.PollEmptyResult, err error) {
	return a.Compat().MembersRemoveJobStatusGet(arg)
}
func (a *ApiTeam) MembersSendWelcomeEmail(arg *team.UserSelectorArg) (err error) {
	return a.Compat().MembersSendWelcomeEmail(arg)
}
func (a *ApiTeam) MembersSetAdminPermissions(arg *team.MembersSetPermissionsArg) (res *team.MembersSetPermissionsResult, err error) {
	return a.Compat().MembersSetAdminPermissions(arg)
}
func (a *ApiTeam) MembersSetProfile(arg *team.MembersSetProfileArg) (res *team.TeamMemberInfo, err error) {
	return a.Compat().MembersSetProfile(arg)
}
func (a *ApiTeam) MembersSuspend(arg *team.MembersDeactivateArg) (err error) {
	return a.Compat().MembersSuspend(arg)
}
func (a *ApiTeam) MembersUnsuspend(arg *team.MembersUnsuspendArg) (err error) {
	return a.Compat().MembersUnsuspend(arg)
}
func (a *ApiTeam) NamespacesList(arg *team.TeamNamespacesListArg) (res *team.TeamNamespacesListResult, err error) {
	return a.Compat().NamespacesList(arg)
}
func (a *ApiTeam) NamespacesListContinue(arg *team.TeamNamespacesListContinueArg) (res *team.TeamNamespacesListResult, err error) {
	return a.Compat().NamespacesListContinue(arg)
}
func (a *ApiTeam) PropertiesTemplateAdd(arg *file_properties.AddTemplateArg) (res *file_properties.AddTemplateResult, err error) {
	return a.Compat().PropertiesTemplateAdd(arg)
}
func (a *ApiTeam) PropertiesTemplateGet(arg *file_properties.GetTemplateArg) (res *file_properties.GetTemplateResult, err error) {
	return a.Compat().PropertiesTemplateGet(arg)
}
func (a *ApiTeam) PropertiesTemplateList() (res *file_properties.ListTemplateResult, err error) {
	return a.Compat().PropertiesTemplateList()
}
func (a *ApiTeam) PropertiesTemplateUpdate(arg *file_properties.UpdateTemplateArg) (res *file_properties.UpdateTemplateResult, err error) {
	return a.Compat().PropertiesTemplateUpdate(arg)
}
func (a *ApiTeam) ReportsGetActivity(arg *team.DateRange) (res *team.GetActivityReport, err error) {
	return a.Compat().ReportsGetActivity(arg)
}
func (a *ApiTeam) ReportsGetDevices(arg *team.DateRange) (res *team.GetDevicesReport, err error) {
	return a.Compat().ReportsGetDevices(arg)
}
func (a *ApiTeam) ReportsGetMembership(arg *team.DateRange) (res *team.GetMembershipReport, err error) {
	return a.Compat().ReportsGetMembership(arg)
}
func (a *ApiTeam) ReportsGetStorage(arg *team.DateRange) (res *team.GetStorageReport, err error) {
	return a.Compat().ReportsGetStorage(arg)
}
func (a *ApiTeam) TeamFolderActivate(arg *team.TeamFolderIdArg) (res *team.TeamFolderMetadata, err error) {
	return a.Compat().TeamFolderActivate(arg)
}
func (a *ApiTeam) TeamFolderArchive(arg *team.TeamFolderArchiveArg) (res *team.TeamFolderArchiveLaunch, err error) {
	return a.Compat().TeamFolderArchive(arg)
}
func (a *ApiTeam) TeamFolderArchiveCheck(arg *async.PollArg) (res *team.TeamFolderArchiveJobStatus, err error) {
	return a.Compat().TeamFolderArchiveCheck(arg)
}
func (a *ApiTeam) TeamFolderCreate(arg *team.TeamFolderCreateArg) (res *team.TeamFolderMetadata, err error) {
	return a.Compat().TeamFolderCreate(arg)
}
func (a *ApiTeam) TeamFolderGetInfo(arg *team.TeamFolderIdListArg) (res []*team.TeamFolderGetInfoItem, err error) {
	return a.Compat().TeamFolderGetInfo(arg)
}
func (a *ApiTeam) TeamFolderList(arg *team.TeamFolderListArg) (res *team.TeamFolderListResult, err error) {
	return a.Compat().TeamFolderList(arg)
}
func (a *ApiTeam) TeamFolderListContinue(arg *team.TeamFolderListContinueArg) (res *team.TeamFolderListResult, err error) {
	return a.Compat().TeamFolderListContinue(arg)
}
func (a *ApiTeam) TeamFolderPermanentlyDelete(arg *team.TeamFolderIdArg) (err error) {
	return a.Compat().TeamFolderPermanentlyDelete(arg)
}
func (a *ApiTeam) TeamFolderRename(arg *team.TeamFolderRenameArg) (res *team.TeamFolderMetadata, err error) {
	return a.Compat().TeamFolderRename(arg)
}
func (a *ApiTeam) TokenGetAuthenticatedAdmin() (res *team.TokenGetAuthenticatedAdminResult, err error) {
	return a.Compat().TokenGetAuthenticatedAdmin()
}

type ApiTeamLog struct {
	Context *ApiContext
}

func (a *ApiTeamLog) Compat() team_log.Client {
	return team_log.New(a.Context.compatConfig())
}

func (a *ApiTeamLog) GetEvents(arg *team_log.GetTeamEventsArg) (res *team_log.GetTeamEventsResult, err error) {
	return a.Compat().GetEvents(arg)
}

func (a *ApiTeamLog) GetEventsContinue(arg *team_log.GetTeamEventsContinueArg) (res *team_log.GetTeamEventsResult, err error) {
	return a.Compat().GetEventsContinue(arg)
}

type GetTeamEventsRawResult struct {
	Cursor  string            `json:"cursor"`
	HasMore bool              `json:"has_more"`
	Events  []json.RawMessage `json:"events,omitempty"`
}

func parseGetTeamEventsRawResult(res *ApiRpcResponse) (r *GetTeamEventsRawResult, err error) {
	err = json.Unmarshal(res.Body, &r)
	return
}

func parseGetEventsAPIError(body []byte) error {
	var apiErr team_log.GetEventsAPIError
	if err := json.Unmarshal(body, &apiErr); err != nil {
		return err
	}
	return apiErr
}

func (a *ApiTeamLog) RawGetEvents(arg *team_log.GetTeamEventsArg) (r *GetTeamEventsRawResult, err error) {
	if res, err := a.Context.NewApiRpcRequest("team_log/get_events", parseGetEventsAPIError, arg).Call(); err != nil {
		return nil, err
	} else {
		return parseGetTeamEventsRawResult(res)
	}
}

func (a *ApiTeamLog) RawGetEventsContinue(arg *team_log.GetTeamEventsContinueArg) (res *GetTeamEventsRawResult, err error) {
	if res, err := a.Context.NewApiRpcRequest("team_log/get_events/continue", parseGetEventsAPIError, arg).Call(); err != nil {
		return nil, err
	} else {
		return parseGetTeamEventsRawResult(res)
	}
}

type ApiFileProperties struct {
	Context *ApiContext
}

func (a *ApiFileProperties) Compat() file_properties.Client {
	return file_properties.New(a.Context.compatConfig())
}

func (a *ApiFileProperties) PropertiesAdd(arg *file_properties.AddPropertiesArg) (err error) {
	return a.Compat().PropertiesAdd(arg)
}
func (a *ApiFileProperties) PropertiesOverwrite(arg *file_properties.OverwritePropertyGroupArg) (err error) {
	return a.Compat().PropertiesOverwrite(arg)
}
func (a *ApiFileProperties) PropertiesRemove(arg *file_properties.RemovePropertiesArg) (err error) {
	return a.Compat().PropertiesRemove(arg)
}
func (a *ApiFileProperties) PropertiesSearch(arg *file_properties.PropertiesSearchArg) (res *file_properties.PropertiesSearchResult, err error) {
	return a.Compat().PropertiesSearch(arg)
}
func (a *ApiFileProperties) PropertiesSearchContinue(arg *file_properties.PropertiesSearchContinueArg) (res *file_properties.PropertiesSearchResult, err error) {
	return a.Compat().PropertiesSearchContinue(arg)
}
func (a *ApiFileProperties) PropertiesUpdate(arg *file_properties.UpdatePropertiesArg) (err error) {
	return a.Compat().PropertiesUpdate(arg)
}
func (a *ApiFileProperties) TemplatesAddForTeam(arg *file_properties.AddTemplateArg) (res *file_properties.AddTemplateResult, err error) {
	return a.Compat().TemplatesAddForTeam(arg)
}
func (a *ApiFileProperties) TemplatesAddForUser(arg *file_properties.AddTemplateArg) (res *file_properties.AddTemplateResult, err error) {
	return a.Compat().TemplatesAddForUser(arg)
}
func (a *ApiFileProperties) TemplatesGetForTeam(arg *file_properties.GetTemplateArg) (res *file_properties.GetTemplateResult, err error) {
	return a.Compat().TemplatesGetForTeam(arg)
}
func (a *ApiFileProperties) TemplatesGetForUser(arg *file_properties.GetTemplateArg) (res *file_properties.GetTemplateResult, err error) {
	return a.Compat().TemplatesGetForUser(arg)
}
func (a *ApiFileProperties) TemplatesListForTeam() (res *file_properties.ListTemplateResult, err error) {
	return a.Compat().TemplatesListForTeam()
}
func (a *ApiFileProperties) TemplatesListForUser() (res *file_properties.ListTemplateResult, err error) {
	return a.Compat().TemplatesListForUser()
}
func (a *ApiFileProperties) TemplatesRemoveForTeam(arg *file_properties.RemoveTemplateArg) (err error) {
	return a.Compat().TemplatesRemoveForTeam(arg)
}
func (a *ApiFileProperties) TemplatesRemoveForUser(arg *file_properties.RemoveTemplateArg) (err error) {
	return a.Compat().TemplatesRemoveForUser(arg)
}
func (a *ApiFileProperties) TemplatesUpdateForTeam(arg *file_properties.UpdateTemplateArg) (res *file_properties.UpdateTemplateResult, err error) {
	return a.Compat().TemplatesUpdateForTeam(arg)
}
func (a *ApiFileProperties) TemplatesUpdateForUser(arg *file_properties.UpdateTemplateArg) (res *file_properties.UpdateTemplateResult, err error) {
	return a.Compat().TemplatesUpdateForUser(arg)
}

type ApiFileRequests struct {
	Context *ApiContext
}

func (a *ApiFileRequests) Compat() file_requests.Client {
	return file_requests.New(a.Context.compatConfig())
}

func (a *ApiFileRequests) Create(arg *file_requests.CreateFileRequestArgs) (res *file_requests.FileRequest, err error) {
	return a.Compat().Create(arg)
}
func (a *ApiFileRequests) Get(arg *file_requests.GetFileRequestArgs) (res *file_requests.FileRequest, err error) {
	return a.Compat().Get(arg)
}
func (a *ApiFileRequests) List() (res *file_requests.ListFileRequestsResult, err error) {
	return a.Compat().List()
}
func (a *ApiFileRequests) Update(arg *file_requests.UpdateFileRequestArgs) (res *file_requests.FileRequest, err error) {
	return a.Compat().Update(arg)
}

type ApiPaper struct {
	Context *ApiContext
}

func (a *ApiPaper) Compat() paper.Client {
	return paper.New(a.Context.compatConfig())
}

func (a *ApiPaper) DocsArchive(arg *paper.RefPaperDoc) (err error) {
	return a.Compat().DocsArchive(arg)
}
func (a *ApiPaper) DocsCreate(arg *paper.PaperDocCreateArgs, content io.Reader) (res *paper.PaperDocCreateUpdateResult, err error) {
	return a.Compat().DocsCreate(arg, content)
}
func (a *ApiPaper) DocsDownload(arg *paper.PaperDocExport) (res *paper.PaperDocExportResult, content io.ReadCloser, err error) {
	return a.Compat().DocsDownload(arg)
}
func (a *ApiPaper) DocsFolderUsersList(arg *paper.ListUsersOnFolderArgs) (res *paper.ListUsersOnFolderResponse, err error) {
	return a.Compat().DocsFolderUsersList(arg)
}
func (a *ApiPaper) DocsFolderUsersListContinue(arg *paper.ListUsersOnFolderContinueArgs) (res *paper.ListUsersOnFolderResponse, err error) {
	return a.Compat().DocsFolderUsersListContinue(arg)
}
func (a *ApiPaper) DocsGetFolderInfo(arg *paper.RefPaperDoc) (res *paper.FoldersContainingPaperDoc, err error) {
	return a.Compat().DocsGetFolderInfo(arg)
}
func (a *ApiPaper) DocsList(arg *paper.ListPaperDocsArgs) (res *paper.ListPaperDocsResponse, err error) {
	return a.Compat().DocsList(arg)
}
func (a *ApiPaper) DocsListContinue(arg *paper.ListPaperDocsContinueArgs) (res *paper.ListPaperDocsResponse, err error) {
	return a.Compat().DocsListContinue(arg)
}
func (a *ApiPaper) DocsPermanentlyDelete(arg *paper.RefPaperDoc) (err error) {
	return a.Compat().DocsPermanentlyDelete(arg)
}
func (a *ApiPaper) DocsSharingPolicyGet(arg *paper.RefPaperDoc) (res *paper.SharingPolicy, err error) {
	return a.Compat().DocsSharingPolicyGet(arg)
}
func (a *ApiPaper) DocsSharingPolicySet(arg *paper.PaperDocSharingPolicy) (err error) {
	return a.Compat().DocsSharingPolicySet(arg)
}
func (a *ApiPaper) DocsUpdate(arg *paper.PaperDocUpdateArgs, content io.Reader) (res *paper.PaperDocCreateUpdateResult, err error) {
	return a.Compat().DocsUpdate(arg, content)
}
func (a *ApiPaper) DocsUsersAdd(arg *paper.AddPaperDocUser) (res []*paper.AddPaperDocUserMemberResult, err error) {
	return a.Compat().DocsUsersAdd(arg)
}
func (a *ApiPaper) DocsUsersList(arg *paper.ListUsersOnPaperDocArgs) (res *paper.ListUsersOnPaperDocResponse, err error) {
	return a.Compat().DocsUsersList(arg)
}
func (a *ApiPaper) DocsUsersListContinue(arg *paper.ListUsersOnPaperDocContinueArgs) (res *paper.ListUsersOnPaperDocResponse, err error) {
	return a.Compat().DocsUsersListContinue(arg)
}
func (a *ApiPaper) DocsUsersRemove(arg *paper.RemovePaperDocUser) (err error) {
	return a.Compat().DocsUsersRemove(arg)
}

type ApiSharing struct {
	Context *ApiContext
}

func (a *ApiSharing) Compat() sharing.Client {
	return sharing.New(a.Context.compatConfig())
}

func (a *ApiSharing) AddFileMember(arg *sharing.AddFileMemberArgs) (res []*sharing.FileMemberActionResult, err error) {
	return a.Compat().AddFileMember(arg)
}
func (a *ApiSharing) AddFolderMember(arg *sharing.AddFolderMemberArg) (err error) {
	return a.Compat().AddFolderMember(arg)
}
func (a *ApiSharing) ChangeFileMemberAccess(arg *sharing.ChangeFileMemberAccessArgs) (res *sharing.FileMemberActionResult, err error) {
	return a.Compat().ChangeFileMemberAccess(arg)
}
func (a *ApiSharing) CheckJobStatus(arg *async.PollArg) (res *sharing.JobStatus, err error) {
	return a.Compat().CheckJobStatus(arg)
}
func (a *ApiSharing) CheckRemoveMemberJobStatus(arg *async.PollArg) (res *sharing.RemoveMemberJobStatus, err error) {
	return a.Compat().CheckRemoveMemberJobStatus(arg)
}
func (a *ApiSharing) CheckShareJobStatus(arg *async.PollArg) (res *sharing.ShareFolderJobStatus, err error) {
	return a.Compat().CheckShareJobStatus(arg)
}
func (a *ApiSharing) CreateSharedLink(arg *sharing.CreateSharedLinkArg) (res *sharing.PathLinkMetadata, err error) {
	return a.Compat().CreateSharedLink(arg)
}
func (a *ApiSharing) CreateSharedLinkWithSettings(arg *sharing.CreateSharedLinkWithSettingsArg) (res sharing.IsSharedLinkMetadata, err error) {
	return a.Compat().CreateSharedLinkWithSettings(arg)
}
func (a *ApiSharing) GetFileMetadata(arg *sharing.GetFileMetadataArg) (res *sharing.SharedFileMetadata, err error) {
	return a.Compat().GetFileMetadata(arg)
}
func (a *ApiSharing) GetFileMetadataBatch(arg *sharing.GetFileMetadataBatchArg) (res []*sharing.GetFileMetadataBatchResult, err error) {
	return a.Compat().GetFileMetadataBatch(arg)
}
func (a *ApiSharing) GetFolderMetadata(arg *sharing.GetMetadataArgs) (res *sharing.SharedFolderMetadata, err error) {
	return a.Compat().GetFolderMetadata(arg)
}
func (a *ApiSharing) GetSharedLinkFile(arg *sharing.GetSharedLinkMetadataArg) (res sharing.IsSharedLinkMetadata, content io.ReadCloser, err error) {
	return a.Compat().GetSharedLinkFile(arg)
}
func (a *ApiSharing) GetSharedLinkMetadata(arg *sharing.GetSharedLinkMetadataArg) (res sharing.IsSharedLinkMetadata, err error) {
	return a.Compat().GetSharedLinkMetadata(arg)
}
func (a *ApiSharing) GetSharedLinks(arg *sharing.GetSharedLinksArg) (res *sharing.GetSharedLinksResult, err error) {
	return a.Compat().GetSharedLinks(arg)
}
func (a *ApiSharing) ListFileMembers(arg *sharing.ListFileMembersArg) (res *sharing.SharedFileMembers, err error) {
	return a.Compat().ListFileMembers(arg)
}
func (a *ApiSharing) ListFileMembersBatch(arg *sharing.ListFileMembersBatchArg) (res []*sharing.ListFileMembersBatchResult, err error) {
	return a.Compat().ListFileMembersBatch(arg)
}
func (a *ApiSharing) ListFileMembersContinue(arg *sharing.ListFileMembersContinueArg) (res *sharing.SharedFileMembers, err error) {
	return a.Compat().ListFileMembersContinue(arg)
}
func (a *ApiSharing) ListFolderMembers(arg *sharing.ListFolderMembersArgs) (res *sharing.SharedFolderMembers, err error) {
	return a.Compat().ListFolderMembers(arg)
}
func (a *ApiSharing) ListFolderMembersContinue(arg *sharing.ListFolderMembersContinueArg) (res *sharing.SharedFolderMembers, err error) {
	return a.Compat().ListFolderMembersContinue(arg)
}
func (a *ApiSharing) ListFolders(arg *sharing.ListFoldersArgs) (res *sharing.ListFoldersResult, err error) {
	return a.Compat().ListFolders(arg)
}
func (a *ApiSharing) ListFoldersContinue(arg *sharing.ListFoldersContinueArg) (res *sharing.ListFoldersResult, err error) {
	return a.Compat().ListFoldersContinue(arg)
}
func (a *ApiSharing) ListMountableFolders(arg *sharing.ListFoldersArgs) (res *sharing.ListFoldersResult, err error) {
	return a.Compat().ListMountableFolders(arg)
}
func (a *ApiSharing) ListMountableFoldersContinue(arg *sharing.ListFoldersContinueArg) (res *sharing.ListFoldersResult, err error) {
	return a.Compat().ListMountableFoldersContinue(arg)
}
func (a *ApiSharing) ListReceivedFiles(arg *sharing.ListFilesArg) (res *sharing.ListFilesResult, err error) {
	return a.Compat().ListReceivedFiles(arg)
}
func (a *ApiSharing) ListReceivedFilesContinue(arg *sharing.ListFilesContinueArg) (res *sharing.ListFilesResult, err error) {
	return a.Compat().ListReceivedFilesContinue(arg)
}
func (a *ApiSharing) ListSharedLinks(arg *sharing.ListSharedLinksArg) (res *sharing.ListSharedLinksResult, err error) {
	return a.Compat().ListSharedLinks(arg)
}
func (a *ApiSharing) ModifySharedLinkSettings(arg *sharing.ModifySharedLinkSettingsArgs) (res sharing.IsSharedLinkMetadata, err error) {
	return a.Compat().ModifySharedLinkSettings(arg)
}
func (a *ApiSharing) MountFolder(arg *sharing.MountFolderArg) (res *sharing.SharedFolderMetadata, err error) {
	return a.Compat().MountFolder(arg)
}
func (a *ApiSharing) RelinquishFileMembership(arg *sharing.RelinquishFileMembershipArg) (err error) {
	return a.Compat().RelinquishFileMembership(arg)
}
func (a *ApiSharing) RelinquishFolderMembership(arg *sharing.RelinquishFolderMembershipArg) (res *async.LaunchEmptyResult, err error) {
	return a.Compat().RelinquishFolderMembership(arg)
}
func (a *ApiSharing) RemoveFileMember(arg *sharing.RemoveFileMemberArg) (res *sharing.FileMemberActionIndividualResult, err error) {
	return a.Compat().RemoveFileMember(arg)
}
func (a *ApiSharing) RemoveFileMember2(arg *sharing.RemoveFileMemberArg) (res *sharing.FileMemberRemoveActionResult, err error) {
	return a.Compat().RemoveFileMember2(arg)
}
func (a *ApiSharing) RemoveFolderMember(arg *sharing.RemoveFolderMemberArg) (res *async.LaunchResultBase, err error) {
	return a.Compat().RemoveFolderMember(arg)
}
func (a *ApiSharing) RevokeSharedLink(arg *sharing.RevokeSharedLinkArg) (err error) {
	return a.Compat().RevokeSharedLink(arg)
}
func (a *ApiSharing) ShareFolder(arg *sharing.ShareFolderArg) (res *sharing.ShareFolderLaunch, err error) {
	return a.Compat().ShareFolder(arg)
}
func (a *ApiSharing) TransferFolder(arg *sharing.TransferFolderArg) (err error) {
	return a.Compat().TransferFolder(arg)
}
func (a *ApiSharing) UnmountFolder(arg *sharing.UnmountFolderArg) (err error) {
	return a.Compat().UnmountFolder(arg)
}
func (a *ApiSharing) UnshareFile(arg *sharing.UnshareFileArg) (err error) {
	return a.Compat().UnshareFile(arg)
}
func (a *ApiSharing) UnshareFolder(arg *sharing.UnshareFolderArg) (res *async.LaunchEmptyResult, err error) {
	return a.Compat().UnshareFolder(arg)
}
func (a *ApiSharing) UpdateFileMember(arg *sharing.UpdateFileMemberArgs) (res *sharing.MemberAccessLevelResult, err error) {
	return a.Compat().UpdateFileMember(arg)
}
func (a *ApiSharing) UpdateFolderMember(arg *sharing.UpdateFolderMemberArg) (res *sharing.MemberAccessLevelResult, err error) {
	return a.Compat().UpdateFolderMember(arg)
}
func (a *ApiSharing) UpdateFolderPolicy(arg *sharing.UpdateFolderPolicyArg) (res *sharing.SharedFolderMetadata, err error) {
	return a.Compat().UpdateFolderPolicy(arg)
}

type ApiUsers struct {
	Context *ApiContext
}

func (a *ApiUsers) Compat() users.Client {
	return users.New(a.Context.compatConfig())
}

func (a *ApiUsers) GetAccount(arg *users.GetAccountArg) (res *users.BasicAccount, err error) {
	return a.Compat().GetAccount(arg)
}
func (a *ApiUsers) GetAccountBatch(arg *users.GetAccountBatchArg) (res []*users.BasicAccount, err error) {
	return a.Compat().GetAccountBatch(arg)
}
func (a *ApiUsers) GetCurrentAccount() (res *users.FullAccount, err error) {
	return a.Compat().GetCurrentAccount()
}
func (a *ApiUsers) GetSpaceUsage() (res *users.SpaceUsage, err error) {
	return a.Compat().GetSpaceUsage()
}
