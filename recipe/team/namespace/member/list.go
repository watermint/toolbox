package member

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_namespace"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_namespace"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_profile"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder_member"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"go.uber.org/zap"
)

type MsgList struct {
	ProgressScan app_msg.Message
}

var (
	MList = app_msg.Apply(&MsgList{}).(*MsgList)
)

type ListVO struct {
}

type ListWorker struct {
	namespace *mo_namespace.Namespace
	ctx       dbx_context.Context // should be with admin team member id.
	rep       rp_model.RowReport
	ctl       app_control.Control
}

func (z *ListWorker) Exec() error {
	ui := z.ctl.UI()
	ui.Progress(MList.ProgressScan.
		With("NamespaceName", z.namespace.Name).
		With("NamespaceId", z.namespace.NamespaceId))
	l := z.ctl.Log().With(zap.Any("namespace", z.namespace))

	members, err := sv_sharedfolder_member.NewBySharedFolderId(z.ctx, z.namespace.NamespaceId).List()
	if err != nil {
		l.Debug("Unable to list namespace member", zap.Error(err))
		return nil
	}

	for _, member := range members {
		z.rep.Row(mo_namespace.NewNamespaceMember(z.namespace, member))
	}
	return nil
}

type List struct {
	Peer            dbx_conn.ConnBusinessFile
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
	l.Debug("Run as admin", zap.Any("admin", admin))

	namespaces, err := sv_namespace.New(z.Peer.Context()).List()
	if err != nil {
		return err
	}

	cta := z.Peer.Context().AsAdminId(admin.TeamMemberId)

	q := c.NewQueue()
	for _, namespace := range namespaces {
		if namespace.NamespaceType != "team_folder" &&
			namespace.NamespaceType != "shared_folder" {
			l.Debug("Skip", zap.Any("namespace", namespace))
			continue
		}

		q.Enqueue(&ListWorker{
			namespace: namespace,
			ctx:       cta,
			rep:       z.NamespaceMember,
			ctl:       c,
		})
	}
	q.Wait()
	return nil
}

func (z *List) Test(c app_control.Control) error {
	if err := rc_exec.Exec(c, &List{}, rc_recipe.NoCustomValues); err != nil {
		return err
	}
	return qt_recipe.TestRows(c, "namespace_member", func(cols map[string]string) error {
		if _, ok := cols["namespace_name"]; !ok {
			return errors.New("`namespace_name` is not found")
		}
		if _, ok := cols["email"]; !ok {
			return errors.New("`email` is not found")
		}
		return nil
	})
}
