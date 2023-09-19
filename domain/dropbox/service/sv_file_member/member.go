package sv_file_member

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_list"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder_member"
	"github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/essentials/http/es_response"
)

type Member interface {
	List(fileId string, includeInherited bool) (member []mo_sharedfolder_member.Member, err error)
}

func New(client dbx_client.Client) Member {
	return &memberImpl{
		client: client,
	}
}

type memberImpl struct {
	client dbx_client.Client
}

func (z memberImpl) List(fileId string, includeInherited bool) (member []mo_sharedfolder_member.Member, err error) {
	member = make([]mo_sharedfolder_member.Member, 0)

	p := struct {
		FileId           string `json:"file"`
		IncludeInherited bool   `json:"include_inherited"`
	}{
		FileId:           fileId,
		IncludeInherited: includeInherited,
	}

	res := z.client.List("sharing/list_file_members", api_request.Param(p)).Call(
		dbx_list.Continue("sharing/list_file_members/continue"),
		dbx_list.OnResponse(func(res es_response.Response) error {
			j, err := res.Success().AsJson()
			if err != nil {
				return err
			}
			if users, found := j.FindArray("users"); found {
				for _, u := range users {
					mu := &mo_sharedfolder_member.User{}
					if err := u.Model(mu); err != nil {
						return err
					}
					member = append(member, mu)
				}
			}
			if groups, found := j.FindArray("groups"); found {
				for _, g := range groups {
					mg := &mo_sharedfolder_member.Group{}
					if err := g.Model(mg); err != nil {
						return err
					}
					member = append(member, mg)
				}
			}
			if invitees, found := j.FindArray("invitees"); found {
				for _, i := range invitees {
					mi := &mo_sharedfolder_member.Invitee{}
					if err := i.Model(mi); err != nil {
						return err
					}
					member = append(member, mi)
				}
			}
			return nil
		}),
	)
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	return member, nil
}
