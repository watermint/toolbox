package file

import (
	"errors"
	"github.com/watermint/toolbox/domain/common/model/mo_string"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_namespace"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_namespace"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_profile"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type MsgList struct {
	ErrorScanFailed app_msg.Message
	ProgressScan    app_msg.Message
}

var (
	MList = app_msg.Apply(&MsgList{}).(*MsgList)
)

type ListWorker struct {
	namespace        *mo_namespace.Namespace
	idToMember       map[string]*mo_member.Member
	ctx              dbx_context.Context
	ctl              app_control.Control
	rep              rp_model.RowReport
	IncludeMediaInfo bool
	IncludeDeleted   bool
}

func (z *ListWorker) Exec() error {
	ui := z.ctl.UI()
	ui.Progress(MList.ProgressScan.
		With("NamespaceName", z.namespace.Name).
		With("NamespaceId", z.namespace.NamespaceId))

	l := z.ctl.Log().With(esl.Any("namespace", z.namespace))

	ctn := z.ctx.WithPath(dbx_context.Namespace(z.namespace.NamespaceId))

	opts := make([]sv_file.ListOpt, 0)
	if z.IncludeDeleted {
		opts = append(opts, sv_file.IncludeDeleted())
	}
	if z.IncludeMediaInfo {
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
		l.Debug("Unable to traverse", esl.Error(err))
		ui.Error(MList.ErrorScanFailed.
			With("NamespaceName", z.namespace.Name).
			With("NamespaceId", z.namespace.NamespaceId).
			With("Error", err.Error()))
		return err
	}
	return nil
}

type List struct {
	Peer                dbx_conn.ConnBusinessFile
	IncludeMediaInfo    bool
	IncludeDeleted      bool
	IncludeMemberFolder bool
	IncludeSharedFolder bool
	IncludeTeamFolder   bool
	Name                mo_string.OptionalString
	NamespaceFile       rp_model.RowReport
}

func (z *List) Preset() {
	z.IncludeTeamFolder = true
	z.IncludeSharedFolder = true
	z.IncludeMemberFolder = false
	z.NamespaceFile.SetModel(
		&mo_namespace.NamespaceEntry{},
		rp_model.HiddenColumns(
			"namespace_id",
			"file_id",
			"revision",
			"content_hash",
			"shared_folder_id",
			"parent_shared_folder_id",
		),
	)
}

func (z *List) Exec(c app_control.Control) error {
	l := c.Log()
	if err := z.NamespaceFile.Open(); err != nil {
		return err
	}

	admin, err := sv_profile.NewTeam(z.Peer.Context()).Admin()
	if err != nil {
		return err
	}
	l.Debug("Run as admin", esl.Any("admin", admin))

	members, err := sv_member.New(z.Peer.Context()).List()
	if err != nil {
		return err
	}
	idToMember := mo_member.MapByTeamMemberId(members)

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
		case z.IncludeMemberFolder && namespace.NamespaceType == "app_folder":
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

		q.Enqueue(&ListWorker{
			namespace:        namespace,
			idToMember:       idToMember,
			ctx:              cta,
			rep:              z.NamespaceFile,
			IncludeDeleted:   z.IncludeDeleted,
			IncludeMediaInfo: z.IncludeMediaInfo,
			ctl:              c,
		})
	}
	q.Wait()
	return nil
}

func (z *List) Test(c app_control.Control) error {
	err := rc_exec.Exec(c, &List{}, func(r rc_recipe.Recipe) {
		rc := r.(*List)
		rc.Name = mo_string.NewOptional(qtr_endtoend.TestTeamFolderName)
	})
	if err != nil {
		return err
	}
	return qtr_endtoend.TestRows(c, "namespace_file", func(cols map[string]string) error {
		if _, ok := cols["namespace_name"]; !ok {
			return errors.New("`namespace_name` is not found")
		}
		if _, ok := cols["path_display"]; !ok {
			return errors.New("`path_display` is not found")
		}
		return nil
	})
}
