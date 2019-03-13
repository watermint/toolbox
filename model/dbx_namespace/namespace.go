package dbx_namespace

import (
	"encoding/json"
	"errors"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_rpc"
)

type Namespace struct {
	NamespaceId   string          `json:"namespace_id"`
	NamespaceType string          `json:"namespace_type"`
	Name          string          `json:"name"`
	TeamMemberId  string          `json:"team_member_id,omitempty"`
	Namespace     json.RawMessage `json:"namespace"`
}

func ParseNamespace(n gjson.Result) (namespace *Namespace, err error) {
	namespaceId := n.Get("namespace_id")
	if !namespaceId.Exists() {
		err = errors.New("required field `namespace_id` not found in the response")
		return nil, err
	}

	ns := &Namespace{
		NamespaceId:   namespaceId.String(),
		NamespaceType: n.Get("namespace_type.\\.tag").String(),
		Name:          n.Get("name").String(),
		Namespace:     json.RawMessage(n.Raw),
	}
	return ns, nil
}

type NamespaceList struct {
	OnError func(err error) bool
	OnEntry func(namespace *Namespace) bool
}

func (z *NamespaceList) List(c *dbx_api.Context) bool {
	list := dbx_rpc.RpcList{
		EndpointList:         "team/namespaces/list",
		EndpointListContinue: "team/namespaces/list/continue",
		UseHasMore:           true,
		ResultTag:            "namespaces",
		OnError:              z.OnError,
		OnEntry: func(namespace gjson.Result) bool {
			n, err := ParseNamespace(namespace)
			if err != nil {
				return z.OnError(err)
			}
			return z.OnEntry(n)
		},
	}
	return list.List(c, nil)
}
