package dbx_team

import (
	"encoding/json"
	"errors"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_rpc"
)

type Namespace struct {
	NamespaceId   string          `json:"namespace_id"`
	NamespaceType string          `json:"namespace_type"`
	Name          string          `json:"name"`
	Namespace     json.RawMessage `json:"namespace"`
}

func ParseNamespace(n gjson.Result) (namespace *Namespace, annotation dbx_api.ErrorAnnotation, err error) {
	namespaceId := n.Get("namespace_id")
	if !namespaceId.Exists() {
		err = errors.New("required field `namespace_id` not found in the response")
		annotation = dbx_api.ErrorAnnotation{
			ErrorType: dbx_api.ErrorUnexpectedDataType,
			Error:     err,
		}
		return nil, annotation, err
	}

	c := Namespace{
		NamespaceId:   namespaceId.String(),
		NamespaceType: n.Get("namespace_type.\\.tag").String(),
		Name:          n.Get("name").String(),
		Namespace:     json.RawMessage(n.Raw),
	}
	return c, dbx_api.Success, nil
}

type NamespaceList struct {
	OnError func(annotation dbx_api.ErrorAnnotation) bool
	OnEntry func(namespace *Namespace) bool
}

func (w *NamespaceList) List(c *dbx_api.Context) bool {
	list := dbx_rpc.RpcList{
		EndpointList:         "team/namespaces/list",
		EndpointListContinue: "team/namespaces/list/continue",
		UseHasMore:           true,
		ResultTag:            "namespaces",
		OnError:              w.OnError,
		OnEntry: func(namespace gjson.Result) bool {
			n, ea, _ := ParseNamespace(namespace)
			if ea.IsSuccess() && w.OnEntry != nil {
				return w.OnEntry(n)
			}
			if ea.IsFailure() && w.OnError != nil {
				return w.OnError(ea)
			}
			return false
		},
	}
	list.List(c, nil)
}
