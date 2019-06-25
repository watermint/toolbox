package api_context_impl

import (
	"crypto/sha256"
	"fmt"
	"github.com/watermint/toolbox/domain/infra/api_async"
	"github.com/watermint/toolbox/domain/infra/api_async_impl"
	"github.com/watermint/toolbox/domain/infra/api_auth"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_list"
	"github.com/watermint/toolbox/domain/infra/api_list_impl"
	"github.com/watermint/toolbox/domain/infra/api_rpc"
	"github.com/watermint/toolbox/domain/infra/api_rpc_impl"
	"github.com/watermint/toolbox/experimental/app_kitchen"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func NewKC(kitchen app_kitchen.Kitchen, token api_auth.TokenContainer) api_context.Context {
	c := &kcImpl{
		kitchen:        kitchen,
		tokenContainer: token,
		client:         &http.Client{},
		noRetryOnError: false,
	}
	return c
}

type kcImpl struct {
	kitchen        app_kitchen.Kitchen
	client         *http.Client
	tokenContainer api_auth.TokenContainer
	noAuth         bool
	asMemberId     string
	asAdminId      string
	basePath       api_context.PathRoot
	retryAfter     time.Time
	lastErrors     []error
	noRetryOnError bool
}

func (z *kcImpl) Capture() *zap.Logger {
	return z.kitchen.Control().Capture()
}

func (z *kcImpl) DoRequest(req api_rpc.Request) (code int, header http.Header, body []byte, err error) {
	httpReq, err := req.Request()
	if err != nil {
		return -1, nil, nil, err
	}
	res, err := z.client.Do(httpReq)

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

func (z *kcImpl) AddError(err error) {
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

func (z *kcImpl) LastErrors() []error {
	if z.lastErrors == nil {
		return make([]error, 0)
	} else {
		return z.lastErrors
	}
}

func (z *kcImpl) RetryAfter() time.Time {
	return z.retryAfter
}

func (z *kcImpl) UpdateRetryAfter(after time.Time) {
	z.retryAfter = after
}

func (z *kcImpl) IsNoRetry() bool {
	return z.noRetryOnError
}

func (z *kcImpl) Log() *zap.Logger {
	return z.kitchen.Log()
}

func (z *kcImpl) Request(endpoint string) api_rpc.Caller {
	return api_rpc_impl.New(z, endpoint, z.asMemberId, z.asAdminId, z.basePath, z.tokenContainer)
}

func (z *kcImpl) List(endpoint string) api_list.List {
	return api_list_impl.New(z, endpoint, z.asMemberId, z.asAdminId, z.basePath)
}

func (z *kcImpl) Async(endpoint string) api_async.Async {
	return api_async_impl.New(z, endpoint, z.asMemberId, z.asAdminId, z.basePath)
}

func (z *kcImpl) AsMemberId(teamMemberId string) api_context.Context {
	return &kcImpl{
		kitchen:        z.kitchen,
		tokenContainer: z.tokenContainer,
		noAuth:         z.noAuth,
		client: &http.Client{
			Timeout: z.client.Timeout,
		},
		asMemberId: teamMemberId,
		asAdminId:  "",
		basePath:   z.basePath,
	}
}

func (z *kcImpl) AsAdminId(teamMemberId string) api_context.Context {
	return &kcImpl{
		kitchen:        z.kitchen,
		tokenContainer: z.tokenContainer,
		noAuth:         z.noAuth,
		client: &http.Client{
			Timeout: z.client.Timeout,
		},
		noRetryOnError: z.noRetryOnError,
		asMemberId:     "",
		asAdminId:      teamMemberId,
		basePath:       z.basePath,
	}
}

func (z *kcImpl) WithPath(pathRoot api_context.PathRoot) api_context.Context {
	return &kcImpl{
		kitchen:        z.kitchen,
		tokenContainer: z.tokenContainer,
		noAuth:         z.noAuth,
		client: &http.Client{
			Timeout: z.client.Timeout,
		},
		noRetryOnError: z.noRetryOnError,
		asMemberId:     z.asMemberId,
		asAdminId:      z.asAdminId,
		basePath:       pathRoot,
	}
}

func (z *kcImpl) NoRetryOnError() api_context.Context {
	return &kcImpl{
		kitchen:        z.kitchen,
		tokenContainer: z.tokenContainer,
		noAuth:         z.noAuth,
		client: &http.Client{
			Timeout: z.client.Timeout,
		},
		noRetryOnError: true,
		asMemberId:     z.asMemberId,
		asAdminId:      z.asAdminId,
		basePath:       z.basePath,
	}
}

func (z *kcImpl) Hash() string {
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
