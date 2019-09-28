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
	"github.com/watermint/toolbox/infra/api/api_rpc"
	"github.com/watermint/toolbox/infra/api/api_rpc_impl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
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
	retryAfter     time.Time
	lastErrors     []error
	lastErrorMutex sync.Mutex
	noRetryOnError bool
}

func (z *ccImpl) Token() api_auth.TokenContainer {
	return z.tokenContainer
}

func (z *ccImpl) Capture() *zap.Logger {
	return z.control.Capture()
}

func (z *ccImpl) DoRequest(req api_rpc.Request) (code int, header http.Header, body []byte, err error) {
	httpReq, err := req.Request()
	if err != nil {
		return -1, nil, nil, err
	}
	client := &http.Client{}
	res, err := client.Do(httpReq)

	if err != nil {
		return -1, nil, nil, err
	}
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		// Do not retry
		z.Log().Debug("Unable to read body", zap.Error(err))
		return -1, nil, nil, err
	}
	if err = res.Body.Close(); err != nil {
		z.Log().Debug("Unable to close body", zap.Error(err))
		// fall through
	}
	return res.StatusCode, res.Header, body, nil
}

func (z *ccImpl) AddError(err error) {
	z.lastErrorMutex.Lock()
	defer z.lastErrorMutex.Unlock()

	if z.lastErrors == nil {
		z.lastErrors = make([]error, 0)
	}
	if err == nil {
		return
	}
	if len(z.lastErrors) > maxLastErrors {
		z.lastErrors = z.lastErrors[1:]
	}
	z.lastErrors = append(z.lastErrors, err)
}

func (z *ccImpl) LastErrors() []error {
	if z.lastErrors == nil {
		return make([]error, 0)
	} else {
		return z.lastErrors
	}
}

func (z *ccImpl) RetryAfter() time.Time {
	return z.retryAfter
}

func (z *ccImpl) UpdateRetryAfter(after time.Time) {
	z.retryAfter = after
}

func (z *ccImpl) IsNoRetry() bool {
	return z.noRetryOnError
}

func (z *ccImpl) Log() *zap.Logger {
	return z.control.Log()
}

func (z *ccImpl) Request(endpoint string) api_rpc.Caller {
	return api_rpc_impl.New(z, endpoint, z.asMemberId, z.asAdminId, z.basePath, z.tokenContainer)
}

func (z *ccImpl) List(endpoint string) api_list.List {
	return api_list_impl.New(z, endpoint, z.asMemberId, z.asAdminId, z.basePath)
}

func (z *ccImpl) Async(endpoint string) api_async.Async {
	return api_async_impl.New(z, endpoint, z.asMemberId, z.asAdminId, z.basePath)
}

func (z *ccImpl) AsMemberId(teamMemberId string) api_context.Context {
	return &ccImpl{
		control:        z.control,
		tokenContainer: z.tokenContainer,
		noAuth:         z.noAuth,
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
	}

	if z.basePath != nil {
		seeds = append(seeds, z.basePath.Header())
	}

	h := sha256.Sum224([]byte(strings.Join(seeds, ",")))
	return fmt.Sprintf("%x", h)
}
