package uc_insight_file

import (
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
)

// 160TB -> 160M

// Session mgmt
// table: `files_scan_completed` has completed session_ids
// table: `files_scan_transaction` has session_id + state (started, finished)

type NamespaceRepository interface {
	AddNamespace(namespace *mo_sharedfolder.SharedFolder) error
	AddEntry(entry mo_file.Entry) error
}

type NamespaceScanner interface {
	Scan() error
}

// NamespaceFiles stores scanned file/folder metadata under the namespace.
type NamespaceFiles interface {
	CountTotalFiles() (numFiles int64, err error)
	CountTotalFolders() (numFolders int64, err error)
	SumTotalFileSize() (sizeFiles int64, err error)

	// Name of the namespace. Returns empty string if the namespace is root namespace.
	Name() string

	// DescendantFolders returns namespace with immediate descendant folder path
	DescendantFolders() (nsf []NamespaceFiles, err error)

	// DescendantNamespaces returns sub-namespaces with relative paths
	DescendantNamespaces() (nsf map[string]NamespaceFiles, err error)
}
