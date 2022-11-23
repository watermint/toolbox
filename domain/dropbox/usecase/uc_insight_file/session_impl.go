package uc_insight_file

type fileScanSessionImpl struct {
	dataPath  string
	sessionId string
}

func (z fileScanSessionImpl) scanNamespace(namespaceId string) error {
	//TODO implement me
	panic("implement me")
}

func (z fileScanSessionImpl) ScanNamespace(namespaceId string) (nsf NamespaceFiles, err error) {
	//TODO implement me
	panic("implement me")
}

func (z fileScanSessionImpl) ScanMember(teamMemberId string) (mf MemberFiles, err error) {
	//TODO implement me
	panic("implement me")
}

func (z fileScanSessionImpl) ScanMembers() (mf map[string]MemberFiles, err error) {
	//TODO implement me
	panic("implement me")
}

func (z fileScanSessionImpl) ScanTeamManagedFolders() (tlf map[string]NamespaceFiles, err error) {
	//TODO implement me
	panic("implement me")
}
