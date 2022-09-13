package sv_account

import (
	"github.com/watermint/toolbox/domain/hellosign/api/hs_client"
	"github.com/watermint/toolbox/domain/hellosign/model/mo_account"
	"github.com/watermint/toolbox/essentials/api/api_request"
)

type Account interface {
	Info(accountId string) (info mo_account.Account, err error)
}

func New(client hs_client.Client) Account {
	return &accountImpl{
		client: client,
	}
}

type accountImpl struct {
	client hs_client.Client
}

func (z accountImpl) Info(accountId string) (info mo_account.Account, err error) {
	q := struct {
		AccountId string `url:"account_id,omitempty"`
	}{
		AccountId: accountId,
	}
	res := z.client.Get("account", api_request.Query(&q))
	if err, fail := res.Failure(); fail {
		return info, err
	}
	if err := res.Success().Json().FindModel("account", &info); err != nil {
		return info, err
	}
	return info, nil
}
