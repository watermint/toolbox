package file

import (
	"errors"
	"github.com/watermint/toolbox/domain/common/model/mo_string"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file_size"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_namespace"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_namespace"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_profile"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_file_size"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type MsgSize struct {
	ProgressScan    app_msg.Message
	ErrorScanFailed app_msg.Message
}

var (
	MSize = app_msg.Apply(&MsgSize{}).(*MsgSize)
)

type SizeWorker struct {
	namespace *mo_namespace.Namespace
	ctx       dbx_context.Context
	ctl       app_control.Control
	rep       rp_model.TransactionReport
	depth     int
}

func (z *SizeWorker) Exec() error {
	ui := z.ctl.UI()
	ui.Progress(MSize.ProgressScan.
		With("NamespaceName", z.namespace.Name).
		With("NamespaceId", z.namespace.NamespaceId))
	l := z.ctl.Log().With(esl.Any("namespace", z.namespace))

	ctn := z.ctx.WithPath(dbx_context.Namespace(z.namespace.NamespaceId))

	var lastErr error
	sizes, errs := uc_file_size.New(ctn, z.ctl).Size(mo_path.NewDropboxPath("/"), z.depth)

	for p, size := range sizes {
		if err, ok := errs[p]; ok {
			l.Debug("Unable to traverse", esl.Error(err))
			ui.Error(MSize.ErrorScanFailed.
				With("NamespaceName", z.namespace.Name).
				With("NamespaceId", z.namespace.NamespaceId).
				With("Error", err.Error()))

			lastErr = err
			z.rep.Failure(err, z.namespace)
		} else {
			z.rep.Success(z.namespace, mo_file_size.NewNamespaceSize(z.namespace, size))
		}
	}

	return lastErr
}

type Size struct {
	Peer                dbx_conn.ConnBusinessFile
	IncludeSharedFolder bool
	IncludeTeamFolder   bool
	IncludeMemberFolder bool
	IncludeAppFolder    bool
	Name                mo_string.OptionalString
	Depth               int
	NamespaceSize       rp_model.TransactionReport
}

func (z *Size) Preset() {
	z.NamespaceSize.SetModel(
		&mo_namespace.Namespace{},
		&mo_file_size.NamespaceSize{},
		rp_model.HiddenColumns(
			"result.namespace_name",
			"result.namespace_id",
			"result.namespace_type",
			"result.owner_team_member_id",
			"input.team_member_id",
			"input.namespace_id",
		),
	)
	z.IncludeSharedFolder = true
	z.IncludeTeamFolder = true
	z.Depth = 1
}

func (z *Size) Exec(c app_control.Control) error {
	l := c.Log()

	if z.Depth < 1 {
		return errors.New("depth should grater than 1")
	}

	if err := z.NamespaceSize.Open(); err != nil {
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

	cta := z.Peer.Context().AsAdminId(admin.TeamMemberId)

	q := c.NewLegacyQueue()
	for _, namespace := range namespaces {
		process := false
		switch {
		case z.IncludeTeamFolder && namespace.NamespaceType == "team_folder":
			process = true
		case z.IncludeSharedFolder && namespace.NamespaceType == "shared_folder":
			process = true
		case z.IncludeMemberFolder && namespace.NamespaceType == "team_member_folder":
			process = true
		case z.IncludeAppFolder && namespace.NamespaceType == "app_folder":
			process = true
		}
		if !process {
			l.Debug("Skip", esl.Any("namespace", namespace))
			continue
		}
		if z.Name.IsExists() && namespace.Name != z.Name.Value() {
			l.Debug("Skip", esl.Any("namespace", namespace), esl.String("filter", z.Name.Value()))
			continue
		}

		q.Enqueue(&SizeWorker{
			namespace: namespace,
			ctx:       cta,
			rep:       z.NamespaceSize,
			depth:     z.Depth,
			ctl:       c,
		})
	}
	q.Wait()
	return nil
}

func (z *Size) Test(c app_control.Control) error {
	err := rc_exec.Exec(c, &Size{}, func(r rc_recipe.Recipe) {
		rc := r.(*Size)
		rc.Name = mo_string.NewOptional(qtr_endtoend.TestTeamFolderName)
		rc.IncludeTeamFolder = false
		rc.Depth = 1

	})
	if err != nil {
		return err
	}
	return qtr_endtoend.TestRows(c, "namespace_size", func(cols map[string]string) error {
		if _, ok := cols["input.namespace_id"]; !ok {
			return errors.New("`namespace_id` is not found")
		}
		if _, ok := cols["result.size"]; !ok {
			return errors.New("`size` is not found")
		}
		return nil
	})
}
