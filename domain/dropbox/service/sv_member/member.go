package sv_member

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
)

var (
	ErrorMemberNotFoundForEmail        = errors.New("member not found for the email")
	ErrorMemberNotFoundForTeamMemberId = errors.New("member not found for the team_member_id")
	ErrorNotFound                      = errors.New("not found")
)

type Member interface {
	Update(member *mo_member.Member, opts ...UpdateOpt) (updated *mo_member.Member, err error)
	UpdateVisibility(email string, visible bool) (updated *mo_member.Member, err error)
	List(opts ...ListOpt) (members []*mo_member.Member, err error)
	ListEach(f func(member *mo_member.Member) bool, opts ...ListOpt) error
	Resolve(teamMemberId string) (member *mo_member.Member, err error)
	ResolveByEmail(email string) (member *mo_member.Member, err error)
	Add(email string, opts ...AddOpt) (member *mo_member.Member, err error)
	Remove(member *mo_member.Member, opts ...RemoveOpt) (err error)
	Suspend(member *mo_member.Member, opts ...SuspendOpt) (err error)
	Unsuspend(member *mo_member.Member) (err error)
}

type userSelectorArg struct {
	Tag          string `json:".tag"`
	Email        string `json:"email,omitempty"`
	TeamMemberId string `json:"team_member_id,omitempty"`
	ExternalId   string `json:"external_id,omitempty"`
}

type SuspendOpt func(opt suspendOpts) suspendOpts

func newSuspendOpt(email string) suspendOpts {
	return suspendOpts{
		User: userSelectorArg{
			Tag:   "email",
			Email: email,
		},
	}
}

func SuspendWipeData(enabled bool) SuspendOpt {
	return func(opt suspendOpts) suspendOpts {
		opt.WipeData = enabled
		return opt
	}
}

type suspendOpts struct {
	User     userSelectorArg `json:"user"`
	WipeData bool            `json:"wipe_data"`
}

func (z suspendOpts) Apply(opts []SuspendOpt) suspendOpts {
	switch len(opts) {
	case 0:
		return z
	case 1:
		return opts[0](z)
	default:
		return opts[0](z).Apply(opts[1:])
	}
}

type UpdateOpt func(opt updateOpts) updateOpts
type updateOpts struct {
	clearExternalid bool
}

func (z updateOpts) Apply(opts ...UpdateOpt) updateOpts {
	switch len(opts) {
	case 0:
		return z
	case 1:
		return opts[0](z)
	default:
		x, y := opts[0], opts[1:]
		return x(z).Apply(y...)
	}
}

func ClearExternalId() UpdateOpt {
	return func(opt updateOpts) updateOpts {
		opt.clearExternalid = true
		return opt
	}
}

type ListOpt func(opt listOpts) listOpts
type listOpts struct {
	includeRemoved bool
}

func (z listOpts) Apply(opts ...ListOpt) listOpts {
	switch len(opts) {
	case 0:
		return z
	case 1:
		return opts[0](z)
	default:
		x, y := opts[0], opts[1:]
		return x(z).Apply(y...)
	}
}

func IncludeDeleted(enabled bool) ListOpt {
	return func(opt listOpts) listOpts {
		opt.includeRemoved = enabled
		return opt
	}
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
func AddWithDirectoryRestricted(enabled bool) AddOpt {
	return func(opt *addOptions) *addOptions {
		opt.isDirectoryRestricted = enabled
		return opt
	}
}

type RemoveOpt func(opt *removeOptions) *removeOptions
type removeOptions struct {
	wipeData           bool
	keepAccount        bool
	retainTeamShares   bool
	transferDestEmail  string
	transferAdminEmail string
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
func TransferDest(email string) RemoveOpt {
	return func(opt *removeOptions) *removeOptions {
		opt.transferDestEmail = email
		return opt
	}
}
func TransferNotifyAdminOnError(email string) RemoveOpt {
	return func(opt *removeOptions) *removeOptions {
		opt.transferAdminEmail = email
		return opt
	}
}

func New(ctx dbx_client.Client) Member {
	return &memberImpl{
		ctx: ctx,
	}
}

func NewCached(ctx dbx_client.Client) Member {
	return &cachedMember{
		impl: &memberImpl{
			ctx: ctx,
		},
	}
}

func newTest(ctx dbx_client.Client) Member {
	return &memberImpl{
		ctx:   ctx,
		limit: 3,
	}
}
