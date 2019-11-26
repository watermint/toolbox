package api_context_impl

import (
	"crypto/sha256"
	"fmt"
	"github.com/watermint/toolbox/infra/api/api_async"
	"github.com/watermint/toolbox/infra/api/api_async_impl"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_list"
	"github.com/watermint/toolbox/infra/api/api_list_impl"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/api/api_request_impl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"go.uber.org/zap"
	"io"
	"strconv"
	"strings"
)

const (
	maxLastErrors = 10
)

func New(control app_control.Control, token api_auth.TokenContainer) api_context.Context {
	c := &ccImpl{
		control:        control,
		tokenContainer: token,
		noRetryOnError: false,
	}
	return c
}

type ccImpl struct {
	control        app_control.Control
	tokenContainer api_auth.TokenContainer
	noAuth         bool
	asMemberId     string
	asAdminId      string
	basePath       api_context.PathRoot
	noRetryOnError bool
}

func (z *ccImpl) Token() api_auth.TokenContainer {
	return z.tokenContainer
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
	return api_request_impl.NewPpcRequest(
		z,
		endpoint,
		z.asMemberId,
		z.asAdminId,
		z.basePath,
		z.tokenContainer,
	)
}

func (z *ccImpl) List(endpoint string) api_list.List {
	return api_list_impl.New(z, endpoint, z.asMemberId, z.asAdminId, z.basePath)
}

func (z *ccImpl) Async(endpoint string) api_async.Async {
	return api_async_impl.New(z, endpoint, z.asMemberId, z.asAdminId, z.basePath)
}

func (z *ccImpl) Upload(endpoint string, content io.Reader) api_request.Request {
	return api_request_impl.NewUploadRequest(
		z,
		endpoint,
		content,
		z.asMemberId,
		z.asAdminId,
		z.basePath,
		z.tokenContainer,
	)
}

func (z *ccImpl) AsMemberId(teamMemberId string) api_context.Context {
	return &ccImpl{
		control:        z.control,
		tokenContainer: z.tokenContainer,
		noAuth:         z.noAuth,
		noRetryOnError: z.noRetryOnError,
		asMemberId:     teamMemberId,
		asAdminId:      "",
		basePath:       z.basePath,
	}
}

func (z *ccImpl) AsAdminId(teamMemberId string) api_context.Context {
	return &ccImpl{
		control:        z.control,
		tokenContainer: z.tokenContainer,
		noAuth:         z.noAuth,
		noRetryOnError: z.noRetryOnError,
		asMemberId:     "",
		asAdminId:      teamMemberId,
		basePath:       z.basePath,
	}
}

func (z *ccImpl) WithPath(pathRoot api_context.PathRoot) api_context.Context {
	return &ccImpl{
		control:        z.control,
		tokenContainer: z.tokenContainer,
		noAuth:         z.noAuth,
		noRetryOnError: z.noRetryOnError,
		asMemberId:     z.asMemberId,
		asAdminId:      z.asAdminId,
		basePath:       pathRoot,
	}
}

func (z *ccImpl) NoRetryOnError() api_context.Context {
	return &ccImpl{
		control:        z.control,
		tokenContainer: z.tokenContainer,
		noAuth:         z.noAuth,
		noRetryOnError: true,
		asMemberId:     z.asMemberId,
		asAdminId:      z.asAdminId,
		basePath:       z.basePath,
	}
}

func (z *ccImpl) Hash() string {
	seeds := []string{
		"m",
		z.asMemberId,
		"a",
		z.asAdminId,
		"p",
		z.tokenContainer.PeerName,
		"t",
		z.tokenContainer.Token,
		"y",
		z.tokenContainer.TokenType,
		"n",
		strconv.FormatBool(z.noRetryOnError),
	}

	if z.basePath != nil {
		seeds = append(seeds, z.basePath.Header())
	}

	h := sha256.Sum224([]byte(strings.Join(seeds, ",")))
	return fmt.Sprintf("%x", h)
}
