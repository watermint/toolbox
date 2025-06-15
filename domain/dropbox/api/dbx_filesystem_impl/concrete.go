package dbx_filesystem_impl

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem"
	"github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/essentials/log/esl"
)

func newConcrete(client dbx_client.Client) dbx_filesystem.RootNamespaceResolver {
	return &rootNamespaceResolver{
		client: client,
	}
}

type rootNamespaceResolver struct {
	client dbx_client.Client
}

func (z rootNamespaceResolver) ResolveIndividual() (namespaceId string, err error) {
	l := z.client.Log()

	res := z.client.Post("users/get_current_account")
	if err, fail := res.Failure(); fail {
		l.Debug("Unable to verify token", esl.Error(err))
		return "", err
	}
	rj := res.Success().Json()
	rootInfo := &dbx_filesystem.RootInfo{}
	if rj.Model(rootInfo) != nil {
		if rootNamespaceId, found := rj.FindString("root_info.root_namespace_id"); found {
			l.Debug("Root namespace ID found", esl.String("rootNamespaceId", rootNamespaceId))
			return rootNamespaceId, nil
		}
	}

	return rootInfo.RootNamespaceId, nil
}

func (z rootNamespaceResolver) ResolveTeamMember(teamMemberId string) (namespaceId string, err error) {
	l := z.client.Log()

	type MIV2 struct {
		Tag          string `json:".tag"`
		TeamMemberId string `json:"team_member_id"`
	}
	p := struct {
		Members []*MIV2 `json:"members"`
	}{
		Members: []*MIV2{
			{
				Tag:          "team_member_id",
				TeamMemberId: teamMemberId,
			},
		},
	}

	res := z.client.Post("team/members/get_info_v2", api_request.Param(&p))
	if err, fail := res.Failure(); fail {
		l.Debug("Unable to verify token", esl.Error(err))
		return "", err
	}
	rj := res.Success().Json()
	memberNamespaceId, ok := rj.FindString("members_info.0.profile.root_folder_id")
	if !ok {
		l.Debug("Unable to find team folder id")
		return "", errors.New("unable to find team folder id")
	}
	return memberNamespaceId, nil
}
