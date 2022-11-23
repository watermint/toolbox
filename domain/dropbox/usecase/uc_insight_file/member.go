package uc_insight_file

type MemberFiles interface {
	CountNamespaces() (numNamespace int64, err error)
	// CountTotalFiles counts total files (excludes mountable namespace files).
	CountTotalFiles() (numFiles int64, err error)
	// CountTotalFolders counts total folders (excludes mountable namespace files).
	CountTotalFolders() (numFolders int64, err error)
	SumTotalFileSize() (sizeFiles int64, err error)

	MountableNamespaces() (nsf []NamespaceFiles, err error)
	MountedNamespaces() (map[string]NamespaceFiles, error)
}
