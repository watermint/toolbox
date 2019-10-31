package file

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_file_size"
	"github.com/watermint/toolbox/domain/model/mo_namespace"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_namespace"
	"github.com/watermint/toolbox/domain/service/sv_profile"
	"github.com/watermint/toolbox/domain/usecase/uc_file_size"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/quality/qt_test"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_report"
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"go.uber.org/zap"
)

type SizeVO struct {
	Peer                app_conn.ConnBusinessFile
	IncludeSharedFolder bool
	IncludeTeamFolder   bool
	Name                string
	Depth               int
}

type SizeWorker struct {
	namespace *mo_namespace.Namespace
	ctx       api_context.Context
	ctl       app_control.Control
	rep       app_report.Report
	vo        *SizeVO
}

func (z *SizeWorker) Exec() error {
	ui := z.ctl.UI()
	ui.Info("recipe.team.namespace.file.size.scan",
		app_msg.P{
			"NamespaceName": z.namespace.Name,
			"NamespaceId":   z.namespace.NamespaceId,
		},
	)
	l := z.ctl.Log().With(zap.Any("namespace", z.namespace))

	ctn := z.ctx.WithPath(api_context.Namespace(z.namespace.NamespaceId))

	sizes, err := uc_file_size.New(ctn).Size(mo_path.NewPath("/"), z.vo.Depth)
	if err != nil {
		l.Debug("Unable to traverse", zap.Error(err))
		ui.Error("recipe.team.namespace.file.size.err.scan_failed",
			app_msg.P{
				"NamespaceName": z.namespace.Name,
				"NamespaceId":   z.namespace.NamespaceId,
				"Error":         err.Error(),
			},
		)
		return err
	}

	for _, size := range sizes {
		z.rep.Row(mo_file_size.NewNamespaceSize(z.namespace, size))
	}

	return nil
}

type Size struct {
}

func (z *Size) Requirement() app_vo.ValueObject {
	return &SizeVO{
		IncludeSharedFolder: true,
		IncludeTeamFolder:   true,
		Depth:               2,
	}
}

func (z *Size) Exec(k app_kitchen.Kitchen) error {
	l := k.Log()
	vo := k.Value().(*SizeVO)

	ctx, err := vo.Peer.Connect(k.Control())
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

	rep, err := k.Report("namespace_size", &mo_file_size.NamespaceSize{})
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

		q.Enqueue(&SizeWorker{
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

func (z *Size) Test(c app_control.Control) error {
	lvo := &SizeVO{
		Name: app_test.TestTeamFolderName,
	}
	if !app_test.ApplyTestPeers(c, lvo) {
		return qt_test.NotEnoughResource()
	}
	if err := z.Exec(app_kitchen.NewKitchen(c, lvo)); err != nil {
		return err
	}
	return app_test.TestRows(c, "namespace_size", func(cols map[string]string) error {
		if _, ok := cols["namespace_id"]; !ok {
			return errors.New("`namespace_id` is not found")
		}
		if _, ok := cols["size"]; !ok {
			return errors.New("`size` is not found")
		}
		return nil
	})
}
