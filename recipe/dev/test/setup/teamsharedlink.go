package setup

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_group_member"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedlink"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedlink"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/security/sc_random"
	"math/rand"
	"time"
)

func SelectOpts(mode string, r *rand.Rand) []sv_sharedlink.LinkOpt {
	switch mode {
	case "public":
		return []sv_sharedlink.LinkOpt{sv_sharedlink.Public()}
	case "team_only":
		return []sv_sharedlink.LinkOpt{sv_sharedlink.TeamOnly()}
	case "with_expire":
		return []sv_sharedlink.LinkOpt{
			sv_sharedlink.Expires(time.Now().Add(time.Duration(r.Intn(365*86400)) * time.Second)),
		}
	case "with_password":
		return []sv_sharedlink.LinkOpt{
			sv_sharedlink.Password(sc_random.MustGetPseudoRandomString(r, 10)),
		}
	default:
		opts := make([]sv_sharedlink.LinkOpt, 0)
		switch r.Intn(4) {
		case 0, 2:
			opts = append(opts, sv_sharedlink.Public())
		case 1, 3:
			opts = append(opts, sv_sharedlink.TeamOnly())
		}
		switch r.Intn(10) {
		case 0, 1, 2:
			opts = append(opts, sv_sharedlink.Password(sc_random.MustGetPseudoRandomString(r, r.Intn(15)+1)))
		}
		switch r.Intn(10) {
		case 0, 1, 2:
			opts = append(opts, sv_sharedlink.Expires(time.Now().Add(time.Duration(r.Intn(365*86400))*time.Second)))
		}
		return opts
	}
}

type Teamsharedlink struct {
	rc_recipe.RemarkSecret
	rc_recipe.RemarkIrreversible
	rc_recipe.RemarkLicenseRequired
	Peer              dbx_conn.ConnScopedTeam
	Group             string
	Query             string
	Visibility        mo_string.SelectString
	NumLinksPerMember int
	Seed              int64
	Created           rp_model.RowReport
}

func (z *Teamsharedlink) Preset() {
	z.NumLinksPerMember = 5
	z.Visibility.SetOptions(
		"random",
		"random",
		"public",
		"team_only",
		"with_expire",
		"with_password",
	)
	z.Created.SetModel(&mo_sharedlink.Metadata{})
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentWrite,
		dbx_auth.ScopeSharingWrite,
		dbx_auth.ScopeGroupsRead,
		dbx_auth.ScopeTeamDataMember,
	)
}

func (z *Teamsharedlink) createForMember(member *mo_group_member.Member, c app_control.Control, r *rand.Rand) error {
	l := c.Log().With(esl.String("member", member.Email))
	files, err := sv_file.NewFiles(z.Peer.Client().AsMemberId(member.TeamMemberId)).Search(
		z.Query,
		sv_file.SearchMaxResults(z.NumLinksPerMember*2),
	)
	if err != nil {
		l.Debug("Unable to search a file", esl.Error(err))
		return err
	}

	linksCreated := 0
	for _, file := range files {
		if file.EntryPathDisplay == "" {
			continue
		}
		link, err := sv_sharedlink.New(z.Peer.Client().AsMemberId(member.TeamMemberId)).Create(
			mo_path.NewDropboxPath(file.EntryPathDisplay),
			SelectOpts(z.Visibility.Value(), r)...,
		)
		if err != nil {
			dbxErr := dbx_error.NewErrors(err)
			if dbxErr.IsSharedLinkAlreadyExists() {
				continue
			}
			return err
		}
		l.Debug("Link created", esl.String("url", link.LinkUrl()))
		z.Created.Row(link.Metadata())
		linksCreated++
		if z.NumLinksPerMember < linksCreated {
			return nil
		}
	}
	return nil
}

func (z *Teamsharedlink) Exec(c app_control.Control) (err error) {
	if err := z.Created.Open(); err != nil {
		return err
	}

	group, err := sv_group.New(z.Peer.Client()).ResolveByName(z.Group)
	if err != nil {
		return err
	}

	members, err := sv_group_member.New(z.Peer.Client(), group).List()
	if err != nil {
		return err
	}

	r := rand.New(rand.NewSource(z.Seed))

	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("create_for_member", z.createForMember, c, r)
		q := s.Get("create_for_member")
		for _, member := range members {
			if member.Profile().EmailVerified {
				q.Enqueue(member)
			}
		}
	})
	return nil
}

func (z *Teamsharedlink) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Teamsharedlink{}, func(r rc_recipe.Recipe) {
		m := r.(*Teamsharedlink)
		m.Group = "watermint"
		m.Query = "toolbox"
	})
}
