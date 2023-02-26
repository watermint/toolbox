package fg_client_impl

import (
	"github.com/watermint/toolbox/domain/figma/api/fg_request"
	"github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"github.com/watermint/toolbox/infra/control/app_control"
)

type clientImpl struct {
	peerName string
	client   nw_client.Rest
	ctl      app_control.Control
	builder  fg_request.Builder
}

func (z clientImpl) Name() string {
	//TODO implement me
	panic("implement me")
}

func (z clientImpl) ClientHash() string {
	//TODO implement me
	panic("implement me")
}

func (z clientImpl) Log() esl.Logger {
	//TODO implement me
	panic("implement me")
}

func (z clientImpl) Capture() esl.Logger {
	//TODO implement me
	panic("implement me")
}

func (z clientImpl) Get(endpoint string, d ...api_request.RequestDatum) (res es_response.Response) {
	//TODO implement me
	panic("implement me")
}

func (z clientImpl) Post(endpoint string, d ...api_request.RequestDatum) (res es_response.Response) {
	//TODO implement me
	panic("implement me")
}

func (z clientImpl) Delete(endpoint string, d ...api_request.RequestDatum) (res es_response.Response) {
	//TODO implement me
	panic("implement me")
}
