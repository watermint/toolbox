package sv_sharedfolder_member

import (
	"errors"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_parser"
	"github.com/watermint/toolbox/domain/infra/api_rpc"
	"github.com/watermint/toolbox/domain/model/mo_group"
	"github.com/watermint/toolbox/domain/model/mo_profile"
	"github.com/watermint/toolbox/domain/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/model/mo_sharedfolder_member"
	"github.com/watermint/toolbox/domain/model/mo_teamfolder"
)

func New(ctx api_context.Context, sf *mo_sharedfolder.SharedFolder) Member {
	return &memberImpl{
		ctx:            ctx,
		sharedFolderId: sf.SharedFolderId,
	}
}

func NewByTeamFolder(ctx api_context.Context, tf *mo_teamfolder.TeamFolder) Member {
	return &memberImpl{
		ctx:            ctx,
		sharedFolderId: tf.TeamFolderId,
	}
}
func NewBySharedFolderId(ctx api_context.Context, sfId string) Member {
	return &memberImpl{
		ctx:            ctx,
		sharedFolderId: sfId,
	}
}

type Member interface {
	List() (member []mo_sharedfolder_member.Member, err error)
	Add(members []MemberAddOption, opts ...AddOption) (err error)
	Remove(member MemberRemoveOption, opts ...RemoveOption) (err error)
}

func AddByEmail(email, accessLevel string) MemberAddOption {
	return func(opt *memberAddOptions) *memberAddOptions {
		opt.memberEmails[email] = accessLevel
		return opt
	}
}
func AddByProfile(profile *mo_profile.Profile, accessLevel string) MemberAddOption {
	return func(opt *memberAddOptions) *memberAddOptions {
		if profile.TeamMemberId != "" {
			opt.memberDropboxId[profile.TeamMemberId] = accessLevel
		} else {
			opt.memberDropboxId[profile.AccountId] = accessLevel
		}
		return opt
	}
}
func AddByTeamMemberId(teamMemberId, accessLevel string) MemberAddOption {
	return func(opt *memberAddOptions) *memberAddOptions {
		opt.memberDropboxId[teamMemberId] = accessLevel
		return opt
	}
}
func AddByGroup(group *mo_group.Group, accessLevel string) MemberAddOption {
	return func(opt *memberAddOptions) *memberAddOptions {
		opt.memberDropboxId[group.GroupId] = accessLevel
		return opt
	}
}
func AddByGroupId(groupId, accessLevel string) MemberAddOption {
	return func(opt *memberAddOptions) *memberAddOptions {
		opt.memberDropboxId[groupId] = accessLevel
		return opt
	}
}

type MemberAddOption func(opt *memberAddOptions) *memberAddOptions

type memberAddOptions struct {
	// email-permission mapping
	memberEmails map[string]string

	// dropboxId-permission mapping
	memberDropboxId map[string]string
}

type AddOption func(opt *addOptions) *addOptions

type addOptions struct {
	quiet         bool
	customMessage string
}

type RemoveOption func(opt *removeOptions) *removeOptions

type removeOptions struct {
	leaveACopy bool
}

func LeaveACopy() RemoveOption {
	return func(opt *removeOptions) *removeOptions {
		opt.leaveACopy = true
		return opt
	}
}

type MemberRemoveOption func(opt *memberRemoveOptions) *memberRemoveOptions

type memberRemoveOptions struct {
	// email
	memberEmails []string

	// dropboxId
	memberDropboxId []string
}

func RemoveByEmail(email string) MemberRemoveOption {
	return func(opt *memberRemoveOptions) *memberRemoveOptions {
		opt.memberEmails = append(opt.memberEmails, email)
		return opt
	}
}

func RemoveByProfile(profile *mo_profile.Profile) MemberRemoveOption {
	return func(opt *memberRemoveOptions) *memberRemoveOptions {
		if profile.TeamMemberId != "" {
			opt.memberDropboxId = append(opt.memberDropboxId, profile.TeamMemberId)
		} else {
			opt.memberDropboxId = append(opt.memberDropboxId, profile.AccountId)
		}
		return opt
	}
}

func RemoveByTeamMemberId(teamMemberId string) MemberRemoveOption {
	return func(opt *memberRemoveOptions) *memberRemoveOptions {
		opt.memberDropboxId = append(opt.memberDropboxId, teamMemberId)
		return opt
	}
}

func RemoveByGroup(group *mo_group.Group) MemberRemoveOption {
	return func(opt *memberRemoveOptions) *memberRemoveOptions {
		opt.memberDropboxId = append(opt.memberDropboxId, group.GroupId)
		return opt
	}
}

func RemoveByGroupId(groupId string) MemberRemoveOption {
	return func(opt *memberRemoveOptions) *memberRemoveOptions {
		opt.memberDropboxId = append(opt.memberDropboxId, groupId)
		return opt
	}
}

type memberImpl struct {
	ctx            api_context.Context
	sharedFolderId string
}

func (z *memberImpl) List() (member []mo_sharedfolder_member.Member, err error) {
	member = make([]mo_sharedfolder_member.Member, 0)

	p := struct {
		SharedFolderId string `json:"shared_folder_id"`
	}{
		SharedFolderId: z.sharedFolderId,
	}

	err = z.ctx.List("sharing/list_folder_members").
		Continue("sharing/list_folder_members/continue").
		Param(p).
		UseHasMore(false).
		OnResponse(func(res api_rpc.Response) error {
			j, err := res.Json()
			if err != nil {
				return err
			}
			users := j.Get("users")
			if users.Exists() && users.IsArray() {
				for _, u := range users.Array() {
					mu := &mo_sharedfolder_member.User{}
					if err := api_parser.ParseModel(mu, u); err != nil {
						return err
					}
					member = append(member, mu)
				}
			}
			groups := j.Get("groups")
			if groups.Exists() && groups.IsArray() {
				for _, g := range groups.Array() {
					mg := &mo_sharedfolder_member.Group{}
					if err := api_parser.ParseModel(mg, g); err != nil {
						return err
					}
					member = append(member, mg)
				}
			}
			invitees := j.Get("invitees")
			if invitees.Exists() && invitees.IsArray() {
				for _, i := range invitees.Array() {
					mi := &mo_sharedfolder_member.Invitee{}
					if err := api_parser.ParseModel(mi, i); err != nil {
						return err
					}
					member = append(member, mi)
				}
			}
			return nil
		}).Call()
	if err != nil {
		return nil, err
	}
	return member, nil
}

func (z *memberImpl) Add(members []MemberAddOption, opts ...AddOption) (err error) {
	mem := &memberAddOptions{
		memberEmails:    make(map[string]string),
		memberDropboxId: make(map[string]string),
	}
	ao := &addOptions{}

	for _, m := range members {
		m(mem)
	}
	for _, o := range opts {
		o(ao)
	}

	type MS struct {
		Tag       string `json:".tag"`
		Email     string `json:"email,omitempty"`
		DropboxId string `json:"dropbox_id,omitempty"`
	}
	type AM struct {
		Member      MS     `json:"member"`
		AccessLevel string `json:"access_level"`
	}
	mems := make([]*AM, 0)
	for m, a := range mem.memberEmails {
		mems = append(mems, &AM{
			Member: MS{
				Tag:   "email",
				Email: m,
			},
			AccessLevel: a,
		})
	}
	for m, a := range mem.memberDropboxId {
		mems = append(mems, &AM{
			Member: MS{
				Tag:       "dropbox_id",
				DropboxId: m,
			},
			AccessLevel: a,
		})
	}
	p := struct {
		SharedFolderId string `json:"shared_folder_id"`
		Members        []*AM  `json:"members"`
		Quiet          bool   `json:"quiet,omitempty"`
		CustomMessage  string `json:"custom_message,omitempty"`
	}{
		SharedFolderId: z.sharedFolderId,
		Members:        mems,
		Quiet:          ao.quiet,
		CustomMessage:  ao.customMessage,
	}

	_, err = z.ctx.Request("sharing/add_folder_member").Param(p).Call()
	if err != nil {
		return err
	}
	return err
}

func (z *memberImpl) Remove(member MemberRemoveOption, opts ...RemoveOption) (err error) {
	mem := &memberRemoveOptions{
		memberDropboxId: make([]string, 0),
		memberEmails:    make([]string, 0),
	}
	ro := &removeOptions{}
	member(mem)
	for _, o := range opts {
		o(ro)
	}

	type MS struct {
		Tag       string `json:".tag"`
		Email     string `json:"email,omitempty"`
		DropboxId string `json:"dropbox_id,omitempty"`
	}
	mems := make([]*MS, 0)
	for _, m := range mem.memberEmails {
		mems = append(mems, &MS{
			Tag:   "email",
			Email: m,
		})
	}
	for _, m := range mem.memberDropboxId {
		mems = append(mems, &MS{
			Tag:       "dropbox_id",
			DropboxId: m,
		})
	}
	if len(mems) != 1 {
		return errors.New("invalid number of member arguments")
	}
	p := struct {
		SharedFolderId string `json:"shared_folder_id"`
		Member         *MS    `json:"member"`
		LeaveACopy     bool   `json:"leave_a_copy"`
	}{
		SharedFolderId: z.sharedFolderId,
		Member:         mems[0],
		LeaveACopy:     ro.leaveACopy,
	}
	_, err = z.ctx.Async("sharing/remove_folder_member").
		Status("sharing/check_remove_member_job_status").
		Param(p).
		Call()
	if err != nil {
		return err
	}
	return nil
}
