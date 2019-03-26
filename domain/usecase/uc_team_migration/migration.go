package uc_team_migration

type Migration interface {
	// Do entire migration process.
	Migrate() (err error)

	// Inspect team status.
	// Ensure both team allow externally sharing shared folders.
	// Ensure which contents are not migrated by this migration.
	Inspect(ctx Context) (err error)

	// Preserve members, groups, and sharing status.
	Preserve(ctx Context) (err error)

	// Bridge shared folders.
	// Share all shared folders to destination admin.
	Bridge(ctx Context) (err error)

	// Mount
	Mount(ctx Context) (err error)

	// Mirror team folders.
	Content(ctx Context) (err error)

	// Verify contents.
	Verify(ctx Context) (err error)

	// Transfer members.
	// Convert accounts into Basic, and invite from destination team.
	Transfer(ctx Context) (err error)

	// Mirror permissions.
	// Create groups, invite members to shared folders or nested folders,
	// leave destination admin from bridged shared folders.
	Permissions(ctx Context) (err error)

	// Unmount
	Unmount(ctx Context) (err error)

	// Restore state.
	// Restore mount state.
	Restore(ctx Context) (err error)

	// Cleanup
	Cleanup(ctx Context) (err error)
}

// Migration context. Migration scope includes mutable states like permissions.
type Context interface {
}
