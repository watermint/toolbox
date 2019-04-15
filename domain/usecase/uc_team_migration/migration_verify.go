package uc_team_migration

// Verify content
func (z *migrationImpl) Verify(ctx Context) (err error) {
	z.log().Info("Verify team folders")
	if err = z.teamFolderMirror.VerifyScope(ctx.ContextTeamFolder()); err != nil {
		return err
	}

	return nil
}
