package uc_team_migration

func (z *migrationImpl) Migrate(ctx Context) (err error) {
	//if err = z.Bridge(ctx); err != nil {
	//	return err
	//}

	if err = z.Transfer(ctx); err != nil {
		return err
	}

	if err = z.Permissions(ctx); err != nil {
		return err
	}

	if err = z.Cleanup(ctx); err != nil {
		return err
	}

	return nil
}
