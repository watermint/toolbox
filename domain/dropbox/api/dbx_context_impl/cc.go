package dbx_context_impl

import (
	"crypto/sha256"
	"fmt"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_async"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_list"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_request"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_response"
	"github.com/watermint/toolbox/infra/api/api_async"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_list"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/api/api_response"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/util/ut_io"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

func New(control app_control.Control, token api_auth.Context) dbx_context.Context {
	c := &ccImpl{
		control:        control,
		token:          token,
		noRetryOnError: false,
	}
	return c
}

type ccImpl struct {
	control        app_control.Control
	token          api_auth.Context
	asMemberId     string
	asAdminId      string
	basePath       dbx_context.PathRoot
	noRetryOnError bool
	hashComputed   string
	hashMutex      sync.Mutex
}

func (z *ccImpl) MakeResponse(req *http.Request, res *http.Response) (api_response.Response, error) {
	return dbx_response.New(z, req, res)
}

func (z *ccImpl) NoAuth() api_context.Context {
	return &ccImpl{
		control:        z.control,
		token:          dbx_auth.NewNoAuth(),
		asMemberId:     z.asMemberId,
		asAdminId:      z.asAdminId,
		basePath:       z.basePath,
		noRetryOnError: z.noRetryOnError,
		hashComputed:   z.hashComputed,
		hashMutex:      sync.Mutex{},
	}
}

func (z *ccImpl) Capture() *zap.Logger {
	return z.control.Capture()
}

func (z *ccImpl) IsNoRetry() bool {
	return z.noRetryOnError
}

func (z *ccImpl) Log() *zap.Logger {
	return z.control.Log()
}

func (z *ccImpl) Rpc(endpoint string) api_request.Request {
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
		control:        z.control,
		token:          z.token,
		noRetryOnError: z.noRetryOnError,
		asMemberId:     teamMemberId,
		asAdminId:      "",
		basePath:       z.basePath,
	}
}

func (z *ccImpl) AsAdminId(teamMemberId string) dbx_context.Context {
	return &ccImpl{
		control:        z.control,
		token:          z.token,
		noRetryOnError: z.noRetryOnError,
		asMemberId:     "",
		asAdminId:      teamMemberId,
		basePath:       z.basePath,
	}
}

func (z *ccImpl) WithPath(pathRoot dbx_context.PathRoot) dbx_context.Context {
	return &ccImpl{
		control:        z.control,
		token:          z.token,
		noRetryOnError: z.noRetryOnError,
		asMemberId:     z.asMemberId,
		asAdminId:      z.asAdminId,
		basePath:       pathRoot,
	}
}

func (z *ccImpl) NoRetryOnError() api_context.Context {
	return &ccImpl{
		control:        z.control,
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
		"m",
		z.asMemberId,
		"a",
		z.asAdminId,
		"p",
		z.token.PeerName(),
		"t",
		z.token.Token().AccessToken,
		"y",
		z.token.Scope(),
		"n",
		strconv.FormatBool(z.noRetryOnError),
	}

	if z.basePath != nil {
		seeds = append(seeds, z.basePath.Header())
	}

	h := sha256.Sum224([]byte(strings.Join(seeds, ",")))
	z.hashComputed = fmt.Sprintf("%x", h)
	return z.hashComputed
}
