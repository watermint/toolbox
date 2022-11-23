package uc_insight_file

// FileScanSession handles scanning session.
// THe implementation may cache results.
type FileScanSession interface {
	// ScanNamespace scans a namespace then returns NamespaceFiles upon success
	ScanNamespace(namespaceId string) (nsf NamespaceFiles, err error)

	// ScanSelf scans my individual Dropbox
	ScanSelf() (mf MemberFiles, err error)

	// ScanMember scans a member then returns MemberFiles upon success
	ScanMember(teamMemberId string) (mf MemberFiles, err error)

	// ScanMembers scans all members then returns team_member_id -> MemberFile map upon success.
	ScanMembers() (mf map[string]MemberFiles, err error)

	// ScanTeamManagedFolders scans team folder or top level folder of team space.
	ScanTeamManagedFolders() (tlf map[string]NamespaceFiles, err error)
}
