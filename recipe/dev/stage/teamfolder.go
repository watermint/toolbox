package stage

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_folder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_profile"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_teamfolder"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"strings"
)

type Teamfolder struct {
	rc_recipe.RemarkSecret
	rc_recipe.RemarkExperimental
	Peer dbx_conn.ConnScopedTeam
}

func (z *Teamfolder) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeGroupsWrite,
		dbx_auth.ScopeFilesContentWrite,
	)
}

func (z *Teamfolder) Exec(c app_control.Control) error {
	teamFolderName := "Tokyo Branch 2"
	nestedFolderPlainName := "Organization"
	nestedFolderSharedName := "Sales"
	nestedFolderRestrictedName := "Report"
	adminGroupName := "toolbox-admin"

	// [Tokyo Branch] (Team folder)
	//  |
	//  +-- [Organization] (plain folder, not_synced)
	//  |
	//  +-- [Sales] (nested folder, not_synced)
	//       |
	//       +-- [Report] (nested folder, do not inherit, no external sharing)

	l := c.Log()

	// find admin
	admin, err := sv_profile.NewTeam(z.Peer.Context()).Admin()
	if err != nil {
		return err
	}

	// create team folder
	tf, err := sv_teamfolder.New(z.Peer.Context()).Create(teamFolderName)
	de := dbx_error.NewErrors(err)
	switch {
	case de == nil:
		l.Info("Team folder created", esl.Any("teamfolder", tf))
		break

	case de.IsFolderNameAlreadyUsed():
		l.Info("The folder already created")
		teamfolders, err := sv_teamfolder.New(z.Peer.Context()).List()
		if err != nil {
			l.Warn("Unable to retrieve team folder list", esl.Error(err))
			return err
		}

		for _, teamfolder := range teamfolders {
			if strings.ToLower(teamfolder.Name) == strings.ToLower(teamFolderName) {
				tf = teamfolder
				break
			}
		}
		if tf == nil {
			l.Warn("Team folder not found")
			return errors.New("team folder not found")
		}

		break

	default:
		l.Warn("Unable to create team folder", esl.Error(err))
		return err
	}

	tfCtx := z.Peer.Context().AsAdminId(admin.TeamMemberId).WithPath(dbx_context.Namespace(tf.TeamFolderId))

	// create sub folder : Organization
	folderOrganization, err := sv_file_folder.New(tfCtx).Create(mo_path.NewDropboxPath("/" + nestedFolderPlainName))
	de = dbx_error.NewErrors(err)
	switch {
	case de == nil:
		l.Info("Team folder created", esl.Any("folder", folderOrganization))
		break

	case de.Path().IsConflict():
		l.Info("The folder already created")
		folderOrganization, err = sv_file.NewFiles(tfCtx).Resolve(mo_path.NewDropboxPath("/" + nestedFolderPlainName))
		if err != nil {
			l.Warn("Unable to identify sub folder", esl.Error(err))
			return err
		}
		break

	default:
		l.Warn("Unable to create team folder", esl.Error(err))
		return err
	}

	// create nested folder : Sales
	folderSales, err := sv_sharedfolder.New(tfCtx).Create(mo_path.NewDropboxPath("/" + nestedFolderSharedName))
	de = dbx_error.NewErrors(err)
	switch {
	case de == nil:
		l.Info("Team folder created", esl.Any("folder", folderSales))
		break

	case de.BadPath().IsAlreadyShared():
		l.Info("The folder is already shared")
		folderSalesMeta, err := sv_file.NewFiles(tfCtx).Resolve(mo_path.NewDropboxPath("/" + nestedFolderSharedName))
		if err != nil {
			l.Warn("Unable to resolve nested folder", esl.Error(err))
			return err
		}

		folderSales, err = sv_sharedfolder.New(tfCtx).Resolve(folderSalesMeta.Concrete().SharedFolderId)
		if err != nil {
			l.Warn("Unable to resolve nested folder", esl.Error(err))
			return err
		}
		l.Info("Nested folder resolved", esl.Any("folder", folderSales))

	default:
		l.Warn("Unable to create team folder", esl.Error(err))
		return err
	}

	// create nested folder : Sales
	folderSalesReport, err := sv_sharedfolder.New(tfCtx).Create(mo_path.NewDropboxPath("/" + nestedFolderSharedName + "/" + nestedFolderRestrictedName))
	de = dbx_error.NewErrors(err)
	switch {
	case de == nil:
		l.Info("Team folder created", esl.Any("folder", folderSalesReport))
		break

	case de.BadPath().IsAlreadyShared():
		l.Info("The folder is already shared")
		folderSalesReportMeta, err := sv_file.NewFiles(tfCtx).Resolve(mo_path.NewDropboxPath("/" + nestedFolderSharedName + "/" + nestedFolderRestrictedName))
		if err != nil {
			l.Warn("Unable to resolve nested folder", esl.Error(err))
			return err
		}

		folderSalesReport, err = sv_sharedfolder.New(tfCtx).Resolve(folderSalesReportMeta.Concrete().SharedFolderId)
		if err != nil {
			l.Warn("Unable to resolve nested folder", esl.Error(err))
			return err
		}
		l.Info("Nested folder resolved", esl.Any("folder", folderSales))

	default:
		l.Warn("Unable to create team folder", esl.Error(err))
		return err
	}

	// Change sync setting
	folderOrganizationMeta, err := sv_file.NewFiles(tfCtx).Resolve(mo_path.NewDropboxPath("/" + nestedFolderPlainName))
	if err != nil {
		l.Warn("Unable to find meta", esl.Error(err))
		return err
	}
	folderSalesMeta, err := sv_file.NewFiles(tfCtx).Resolve(mo_path.NewDropboxPath("/" + nestedFolderSharedName))
	if err != nil {
		l.Warn("Unable to find meta", esl.Error(err))
		return err
	}

	updated, err := sv_teamfolder.New(z.Peer.Context()).UpdateSyncSetting(tf,
		sv_teamfolder.AddNestedSetting(folderOrganizationMeta, sv_teamfolder.SyncSettingNotSynced),
		sv_teamfolder.AddNestedSetting(folderSalesMeta, sv_teamfolder.SyncSettingNotSynced),
	)
	if err != nil {
		l.Warn("Unable to change", esl.Error(err))
		return err
	}
	l.Info("Sync settings updated", esl.Any("updated", updated))

	// Create toolbox admin group
	adminGroup, err := sv_group.New(z.Peer.Context()).Create(adminGroupName)
	de = dbx_error.NewErrors(err)
	switch {
	case de == nil:
		l.Info("The admin group created", esl.Any("group", adminGroup))

	case de.IsGroupNameAlreadyUsed():
		l.Info("The admin group already created")
		adminGroup, err = sv_group.New(z.Peer.Context()).ResolveByName(adminGroupName)
		if err != nil {
			l.Warn("Unable to find the admin group", esl.Error(err))
			return err
		}

	default:
		l.Warn("Unable to create the admin group", esl.Error(err))
		return err
	}

	// Add admin group to the team folder
	err = sv_sharedfolder_member.NewByTeamFolder(z.Peer.Context().AsAdminId(admin.TeamMemberId), tf).Add(
		sv_sharedfolder_member.AddByGroup(adminGroup, "editor"),
	)
	de = dbx_error.NewErrors(err)
	switch {
	case de == nil:
		l.Info("The admin group added to the team folder as editor")

	}

	// Do not inherit permission from parent : Sales/Report
	updatedFolderSalesReport, err := sv_sharedfolder.New(z.Peer.Context().AsMemberId(admin.TeamMemberId)).UpdateInheritance(folderSalesReport.SharedFolderId, sv_sharedfolder.AccessInheritanceNoInherit)
	if err != nil {
		l.Warn("Unable to change", esl.Error(err))
		return err
	}
	l.Info("Sync access inheritance updated", esl.Any("updated", updatedFolderSalesReport))

	return nil
}

func (z *Teamfolder) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Teamfolder{}, rc_recipe.NoCustomValues)
}
