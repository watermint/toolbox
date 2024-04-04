package sv_graphql

import (
	"github.com/watermint/toolbox/domain/github/api/gh_client"
	"github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
)

type Query interface {
	Query(query string) (result es_json.Json, err error)
}

func NewQuery(client gh_client.Client) Query {
	return &queryImpl{
		client: client,
	}
}

type queryImpl struct {
	client gh_client.Client
}

func (z queryImpl) Query(query string) (result es_json.Json, err error) {
	q := struct {
		Query string `json:"query"`
	}{
		Query: query,
	}
	res := z.client.Post("graphql", api_request.Param(&q))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	return res.Success().Json(), nil
}
