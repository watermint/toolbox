package filerequest

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_filerequest"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_filerequest"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type List struct {
	Peer        dbx_conn.ConnScopedTeam
	FileRequest rp_model.RowReport
	BasePath    mo_string.SelectString
}

func (z *List) Preset() {
	z.FileRequest.SetModel(
		&mo_filerequest.MemberFileRequest{},
		rp_model.HiddenColumns(
			"account_id",
			"team_member_id",
			"file_request_id",
		),
	)
	z.Peer.SetScopes(
		dbx_auth.ScopeFileRequestsRead,
		dbx_auth.ScopeMembersRead,
		dbx_auth.ScopeTeamDataMember,
	)
	z.BasePath.SetOptions(
		dbx_filesystem.BaseNamespaceDefaultInString,
		dbx_filesystem.BaseNamespaceTypesInString...,
	)
}

func (z *List) scanMember(member *mo_member.Member, c app_control.Control) error {
	l := c.Log().With(esl.Any("mmeber", member))
	mc := z.Peer.Client().AsMemberId(member.TeamMemberId, dbx_filesystem.AsNamespaceType(z.BasePath.Value()))
	reqs, err := sv_filerequest.New(mc).List()
	if err != nil {
		l.Debug("Unable to retrieve file requests for the member", esl.Error(err))
		return err
	}
	for _, req := range reqs {
		fm := mo_filerequest.NewMemberFileRequest(req, member)
		z.FileRequest.Row(fm)
	}
	return nil
}

func (z *List) Exec(c app_control.Control) error {
	members, err := sv_member.New(z.Peer.Client()).List()
	if err != nil {
		return err
	}

	if err := z.FileRequest.Open(); err != nil {
		return err
	}

	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("scan", z.scanMember, c)
		scan := s.Get("scan")
		for _, member := range members {
			scan.Enqueue(member)
		}
	})

	return nil
}

func (z *List) Test(c app_control.Control) error {
	if err := rc_exec.Exec(c, &List{}, rc_recipe.NoCustomValues); err != nil {
		return err
	}
	return qtr_endtoend.TestRows(c, "file_request", func(cols map[string]string) error {
		if _, ok := cols["url"]; !ok {
			return errors.New("`url` is not found")
		}
		if _, ok := cols["email"]; !ok {
			return errors.New("`email` is not found")
		}
		return nil
	})
}
