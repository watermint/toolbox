package cmd_teamfolder

import (
	"github.com/watermint/toolbox/dbx_api/dbx_file"
	"github.com/watermint/toolbox/dbx_api/dbx_namespace"
	"go.uber.org/zap"
	"testing"
)

func TestNamespaceSizes_OnFile(t *testing.T) {
	log, err := zap.NewDevelopment()
	if err != nil {
		t.Error(err)
	}

	ns := NamespaceSizes{
		Depth: 2,
	}
	ns.Init()

	ns.OnFile(
		&dbx_namespace.NamespaceFile{
			Namespace: &dbx_namespace.Namespace{
				NamespaceId: "ns:100",
			},
			File: &dbx_file.File{
				Name:      "park.jpg",
				FileId:    "id:1000",
				PathLower: "/park.jpg",
				Size:      1,
			},
		},
	)
	ns.OnFile(
		&dbx_namespace.NamespaceFile{
			Namespace: &dbx_namespace.Namespace{
				NamespaceId: "ns:100",
			},
			File: &dbx_file.File{
				Name:      "park.jpg",
				FileId:    "id:1001",
				PathLower: "/a/park.jpg",
				Size:      2,
			},
		},
	)
	ns.OnFile(
		&dbx_namespace.NamespaceFile{
			Namespace: &dbx_namespace.Namespace{
				NamespaceId: "ns:100",
			},
			File: &dbx_file.File{
				Name:      "park.jpg",
				FileId:    "id:1002",
				PathLower: "/a/b/park.jpg",
				Size:      3,
			},
		},
	)
	ns.OnFile(
		&dbx_namespace.NamespaceFile{
			Namespace: &dbx_namespace.Namespace{
				NamespaceId: "ns:100",
			},
			File: &dbx_file.File{
				Name:      "park.jpg",
				FileId:    "id:1003",
				PathLower: "/a/b/c/park.jpg",
				Size:      5,
			},
		},
	)
	ns.OnFile(
		&dbx_namespace.NamespaceFile{
			Namespace: &dbx_namespace.Namespace{
				NamespaceId: "ns:100",
			},
			File: &dbx_file.File{
				Name:      "park.jpg",
				FileId:    "id:1004",
				PathLower: "/a/b/c/d/park.jpg",
				Size:      7,
			},
		},
	)

	if ns100, ok := ns.Sizes["ns:100/"]; !ok {
		t.Error("expected ns not found")
	} else if ns100.Size != 18 {
		t.Error("invalid sum of size")
	} else if ns100.FileCount != 5 {
		t.Error("invalid file count")
	} else if ns100.FolderCount != 0 {
		t.Error("invalid folder count")
	} else if ns100.DescendantCount != 5 {
		t.Error("invalid descendant count")
	}

	ns.OnFolder(
		&dbx_namespace.NamespaceFolder{
			Namespace: &dbx_namespace.Namespace{
				NamespaceId: "ns:100",
			},
			Folder: &dbx_file.Folder{
				Name:      "park.jpg",
				FolderId:  "id:2000",
				PathLower: "/a",
			},
		},
	)

	ns.OnFolder(
		&dbx_namespace.NamespaceFolder{
			Namespace: &dbx_namespace.Namespace{
				NamespaceId: "ns:100",
			},
			Folder: &dbx_file.Folder{
				Name:      "park.jpg",
				FolderId:  "id:2001",
				PathLower: "/a/b",
			},
		},
	)

	ns.OnFolder(
		&dbx_namespace.NamespaceFolder{
			Namespace: &dbx_namespace.Namespace{
				NamespaceId: "ns:100",
			},
			Folder: &dbx_file.Folder{
				Name:      "park.jpg",
				FolderId:  "id:2002",
				PathLower: "/a/b/c",
			},
		},
	)

	if ns100, ok := ns.Sizes["ns:100/"]; !ok {
		t.Error("expected ns not found")
	} else if ns100.Size != 18 {
		t.Error("invalid sum of size")
	} else if ns100.FileCount != 5 {
		t.Error("invalid file count")
	} else if ns100.FolderCount != 3 {
		t.Error("invalid folder count")
	} else if ns100.DescendantCount != 8 {
		t.Error("invalid descendant count")
	}

	ns.OnFolder(
		&dbx_namespace.NamespaceFolder{
			Namespace: &dbx_namespace.Namespace{
				NamespaceId: "ns:100",
			},
			Folder: &dbx_file.Folder{
				Name:      "park.jpg",
				FolderId:  "id:2003",
				PathLower: "/a/b/c/d",
			},
		},
	)

	for p, ns := range ns.Sizes {
		log.Info("Size",
			zap.String("path", p),
			zap.Int64("descendant", ns.DescendantCount),
			zap.Int64("folder", ns.FolderCount),
			zap.Int64("file", ns.FileCount),
			zap.Int64("size", ns.Size),
		)
	}

	if ns100, ok := ns.Sizes["ns:100/"]; !ok {
		t.Error("expected ns not found")
	} else if ns100.Size != 18 {
		t.Error("invalid sum of size")
	} else if ns100.FileCount != 5 {
		t.Error("invalid file count")
	} else if ns100.FolderCount != 4 {
		t.Error("invalid folder count")
	} else if ns100.DescendantCount != 9 {
		t.Error("invalid descendant count")
	}

	if ns100, ok := ns.Sizes["ns:100/a"]; !ok {
		t.Error("expected ns not found")
	} else if ns100.Size != 17 {
		t.Error("invalid sum of size")
	} else if ns100.FileCount != 4 {
		t.Error("invalid file count")
	} else if ns100.FolderCount != 3 {
		t.Error("invalid folder count")
	} else if ns100.DescendantCount != 7 {
		t.Error("invalid descendant count")
	}

	if ns100, ok := ns.Sizes["ns:100/a/b"]; !ok {
		t.Error("expected ns not found")
	} else if ns100.Size != 15 {
		t.Error("invalid sum of size")
	} else if ns100.FileCount != 3 {
		t.Error("invalid file count")
	} else if ns100.FolderCount != 2 {
		t.Error("invalid folder count")
	} else if ns100.DescendantCount != 5 {
		t.Error("invalid descendant count")
	}
}
