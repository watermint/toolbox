package dbx_sharing

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_rpc"
	"go.uber.org/zap"
	"time"
)

type SharedLink struct {
	Raw                          json.RawMessage
	Kind                         string `path:"\\.tag" json:"kind"`
	SharedLinkId                 string `path:"id" json:"shared_link_id"`
	Url                          string `path:"url" json:"url"`
	Name                         string `path:"name" json:"name"`
	PathLower                    string `path:"path_lower" json:"path_lower"`
	ClientModified               string `path:"client_modified" json:"client_modified"`
	ServerModified               string `path:"server_modified" json:"server_modified"`
	Revision                     string `path:"rev" json:"revision"`
	Size                         uint64 `path:"size" json:"size,omitempty"`
	Expires                      string `path:"expires" json:"expires"`
	TeamId                       string `path:"team_member_info.team_info.id" json:"team_id"`
	TeamName                     string `path:"team_member_info.team_info.name" json:"team_name"`
	TeamMemberId                 string `path:"team_member_info.member_id" json:"team_member_id"`
	TeamMemberName               string `path:"team_member_info.display_name" json:"team_member_name"`
	ContentOwnerTeamId           string `path:"content_owner_team_info.id" json:"content_owner_team_id"`
	ContentOwnerTeamName         string `path:"content_owner_team_info.name" json:"content_owner_team_name"`
	PermissionResolvedVisibility string `path:"link_permissions.resolved_visibility.\\.tag" json:"permission_resolved_visibility"`
	PermissionAllowDownload      bool   `path:"link_permissions.allow_download" json:"permission_allow_download"`
}

func (a *SharedLink) UpdateExpire(c *dbx_api.Context, newExpire time.Time) (newLInk *SharedLink, annotation dbx_api.ErrorAnnotation, err error) {
	link := string(a.Raw)
	expires := gjson.Get(link, "expires").String()
	var origTime time.Time
	if expires != "" {
		var err error
		origTime, err = time.Parse(dbx_api.DateTimeFormat, expires)
		if err != nil {
			annotation = dbx_api.ErrorAnnotation{
				ErrorType: dbx_api.ErrorUnexpectedDataType,
				Error:     err,
			}
			return nil, annotation, err
		}
	}
	if origTime.IsZero() || origTime.After(newExpire) {
		return a.OverwriteExpire(c, newExpire)
	} else {
		c.Log().Debug(
			"skip updating link",
			zap.String("shared_link_id", a.SharedLinkId),
			zap.Time("orig_time", origTime),
			zap.Time("target_time", newExpire),
		)
	}
	return nil, dbx_api.Success, nil
}

func (a *SharedLink) OverwriteExpire(c *dbx_api.Context, newExpire time.Time) (newLink *SharedLink, annotation dbx_api.ErrorAnnotation, err error) {
	url := gjson.Get(string(a.Raw), "url").String()

	type SettingsParam struct {
		Expires string `json:"expires"`
	}
	type UpdateParam struct {
		Url      string        `json:"url"`
		Settings SettingsParam `json:"settings"`
	}

	up := UpdateParam{
		Url: url,
		Settings: SettingsParam{
			Expires: dbx_api.RebaseTimeForAPI(newExpire).Format(dbx_api.DateTimeFormat),
		},
	}

	req := dbx_rpc.RpcRequest{
		Endpoint:   "sharing/modify_shared_link_settings",
		Param:      up,
		AsMemberId: a.TeamMemberId,
	}
	res, ea, err := req.Call(c)
	c.Log().Debug("shared_link_response", zap.String("body", res.Body))
	if ea.IsFailure() {
		return nil, ea, err
	}

	c.Log().Debug("shared_link_response", zap.String("body", res.Body))

	newLink = &SharedLink{}
	err = c.ParseModel(newLink, gjson.Parse(res.Body))
	if err == nil {
		return newLink, dbx_api.Success, nil
	} else {
		return nil, dbx_api.ErrorAnnotation{
			ErrorType: dbx_api.ErrorUnexpectedDataType,
			Error:     err,
		}, err
	}
}

type SharedLinkList struct {
	Path          string
	AsMemberId    string
	AsMemberEmail string
	OnError       func(annotation dbx_api.ErrorAnnotation) bool
	OnEntry       func(link *SharedLink) bool
}

func (a *SharedLinkList) List(c *dbx_api.Context) bool {
	type ListParam struct {
		Path string `json:"path,omitempty"`
	}
	lp := ListParam{
		Path: a.Path,
	}
	list := dbx_rpc.RpcList{
		EndpointList:         "sharing/list_shared_links",
		EndpointListContinue: "sharing/list_shared_links",
		AsMemberId:           a.AsMemberId,
		UseHasMore:           false,
		ResultTag:            "links",
		OnError:              a.OnError,
		OnEntry: func(link gjson.Result) bool {
			if a.OnEntry == nil {
				return true
			}

			s := SharedLink{}
			err := c.ParseModel(&s, link)
			if err == nil {
				return a.OnEntry(&s)
			}

			if a.OnError != nil {
				return a.OnError(dbx_api.ErrorAnnotation{
					ErrorType: dbx_api.ErrorUnexpectedDataType,
					Error:     err,
				})
			}
			return false
		},
	}

	return list.List(c, lp)
}
