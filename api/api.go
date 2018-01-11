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
	return a.FilesImpl()
}

func (a *ApiContext) FilesImpl() *ApiFiles {
	return &ApiFiles{
		Context: a,
	}
}

func (a *ApiContext) Team() team.Client {
	return a.TeamImpl()
}

func (a *ApiContext) TeamImpl() *ApiTeam {
	return &ApiTeam{
		Context: a,
	}
}

func (a *ApiContext) TeamLog() team_log.Client {
	return a.TeamLogImpl()
}

func (a *ApiContext) TeamLogImpl() *ApiTeamLog {
	return &ApiTeamLog{
		Context: a,
	}
}

func (a *ApiContext) FileProperties() file_properties.Client {
	return a.FilePropertiesImpl()
}

func (a *ApiContext) FilePropertiesImpl() *ApiFileProperties {
	return &ApiFileProperties{
		Context: a,
	}
}

func (a *ApiContext) FileRequests() file_requests.Client {
	return a.FileRequestsImpl()
}

func (a *ApiContext) FileRequestsImpl() *ApiFileRequests {
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
	defaultErrorParser := func(body []byte) error {
		// Try parse as APIError
		var apiErr dropbox.APIError
		err = json.Unmarshal(body, &apiErr)
		if err != nil {
			seelog.Debugf("Route[%s] unknown or server error. response body[%s], unmarshal err[%s]", a.Route, string(body), err)
			return err
		}
		seelog.Debugf("Route[%s] unknown or server error[%s]", a.Route, err)
		return apiErr
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
			if a.EndpointErrorParser != nil {
				return nil, a.EndpointErrorParser(body)
			} else {
				return nil, defaultErrorParser(body)
			}

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
			return nil, defaultErrorParser(body)
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
