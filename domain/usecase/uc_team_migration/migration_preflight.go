package uc_team_migration

func (z *migrationImpl) Preflight(ctx Context) (err error) {
	if err = z.Inspect(ctx); err != nil {
		return err
	}

	if err = z.Preserve(ctx); err != nil {
		return err
	}

	return nil
}
