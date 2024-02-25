package file

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type List struct {
	Peer                         dbx_conn.ConnScopedTeam
	Path                         mo_path.DropboxPath
	MemberEmail                  string
	Recursive                    bool
	IncludeDeleted               bool
	IncludeMountedFolders        bool
	IncludeExplicitSharedMembers bool
	FileList                     rp_model.RowReport
}

func (z *List) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeMembersRead,
		dbx_auth.ScopeFilesMetadataRead,
		dbx_auth.ScopeTeamDataMember,
	)
	z.FileList.SetModel(
		&mo_file.ConcreteEntry{},
		rp_model.HiddenColumns(
			"id",
			"path_lower",
			"revision",
			"content_hash",
			"shared_folder_id",
			"parent_shared_folder_id",
		),
	)
}

func (z *List) Exec(c app_control.Control) error {
	if err := z.FileList.Open(); err != nil {
		return err
	}
	member, err := sv_member.New(z.Peer.Client()).ResolveByEmail(z.MemberEmail)
	if err != nil {
		return err
	}

	opts := make([]sv_file.ListOpt, 0)
	opts = append(opts, sv_file.IncludeDeleted(z.IncludeDeleted))
	opts = append(opts, sv_file.Recursive(z.Recursive))
	opts = append(opts, sv_file.IncludeHasExplicitSharedMembers(z.IncludeExplicitSharedMembers))
	opts = append(opts, sv_file.IncludeMountedFolders(z.IncludeMountedFolders))

	return sv_file.NewFiles(z.Peer.Client().AsMemberId(member.TeamMemberId)).ListEach(
		z.Path,
		func(entry mo_file.Entry) {
			z.FileList.Row(entry.Concrete())
		},
		opts...,
	)
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &List{}, func(r rc_recipe.Recipe) {
		m := r.(*List)
		m.MemberEmail = "john@example.com"
		m.Path = qtr_endtoend.NewTestDropboxFolderPath()
	})
}
