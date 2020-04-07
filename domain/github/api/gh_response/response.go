package gh_response

import (
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_response"
	"github.com/watermint/toolbox/infra/network/nw_monitor"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

func New(ctx api_context.Context, req *http.Request, res *http.Response) (api_response.Response, error) {
	l := ctx.Log()
	defer nw_monitor.Log(req, res)
	if res == nil {
		l.Debug("Null response")
		return nil, api_response.ErrorNoResponse
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		l.Debug("Unable to read body", zap.Error(err))
		return nil, err
	}
	res.ContentLength = int64(len(body))

	return api_response.New(res, body), nil
}
