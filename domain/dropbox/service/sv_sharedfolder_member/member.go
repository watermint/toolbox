package sv_sharedfolder_member

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_group"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_profile"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder_member"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_teamfolder"
	"github.com/watermint/toolbox/essentials/http/response"
)

const (
	LevelOwner           = "owner"
	LevelEditor          = "editor"
	LevelViewer          = "viewer"
	LevelViewerNoComment = "viewer_no_comment"
)

func New(ctx dbx_context.Context, sf *mo_sharedfolder.SharedFolder) Member {
	return &memberImpl{
		ctx:            ctx,
		sharedFolderId: sf.SharedFolderId,
	}
}

func NewByTeamFolder(ctx dbx_context.Context, tf *mo_teamfolder.TeamFolder) Member {
	return &memberImpl{
		ctx:            ctx,
		sharedFolderId: tf.TeamFolderId,
	}
}
func NewBySharedFolderId(ctx dbx_context.Context, sfId string) Member {
	return &memberImpl{
		ctx:            ctx,
		sharedFolderId: sfId,
	}
}
func NewCached(ctx dbx_context.Context, sfId string) Member {
	return &cachedMember{
		impl: &memberImpl{
			ctx:            ctx,
			sharedFolderId: sfId,
		},
	}
}

type Member interface {
	List() (member []mo_sharedfolder_member.Member, err error)
	Add(member MemberAddOption, opts ...AddOption) (err error)
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

func AddQuiet() AddOption {
	return func(opt *addOptions) *addOptions {
		opt.quiet = true
		return opt
	}
}
func AddCustomMessage(message string) AddOption {
	return func(opt *addOptions) *addOptions {
		opt.customMessage = message
		return opt
	}
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

type cachedMember struct {
	impl    Member
	members []mo_sharedfolder_member.Member
}

func (z *cachedMember) List() (member []mo_sharedfolder_member.Member, err error) {
	if z.members == nil {
		z.members, err = z.impl.List()
		if err != nil {
			return nil, err
		}
	}
	return z.members, nil
}

func (z *cachedMember) Add(member MemberAddOption, opts ...AddOption) (err error) {
	z.members = nil // invalidate
	return z.impl.Add(member, opts...)
}

func (z *cachedMember) Remove(member MemberRemoveOption, opts ...RemoveOption) (err error) {
	z.members = nil // invalidate
	return z.impl.Remove(member, opts...)
}

type memberImpl struct {
	ctx            dbx_context.Context
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
		OnResponse(func(res response.Response) error {
			j, err := res.Success().AsJson()
			if err != nil {
				return err
			}
			if users, found := j.FindArray("users"); found {
				for _, u := range users {
					mu := &mo_sharedfolder_member.User{}
					if _, err := u.Model(mu); err != nil {
						return err
					}
					member = append(member, mu)
				}
			}
			if groups, found := j.FindArray("groups"); found {
				for _, g := range groups {
					mg := &mo_sharedfolder_member.Group{}
					if _, err := g.Model(mg); err != nil {
						return err
					}
					member = append(member, mg)
				}
			}
			if invitees, found := j.FindArray("invitees"); found {
				for _, i := range invitees {
					mi := &mo_sharedfolder_member.Invitee{}
					if _, err := i.Model(mi); err != nil {
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

func (z *memberImpl) Add(member MemberAddOption, opts ...AddOption) (err error) {
	mem := &memberAddOptions{
		memberEmails:    make(map[string]string),
		memberDropboxId: make(map[string]string),
	}
	ao := &addOptions{}

	member(mem)
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

	_, err = z.ctx.Post("sharing/add_folder_member").Param(p).Call()
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
