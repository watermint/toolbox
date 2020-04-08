package dbx_context_impl

import (
	"bufio"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_async"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_list"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_request"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/infra/api/api_async"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_list"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/api/api_response"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/network/nw_bandwidth"
	"github.com/watermint/toolbox/infra/network/nw_monitor"
	"github.com/watermint/toolbox/infra/util/ut_io"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func New(ctl app_control.Control, token api_auth.Context) dbx_context.Context {
	c := &ccImpl{
		ctl:            ctl,
		token:          token,
		noRetryOnError: false,
	}
	return c
}

type ccImpl struct {
	ctl            app_control.Control
	token          api_auth.Context
	asMemberId     string
	asAdminId      string
	basePath       dbx_context.PathRoot
	noRetryOnError bool
	hashComputed   string
	hashMutex      sync.Mutex
}

func (z *ccImpl) MakeResponse(req *http.Request, res *http.Response) (api_response.Response, error) {
	return NewResponse(z, req, res)
}

func (z *ccImpl) Capture() *zap.Logger {
	return z.ctl.Capture()
}

func (z *ccImpl) IsNoRetry() bool {
	return z.noRetryOnError
}

func (z *ccImpl) Log() *zap.Logger {
	return z.ctl.Log()
}

func (z *ccImpl) Post(endpoint string) api_request.Request {
	return dbx_request.NewPpcRequest(
		z,
		endpoint,
		z.asMemberId,
		z.asAdminId,
		z.basePath,
		z.token,
		dbx_request.RpcEndpoint,
	)
}

func (z *ccImpl) Notify(endpoint string) api_request.Request {
	return dbx_request.NewPpcRequest(
		z,
		endpoint,
		z.asMemberId,
		z.asAdminId,
		z.basePath,
		z.token,
		dbx_request.NotifyEndpoint,
	)
}

func (z *ccImpl) List(endpoint string) api_list.List {
	return dbx_list.New(z, endpoint, z.asMemberId, z.asAdminId, z.basePath)
}

func (z *ccImpl) Async(endpoint string) api_async.Async {
	return dbx_async.New(z, endpoint, z.asMemberId, z.asAdminId, z.basePath)
}

func (z *ccImpl) Upload(endpoint string, content ut_io.ReadRewinder) api_request.Request {
	return dbx_request.NewUploadRequest(
		z,
		endpoint,
		content,
		z.asMemberId,
		z.asAdminId,
		z.basePath,
		z.token,
	)
}

func (z *ccImpl) Download(endpoint string) api_request.Request {
	return dbx_request.NewDownloadRequest(
		z,
		endpoint,
		z.asMemberId,
		z.asAdminId,
		z.basePath,
		z.token,
	)
}

func (z *ccImpl) AsMemberId(teamMemberId string) dbx_context.Context {
	return &ccImpl{
		ctl:            z.ctl,
		token:          z.token,
		noRetryOnError: z.noRetryOnError,
		asMemberId:     teamMemberId,
		asAdminId:      "",
		basePath:       z.basePath,
	}
}

func (z *ccImpl) AsAdminId(teamMemberId string) dbx_context.Context {
	return &ccImpl{
		ctl:            z.ctl,
		token:          z.token,
		noRetryOnError: z.noRetryOnError,
		asMemberId:     "",
		asAdminId:      teamMemberId,
		basePath:       z.basePath,
	}
}

func (z *ccImpl) WithPath(pathRoot dbx_context.PathRoot) dbx_context.Context {
	return &ccImpl{
		ctl:            z.ctl,
		token:          z.token,
		noRetryOnError: z.noRetryOnError,
		asMemberId:     z.asMemberId,
		asAdminId:      z.asAdminId,
		basePath:       pathRoot,
	}
}

func (z *ccImpl) NoRetryOnError() api_context.Context {
	return &ccImpl{
		ctl:            z.ctl,
		token:          z.token,
		noRetryOnError: true,
		asMemberId:     z.asMemberId,
		asAdminId:      z.asAdminId,
		basePath:       z.basePath,
	}
}

func (z *ccImpl) ClientHash() string {
	z.hashMutex.Lock()
	defer z.hashMutex.Unlock()

	if z.hashComputed != "" {
		return z.hashComputed
	}

	seeds := []string{
		"m", z.asMemberId,
		"a", z.asAdminId,
		"p", z.token.PeerName(),
		"t", z.token.Token().AccessToken,
		"y", z.token.Scope(),
		"n", strconv.FormatBool(z.noRetryOnError),
	}
	if z.basePath != nil {
		seeds = append(seeds, z.basePath.Header())
	}
	z.hashComputed = ClientHash(seeds)
	return z.hashComputed
}

func NewResponse(ctx api_context.Context, req *http.Request, res *http.Response) (api_response.Response, error) {
	l := ctx.Log()
	defer nw_monitor.Log(req, res)

	if res == nil {
		l.Debug("Null response")
		return nil, api_response.ErrorNoResponse
	}

	result := ""
	if res.Header != nil {
		result = res.Header.Get(api_response.DropboxApiResHeaderResult)
	}
	if result != "" {
		resFile, err := ioutil.TempFile("", ctx.ClientHash())
		if err != nil {
			l.Debug("Unable to create temp file to store download", zap.Error(err))
			return nil, err
		}
		defer resFile.Close()
		buf := make([]byte, 4096)
		resBody := nw_bandwidth.WrapReader(res.Body)
		resWrite := bufio.NewWriter(resFile)
		var loadedLength int64

		for {
			n, err := resBody.Read(buf)
			if n > 0 {
				l.Debug("writing", zap.Int("writing", n))
				if _, wErr := resWrite.Write(buf[:n]); wErr != nil {
					l.Debug("Error on writing body to the file", zap.Error(wErr))
					return nil, err
				}
				loadedLength += int64(n)
			}
			if err == io.EOF {
				break
			}
			if err != nil {
				l.Debug("Error on reading body", zap.Error(err))
				return nil, err
			}
			if n == 0 {
				// Wait for throttling
				time.Sleep(50 * time.Millisecond)
				continue
			}
		}
		resWrite.Flush()
		res.ContentLength = loadedLength

		return api_response.NewDownload(res, mo_path.NewFileSystemPath(resFile.Name()), result, loadedLength), nil

	} else {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			l.Debug("Unable to read body", zap.Error(err))
			return nil, err
		}
		res.ContentLength = int64(len(body))

		return api_response.New(res, body), nil
	}
}
