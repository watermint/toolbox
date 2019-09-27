package member

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_namespace"
	"github.com/watermint/toolbox/domain/service/sv_namespace"
	"github.com/watermint/toolbox/domain/service/sv_profile"
	"github.com/watermint/toolbox/domain/service/sv_sharedfolder_member"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_report"
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"go.uber.org/zap"
)

type ListVO struct {
	PeerName app_conn.ConnBusinessFile
}

type ListWorker struct {
	namespace *mo_namespace.Namespace
	ctx       api_context.Context // should be with admin team member id.
	rep       app_report.Report
	ctl       app_control.Control
}

func (z *ListWorker) Exec() error {
	ui := z.ctl.UI()
	ui.Info("recipe.team.namespace.member.list.scan",
		app_msg.P{
			"NamespaceName": z.namespace.Name,
			"NamespaceId":   z.namespace.NamespaceId,
		})
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
}

func (z *List) Requirement() app_vo.ValueObject {
	return &ListVO{}
}

func (z *List) Exec(k app_kitchen.Kitchen) error {
	l := k.Log()
	vo := k.Value().(*ListVO)

	ctx, err := vo.PeerName.Connect(k.Control())
	if err != nil {
		return err
	}

	admin, err := sv_profile.NewTeam(ctx).Admin()
	if err != nil {
		return err
	}
	l.Debug("Run as admin", zap.Any("admin", admin))

	namespaces, err := sv_namespace.New(ctx).List()
	if err != nil {
		return err
	}

	cta := ctx.AsAdminId(admin.TeamMemberId)

	rep, err := k.Report("namespace_member", &mo_namespace.NamespaceMember{})
	if err != nil {
		return err
	}
	defer rep.Close()

	q := k.NewQueue()
	for _, namespace := range namespaces {
		if namespace.NamespaceType != "team_folder" &&
			namespace.NamespaceType != "shared_folder" {
			l.Debug("Skip", zap.Any("namespace", namespace))
			continue
		}

		q.Enqueue(&ListWorker{
			namespace: namespace,
			ctx:       cta,
			rep:       rep,
			ctl:       k.Control(),
		})
	}
	q.Wait()
	return nil
}

func (z *List) Test(c app_control.Control) error {
	lvo := &ListVO{}
	if !app_test.ApplyTestPeers(c, lvo) {
		return nil
	}
	if err := z.Exec(app_kitchen.NewKitchen(c, lvo)); err != nil {
		return err
	}
	return app_test.TestRows(c, "namespace_member", func(cols map[string]string) error {
		if _, ok := cols["namespace_id"]; !ok {
			return errors.New("`namespace_id` is not found")
		}
		if _, ok := cols["email"]; !ok {
			return errors.New("`email` is not found")
		}
		return nil
	})
}
