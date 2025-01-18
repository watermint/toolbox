package partial

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_profile"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_team"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_teamfolder"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_file_mirror"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type Replication struct {
	rc_recipe.RemarkIrreversible
	Src                           dbx_conn.ConnScopedTeam
	Dst                           dbx_conn.ConnScopedTeam
	SrcTeamFolderName             string
	DstTeamFolderName             string
	SrcPath                       mo_path.DropboxPath
	DstPath                       mo_path.DropboxPath
	ErrSrcTeamFolderNotFound      app_msg.Message
	ErrDstTeamFolderNotFound      app_msg.Message
	ErrorTeamSpaceNotSupportedSrc app_msg.Message
	ErrorTeamSpaceNotSupportedDst app_msg.Message
	BasePath                      mo_string.SelectString
}

func (z *Replication) Preset() {
	z.Src.SetPeerName("src")
	z.Dst.SetPeerName("dst")
	z.Src.SetScopes(
		dbx_auth.ScopeFilesContentRead,
		dbx_auth.ScopeFilesContentWrite,
		dbx_auth.ScopeTeamDataMember,
		dbx_auth.ScopeTeamDataTeamSpace,
		dbx_auth.ScopeTeamInfoRead,
	)
	z.Dst.SetScopes(
		dbx_auth.ScopeAccountInfoWrite,
		dbx_auth.ScopeFilesContentRead,
		dbx_auth.ScopeFilesContentWrite,
		dbx_auth.ScopeTeamDataMember,
		dbx_auth.ScopeTeamDataTeamSpace,
		dbx_auth.ScopeTeamInfoRead,
	)
	z.BasePath.SetOptions(
		dbx_filesystem.BaseNamespaceDefaultInString,
		dbx_filesystem.BaseNamespaceTypesInString...,
	)
}

func (z *Replication) Exec(c app_control.Control) error {
	if ok, _ := sv_team.UnlessTeamFolderApiSupported(z.Src.Client()); ok {
		c.UI().Error(z.ErrorTeamSpaceNotSupportedSrc)
		return errors.New("team space is not supported by this command")
	}
	if ok, _ := sv_team.UnlessTeamFolderApiSupported(z.Dst.Client()); ok {
		c.UI().Error(z.ErrorTeamSpaceNotSupportedDst)
		return errors.New("team space is not supported by this command")
	}

	l := c.Log()
	ui := c.UI()
	srcAdmin, err := sv_profile.NewTeam(z.Src.Client()).Admin()
	if err != nil {
		l.Debug("Unable to resolve src admin", esl.Error(err))
		return err
	}
	srcTeamFolder, err := sv_teamfolder.New(z.Src.Client()).ResolveByName(z.SrcTeamFolderName)
	if err != nil {
		l.Debug("Unable to find the src team folder", esl.Error(err))
		ui.Error(z.ErrSrcTeamFolderNotFound.With("Name", z.SrcTeamFolderName).With("Error", err))
		return err
	}
	l.Debug("Source team folder found", esl.Any("srcTeamFolder", srcTeamFolder))

	dstAdmin, err := sv_profile.NewTeam(z.Dst.Client()).Admin()
	if err != nil {
		l.Debug("Unable to resolve dst admin", esl.Error(err))
		return err
	}
	dstTeamFolder, err := sv_teamfolder.New(z.Dst.Client()).ResolveByName(z.DstTeamFolderName)
	if err != nil {
		l.Debug("Unable to find the dst team folder", esl.Error(err))
		ui.Error(z.ErrDstTeamFolderNotFound.With("Name", z.DstTeamFolderName).With("Error", err))
		return err
	}
	l.Debug("Dest team folder found", esl.Any("dstTeamFolder", dstTeamFolder))

	srcCtx := z.Src.Client().
		AsMemberId(srcAdmin.TeamMemberId, dbx_filesystem.AsNamespaceType(z.BasePath.Value())).
		WithPath(dbx_client.Namespace(srcTeamFolder.TeamFolderId))
	dstCtx := z.Dst.Client().
		AsMemberId(dstAdmin.TeamMemberId, dbx_filesystem.AsNamespaceType(z.BasePath.Value())).
		WithPath(dbx_client.Namespace(dstTeamFolder.TeamFolderId))

	mirror := uc_file_mirror.New(srcCtx, dstCtx)
	return mirror.Mirror(z.SrcPath, z.DstPath)
}

func (z *Replication) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Replication{}, func(r rc_recipe.Recipe) {
		m := r.(*Replication)
		m.DstTeamFolderName = "Sales"
		m.DstPath = mo_path.NewDropboxPath("/")
		m.SrcTeamFolderName = "Tokyo Office"
		m.SrcPath = mo_path.NewDropboxPath("/Sales")
	})
}
