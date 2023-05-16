package sv_legalhold

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_list"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_legalhold"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"time"
)

type LegalHold interface {
	Create(name, description string, startDate, endDate time.Time, members []*mo_member.Member) (policy *mo_legalhold.Policy, err error)
	Info(id string) (policy *mo_legalhold.Policy, err error)
	List(includeReleased bool) (policies []*mo_legalhold.Policy, err error)
	Release(policyId string) error
	Revisions(policyId string, dateAfter time.Time, onEntry func(rev *mo_legalhold.Revision)) error
	UpdateName(policyId string, name string) (policy *mo_legalhold.Policy, err error)
	UpdateDesc(policyId string, desc string) (policy *mo_legalhold.Policy, err error)
	UpdateMember(policyId string, members []*mo_member.Member) (policy *mo_legalhold.Policy, err error)
}

func New(client dbx_client.Client) LegalHold {
	return &legalHoldImpl{
		client: client,
	}
}

type CreateParams struct {
	Name        string   `json:"name"`
	StartDate   *string  `json:"start_date,omitempty"`
	EndDate     *string  `json:"end_date,omitempty"`
	Description *string  `json:"description,omitempty"`
	Members     []string `json:"members,omitempty"`
}

type UpdateParams struct {
	Id          string   `json:"id"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Members     []string `json:"members,omitempty"`
}

type PolicyIdParam struct {
	Id string `json:"id"`
}

type legalHoldImpl struct {
	client dbx_client.Client
}

func (z legalHoldImpl) Create(name, description string, startDate, endDate time.Time, members []*mo_member.Member) (policy *mo_legalhold.Policy, err error) {
	params := CreateParams{
		Name:        name,
		Description: &description,
		Members:     make([]string, 0),
	}
	for _, member := range members {
		params.Members = append(params.Members, member.TeamMemberId)
	}
	if !startDate.IsZero() {
		t := dbx_util.ToApiTimeString(startDate)
		params.StartDate = &t
	}
	if !endDate.IsZero() {
		t := dbx_util.ToApiTimeString(endDate)
		params.EndDate = &t
	}
	res := z.client.Post("team/legal_holds/create_policy", api_request.Param(&params))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	policy = &mo_legalhold.Policy{}
	err = res.Success().Json().Model(policy)
	return
}

func (z legalHoldImpl) Info(id string) (policy *mo_legalhold.Policy, err error) {
	p := PolicyIdParam{
		Id: id,
	}
	res := z.client.Post("team/legal_holds/get_policy", api_request.Param(&p))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	policy = &mo_legalhold.Policy{}
	err = res.Success().Json().Model(policy)
	return
}

func (z legalHoldImpl) List(includeReleased bool) (policies []*mo_legalhold.Policy, err error) {
	p := struct {
		IncludeReleased bool `json:"include_released"`
	}{
		IncludeReleased: includeReleased,
	}
	res := z.client.Post("team/legal_holds/list_policies", api_request.Param(&p))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	policies = make([]*mo_legalhold.Policy, 0)
	err = res.Success().Json().FindArrayEach("policies", func(e es_json.Json) error {
		policy := &mo_legalhold.Policy{}
		if me := e.Model(policy); me != nil {
			return me
		}
		policies = append(policies, policy)
		return nil
	})
	return
}

func (z legalHoldImpl) Release(policyId string) error {
	p := PolicyIdParam{Id: policyId}
	res := z.client.Post("team/legal_holds/release_policy", api_request.Param(&p))
	err, _ := res.Failure()
	return err
}

func (z legalHoldImpl) Revisions(policyId string, dateAfter time.Time, onEntry func(rev *mo_legalhold.Revision)) error {
	p := PolicyIdParam{Id: policyId}
	errEutOfRange := errors.New("out of range")

	res := z.client.List("team/legal_holds/list_held_revisions", api_request.Param(&p)).Call(
		dbx_list.UseHasMore(),
		dbx_list.Continue("team/legal_holds/list_held_revisions_continue"),
		dbx_list.ResultTag("entries"),
		dbx_list.OnEntry(func(entry es_json.Json) error {
			rev := &mo_legalhold.Revision{}
			if err := entry.Model(rev); err != nil {
				return err
			}
			t, err := dbx_util.Parse(rev.ServerModified)
			if err != nil {
				return err
			}
			onEntry(rev)
			if dateAfter.After(t) {
				return errEutOfRange
			}
			return nil
		}),
	)
	if err, fail := res.Failure(); fail {
		if err == errEutOfRange {
			return nil
		}
		return err
	}
	return nil
}

func (z legalHoldImpl) UpdateName(policyId string, name string) (policy *mo_legalhold.Policy, err error) {
	params := UpdateParams{
		Id:   policyId,
		Name: name,
	}
	res := z.client.Post("team/legal_holds/update_policy", api_request.Param(&params))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	policy = &mo_legalhold.Policy{}
	err = res.Success().Json().Model(policy)
	return
}

func (z legalHoldImpl) UpdateDesc(policyId string, desc string) (policy *mo_legalhold.Policy, err error) {
	params := UpdateParams{
		Id:          policyId,
		Description: desc,
	}
	res := z.client.Post("team/legal_holds/update_policy", api_request.Param(&params))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	policy = &mo_legalhold.Policy{}
	err = res.Success().Json().Model(policy)
	return
}

func (z legalHoldImpl) UpdateMember(policyId string, members []*mo_member.Member) (policy *mo_legalhold.Policy, err error) {
	params := UpdateParams{
		Id:      policyId,
		Members: make([]string, 0),
	}
	for _, member := range members {
		params.Members = append(params.Members, member.TeamMemberId)
	}

	res := z.client.Post("team/legal_holds/update_policy", api_request.Param(&params))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	policy = &mo_legalhold.Policy{}
	err = res.Success().Json().Model(policy)
	return
}
