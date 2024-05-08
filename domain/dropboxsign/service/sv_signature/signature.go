package sv_signature

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropboxsign/api/hs_client"
	"github.com/watermint/toolbox/domain/dropboxsign/model/mo_signature"
	"github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/essentials/log/esl"
)

type Signature interface {
	// List all signatures, and call handler for each signature.
	// If handler returns false, the iteration stops.
	List(accountId string, handler func(request *mo_signature.Request) bool) (err error)
}

func New(client hs_client.Client) Signature {
	return &signatureImpl{
		client: client,
	}
}

type signatureImpl struct {
	client hs_client.Client
}

func (z signatureImpl) List(accountId string, handler func(request *mo_signature.Request) bool) (err error) {
	l := z.client.Log().With(esl.String("accountId", accountId))
	q := struct {
		AccountId string `url:"account_id,omitempty"`
		Page      int    `url:"page,omitempty"`
		PageSize  int    `url:"page_size,omitempty"`
	}{
		AccountId: accountId,
		Page:      1,
		PageSize:  100,
	}

	for {
		res := z.client.Get("signature_request/list", api_request.Query(&q))
		if err, fail := res.Failure(); fail {
			l.Debug("Unable to list signatures", esl.Error(err))
			return err
		}
		reqList := mo_signature.RequestList{}
		if err := json.Unmarshal(res.Success().Body(), &reqList); err != nil {
			l.Debug("Unable to unmarshal response", esl.Error(err))
			return err
		}

		for _, req := range reqList.SignatureRequests {
			if !handler(req) {
				return nil
			}
		}
		q.Page++
		if reqList.ListInfo.NumPages <= q.Page {
			l.Debug("No more pages. Stop iteration.", esl.Int("page", reqList.ListInfo.NumPages))
			return nil
		}
		l.Debug("Next page", esl.Int("page", q.Page))
	}
}
