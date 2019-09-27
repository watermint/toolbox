package file

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_namespace"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_file"
	"github.com/watermint/toolbox/domain/service/sv_namespace"
	"github.com/watermint/toolbox/domain/service/sv_profile"
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
	PeerName            app_conn.ConnBusinessFile
	IncludeMediaInfo    bool
	IncludeDeleted      bool
	IncludeSharedFolder bool
	IncludeTeamFolder   bool
	Name                string
}

type ListWorker struct {
	namespace *mo_namespace.Namespace
	ctx       api_context.Context
	ctl       app_control.Control
	rep       app_report.Report
	vo        *ListVO
}

func (z *ListWorker) Exec() error {
	ui := z.ctl.UI()
	ui.Info("recipe.team.namespace.file.list.scan",
		app_msg.P{
			"NamespaceName": z.namespace.Name,
			"NamespaceId":   z.namespace.NamespaceId,
		},
	)
	l := z.ctl.Log().With(zap.Any("namespace", z.namespace))

	ctn := z.ctx.WithPath(api_context.Namespace(z.namespace.NamespaceId))

	opts := make([]sv_file.ListOpt, 0)
	if z.vo.IncludeDeleted {
		opts = append(opts, sv_file.IncludeDeleted())
	}
	if z.vo.IncludeMediaInfo {
		opts = append(opts, sv_file.IncludeMediaInfo())
	}
	opts = append(opts, sv_file.IncludeHasExplicitSharedMembers())
	opts = append(opts, sv_file.Recursive())

	err := sv_file.NewFiles(ctn).ListChunked(mo_path.NewPath(""), func(entry mo_file.Entry) {
		z.rep.Row(mo_namespace.NewNamespaceEntry(z.namespace, entry.Concrete()))
	}, opts...)

	if err != nil {
		l.Debug("Unable to traverse", zap.Error(err))
		ui.Error("recipe.team.namespace.file.list.err.scan_failed",
			app_msg.P{
				"NamespaceName": z.namespace.Name,
				"NamespaceId":   z.namespace.NamespaceId,
				"Error":         err.Error(),
			},
		)
		return err
	}
	return nil
}

type List struct {
}

func (z *List) Requirement() app_vo.ValueObject {
	return &ListVO{
		IncludeTeamFolder:   true,
		IncludeSharedFolder: true,
	}
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

	rep, err := k.Report("namespace_file", &mo_namespace.NamespaceEntry{})
	if err != nil {
		return err
	}
	defer rep.Close()

	q := k.NewQueue()
	for _, namespace := range namespaces {
		process := false
		switch {
		case vo.IncludeTeamFolder && namespace.NamespaceType == "team_folder":
			process = true
		case vo.IncludeSharedFolder && namespace.NamespaceType == "shared_folder":
			process = true
		}
		if !process {
			l.Debug("Skip", zap.Any("namespace", namespace))
			continue
		}
		if vo.Name != "" && namespace.Name != vo.Name {
			l.Debug("Skip", zap.Any("namespace", namespace), zap.String("filter", vo.Name))
			continue
		}

		q.Enqueue(&ListWorker{
			namespace: namespace,
			ctx:       cta,
			rep:       rep,
			vo:        vo,
			ctl:       k.Control(),
		})
	}
	q.Wait()
	return nil
}

func (z *List) Test(c app_control.Control) error {
	lvo := &ListVO{
		Name: app_test.TestTeamFolderName,
	}
	if !app_test.ApplyTestPeers(c, lvo) {
		return nil
	}
	if err := z.Exec(app_kitchen.NewKitchen(c, lvo)); err != nil {
		return err
	}
	return app_test.TestRows(c, "namespace_file", func(cols map[string]string) error {
		if _, ok := cols["namespace_id"]; !ok {
			return errors.New("`namespace_id` is not found")
		}
		if _, ok := cols["id"]; !ok {
			return errors.New("`id` is not found")
		}
		if _, ok := cols["path_display"]; !ok {
			return errors.New("`path_display` is not found")
		}
		return nil
	})
}
