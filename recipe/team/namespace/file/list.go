package file

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/domain/model/mo_namespace"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_file"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/domain/service/sv_namespace"
	"github.com/watermint/toolbox/domain/service/sv_profile"
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

type ListVO struct {
	Peer                rc_conn.ConnBusinessFile
	IncludeMediaInfo    bool
	IncludeDeleted      bool
	IncludeMemberFolder bool
	IncludeSharedFolder bool
	IncludeTeamFolder   bool
	Name                string
}

type ListWorker struct {
	namespace  *mo_namespace.Namespace
	idToMember map[string]*mo_member.Member
	ctx        api_context.Context
	ctl        app_control.Control
	rep        rp_model.Report
	vo         *ListVO
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

	err := sv_file.NewFiles(ctn).ListChunked(mo_path.NewDropboxPath(""), func(entry mo_file.Entry) {
		ne := mo_namespace.NewNamespaceEntry(z.namespace, entry.Concrete())
		if m, e := z.idToMember[z.namespace.TeamMemberId]; e {
			ne.NamespaceMemberEmail = m.Email
		}
		z.rep.Row(ne)
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

const (
	reportList = "namespace_file"
)

type List struct {
}

func (z *List) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{
		rp_spec_impl.Spec(reportList, &mo_namespace.NamespaceEntry{}),
	}
}

func (z *List) Requirement() rc_vo.ValueObject {
	return &ListVO{
		IncludeTeamFolder:   true,
		IncludeSharedFolder: true,
		IncludeMemberFolder: false,
	}
}

func (z *List) Exec(k rc_kitchen.Kitchen) error {
	l := k.Log()
	vo := k.Value().(*ListVO)

	ctx, err := vo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}

	admin, err := sv_profile.NewTeam(ctx).Admin()
	if err != nil {
		return err
	}
	l.Debug("Run as admin", zap.Any("admin", admin))

	members, err := sv_member.New(ctx).List()
	if err != nil {
		return err
	}
	idToMember := mo_member.MapByTeamMemberId(members)

	namespaces, err := sv_namespace.New(ctx).List()
	if err != nil {
		return err
	}

	cta := ctx.AsAdminId(admin.TeamMemberId)

	rep, err := rp_spec_impl.New(z, k.Control()).Open(reportList)
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
		case vo.IncludeMemberFolder && namespace.NamespaceType == "app_folder":
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
			namespace:  namespace,
			idToMember: idToMember,
			ctx:        cta,
			rep:        rep,
			vo:         vo,
			ctl:        k.Control(),
		})
	}
	q.Wait()
	return nil
}

func (z *List) Test(c app_control.Control) error {
	lvo := &ListVO{
		Name: qt_recipe.TestTeamFolderName,
	}
	if !qt_recipe.ApplyTestPeers(c, lvo) {
		return qt_recipe.NotEnoughResource()
	}
	if err := z.Exec(rc_kitchen.NewKitchen(c, lvo)); err != nil {
		return err
	}
	return qt_recipe.TestRows(c, "namespace_file", func(cols map[string]string) error {
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
