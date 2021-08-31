package member

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_namespace"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_namespace"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_profile"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder_member"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type List struct {
	Peer            dbx_conn.ConnScopedTeam
	AllColumns      bool
	NamespaceMember rp_model.RowReport
}

func (z *List) Preset() {
	z.NamespaceMember.SetModel(&mo_namespace.NamespaceMember{}, rp_model.HiddenColumns(
		"account_id",
		"group_id",
		"namespace_team_member_id",
		"team_member_id",
		"namespace_id",
	))
	z.Peer.SetScopes(
		dbx_auth.ScopeSharingRead,
		dbx_auth.ScopeTeamDataMember,
		dbx_auth.ScopeTeamInfoRead,
	)
}

func (z *List) scanNamespace(namespace *mo_namespace.Namespace, c app_control.Control, ctx dbx_context.Context) error {
	l := c.Log().With(esl.Any("namespace", namespace))

	members, err := sv_sharedfolder_member.NewBySharedFolderId(ctx, namespace.NamespaceId).List()
	if err != nil {
		l.Debug("Unable to list namespace member", esl.Error(err))
		return nil
	}

	for _, member := range members {
		z.NamespaceMember.Row(mo_namespace.NewNamespaceMember(namespace, member))
	}
	return nil
}

func (z *List) Exec(c app_control.Control) error {
	l := c.Log()
	if err := z.NamespaceMember.Open(); err != nil {
		return err
	}

	admin, err := sv_profile.NewTeam(z.Peer.Context()).Admin()
	if err != nil {
		return err
	}
	l.Debug("Run as admin", esl.Any("admin", admin))

	namespaces, err := sv_namespace.New(z.Peer.Context()).List()
	if err != nil {
		return err
	}

	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("scan_namespace", z.scanNamespace, c, z.Peer.Context().AsAdminId(admin.TeamMemberId))
		q := s.Get("scan_namespace")
		for _, namespace := range namespaces {
			if namespace.NamespaceType != "team_folder" &&
				namespace.NamespaceType != "shared_folder" {
				l.Debug("Skip", esl.Any("namespace", namespace))
				continue
			}
			q.Enqueue(namespace)
		}
	})

	return nil
}

func (z *List) Test(c app_control.Control) error {
	if err := rc_exec.Exec(c, &List{}, rc_recipe.NoCustomValues); err != nil {
		return err
	}
	return qtr_endtoend.TestRows(c, "namespace_member", func(cols map[string]string) error {
		if _, ok := cols["namespace_name"]; !ok {
			return errors.New("`namespace_name` is not found")
		}
		if _, ok := cols["email"]; !ok {
			return errors.New("`email` is not found")
		}
		return nil
	})
}
