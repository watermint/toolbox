package api_context_impl

import (
	"crypto/sha256"
	"fmt"
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/app/app_ui"
	"github.com/watermint/toolbox/domain/infra/api_async"
	"github.com/watermint/toolbox/domain/infra/api_async_impl"
	"github.com/watermint/toolbox/domain/infra/api_auth"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_list"
	"github.com/watermint/toolbox/domain/infra/api_list_impl"
	"github.com/watermint/toolbox/domain/infra/api_rpc"
	"github.com/watermint/toolbox/domain/infra/api_rpc_impl"
	"github.com/watermint/toolbox/domain/infra/api_util"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func New(ec *app.ExecContext, token api_auth.TokenContainer) api_context.Context {
	c := &contextImpl{
		ec:             ec,
		tokenContainer: token,
		client:         &http.Client{},
	}
	return c
}

func NewContextNoAuth(ec *app.ExecContext) api_context.Context {
	DefaultClientTimeout := time.Duration(60) * time.Second
	c := contextImpl{
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

type contextImpl struct {
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

func (z *contextImpl) Hash() string {
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

func (z *contextImpl) IsNoRetry() bool {
	return z.noRetryOnError
}

func (z *contextImpl) NoRetryOnError() api_context.Context {
	c := &contextImpl{
		ec:             z.ec,
		tokenContainer: z.tokenContainer,
		client:         &http.Client{},
		noRetryOnError: true,
	}
	return c
}

func (z *contextImpl) Token() api_auth.TokenContainer {
	return z.tokenContainer
}

func (z *contextImpl) DoRequest(req api_rpc.Request) (code int, header http.Header, body []byte, err error) {
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

func (z *contextImpl) UpdateRetryAfter(after time.Time) {
	z.retryAfter = after
}

func (z *contextImpl) AddError(err error) {
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

func (z *contextImpl) LastErrors() []error {
	if z.lastErrors == nil {
		return make([]error, 0)
	} else {
		return z.lastErrors
	}
}

func (z *contextImpl) RetryAfter() time.Time {
	return z.retryAfter
}

func (z *contextImpl) AsMemberId(teamMemberId string) api_context.Context {
	return &contextImpl{
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

func (z *contextImpl) AsAdminId(teamMemberId string) api_context.Context {
	return &contextImpl{
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

func (z *contextImpl) WithPath(pathRoot api_context.PathRoot) api_context.Context {
	return &contextImpl{
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

func (z *contextImpl) ErrorMsg(err error) app_ui.UIMessage {
	if err == nil {
		return z.ec.Msg(app.MsgNoError)
	}
	summary := api_util.ErrorSummary(err)
	if summary == "" {
		return z.ec.Msg("app.common.api.err.general_error").WithData(struct {
			Error string
		}{
			Error: err.Error(),
		})
	} else {
		errMsgKey := "dbx.err." + summary
		userMessage := api_util.ErrorUserMessage(err)

		if z.ec.MessageContainer().MsgExists(errMsgKey) {
			errDesc := z.ec.Msg(errMsgKey).T()
			return z.ec.Msg("app.common.api.err.api_error").WithData(struct {
				Tag   string
				Error string
			}{
				Tag:   summary,
				Error: errDesc,
			})
		}

		return z.ec.Msg("app.common.api.err.api_error").WithData(struct {
			Tag   string
			Error string
		}{
			Tag:   summary,
			Error: userMessage,
		})
	}
}

func (z *contextImpl) ClientTimeout(second int) {
	z.client.Timeout = time.Duration(second) * time.Second
}

func (z *contextImpl) Log() *zap.Logger {
	return z.ec.Log()
}

func (z *contextImpl) Msg(key string) app_ui.UIMessage {
	return z.ec.Msg(key)
}

func (z *contextImpl) Request(endpoint string) api_rpc.Caller {
	return api_rpc_impl.New(z, endpoint, z.asMemberId, z.asAdminId, z.basePath, z.tokenContainer)
}

func (z *contextImpl) List(endpoint string) api_list.List {
	return api_list_impl.New(z, endpoint, z.asMemberId, z.asAdminId, z.basePath)
}

func (z *contextImpl) Async(endpoint string) api_async.Async {
	return api_async_impl.New(z, endpoint, z.asMemberId, z.asAdminId, z.basePath)
}
