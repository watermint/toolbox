package sv_member

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_async"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_list"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/infra/api/api_request"
	"go.uber.org/zap"
	"strings"
)

var (
	ErrorMemberNotFoundForEmail        = errors.New("member not found for the email")
	ErrorMemberNotFoundForTeamMemberId = errors.New("member not found for the team_member_id")
	ErrorNotFound                      = errors.New("not found")
)

type Member interface {
	Update(member *mo_member.Member) (updated *mo_member.Member, err error)
	List() (members []*mo_member.Member, err error)
	Resolve(teamMemberId string) (member *mo_member.Member, err error)
	ResolveByEmail(email string) (member *mo_member.Member, err error)
	Add(email string, opts ...AddOpt) (member *mo_member.Member, err error)
	Remove(member *mo_member.Member, opts ...RemoveOpt) (err error)
}

type AddOpt func(opt *addOptions) *addOptions
type addOptions struct {
	givenName             string
	surname               string
	externalId            string
	sendWelcomeEmail      bool
	role                  string
	isDirectoryRestricted bool
}

func AddWithGivenName(givenName string) AddOpt {
	return func(opt *addOptions) *addOptions {
		opt.givenName = givenName
		return opt
	}
}
func AddWithSurname(surname string) AddOpt {
	return func(opt *addOptions) *addOptions {
		opt.surname = surname
		return opt
	}
}
func AddWithExternalId(externalId string) AddOpt {
	return func(opt *addOptions) *addOptions {
		opt.externalId = externalId
		return opt
	}
}

// Use silent provisioning.
// (required to verify domain first)
// https://help.dropbox.com/business/domain-verification-invite-enforcement
func AddWithoutSendWelcomeEmail() AddOpt {
	return func(opt *addOptions) *addOptions {
		opt.sendWelcomeEmail = false
		return opt
	}
}
func AddWithRole(role string) AddOpt {
	return func(opt *addOptions) *addOptions {
		opt.role = role
		return opt
	}
}
func AddWithDirectoryRestricted() AddOpt {
	return func(opt *addOptions) *addOptions {
		opt.isDirectoryRestricted = true
		return opt
	}
}

type RemoveOpt func(opt *removeOptions) *removeOptions
type removeOptions struct {
	wipeData         bool
	keepAccount      bool
	retainTeamShares bool
}

// Downgrade the member to a Basic account.
func Downgrade() RemoveOpt {
	return func(opt *removeOptions) *removeOptions {
		opt.wipeData = false
		opt.keepAccount = true
		return opt
	}
}
func RemoveWipeData() RemoveOpt {
	return func(opt *removeOptions) *removeOptions {
		opt.wipeData = true
		return opt
	}
}
func RetainTeamShares() RemoveOpt {
	return func(opt *removeOptions) *removeOptions {
		opt.retainTeamShares = true
		return opt
	}
}

func New(ctx dbx_context.Context) Member {
	return &memberImpl{
		ctx: ctx,
	}
}

func NewCached(ctx dbx_context.Context) Member {
	return &cachedMember{
		impl: &memberImpl{
			ctx: ctx,
		},
	}
}

func newTest(ctx dbx_context.Context) Member {
	return &memberImpl{
		ctx:   ctx,
		limit: 3,
	}
}

type cachedMember struct {
	impl    Member
	members []*mo_member.Member
}

func (z *cachedMember) Update(member *mo_member.Member) (updated *mo_member.Member, err error) {
	z.members = nil // invalidate
	return z.impl.Update(member)
}

func (z *cachedMember) List() (members []*mo_member.Member, err error) {
	if z.members == nil {
		z.members, err = z.impl.List()
		if err != nil {
			return nil, err
		}
	}
	return z.members, nil
}

func (z *cachedMember) Resolve(teamMemberId string) (member *mo_member.Member, err error) {
	if z.members == nil {
		z.members, err = z.impl.List()
		if err != nil {
			return nil, err
		}
	}
	for _, m := range z.members {
		if m.TeamMemberId == teamMemberId {
			return m, nil
		}
	}
	return nil, ErrorMemberNotFoundForTeamMemberId
}

func (z *cachedMember) ResolveByEmail(email string) (member *mo_member.Member, err error) {
	if z.members == nil {
		z.members, err = z.impl.List()
		if err != nil {
			return nil, err
		}
	}
	em := strings.ToLower(email)
	for _, m := range z.members {
		if strings.ToLower(m.Email) == em {
			return m, nil
		}
	}
	return nil, ErrorMemberNotFoundForEmail
}

func (z *cachedMember) Add(email string, opts ...AddOpt) (member *mo_member.Member, err error) {
	z.members = nil // invalidate cache
	return z.impl.Add(email, opts...)
}

func (z *cachedMember) Remove(member *mo_member.Member, opts ...RemoveOpt) (err error) {
	z.members = nil // invalidate cache
	return z.impl.Remove(member, opts...)
}

type memberImpl struct {
	ctx            dbx_context.Context
	includeDeleted bool
	limit          int
}

func (z *memberImpl) Add(email string, opts ...AddOpt) (member *mo_member.Member, err error) {
	ao := &addOptions{
		sendWelcomeEmail: true,
	}
	for _, o := range opts {
		o(ao)
	}
	type NM struct {
		MemberEmail      string `json:"member_email"`
		MemberGivenName  string `json:"member_given_name,omitempty"`
		MemberSurname    string `json:"member_surname,omitempty"`
		MemberExternalId string `json:"member_external_id,omitempty"`
		SendWelcomeEmail bool   `json:"send_welcome_email"`
		Role             string `json:"role,omitempty"`
	}
	p := struct {
		NewMembers []*NM `json:"new_members"`
	}{
		NewMembers: []*NM{
			{
				MemberEmail:      email,
				MemberGivenName:  ao.givenName,
				MemberSurname:    ao.surname,
				MemberExternalId: ao.externalId,
				SendWelcomeEmail: ao.sendWelcomeEmail,
				Role:             ao.role,
			},
		},
	}

	res := z.ctx.Async("team/members/add", api_request.Param(p)).Call(
		dbx_async.Status("team/members/add/job_status/get"),
	)
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	member = &mo_member.Member{}
	err = res.Success().Json().FindModel(es_json.PathArrayFirst, member)
	return
}

func (z *memberImpl) Remove(member *mo_member.Member, opts ...RemoveOpt) (err error) {
	ro := &removeOptions{}
	for _, o := range opts {
		o(ro)
	}
	type US struct {
		Tag          string `json:".tag"`
		TeamMemberId string `json:"team_member_id"`
	}
	p := struct {
		User             US   `json:"user"`
		WipeData         bool `json:"wipe_data"`
		KeepAccount      bool `json:"keep_account"`
		RetainTeamShares bool `json:"retain_team_shares"`
	}{
		User: US{
			Tag:          "team_member_id",
			TeamMemberId: member.TeamMemberId,
		},
		WipeData:         ro.wipeData,
		KeepAccount:      ro.keepAccount,
		RetainTeamShares: ro.retainTeamShares,
	}

	res := z.ctx.Async("team/members/remove", api_request.Param(p)).Call(
		dbx_async.Status("team/members/remove/job_status/get"))
	if err, fail := res.Failure(); fail {
		return err
	}
	return nil
}

func (z *memberImpl) Update(member *mo_member.Member) (updated *mo_member.Member, err error) {
	type US struct {
		Tag          string `json:".tag"`
		TeamMemberId string `json:"team_member_id"`
	}
	p := struct {
		User            US     `json:"user"`
		NewEmail        string `json:"new_email,omitempty"`
		NewExternalId   string `json:"new_external_id"`
		NewGivenName    string `json:"new_given_name,omitempty"`
		NewSurname      string `json:"new_surname,omitempty"`
		NewPersistentId string `json:"new_persistent_id,omitempty"`
	}{
		User: US{
			Tag:          "team_member_id",
			TeamMemberId: member.TeamMemberId,
		},
		NewEmail:        member.Email,
		NewExternalId:   member.ExternalId,
		NewGivenName:    member.GivenName,
		NewSurname:      member.Surname,
		NewPersistentId: member.PersistentId,
	}
	res := z.ctx.Post("team/members/set_profile", api_request.Param(p))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	updated = &mo_member.Member{}
	err = res.Success().Json().Model(updated)
	return
}

func (z *memberImpl) Resolve(teamMemberId string) (member *mo_member.Member, err error) {
	type US struct {
		Tag          string `json:".tag"`
		TeamMemberId string `json:"team_member_id"`
	}
	p := struct {
		Members []US `json:"members"`
	}{
		Members: []US{
			{
				Tag:          "team_member_id",
				TeamMemberId: teamMemberId,
			},
		},
	}
	res := z.ctx.Post("team/members/get_info", api_request.Param(p))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	return z.parseOneMember(res)
}

func (z *memberImpl) parseOneMember(res es_response.Response) (member *mo_member.Member, err error) {
	ba, found := res.Success().Json().Array()
	if !found || len(ba) < 1 {
		return nil, ErrorNotFound
	}
	a := ba[0]

	// id_not_found response:
	// {".tag": "id_not_found", "id_not_found": "xxx+xxxxx@xxxxxxxxx.xxx"}
	if id, found := a.FindString("id_not_found"); found {
		z.ctx.Log().Debug("`id_not_found`", zap.String("id", id))
		return nil, dbx_error.DropboxError{
			ErrorTag:     "id_not_found",
			ErrorSummary: "id_not_found",
		}
	}

	member = &mo_member.Member{}
	err = a.Model(member)
	return
}

func (z *memberImpl) ResolveByEmail(email string) (member *mo_member.Member, err error) {
	type US struct {
		Tag   string `json:".tag"`
		Email string `json:"email"`
	}
	p := struct {
		Members []US `json:"members"`
	}{
		Members: []US{
			{
				Tag:   "email",
				Email: email,
			},
		},
	}
	member = &mo_member.Member{}
	res := z.ctx.Post("team/members/get_info", api_request.Param(p))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	return z.parseOneMember(res)
}

func (z *memberImpl) List() (members []*mo_member.Member, err error) {
	members = make([]*mo_member.Member, 0)
	p := struct {
		IncludeRemoved bool `json:"include_removed,omitempty"`
		Limit          int  `json:"limit,omitempty"`
	}{
		IncludeRemoved: z.includeDeleted,
		Limit:          z.limit,
	}

	res := z.ctx.List("team/members/list", api_request.Param(p)).Call(
		dbx_list.Continue("team/members/list/continue"),
		dbx_list.UseHasMore(),
		dbx_list.ResultTag("members"),
		dbx_list.OnEntry(func(entry es_json.Json) error {
			m := &mo_member.Member{}
			if err := entry.Model(m); err != nil {
				return err
			}
			members = append(members, m)
			return nil
		}),
	)
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	return members, nil
}
