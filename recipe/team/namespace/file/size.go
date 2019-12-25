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
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/recipe/rc_vo"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/report/rp_spec_impl"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"go.uber.org/zap"
)

type SizeVO struct {
	Peer                rc_conn.OldConnBusinessFile
	IncludeSharedFolder bool
	IncludeTeamFolder   bool
	IncludeMemberFolder bool
	IncludeAppFolder    bool
	Name                string
	Depth               int
}

type SizeWorker struct {
	namespace *mo_namespace.Namespace
	ctx       api_context.Context
	ctl       app_control.Control
	rep       rp_model.SideCarReport
	vo        *SizeVO
	k         rc_kitchen.Kitchen
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

	var lastErr error
	sizes, errs := uc_file_size.New(ctn, z.k).Size(mo_path.NewDropboxPath("/"), z.vo.Depth)

	for p, size := range sizes {
		if err, ok := errs[p]; ok {
			l.Debug("Unable to traverse", zap.Error(err))
			ui.Error("recipe.team.namespace.file.size.err.scan_failed",
				app_msg.P{
					"NamespaceName": z.namespace.Name,
					"NamespaceId":   z.namespace.NamespaceId,
					"Error":         err.Error(),
				},
			)
			lastErr = err
			z.rep.Failure(err, z.namespace)
		} else {
			z.rep.Success(z.namespace, mo_file_size.NewNamespaceSize(z.namespace, size))
		}
	}

	return lastErr
}

const (
	reportSize = "namespace_size"
)

type Size struct {
}

func (z *Size) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{
		rp_spec_impl.Spec(
			reportSize,
			rp_model.TransactionHeader(
				&mo_namespace.Namespace{},
				&mo_file_size.NamespaceSize{},
			),
			rp_model.HiddenColumns(
				"result.namespace_name",
				"result.namespace_id",
				"result.namespace_type",
				"result.owner_team_member_id",
			),
		),
	}
}

func (z *Size) Requirement() rc_vo.ValueObject {
	return &SizeVO{
		IncludeSharedFolder: true,
		IncludeTeamFolder:   true,
		Depth:               1,
	}
}

func (z *Size) Exec(k rc_kitchen.Kitchen) error {
	l := k.Log()
	vo := k.Value().(*SizeVO)

	if vo.Depth < 1 {
		return errors.New("depth should grater than 1")
	}

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

	rep, err := rp_spec_impl.New(z, k.Control()).Open(reportSize)
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
		case vo.IncludeMemberFolder && namespace.NamespaceType == "team_member_folder":
			process = true
		case vo.IncludeAppFolder && namespace.NamespaceType == "app_folder":
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
			k:         k,
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
		Name:              qt_recipe.TestTeamFolderName,
		IncludeTeamFolder: false,
		Depth:             1,
	}
	if !qt_recipe.ApplyTestPeers(c, lvo) {
		return qt_recipe.NotEnoughResource()
	}
	if err := z.Exec(rc_kitchen.NewKitchen(c, lvo)); err != nil {
		return err
	}
	return qt_recipe.TestRows(c, "namespace_size", func(cols map[string]string) error {
		if _, ok := cols["input.namespace_id"]; !ok {
			return errors.New("`namespace_id` is not found")
		}
		if _, ok := cols["result.size"]; !ok {
			return errors.New("`size` is not found")
		}
		return nil
	})
}
