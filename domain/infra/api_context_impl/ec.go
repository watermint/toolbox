package api_context_impl

import (
	"crypto/sha256"
	"fmt"
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/domain/infra/api_async"
	"github.com/watermint/toolbox/domain/infra/api_async_impl"
	"github.com/watermint/toolbox/domain/infra/api_auth"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_list"
	"github.com/watermint/toolbox/domain/infra/api_list_impl"
	"github.com/watermint/toolbox/domain/infra/api_rpc"
	"github.com/watermint/toolbox/domain/infra/api_rpc_impl"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func New(ec *app.ExecContext, token api_auth.TokenContainer) api_context.Context {
	c := &ecImpl{
		ec:             ec,
		tokenContainer: token,
		client:         &http.Client{},
	}
	return c
}

func NewContextNoAuth(ec *app.ExecContext) api_context.Context {
	DefaultClientTimeout := time.Duration(60) * time.Second
	c := ecImpl{
		ec:     ec,
		noAuth: true,
		client: &http.Client{
			Timeout: DefaultClientTimeout,
		},
	}
	return &c
}

const (
	maxLastErrors = 10
)

type ecImpl struct {
	ec             *app.ExecContext
	tokenContainer api_auth.TokenContainer
	noAuth         bool
	client         *http.Client
	asMemberId     string
	asAdminId      string
	basePath       api_context.PathRoot
	retryAfter     time.Time
	lastErrors     []error
	noRetryOnError bool
}

func (z *ecImpl) Hash() string {
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

func (z *ecImpl) IsNoRetry() bool {
	return z.noRetryOnError
}

func (z *ecImpl) NoRetryOnError() api_context.Context {
	c := &ecImpl{
		ec:             z.ec,
		tokenContainer: z.tokenContainer,
		client:         &http.Client{},
		noRetryOnError: true,
	}
	return c
}

func (z *ecImpl) Token() api_auth.TokenContainer {
	return z.tokenContainer
}

func (z *ecImpl) DoRequest(req api_rpc.Request) (code int, header http.Header, body []byte, err error) {
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

func (z *ecImpl) UpdateRetryAfter(after time.Time) {
	z.retryAfter = after
}

func (z *ecImpl) AddError(err error) {
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

func (z *ecImpl) LastErrors() []error {
	if z.lastErrors == nil {
		return make([]error, 0)
	} else {
		return z.lastErrors
	}
}

func (z *ecImpl) RetryAfter() time.Time {
	return z.retryAfter
}

func (z *ecImpl) AsMemberId(teamMemberId string) api_context.Context {
	return &ecImpl{
		ec:             z.ec,
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

func (z *ecImpl) AsAdminId(teamMemberId string) api_context.Context {
	return &ecImpl{
		ec:             z.ec,
		tokenContainer: z.tokenContainer,
		noAuth:         z.noAuth,
		client: &http.Client{
			Timeout: z.client.Timeout,
		},
		asMemberId: "",
		asAdminId:  teamMemberId,
		basePath:   z.basePath,
	}
}

func (z *ecImpl) WithPath(pathRoot api_context.PathRoot) api_context.Context {
	return &ecImpl{
		ec:             z.ec,
		tokenContainer: z.tokenContainer,
		noAuth:         z.noAuth,
		client: &http.Client{
			Timeout: z.client.Timeout,
		},
		asMemberId: z.asMemberId,
		asAdminId:  z.asAdminId,
		basePath:   pathRoot,
	}
}

func (z *ecImpl) ClientTimeout(second int) {
	z.client.Timeout = time.Duration(second) * time.Second
}

func (z *ecImpl) Log() *zap.Logger {
	return z.ec.Log()
}

func (z *ecImpl) Request(endpoint string) api_rpc.Caller {
	return api_rpc_impl.New(z, endpoint, z.asMemberId, z.asAdminId, z.basePath, z.tokenContainer)
}

func (z *ecImpl) List(endpoint string) api_list.List {
	return api_list_impl.New(z, endpoint, z.asMemberId, z.asAdminId, z.basePath)
}

func (z *ecImpl) Async(endpoint string) api_async.Async {
	return api_async_impl.New(z, endpoint, z.asMemberId, z.asAdminId, z.basePath)
}
