package dbx_request

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/essentials/io/es_rewinder"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"net/http"
	"strings"
)

func NewBuilder(ctl app_control.Control, entity api_auth.OAuthEntity, resolver dbx_filesystem.RootNamespaceResolver) Builder {
	return &builderImpl{
		ctl:          ctl,
		entity:       entity,
		rootResolver: resolver,
	}
}

type Builder interface {
	api_request.Builder
	AsMemberId(teamMemberId string) Builder
	AsAdminId(teamMemberId string) Builder
	WithPath(pathRoot dbx_client.PathRoot) Builder
	With(method, url string, data api_request.RequestData) Builder
	BaseNamespace(baseNamespace dbx_filesystem.BaseNamespaceType) Builder
	NoAuth() Builder
}

type builderImpl struct {
	ctl           app_control.Control
	entity        api_auth.OAuthEntity
	asMemberId    string
	asAdminId     string
	basePath      dbx_client.PathRoot
	baseNamespace dbx_filesystem.BaseNamespaceType
	rootResolver  dbx_filesystem.RootNamespaceResolver
	method        string
	data          api_request.RequestData
	url           string
}

func (z builderImpl) BaseNamespace(baseNamespace dbx_filesystem.BaseNamespaceType) Builder {
	z.baseNamespace = baseNamespace
	return z
}

func (z builderImpl) WithData(data api_request.RequestDatum) api_request.Builder {
	z.data = z.data.WithDatum(data)
	return z
}

func (z builderImpl) Endpoint() string {
	return z.url
}

func (z builderImpl) NoAuth() Builder {
	z.entity = api_auth.NewNoAuthOAuthEntity()
	return z
}
func (z builderImpl) AsMemberId(teamMemberId string) Builder {
	z.asMemberId = teamMemberId
	return z
}
func (z builderImpl) AsAdminId(teamMemberId string) Builder {
	z.asAdminId = teamMemberId
	return z
}
func (z builderImpl) WithPath(pathRoot dbx_client.PathRoot) Builder {
	z.basePath = pathRoot
	return z
}

func (z builderImpl) With(method, url string, data api_request.RequestData) Builder {
	z.method = method
	z.url = url
	z.data = data
	return z
}

func (z builderImpl) Log() esl.Logger {
	l := z.ctl.Log()
	if z.asMemberId != "" {
		l = l.With(esl.String("asMemberId", z.asMemberId))
	}
	if z.asAdminId != "" {
		l = l.With(esl.String("asAdminId", z.asAdminId))
	}
	if z.basePath != nil {
		l = l.With(esl.Any("basePath", z.basePath))
	}
	return l
}

func (z builderImpl) ClientHash() string {
	var ss, sr, sp []string
	sr = []string{
		"m", z.method,
		"u", z.url,
		"b", string(z.baseNamespace),
	}
	ss = []string{
		"m", z.asMemberId,
		"a", z.asAdminId,
	}
	if z.basePath != nil {
		sp = []string{"p", z.basePath.Header()}
	}
	return nw_client.ClientHash(ss, sr, z.entity.HashSeed(), sp)
}

func (z builderImpl) Build() (*http.Request, error) {
	l := z.Log().With(esl.String("method", z.method), esl.String("url", z.url))
	rc := z.reqContent()
	req, err := nw_client.NewHttpRequest(z.method, z.url, rc)
	if err != nil {
		l.Debug("Unable to make request")
		return nil, err
	}
	headers := z.reqHeaders()
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	req.ContentLength = rc.Length()

	return req, nil
}

func (z builderImpl) reqHeaders() map[string]string {
	l := z.Log()

	headers := make(map[string]string)
	headers[api_request.ReqHeaderUserAgent] = app_definitions.UserAgent()
	//if !z.entity.IsNoAuth() {
	//	headers[api_request.ReqHeaderAuthorization] = "Bearer " + z.entity.Token.AccessToken
	//}
	if z.asAdminId != "" {
		headers[api_request.ReqHeaderDropboxApiSelectAdmin] = z.asAdminId
	}
	if z.asMemberId != "" {
		headers[api_request.ReqHeaderDropboxApiSelectUser] = z.asMemberId
	}
	if z.basePath != nil {
		p, err := dbx_util.HeaderSafeJson(z.basePath)
		if err != nil {
			l.Debug("Unable to marshal base path", esl.Error(err))
		} else {
			headers[api_request.ReqHeaderDropboxApiPathRoot] = p
		}
	} else {
		switch {
		case z.baseNamespace == dbx_filesystem.BaseNamespaceHome:
			// do nothing
			break

		case z.ctl.Feature().Experiment(app_definitions.ExperimentDbxDisableAutoPathRoot):
			l.Debug("Auto path root disabled")

		case strings.Contains(z.url, "2/files"),
			strings.Contains(z.url, "2/sharing"):
			var rootNamespaceId string
			var err error
			if dbx_auth.IsTeamAppKey(z.entity.KeyName) && z.asMemberId != "" {
				rootNamespaceId, err = z.rootResolver.ResolveTeamMember(z.asMemberId)
				if err != nil {
					l.Warn("Unable to resolve team member namespace", esl.Error(err))
					break
				}
			}
			if dbx_auth.IsIndividualAppKey(z.entity.KeyName) {
				rootNamespaceId, err = z.rootResolver.ResolveIndividual()
				if err != nil {
					l.Warn("Unable to resolve individual namespace", esl.Error(err))
					break
				}
			}

			if rootNamespaceId == "" {
				l.Debug("Root namespace is empty")
			} else {
				p, err := dbx_util.HeaderSafeJson(dbx_client.Namespace(rootNamespaceId))
				if err != nil {
					l.Debug("Unable to marshal root namespace", esl.Error(err))
				} else {
					headers[api_request.ReqHeaderDropboxApiPathRoot] = p
				}
			}
		}
	}
	if z.data.Content() != nil {
		p, err := dbx_util.HeaderSafeJson(z.data.Param())
		if err != nil {
			l.Debug("Unable to marshal params", esl.Error(err))
		} else {
			headers[api_request.ReqHeaderDropboxApiArg] = p
		}
		headers[api_request.ReqHeaderContentType] = "application/octet-stream"
	} else if len(z.data.ParamJson()) > 0 {
		headers[api_request.ReqHeaderContentType] = "application/json"
	}
	for k, v := range z.data.Headers() {
		headers[k] = v
	}
	return headers
}

func (z builderImpl) reqContent() es_rewinder.ReadRewinder {
	if z.data.Content() != nil {
		return z.data.Content()
	}
	return es_rewinder.NewReadRewinderOnMemory(z.data.ParamJson())
}

func (z builderImpl) Param() string {
	return string(z.data.ParamJson())
}
